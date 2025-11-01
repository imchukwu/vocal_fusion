package repository

import (
	"gorm.io/gorm"
	"vocal_fusion/internals/models"
)

type MediaRepository interface {
	CreateMedia(media *models.Media) error
	GetAllMedia() ([]models.Media, error)
	GetMediaByID(id int) (*models.Media, error)
	UpdateMedia(media *models.Media) error
	DeleteMedia(id int) error
}

type mediaRepository struct {
	DB *gorm.DB
}

func NewMediaRepository(db *gorm.DB) MediaRepository {
	return &mediaRepository{DB: db}
}

func (r *mediaRepository) CreateMedia(media *models.Media) error {
	return r.DB.Create(media).Error
}

func (r *mediaRepository) GetAllMedia() ([]models.Media, error) {
	var media []models.Media
	err := r.DB.Order("created_at desc").Find(&media).Error
	return media, err
}

func (r *mediaRepository) GetMediaByID(id int) (*models.Media, error) {
	var media models.Media
	err := r.DB.First(&media, id).Error
	return &media, err
}

func (r *mediaRepository) UpdateMedia(media *models.Media) error {
	return r.DB.Save(media).Error
}

func (r *mediaRepository) DeleteMedia(id int) error {
	return r.DB.Delete(&models.Media{}, id).Error
}
