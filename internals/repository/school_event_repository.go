package repository

import (
	"gorm.io/gorm"
	"vocal_fusion/internals/models"
)

type SchoolEventRepository interface {
	RegisterSchoolForEvent(reg *models.SchoolEvent) error
	GetRegistrationsByEvent(eventID int) ([]models.SchoolEvent, error)
	GetRegistrationsBySchool(schoolID int) ([]models.SchoolEvent, error)
	UnregisterSchool(eventID int, schoolID int) error
}

type schoolEventRepository struct {
	db *gorm.DB
}

func NewSchoolEventRepository(db *gorm.DB) SchoolEventRepository {
	return &schoolEventRepository{db: db}
}

func (r *schoolEventRepository) RegisterSchoolForEvent(reg *models.SchoolEvent) error {
	return r.db.Create(reg).Error
}

func (r *schoolEventRepository) GetRegistrationsByEvent(eventID int) ([]models.SchoolEvent, error) {
	var regs []models.SchoolEvent
	err := r.db.
		Preload("School").
		Preload("Event").
		Where("event_id = ?", eventID).
		Find(&regs).Error
	return regs, err
}

func (r *schoolEventRepository) GetRegistrationsBySchool(schoolID int) ([]models.SchoolEvent, error) {
	var regs []models.SchoolEvent
	err := r.db.
		Preload("Event").
		Preload("School").
		Where("school_id = ?", schoolID).
		Find(&regs).Error
	return regs, err
}

func (r *schoolEventRepository) UnregisterSchool(eventID int, schoolID int) error {
	return r.db.
		Where("event_id = ? AND school_id = ?", eventID, schoolID).
		Delete(&models.SchoolEvent{}).
		Error
}
