package global

import (
	"context"
	"log/slog"
	"os"

	"github.com/google/uuid"
)

var Slog *slog.Logger

const LogIdKey = "logId"

// ContextHandler log handler的实现
type ContextHandler struct {
	slog.Handler
}

func (h *ContextHandler) Handle(ctx context.Context, record slog.Record) error {
	// 从context中获取logId
	if logId, ok := ctx.Value(LogIdKey).(string); ok {
		record.Add(LogIdKey, logId)
	}

	return slog.Default().Handler().Handle(ctx, record)
}

func createGlobalLog() *slog.Logger {
	textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelError,
	})
	// 创建一个自定义的ContextHandler
	contextHandler := &ContextHandler{Handler: textHandler}

	// 创建一个Logger
	logger := slog.New(contextHandler)
	return logger
}

func NewLogId() string {
	newUUID := uuid.New()
	logId := newUUID.String()

	return logId
}

func CreateLogContextByLogId(ctx context.Context, logId string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	// TODO:判断是否已经有logId
	return context.WithValue(ctx, LogIdKey, logId)
}

func init() {
	Slog = createGlobalLog()
}
