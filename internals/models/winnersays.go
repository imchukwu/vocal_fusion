package models

import "time"

// WinnerSays represents a testimonial or quote from a winner/participant.
type WinnerSays struct {
    ID          int       `json:"id" gorm:"primaryKey"`
    UserID      int       `json:"user_id" gorm:"not null"` // Foreign Key to User (the person giving the testimonial)
    WinnerYear  int       `json:"winner_year"`             // Year of their winning/participation
    Message     string    `json:"message" gorm:"type:text"`// The testimonial content
    
    // GORM tags for relation
    User        User      `json:"user" gorm:"foreignKey:UserID"`

    CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}