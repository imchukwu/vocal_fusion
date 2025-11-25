package models

import "time"

type SchoolEvent struct {
    ID        int       `json:"id" gorm:"primaryKey"`

    SchoolID  int       `json:"school_id" gorm:"not null"`
    School    School    `json:"school" gorm:"foreignKey:SchoolID"`

    EventID   int       `json:"event_id" gorm:"not null"`
    Event     Event     `json:"event" gorm:"foreignKey:EventID"`

    Status    string    `json:"status" gorm:"default:'Registered'"` 
    // Optional extra fields
    Notes     string    `json:"notes"`

    CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
