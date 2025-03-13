package chat

import (
	"context"
	"dp_client/global"
	"dp_client/storage"
	"dp_client/storage/model"
	"log/slog"

	"github.com/spf13/viper"
)

type ReqParamBuilder struct {
	// TODO: 上下文
	Model          string
	ConversationId uint64
	ctx            context.Context
}

func NewReqParamBuilder(ctx context.Context, modelName string, conversationId uint64) *ReqParamBuilder {
	return &ReqParamBuilder{
		ctx:            ctx,
		Model:          modelName,
		ConversationId: conversationId,
	}
}

type ChatReqParam struct {
	Model    string
	Messages []*ChatMessage
	Stream   bool
}

func (receiver *ReqParamBuilder) buildChatMessage(question string, prompt string) (retMessage []*ChatMessage) {

	message, err := storage.NewMessageStorage(receiver.ctx).QueryMessages(receiver.ConversationId, viper.GetInt("llm_context.max_limit"))
	if err != nil {
		global.Slog.Error("QueryMessages failed", slog.Any("err", err))
		return nil
	}

	// prompt
	retMessage = append(retMessage, &ChatMessage{
		Role:    string(global.RoleUser),
		Content: prompt,
	})

	//上下文
	retMessage = append(retMessage, convertMessageModel2ChatMessage(message)...)
	// 新问题
	retMessage = append(retMessage, &ChatMessage{
		Role:    string(global.RoleUser),
		Content: question,
	})

	return retMessage
}

// convertMessageModel2ChatMessage 需要翻转
func convertMessageModel2ChatMessage(message []*model.Message) []*ChatMessage {
	ret := make([]*ChatMessage, 0, len(message))
	n := len(message)
	for i := 0; i < n; i++ {
		ret = append(ret, &ChatMessage{
			Role:    message[n-i-1].Role,
			Content: message[n-i-1].Content,
		})
	}

	return ret
}

func (receiver *ReqParamBuilder) buildReq(question string, prompts string) *ChatReqParam {

	message := receiver.buildChatMessage(question, prompts)
	return &ChatReqParam{
		Model:    receiver.Model,
		Messages: message,
		Stream:   true,
	}
}
