package repository

import (
	"memoria-backend/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetByID(id uint64) (*models.User, error)
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(id uint64) (uint64, error)
}

type GormUserRepository struct {
	db *gorm.DB
}

func (r *GormUserRepository) GetAll() ([]models.User, error) {
	var users []models.User

	result := r.db.Find(&users)
	return users, result.Error
}

func (r *GormUserRepository) GetByID(id uint64) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, id)
	return &user, result.Error
}

func (r *GormUserRepository) Create(user *models.User) (*models.User, error) {
	var createdUser models.User
	result := r.db.Create(&user)
	return &createdUser, result.Error
}

func (r *GormUserRepository) Update(user *models.User) (*models.User, error) {
	var updatedUser models.User
	result := r.db.Save(&user)
	return &updatedUser, result.Error
}

func (r *GormUserRepository) Delete(id uint64) (uint64, error) {
	var user models.User
	result := r.db.Delete(&user, id)
	return id, result.Error
}
