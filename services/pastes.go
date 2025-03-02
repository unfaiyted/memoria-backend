package services

import (
	"memoria-backend/models"
	"memoria-backend/repository"
)

type PasteService struct {
	repo repository.PasteRepository
}

func (s *PasteService) GetAll() ([]models.Paste, error) {
	pastes, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return pastes, nil
}

func (s *PasteService) GetByID(id uint64) (*models.Paste, error) {
	paste, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return paste, nil
}

func (s *PasteService) Create(newPaste models.CreatePasteRequest) (*models.Paste, error) {
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

func (s *PasteService) Update(updatedPaste *models.UpdatePasteRequest) (*models.Paste, error) {
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

func (s *PasteService) Delete(id uint64) (uint64, error) {
	deletedID, err := s.repo.Delete(id)
	if err != nil {
		return 0, err
	}

	return deletedID, nil
}
