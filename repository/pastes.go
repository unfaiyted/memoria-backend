package repository

import (
	"memoria-backend/models"

	"gorm.io/gorm"
)

type PasteRepository interface {
	GetAll() ([]models.Paste, error)
	GetByID(id uint) (*models.Paste, error)
	Create(paste *models.Paste) error
	Update(paste *models.Paste) error
	Delete(id uint) error
}

type GormPasteRepository struct {
	db *gorm.DB
}

func (r *GormPasteRepository) GetAll() ([]models.Paste, error) {
	var pastes []models.Paste

	result := r.db.Find(&pastes)
	return pastes, result.Error
}

func (r *GormPasteRepository) GetByID(id uint) (*models.Paste, error) {
	var paste models.Paste
	result := r.db.First(&paste, id)
	return &paste, result.Error
}

func (r *GormPasteRepository) Create(paste *models.Paste) (*models.Paste, error) {
	var createdPaste models.Paste
	result := r.db.Create(&paste)
	return &createdPaste, result.Error
}

func (r *GormPasteRepository) Update(paste *models.Paste) (*models.Paste, error) {
	var updatedPaste models.Paste
	result := r.db.Save(&paste)
	return &updatedPaste, result.Error
}

func (r *GormPasteRepository) Delete(id uint) (uint, error) {
	var paste models.Paste
	result := r.db.Delete(&paste, id)
	// TODO: Look into optimal way to return that something was deleted.
	return id, result.Error
}
