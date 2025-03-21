package storage

import (
	"context"
	"dp_client/global"
	"dp_client/storage/model"
	"dp_client/storage/query"
	"errors"
	"log/slog"

	"gorm.io/gorm"
)

type AgentStorage struct {
	ctx context.Context
}

func NewAgentStorage(ctx context.Context) *AgentStorage {
	return &AgentStorage{
		ctx: ctx,
	}
}

func (storage *AgentStorage) CreateAgent(agent *model.Agent) (err error) {
	if agent == nil {
		err = errors.New("agent is nil")
		return
	}

	err = query.Agent.Create(agent)
	if err != nil {
		global.Slog.ErrorContext(storage.ctx, "CreateAgent failed ", slog.Any("err", err))
		return
	}
	return
}

func (storage *AgentStorage) QueryAgent(id uint64) (agent *model.Agent, err error) {
	agent, err = query.Agent.WithContext(storage.ctx).Where(query.Agent.ID.Eq(uint(id))).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		global.Slog.Error("QueryAgent failed", slog.Any("err", err))
		return
	}

	return
}

func (storage *AgentStorage) UpdateConversationId(ctx context.Context, id, conversationId uint64) error {
	_, err := query.Agent.WithContext(ctx).Where(query.Agent.ID.Eq(uint(id))).Update(query.Agent.ConversationId, conversationId)
	if err != nil {
		global.Slog.ErrorContext(ctx, "UpdateConversationId failed ", slog.Any("err", err))
		return err
	}

	return nil
}
