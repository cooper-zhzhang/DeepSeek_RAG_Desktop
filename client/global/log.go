package global

import (
	"context"
	"log/slog"
	"os"

	"github.com/spf13/viper"

	"github.com/google/uuid"
)

var Slog *slog.Logger

const LogIdKey = "logId"

// ContextHandler log handler的实现
type ContextHandler struct {
	handler slog.Handler
}

func (h *ContextHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *ContextHandler) Handle(ctx context.Context, record slog.Record) error {
	// 添加 log ID
	if logId, ok := ctx.Value(LogIdKey).(string); ok {
		record.Add(LogIdKey, logId)
	}
	return h.handler.Handle(ctx, record)
}

func (h *ContextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ContextHandler{handler: h.handler.WithAttrs(attrs)}
}

func (h *ContextHandler) WithGroup(name string) slog.Handler {
	return &ContextHandler{handler: h.handler.WithGroup(name)}
}

func NewContextHandler(handler slog.Handler) slog.Handler {
	return &ContextHandler{handler: handler}
}

func getLogLevel() slog.Level {
	switch viper.GetString("log.log_level") {
	case "DEBUG":
		return slog.LevelDebug
	case "WARN":
		return slog.LevelWarn
	case "INFO":
		return slog.LevelInfo
	default: //"ERROR"
		return slog.LevelError
	}
}

func createGlobalLog() *slog.Logger {

	baseHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true, // 添加源文件和行号
		Level:     getLogLevel(),
	})

	// 添加自定义处理器
	handler := NewContextHandler(baseHandler)
	logger := slog.New(handler)
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
	} else {
		if _, ok := ctx.Value(LogIdKey).(string); ok {
			return ctx
		}
	}

	return context.WithValue(ctx, LogIdKey, logId)
}

func initLog() {
	Slog = createGlobalLog()
}
