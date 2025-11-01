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
    PaymentStatus    string    `json:"payment_status"`   // e.g., "Pending", "Paid", "Canceled"
    // Assuming MediaList is a comma-separated string or a JSON array string in the DB
    MediaList        string    `json:"media_list"`
    ConfirmStatus    string    `json:"confirm_status"`   // e.g., "Confirmed", "Pending"
    CreatedAt        time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt        time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}