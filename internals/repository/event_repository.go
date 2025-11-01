package repository

import (
	"vocal_fusion/internals/models"

	"gorm.io/gorm"
)

type EventRepository interface {
	CreateEvent(event *models.Event) error
	GetAllEvents() ([]models.Event, error)
	GetEventByID(id uint) (*models.Event, error)
	UpdateEvent(event *models.Event) error
	DeleteEvent(id uint) error
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db}
}

func (r *eventRepository) CreateEvent(event *models.Event) error {
	return r.db.Create(event).Error
}

func (r *eventRepository) GetAllEvents() ([]models.Event, error) {
	var events []models.Event
	err := r.db.Find(&events).Error
	return events, err
}

func (r *eventRepository) GetEventByID(id uint) (*models.Event, error) {
	var event models.Event
	err := r.db.First(&event, id).Error
	return &event, err
}

// ✅ Update Event
func (r *eventRepository) UpdateEvent(event *models.Event) error {
	return r.db.Save(event).Error
}

func (r *eventRepository) DeleteEvent(id uint) error {
	return r.db.Delete(&models.Event{}, id).Error
}
