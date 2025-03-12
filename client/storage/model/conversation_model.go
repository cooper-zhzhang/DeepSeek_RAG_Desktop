package model

import (
	"gorm.io/gorm"
)

type Conversation struct {
	AgentId         int64
	ConversationUId string
	gorm.Model
}
