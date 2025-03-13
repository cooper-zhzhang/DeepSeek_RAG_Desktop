package service

import (
	"context"
	"dp_client/global"
	"dp_client/storage"
	"log/slog"

	"github.com/google/uuid"
)

type ConversationService struct {
}

func NewConversationService() *ConversationService {
	return &ConversationService{}
}

func (receiver *ConversationService) Query(ctx context.Context, id uint64) (uint64, error) {

	conversation, err := storage.NewConversationStorage(ctx).QueryConversationById(id)
	if err != nil {
		global.Slog.ErrorContext(ctx, "QueryConversationById failed ", slog.Any("err", err))
		return 0, err
	}

	return uint64(conversation.ID), nil
}

func (receiver *ConversationService) Create(ctx context.Context, agentId uint64) (uint64, error) {
	uid, _ := uuid.NewUUID()
	id, err := storage.NewConversationStorage(ctx).Create(agentId, uid.String())
	if err != nil {
		global.Slog.ErrorContext(ctx, "CreateConversation failed ", slog.Any("err", err))
		return 0, err
	}
	return id, nil
}
