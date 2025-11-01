package models

import "time"

// FAQ represents a frequently asked question and its answer.
type FAQ struct {
    ID        int       `json:"id" gorm:"primaryKey"`
    Subject   string    `json:"subject" gorm:"not null"` // The Question
    Message   string    `json:"message" gorm:"type:text"` // The Answer
    
    CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}