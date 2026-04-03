package models

import "time"

type Event struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Date        time.Time `json:"date" gorm:"not null"`
	Time        string    `json:"time"`
	Location    string    `json:"location"`

	// Many-to-Many Pivot Relationship
	Schools []School `json:"schools" gorm:"many2many:school_events;"`

	// One-to-Many Relationship
	Media []Media `json:"media" gorm:"foreignKey:EventID"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
