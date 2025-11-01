package repository

import (
	"gorm.io/gorm"
	"vocal_fusion/internals/models"
)

type SchoolRepository interface {
	CreateSchool(school *models.School) error
	GetAllSchools() ([]models.School, error)
	GetSchoolByID(id int) (*models.School, error)
	UpdateSchool(school *models.School) error
	DeleteSchool(id int) error
}

type schoolRepository struct {
	DB *gorm.DB
}

func NewSchoolRepository(db *gorm.DB) SchoolRepository {
	return &schoolRepository{DB: db}
}

func (r *schoolRepository) CreateSchool(school *models.School) error {
	return r.DB.Create(school).Error
}

func (r *schoolRepository) GetAllSchools() ([]models.School, error) {
	var schools []models.School
	err := r.DB.Order("created_at desc").Find(&schools).Error
	return schools, err
}

func (r *schoolRepository) GetSchoolByID(id int) (*models.School, error) {
	var school models.School
	err := r.DB.First(&school, id).Error
	return &school, err
}

func (r *schoolRepository) UpdateSchool(school *models.School) error {
	return r.DB.Save(school).Error
}

func (r *schoolRepository) DeleteSchool(id int) error {
	return r.DB.Delete(&models.School{}, id).Error
}
