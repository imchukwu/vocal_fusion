package repository

import (
	"vocal_fusion/internals/models"
	"gorm.io/gorm"
)

type SettingsRepository interface {
	GetSettings() (*models.Settings, error)
	UpdateSettings(settings *models.Settings) error
}

type settingsRepository struct {
	DB *gorm.DB
}

func NewSettingsRepository(db *gorm.DB) SettingsRepository {
	return &settingsRepository{DB: db}
}

func (r *settingsRepository) GetSettings() (*models.Settings, error) {
	var settings models.Settings
	// We assume there's only one row for settings (ID=1)
	err := r.DB.FirstOrCreate(&settings, models.Settings{ID: 1}).Error
	return &settings, err
}

func (r *settingsRepository) UpdateSettings(settings *models.Settings) error {
	settings.ID = 1 // Ensure we always update the same row
	return r.DB.Save(settings).Error
}
