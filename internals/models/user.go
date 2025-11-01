package models

import "time"

type User struct {
	ID         int        `json:"id" gorm:"primaryKey"`
	Name       string     `json:"name" gorm:"not null"`
	Email      string     `json:"email" gorm:"unique;not null"`
	Phone      string     `json:"phone"`
	Role       string     `json:"role" gorm:"default:school_user"`
	LastLogin  *time.Time `json:"last_login"`
	Status     string     `json:"status" gorm:"default:active"`
	MediaID    *int       `json:"media_id"` // ✅ now nullable!
	Position   string     `json:"position"`

	Media      Media      `json:"media" gorm:"foreignKey:MediaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt  time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}
