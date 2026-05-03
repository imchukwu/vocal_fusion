package models

import "time"

// Settings represents global site configuration.
type Settings struct {
    ID             int       `json:"id" gorm:"primaryKey"`
    SiteName       string    `json:"siteName"`
    ContactEmail   string    `json:"contactEmail"`
    ContactPhone   string    `json:"contactPhone"`
    Address        string    `json:"address"`
    FacebookURL    string    `json:"facebookUrl"`
    TwitterURL     string    `json:"twitterUrl"`
    InstagramURL   string    `json:"instagramUrl"`
    YoutubeURL     string    `json:"youtubeUrl"`
    RegistrationFee float64   `json:"registrationFee"`
    
    UpdatedAt      time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
