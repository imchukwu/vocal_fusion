package models

import "time"

// School represents a school entity in the system.
type School struct {
    ID               int       `json:"id" gorm:"primaryKey"`
    Name             string    `json:"name" gorm:"not null"`
    Address          string    `json:"address"`
    Email            string    `json:"email" gorm:"unique"`
    State            string    `json:"state"`
    City             string    `json:"city"`
    PrincipalName    string    `json:"principal_name"`
    CoordinationName string    `json:"coordination_name"`
    PaymentStatus    string    `json:"payment_status"`
    MediaList        string    `json:"media_list"`
    ConfirmStatus    string    `json:"confirm_status"`

    // Many-to-Many Pivot Relationship
    Events []Event `json:"events" gorm:"many2many:school_events;"`

    CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
