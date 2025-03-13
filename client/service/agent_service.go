package service

import (
	"context"
	"dp_client/global"
	"dp_client/storage"
	"dp_client/storage/model"
	"log/slog"
)

type AgentService struct {
	//ConversationId int64
	AgentModel *model.Agent

	// TODO:包含chatservice conver 模型信息 prompt etc
	//以及用户信息
	//这是一个对外的类，包含
}

func NewAgentService() *AgentService {
	return &AgentService{
		//ConversationId: 1,
	}
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

func (receiver *AgentService) CreateAgent(ctx context.Context, LLMModelName string, Prompt string) {
	agentModel := model.Agent{
		LLMModelName: LLMModelName,
		Prompt:       Prompt,
	}

	if err := storage.NewAgentStorage(ctx).CreateMessage(&agentModel); err != nil {
		global.Slog.Error("CreateMessage failed", slog.Any("err", err))
		return
	}

	receiver.AgentModel = &agentModel
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
		err = storage.NewAgentStorage(ctx).UpdateConversationId(receiver.AgentModel.ID, conversationId)
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
