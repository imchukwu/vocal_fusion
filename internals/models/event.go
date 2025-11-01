package models

import "time"


type Event struct {
    ID        int       `json:"id" gorm:"primaryKey"`
    Title     string    `json:"title" gorm:"not null"`
    Type      string    `json:"type"` 
    Date      time.Time `json:"date" gorm:"not null"`
    Time      string    `json:"time"` 
    Location  string    `json:"location"`
    
    CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}