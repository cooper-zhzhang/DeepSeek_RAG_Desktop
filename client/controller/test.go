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

	fileService, err := document.NewFileService(ctx, document.TextFileType, "西游记.txt")
	defer fileService.CloseFile(ctx)
	/*_, err = fileService.TextToChunks(ctx)
	if err != nil {
		global.Slog.Error("TextToChunks failed", slog.Any("err", err))
		return
	}
	err = fileService.StoreDocs(ctx, "xiyouji", nil)
	if err != nil {
		global.Slog.Error("StoreDocs failed", slog.Any("err", err))
		return
	}*/
	agent := service.NewAgentService()
	err = agent.CreateAgent(ctx, global.MODEL_NAME_1_5,
		"你扮演一个回答问题的机器人，使用很热情的语句回答问题,尽量使用中文", 10007)
	if err != nil {
		global.Slog.Error("CreateAgent failed", slog.Any("err", err))
		return
	}

	for {
		var input string
		fmt.Println("请输入您的问题：>")
		fmt.Scan(&input)
		ctx = global.CreateLogContextByLogId(nil, global.NewLogId())

		str, err := agent.ChatByRag(ctx, input, nil)
		if err != nil {
			global.Slog.Error("GetAnswer failed", slog.Any("err", err))
			return
		}

		fmt.Println(str)
	}

}
