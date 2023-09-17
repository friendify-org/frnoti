package models

import "gorm.io/gorm"

type Email struct {
	gorm.Model
	SendTo  string `json:"sendTo"` // May be "all", "user1Id,user2Id" -> process latter
	Content string `json:"content"`
	Subject string `json:"subject"`
}
