package models

import "time"

type Message struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	SenderName string `json:"sender_name"`
	Email      string `json:"email"`
	Subject    string `json:"subject"`
	SchoolName string `json:"school_name"`
	Status     string `json:"status" gorm:"default:unread"`
	Phone      string `json:"phone"`
	Content    string `json:"content" gorm:"type:text"`
	ReplyToID  *int   `json:"reply_to_id"` // Pointer to parent message ID

	// GORM tags for relation
	ParentMessage *Message `json:"parent_message" gorm:"foreignKey:ReplyToID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
