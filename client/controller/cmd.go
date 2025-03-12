package controller

import (
	"dp_client/global"
	"dp_client/service/chat"
	"fmt"
)

type CMD struct {
}

func (receiver *CMD) Run() {
	ctx := global.CreateLogContextByLogId(nil, global.NewLogId())

	conversationId := uint64(1001)
	chatService := chat.NewChatService(global.CHAT_URL, conversationId, false)

	for {
		var content string
		fmt.Println("请输入您的问题：>")
		fmt.Scan(&content)
		err := chatService.Run(ctx, content)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

}
