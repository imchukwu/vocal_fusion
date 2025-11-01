package repository

import (
	"gorm.io/gorm"
	"vocal_fusion/internals/models"
)

type FAQRepository interface {
	CreateFAQ(faq *models.FAQ) error
	GetAllFAQs() ([]models.FAQ, error)
	GetFAQByID(id int) (*models.FAQ, error)
	UpdateFAQ(faq *models.FAQ) error
	DeleteFAQ(id int) error
}

type faqRepository struct {
	DB *gorm.DB
}

func NewFAQRepository(db *gorm.DB) FAQRepository {
	return &faqRepository{DB: db}
}

func (r *faqRepository) CreateFAQ(faq *models.FAQ) error {
	return r.DB.Create(faq).Error
}

func (r *faqRepository) GetAllFAQs() ([]models.FAQ, error) {
	var faqs []models.FAQ
	err := r.DB.Find(&faqs).Error
	return faqs, err
}

func (r *faqRepository) GetFAQByID(id int) (*models.FAQ, error) {
	var faq models.FAQ
	err := r.DB.First(&faq, id).Error
	return &faq, err
}

func (r *faqRepository) UpdateFAQ(faq *models.FAQ) error {
	return r.DB.Save(faq).Error
}

func (r *faqRepository) DeleteFAQ(id int) error {
	return r.DB.Delete(&models.FAQ{}, id).Error
}
