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

	input := "小明认识谁"
	fileService, _ := document.NewFileService(ctx, document.TextFileType, "text.txt")
	docs, err := fileService.TextToChunks(ctx)
	if err != nil {
		global.Slog.Error("TextToChunks failed", slog.Any("err", err))
		return
	}

	err = fileService.StoreDocs(ctx, docs)
	if err != nil {
		global.Slog.Error("StoreDocs failed", slog.Any("err", err))
		return
	}

	docs, err = fileService.UseRetriever(ctx, input, 10)
	if err != nil {
		global.Slog.Error("UseRetriever failed", slog.Any("err", err))
		return
	}

	result, err := service.GetAnswer(ctx, ollama_agent.GetLLMClient(ollama_agent.DEEP_SEEK_MODEL_1_DOT_5), docs, input)
	if err != nil {
		global.Slog.Error("GetAnswer failed", slog.Any("err", err))
		return
	}
	fmt.Println(result)
}
