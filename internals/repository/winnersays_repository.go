package repository

import (
    "vocal_fusion/internals/models"
    "gorm.io/gorm"
)

type WinnerSaysRepository struct {
    db *gorm.DB
}

func NewWinnerSaysRepository(db *gorm.DB) *WinnerSaysRepository {
    return &WinnerSaysRepository{db: db}
}

func (r *WinnerSaysRepository) Create(w *models.WinnerSays) error {
    return r.db.Create(w).Error
}

func (r *WinnerSaysRepository) GetAll() ([]models.WinnerSays, error) {
    var winners []models.WinnerSays
    err := r.db.Preload("User").Find(&winners).Error
    return winners, err
}

func (r *WinnerSaysRepository) GetByID(id int) (*models.WinnerSays, error) {
    var winner models.WinnerSays
    err := r.db.Preload("User").First(&winner, id).Error
    return &winner, err
}

func (r *WinnerSaysRepository) Delete(id int) error {
    return r.db.Delete(&models.WinnerSays{}, id).Error
}
