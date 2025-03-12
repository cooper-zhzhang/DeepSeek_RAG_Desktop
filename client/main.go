package main

import (
	"dp_client/global"
)

func main() {
	// TODO:添加一个conversation id
	ctx := global.CreateLogContextByLogId(nil, global.NewLogId())

	receiver := NewChatService(CHAT_URL, false)
	err := receiver.Run(ctx, "Do you know my name")
	if err != nil {
		return
	}
}
