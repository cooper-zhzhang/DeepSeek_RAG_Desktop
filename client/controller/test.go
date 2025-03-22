package controller

import (
	"context"
	"dp_client/global"
	"dp_client/service"
	"dp_client/service/document"
	"dp_client/service/ollama_agent"
	"fmt"
	"log/slog"

	"github.com/tmc/langchaingo/llms"
)

type TestConsole struct {
}

func (receiver *TestConsole) Run() {
	ctx := global.CreateLogContextByLogId(nil, global.NewLogId())
	receiver.RAG(ctx)
}

func (receiver *TestConsole) CallByLLM(ctx context.Context) {
	text := llms.TextContent{"who are you"}
	part := llms.ContentPart(text)
	result, err := ollama_agent.GetLLMClient(ollama_agent.DEEP_SEEK_MODEL_7).GenerateContent(ctx, []llms.MessageContent{{
		Role:  llms.ChatMessageTypeHuman,
		Parts: []llms.ContentPart{part},
	}})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func (receiver *TestConsole) RAG(ctx context.Context) {

	fileService, _ := document.NewFileService(ctx, document.TextFileType, "text.txt")
	agent := service.NewAgentService()
	err := agent.CreateAgent(ctx, global.MODEL_NAME,
		"你扮演一个回答问题的机器人，使用很热情的语句回答问题,尽量使用中文", 10004)
	if err != nil {
		global.Slog.Error("CreateAgent failed", slog.Any("err", err))
		return
	}

	for {
		var input string
		fmt.Println("请输入您的问题：>")
		fmt.Scan(&input)
		ctx = global.CreateLogContextByLogId(nil, global.NewLogId())
		docs, err := fileService.UseRetriever(ctx, input, 10)
		if err != nil {
			global.Slog.Error("UseRetriever failed", slog.Any("err", err))
			return
		}

		if len(docs) > 0 {
			docs = nil
		}

		err = agent.ChatStream(ctx, input, docs, func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		})
		/*_, err = agent.ChatStreamByChains(ctx, input, docs, func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		})*/
		if err != nil {
			global.Slog.Error("GetAnswer failed", slog.Any("err", err))
			return
		}

		fmt.Println()
	}

}
