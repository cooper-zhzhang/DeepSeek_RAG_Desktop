package controller

import (
	"dp_client/global"
	"dp_client/service"
	"dp_client/service/chat"
	"fmt"
	"log/slog"
)

type CMD struct {
}

func (receiver *CMD) Run() {

	agent := service.NewAgentService()
	ctx := global.CreateLogContextByLogId(nil, global.NewLogId())

	err := agent.CreateAgent(ctx, global.MODEL_NAME, "你扮演一个回答问题的机器人，使用很热情的语句回答问题,尽量使用中文")
	if err != nil {
		global.Slog.ErrorContext(ctx, "CreateAgent failed", slog.Any("err", err))
		return
	}

	conversationId := uint64(1002)
	chatService := chat.NewChatService(global.CHAT_URL, conversationId, false)
	for {
		var content string
		fmt.Println("请输入您的问题：>")
		fmt.Scan(&content)
		ctx = global.CreateLogContextByLogId(nil, global.NewLogId())
		//agent.Chat(ctx) TODO:
		err := chatService.Run(ctx, content)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
