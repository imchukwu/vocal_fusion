package models

import "time"

// Media represents a file (image, video, document) uploaded to the system.
type Media struct {
    ID        int       `json:"id" gorm:"primaryKey"`
    Thumbnail string    `json:"thumbnail"` // URL or path to the thumbnail (if applicable)
    Type      string    `json:"type" gorm:"not null"` // e.g., "image/jpeg", "video/mp4"
    Tag       string    `json:"tag"`      // e.g., "school_logo", "event_photo"
    Caption   string    `json:"caption"`
    Date      time.Time `json:"date" gorm:"not null"`
    
    CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}