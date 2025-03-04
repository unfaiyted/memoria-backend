package repository

import (
	"memoria-backend/models"

	"gorm.io/gorm"
)

type PasteRepository interface {
	GetAll() ([]models.Paste, error)
	GetByID(id uint64) (*models.Paste, error)
	Create(paste *models.Paste) (*models.Paste, error)
	Update(paste *models.Paste) (*models.Paste, error)
	Delete(id uint64) (uint64, error)
}

type pasteRepository struct {
	db *gorm.DB
}

func NewPasteRepository(db *gorm.DB) PasteRepository {
	return &pasteRepository{
		db: db,
	}
}

func (r *pasteRepository) GetAll() ([]models.Paste, error) {
	var pastes []models.Paste

	result := r.db.Find(&pastes)
	return pastes, result.Error
}

func (r *pasteRepository) GetByID(id uint64) (*models.Paste, error) {
	var paste models.Paste
	result := r.db.First(&paste, id)
	return &paste, result.Error
}

func (r *pasteRepository) Create(paste *models.Paste) (*models.Paste, error) {
	var createdPaste models.Paste
	result := r.db.Create(&paste)
	return &createdPaste, result.Error
}

func (r *pasteRepository) Update(paste *models.Paste) (*models.Paste, error) {
	var updatedPaste models.Paste
	result := r.db.Save(&paste)
	return &updatedPaste, result.Error
}

func (r *pasteRepository) Delete(id uint64) (uint64, error) {
	var paste models.Paste
	result := r.db.Delete(&paste, id)
	return id, result.Error
}
