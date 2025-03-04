package services

import (
	"context"
	"memoria-backend/models"
	"memoria-backend/repository"
	"memoria-backend/utils"
)

type PasteService interface {
	GetAll(ctx context.Context) ([]models.Paste, error)
	GetByID(ctx context.Context, id uint64) (*models.Paste, error)
	Create(ctx context.Context, newPaste *models.CreatePasteRequest) (*models.Paste, error)
	Update(ctx context.Context, updatedPaste *models.UpdatePasteRequest) (*models.Paste, error)
	Delete(ctx context.Context, id uint64) (uint64, error)
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

func (s *pasteService) GetAll(ctx context.Context) ([]models.Paste, error) {
	pastes, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return pastes, nil
}

func (s *pasteService) GetByID(ctx context.Context, id uint64) (*models.Paste, error) {
	paste, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return paste, nil
}

func (s *pasteService) Create(ctx context.Context, newPaste *models.CreatePasteRequest) (*models.Paste, error) {
	log := utils.LoggerFromContext(ctx)
	paste := &models.Paste{
		Title:           newPaste.Title,
		Content:         newPaste.Content,
		SyntaxHighlight: newPaste.SyntaxHighlight,
		ExpiresAt:       newPaste.ExpiresAt,
		Privacy:         newPaste.Privacy,
	}

	createdPaste, err := s.repo.Create(ctx, paste)

	log.Info().
		Str("title", createdPaste.Title).
		Str("content_preview", utils.Truncate(createdPaste.Content, 50)).
		Msg("Repo returned new paste")

	if err != nil {
		return nil, err
	}

	return createdPaste, nil
}

func (s *pasteService) Update(ctx context.Context, updatedPaste *models.UpdatePasteRequest) (*models.Paste, error) {
	paste := &models.Paste{
		ID:              updatedPaste.ID,
		Title:           updatedPaste.Title,
		Content:         updatedPaste.Content,
		SyntaxHighlight: updatedPaste.SyntaxHighlight,
		ExpiresAt:       updatedPaste.ExpiresAt,
		Privacy:         updatedPaste.Privacy,
	}

	savedPaste, err := s.repo.Update(ctx, paste)
	if err != nil {
		return nil, err
	}

	return savedPaste, nil
}

func (s *pasteService) Delete(ctx context.Context, id uint64) (uint64, error) {
	deletedID, err := s.repo.Delete(ctx, id)
	if err != nil {
		return 0, err
	}

	return deletedID, nil
}
