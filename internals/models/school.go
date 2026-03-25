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
    PrincipalName    string    `json:"principalName"`
    CoordinationName string    `json:"choirCoordinator"`
    PaymentStatus    string    `json:"paymentStatus"`
    MediaList        string    `json:"mediaList"`
    ConfirmStatus    string    `json:"confirmationStatus"`
    Position         string    `json:"position"`
    PhoneNumber      string    `json:"phoneNumber"`
    ChoirSize        string    `json:"choirSize"`

    // Many-to-Many Pivot Relationship
    Events []Event `json:"events" gorm:"many2many:school_events;"`

    CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
