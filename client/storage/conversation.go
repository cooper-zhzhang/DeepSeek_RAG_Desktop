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

type ConversationStorage struct {
	ctx context.Context
}

// TODO: storage可以做成单例
func NewConversationStorage(ctx context.Context) *ConversationStorage {
	return &ConversationStorage{
		ctx: ctx,
	}
}

func (storage *ConversationStorage) Create(agentId uint64, uid string) (uint64, error) {
	conversation := &model.Conversation{
		AgentId:         int64(agentId),
		ConversationUId: uid,
	}

	err := query.Conversation.WithContext(storage.ctx).Create(conversation)
	if err != nil {
		global.Slog.ErrorContext(storage.ctx, "CreateConversation failed ", slog.Any("err", err))
		return 0, err
	}

	return uint64(conversation.ID), nil
}

func (storage *ConversationStorage) QueryConversationById(id uint64) (conversation *model.Conversation, err error) {
	conversation, err = query.Conversation.WithContext(storage.ctx).Debug().Where(query.Conversation.ID.Eq(uint(id))).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		global.Slog.Error("QueryConversationById failed", slog.Any("err", err))
		return nil, err
	}

	return conversation, nil
}
