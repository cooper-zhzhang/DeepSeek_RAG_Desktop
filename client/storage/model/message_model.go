package model

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	Content        string
	Role           string
	ConversationId uint64 `gorm:"index:conv_id_created_at"`
	gorm.Model
	CreatedAt time.Time `gorm:"index:conv_id_created_at"`
}
