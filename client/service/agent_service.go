package service

import (
	"context"
	"dp_client/global"
	"dp_client/service/ollama_agent"
	inn_prompts "dp_client/service/prompts"
	"dp_client/storage"
	"dp_client/storage/model"
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"github.com/spf13/viper"

	"github.com/tmc/langchaingo/prompts"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/schema"
)

type AgentService struct {
	AgentModel *model.Agent

	// TODO:包含chatservice conver 模型信息 prompt etc
	//以及用户信息
	//这是一个对外的类，包含
}

func NewAgentService() *AgentService {
	return &AgentService{}
}

func (receiver *AgentService) GetAgent(ctx context.Context, id uint64) (bool, error) {
	agentModel, err := storage.NewAgentStorage(ctx).QueryAgent(id)
	if err != nil {
		global.Slog.Error("QueryAgent failed", slog.Any("err", err))
		return false, err
	}

	receiver.AgentModel = agentModel

	return false, nil
}

func (receiver *AgentService) SetDataset(ctx context.Context, datasetId uint64) error {

	return nil
}

func (receiver *AgentService) CreateAgent(ctx context.Context, LLMModelName string, Prompt string, conversationId uint64) error {
	agentModel := model.Agent{
		LLMModelName:   LLMModelName,
		Prompt:         Prompt,
		ConversationId: conversationId,
	}

	if err := storage.NewAgentStorage(ctx).CreateAgent(&agentModel); err != nil {
		global.Slog.Error("CreateAgent failed", slog.Any("err", err))
		return err
	}

	receiver.AgentModel = &agentModel
	return nil
}

func (receiver *AgentService) GetConversation(ctx context.Context) bool {

	if receiver.AgentModel == nil {
		return false
	}

	if receiver.AgentModel.ConversationId == 0 {
		conversationId, err := NewConversationService().Create(ctx, uint64(receiver.AgentModel.ID))
		if err != nil {
			global.Slog.ErrorContext(ctx, "CreateConversation failed ", slog.Any("err", err))
			return false
		}
		err = storage.NewAgentStorage(ctx).UpdateConversationId(ctx, uint64(receiver.AgentModel.ID), conversationId)
		if err != nil {
			global.Slog.ErrorContext(ctx, "UpdateConversationId failed", slog.Any("err", err))
			return false
		}
		receiver.AgentModel.ConversationId = conversationId
		return false
	}

	return true

}

func (receiver *AgentService) buildMessageHistory(ctx context.Context, docRetrieved []schema.Document) (*memory.ChatMessageHistory, error) {
	message, err := storage.NewMessageStorage(ctx).QueryMessages(receiver.AgentModel.ConversationId,
		viper.GetInt("llm_context.max_limit"))

	if err != nil {
		global.Slog.Error("QueryMessages failed", slog.Any("err", err))
		return nil, err
	}

	// 创建一个新的聊天消息历史记录
	history := memory.NewChatMessageHistory()
	// 将检索到的文档添加到历史记录中
	for _, doc := range docRetrieved {
		//ChatMessageTypeSystem
		mes := llms.SystemChatMessage{
			Content: doc.PageContent,
		}
		history.AddMessage(ctx, mes)
	}

	for i := len(message) - 1; i > 0; i-- {
		doc := message[i]
		//目前deepseek只有 两种角色，human 和 ai
		if doc.Role == string(model.RoleHuman) {
			_ = history.AddUserMessage(ctx, doc.Content)
		} else {
			_ = history.AddAIMessage(ctx, doc.Content)
		}
	}

	return history, nil
}

func (receiver *AgentService) retrieved(input string) {
	// TODO:实现从向量数据库中检索

	//receiver.AgentModel.DatasetID

}

func (receiver *AgentService) convertHistory2ChatMessages(ctx context.Context, history *memory.ChatMessageHistory, input string) (messages []llms.MessageContent) {
	if history == nil {
		return
	}

	hisMessages, _ := history.Messages(ctx)
	for _, message := range hisMessages {
		//if message.GetType() == llms.ChatMessageTypeHuman {
		messages = append(messages, llms.MessageContent{
			Role:  message.GetType(),
			Parts: []llms.ContentPart{llms.ContentPart(llms.TextContent{Text: message.GetContent()})},
		})
		//}
	}

	if input != "" {
		messages = append(messages, llms.MessageContent{
			Role:  llms.ChatMessageTypeHuman,
			Parts: []llms.ContentPart{llms.ContentPart(llms.TextContent{Text: input})},
		})
	}

	return
}

func (receiver *AgentService) ChatStream(ctx context.Context, input string, docRetrieved []schema.Document, readChunk func(ctx context.Context, chunk []byte) error) (err error) {

	llm := ollama_agent.GetLLMClient(ollama_agent.LLMName(receiver.AgentModel.LLMModelName))
	if llm == nil {
		err = errors.New("not found llm " + receiver.AgentModel.LLMModelName)
		global.Slog.ErrorContext(ctx, "GetLLMClient failed", slog.Any("err", err))
		return err
	}

	history, err := receiver.buildMessageHistory(ctx, docRetrieved)
	if err != nil {
		global.Slog.Error("buildMessageHistory failed", slog.Any("err", err))
		return err
	}

	messages := receiver.convertHistory2ChatMessages(ctx, history, input)

	resp, err := llm.GenerateContent(ctx, messages, llms.WithStreamingFunc(readChunk))
	if err != nil {
		global.Slog.ErrorContext(ctx, "GenerateContent failed", slog.Any("err", err))
		return err
	}

	sb := strings.Builder{}
	for _, choice := range resp.Choices {
		sb.WriteString(choice.Content)
	}

	err = receiver.RecordMessage(ctx, input, sb.String())
	if err != nil {
		global.Slog.ErrorContext(ctx, "RecordMessage failed", slog.Any("err", err))
		return err
	}

	return nil
}
func (receiver *AgentService) Chat(ctx context.Context, input string, docRetrieved []schema.Document) (response string, err error) {

	history, err := receiver.buildMessageHistory(ctx, docRetrieved)
	if err != nil {
		global.Slog.Error("buildMessageHistory failed", slog.Any("err", err))
		return "", err
	}
	conversation := memory.NewConversationBuffer(memory.WithChatHistory(history))
	// TODO: 放在 prompts文件夹下
	Prompt := prompts.PromptTemplate{
		Template:       receiver.AgentModel.Prompt + inn_prompts.DefaultPromptSuffix,
		TemplateFormat: prompts.TemplateFormatGoTemplate,
		InputVariables: []string{"input", "agent_scratchpad"},
		PartialVariables: map[string]any{
			"history": "",
		},
	}

	llm := ollama_agent.GetLLMClient(ollama_agent.LLMName(receiver.AgentModel.LLMModelName))
	executor := agents.NewExecutor(
		agents.NewConversationalAgent(llm, nil,
			agents.WithCallbacksHandler(ollama_agent.GetLLMCallBackHandler()),
			agents.WithPrompt(Prompt)),
		agents.WithMemory(conversation),
	)

	// 设置链调用选项
	options := []chains.ChainCallOption{
		chains.WithTemperature(0.8),
		chains.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Println(string(chunk))
			return nil
		}),
	}
	response, err = chains.Run(ctx, executor, input, options...)
	if err != nil {
		global.Slog.ErrorContext(ctx, "chains.Run failed", slog.Any("err", err))
		return "", err
	}

	err = receiver.RecordMessage(ctx, input, response)
	if err != nil {
		global.Slog.ErrorContext(ctx, "RecordMessage failed", slog.Any("err", err))
		return "", err
	}

	return response, nil
}

func (receiver *AgentService) RecordMessage(ctx context.Context, request, response string) error {
	var messages []*model.Message
	messages = append(messages, &model.Message{
		Content:        request,
		Role:           string(model.RoleHuman),
		ConversationId: receiver.AgentModel.ConversationId,
	}, &model.Message{
		Content:        response,
		Role:           string(model.RoleAI),
		ConversationId: receiver.AgentModel.ConversationId,
	})

	err := storage.NewMessageStorage(ctx).CreateMessages(messages)
	if err != nil {
		global.Slog.ErrorContext(ctx, "insertMessage failed", slog.Any("err", err))
		return err
	}

	return nil
}
