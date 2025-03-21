package service

import (
	"context"
	"dp_client/global"
	"dp_client/service/ollama_agent"
	"dp_client/storage"
	"dp_client/storage/model"
	"log/slog"

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

func (receiver *AgentService) CreateAgent(ctx context.Context, LLMModelName string, Prompt string) error {
	agentModel := model.Agent{
		LLMModelName: LLMModelName,
		Prompt:       Prompt,
	}

	if err := storage.NewAgentStorage(ctx).CreateMessage(&agentModel); err != nil {
		global.Slog.Error("CreateMessage failed", slog.Any("err", err))
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

func (receiver *AgentService) Chat(ctx context.Context, content string) {

}

func GetAnswer(ctx context.Context, llm llms.Model, docRetrieved []schema.Document, input string) (string, error) {

	// 创建一个新的聊天消息历史记录
	history := memory.NewChatMessageHistory()
	// 将检索到的文档添加到历史记录中
	for _, doc := range docRetrieved {
		history.AddAIMessage(ctx, doc.PageContent)
	}
	// 使用历史记录创建一个新的对话缓冲区
	conversation := memory.NewConversationBuffer(memory.WithChatHistory(history))

	promptPrefix := `
	你是问答机器人，回答用户的问题
	\n\n
	`
	promptSuffix := `
	Begin!
	Chat History:
	{{.history}}
	Follow Up Input: {{.input}}
	Standalone question:`

	Prompt := prompts.PromptTemplate{
		Template:       promptPrefix + promptSuffix,
		TemplateFormat: prompts.TemplateFormatGoTemplate,
		InputVariables: []string{"input", "agent_scratchpad"},
		PartialVariables: map[string]any{
			"history": "",
		},
	}

	executor := agents.NewExecutor(
		agents.NewConversationalAgent(llm, nil,
			agents.WithCallbacksHandler(ollama_agent.GetLLMCallBackHandler()),
			agents.WithPrompt(Prompt)),
		agents.WithMemory(conversation),
	)

	// 设置链调用选项
	options := []chains.ChainCallOption{
		chains.WithTemperature(0.8),
	}
	res, err := chains.Run(ctx, executor, input, options...)
	if err != nil {
		return "", err
	}

	return res, nil
}
