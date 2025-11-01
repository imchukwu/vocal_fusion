package repository

import (
	"gorm.io/gorm"
	"vocal_fusion/internals/models"
)

type MessageRepository interface {
	CreateMessage(msg *models.Message) error
	GetAllMessages() ([]models.Message, error)
	GetMessageByID(id int) (*models.Message, error)
	UpdateMessageStatus(id int, status string) error
	DeleteMessage(id int) error
}

type messageRepository struct {
	DB *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{DB: db}
}

func (r *messageRepository) CreateMessage(msg *models.Message) error {
	return r.DB.Create(msg).Error
}

func (r *messageRepository) GetAllMessages() ([]models.Message, error) {
	var messages []models.Message
	err := r.DB.Order("created_at desc").Find(&messages).Error
	return messages, err
}

func (r *messageRepository) GetMessageByID(id int) (*models.Message, error) {
	var msg models.Message
	err := r.DB.Preload("School").First(&msg, id).Error
	return &msg, err
}

func (r *messageRepository) UpdateMessageStatus(id int, status string) error {
	return r.DB.Model(&models.Message{}).Where("id = ?", id).Update("status", status).Error
}

func (r *messageRepository) DeleteMessage(id int) error {
	return r.DB.Delete(&models.Message{}, id).Error
}
