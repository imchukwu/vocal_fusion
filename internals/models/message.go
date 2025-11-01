package models

import "time"

type Message struct {
    ID          int       `json:"id" gorm:"primaryKey"`
    SenderName  string    `json:"sender_name"`
    Email       string    `json:"email"`
    Subject     string    `json:"subject"`
    Date        time.Time `json:"date" gorm:"autoCreateTime"`
    Status      string    `json:"status" gorm:"default:unread"`
    Phone       string    `json:"phone"`
    Content     string    `json:"content" gorm:"type:text"`
    SchoolID    *int      `json:"school_id"` // ✅ Pointer makes it nullable

    // GORM tags for relation
    School      School    `json:"school" gorm:"foreignKey:SchoolID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

    CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
