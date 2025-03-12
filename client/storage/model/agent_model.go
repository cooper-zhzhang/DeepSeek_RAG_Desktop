package model

import "gorm.io/gorm"

type Agent struct {
	ConversationId int64
	LLMModelName   string
	Prompt         string
	gorm.Model
}
