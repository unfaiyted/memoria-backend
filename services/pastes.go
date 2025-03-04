package services

import (
	"memoria-backend/models"
	"memoria-backend/repository"
)

type PasteService interface {
	GetAll() ([]models.Paste, error)
	GetByID(id uint64) (*models.Paste, error)
	Create(newPaste *models.CreatePasteRequest) (*models.Paste, error)
	Update(updatedPaste *models.UpdatePasteRequest) (*models.Paste, error)
	Delete(id uint64) (uint64, error)
}

type pasteService struct {
	repo repository.PasteRepository
}

// NewConfigService creates a new configuration service
func NewPasteService(pasteRepo repository.PasteRepository) PasteService {
	return &pasteService{
		repo: pasteRepo,
	}
}

func (s *pasteService) GetAll() ([]models.Paste, error) {
	pastes, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return pastes, nil
}

func (s *pasteService) GetByID(id uint64) (*models.Paste, error) {
	paste, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return paste, nil
}

func (s *pasteService) Create(newPaste *models.CreatePasteRequest) (*models.Paste, error) {
	paste := &models.Paste{
		Title:           newPaste.Title,
		Content:         newPaste.Content,
		SyntaxHighlight: newPaste.SyntaxHighlight,
		ExpiresAt:       newPaste.ExpiresAt,
		Privacy:         newPaste.Privacy,
	}

	createdPaste, err := s.repo.Create(paste)
	if err != nil {
		return nil, err
	}

	return createdPaste, nil
}

func (s *pasteService) Update(updatedPaste *models.UpdatePasteRequest) (*models.Paste, error) {
	paste := &models.Paste{
		ID:              updatedPaste.ID,
		Content:         updatedPaste.Content,
		SyntaxHighlight: updatedPaste.SyntaxHighlight,
		ExpiresAt:       updatedPaste.ExpiresAt,
		Privacy:         updatedPaste.Privacy,
	}

	savedPaste, err := s.repo.Update(paste)
	if err != nil {
		return nil, err
	}

	return savedPaste, nil
}

func (s *pasteService) Delete(id uint64) (uint64, error) {
	deletedID, err := s.repo.Delete(id)
	if err != nil {
		return 0, err
	}

	return deletedID, nil
}
