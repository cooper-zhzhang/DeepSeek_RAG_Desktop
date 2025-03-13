package storage

import (
	"context"
	"dp_client/global"
	"dp_client/storage/model"
	"dp_client/storage/query"
	"errors"
	"log/slog"
)

type MessageStorage struct {
	ctx context.Context
}

func NewMessageStorage(ctx context.Context) *MessageStorage {
	return &MessageStorage{
		ctx: ctx,
	}
}

func (storage *MessageStorage) QueryMessages(conversationId uint64, limit int) (messages []*model.Message, err error) {

	messages, err = query.Message.WithContext(storage.ctx).Debug().
		Where(query.Message.ConversationId.Eq(conversationId)).Limit(limit).Order(query.Message.CreatedAt.Desc()).Find()

	if err != nil {
		global.Slog.Error("QueryMessages failed", slog.Any("err", err))
		return
	}

	return
}

func (storage *MessageStorage) CreateMessage(message *model.Message) (err error) {

	if message == nil {
		err = errors.New("message is nil")
		return
	}

	err = query.Message.WithContext(storage.ctx).Create(message)
	if err != nil {
		global.Slog.ErrorContext(storage.ctx, "CreateMessage failed ", slog.Any("err", err))
		return
	}

	return
}
