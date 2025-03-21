package model

import "gorm.io/gorm"

type Agent struct {
	ConversationId uint64
	LLMModelName   string
	Prompt         string
	DatasetID      uint64
	gorm.Model
}
