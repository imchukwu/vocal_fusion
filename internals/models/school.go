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
    ProofOfPayment   string    `json:"proofOfPayment"`
    ConfirmStatus    bool      `json:"confirmationStatus"`
    Position         string    `json:"position"`
    PhoneNumber      string    `json:"phoneNumber"`
    ChoirSize        int       `json:"choirSize"`

    // Many-to-Many Pivot Relationship
    Events []Event `json:"events" gorm:"many2many:school_events;"`
    Registrations []SchoolEvent `json:"registrations" gorm:"foreignKey:SchoolID"`

    CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
