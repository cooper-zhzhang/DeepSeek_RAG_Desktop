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
	//receiver.CallByLLM(ctx)

	receiver.RAG(ctx)
	//result, err := ollama_agent.GetLLMClient().Call(ctx, "who are you")

	//llm.GenerateContent(context.Background(), []llms.MessageContent{})

	/*	*/

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
	fileService := document.NewFileService()
	docs, err := fileService.TextToChunks(ctx, "text.txt")
	if err != nil {
		global.Slog.Error("TextToChunks failed", slog.Any("err", err))
		return
	}

	result, err := service.GetAnswer(ctx, ollama_agent.GetLLMClient(ollama_agent.DEEP_SEEK_MODEL_7), docs, "小明认识谁")
	if err != nil {
		global.Slog.Error("GetAnswer failed", slog.Any("err", err))
		return
	}
	fmt.Println(result)
}
