package services

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"memoria-backend/models"
	"memoria-backend/repository"
	"memoria-backend/utils"

	"golang.org/x/crypto/bcrypt"
)

type PasteService interface {
	GetAll(ctx context.Context) ([]models.Paste, error)
	GetByID(ctx context.Context, id uint64) (*models.Paste, error)
	GetByPrivateAccessID(ctx context.Context, privateAccessID string) (*models.Paste, error)
	Create(ctx context.Context, newPaste *models.CreatePasteRequest) (*models.Paste, error)
	Update(ctx context.Context, updatedPaste *models.UpdatePasteRequest) (*models.Paste, error)
	Delete(ctx context.Context, id uint64) (uint64, error)
	VerifyPassword(ctx context.Context, id uint64, providedPassword string) (bool, error)
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

// generatePrivateAccessID creates a secure random ID for private pastes
func generatePrivateAccessID() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// hashPassword securely hashes a password using bcrypt
func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
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

func (s *pasteService) GetByPrivateAccessID(ctx context.Context, privateAccessID string) (*models.Paste, error) {
	// You'll need to add this method to the repository interface and implement it
	paste, err := s.repo.GetByPrivateAccessID(ctx, privateAccessID)
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
		EditorType:      newPaste.EditorType,
		ExpiresAt:       newPaste.ExpiresAt,
		Privacy:         newPaste.Privacy,
	}

	if newPaste.Privacy == "private" {
		privateID, err := generatePrivateAccessID()
		if err != nil {
			log.Error().Err(err).Msg("Failed to generatea a Private Access ID")
			return nil, err
		}
		paste.PrivateAccessID = privateID
		log.Info().Str("privateAccessId", privateID).Msg("Generated PrivateAccessId for private paste")
	}

	if newPaste.Password != "" {
		hashedPassword, err := hashPassword(newPaste.Password)
		if err != nil {
			log.Error().Err(err).Msg("Failed to hash password")
			return nil, err
		}
		paste.Password = hashedPassword
		log.Info().Msg("Password successfully hashed for paste")
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
	log := utils.LoggerFromContext(ctx)

	// First, get the existing paste
	existingPaste, err := s.repo.GetByID(ctx, updatedPaste.ID)
	if err != nil {
		return nil, err
	}

	// Update the paste fields
	existingPaste.Title = updatedPaste.Title
	existingPaste.Content = updatedPaste.Content
	existingPaste.SyntaxHighlight = updatedPaste.SyntaxHighlight
	existingPaste.EditorType = updatedPaste.EditorType
	existingPaste.ExpiresAt = updatedPaste.ExpiresAt
	existingPaste.Privacy = updatedPaste.Privacy

	// Handle privacy changes
	if updatedPaste.Privacy == "private" && existingPaste.PrivateAccessID == "" {
		privateID, err := generatePrivateAccessID()
		if err != nil {
			log.Error().Err(err).Msg("Failed to generate PrivateAccessID")
			return nil, err
		}
		existingPaste.PrivateAccessID = privateID
		log.Info().Str("private_access_id", privateID).Msg("Generated PrivateAccessID for private paste")
	}

	// Handle password changes
	if updatedPaste.Password != "" {
		hashedPassword, err := hashPassword(updatedPaste.Password)
		if err != nil {
			log.Error().Err(err).Msg("Failed to hash password")
			return nil, err
		}
		existingPaste.Password = hashedPassword
		log.Info().Msg("Password successfully updated for paste")
	}

	savedPaste, err := s.repo.Update(ctx, existingPaste)
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

func (s *pasteService) VerifyPassword(ctx context.Context, id uint64, providedPassword string) (bool, error) {
	log := utils.LoggerFromContext(ctx)

	// Retrieve the paste from the repository
	paste, err := s.repo.GetByID(ctx, id)
	if err != nil {
		log.Error().Err(err).Uint64("id", id).Msg("Failed to retrieve paste for password verification")
		return false, err
	}

	// If the paste doesn't have a password, verification fails
	if paste.Password == "" {
		log.Debug().Uint64("id", id).Msg("Paste has no password set")
		return false, nil
	}

	// Compare the provided password with the stored hash
	err = bcrypt.CompareHashAndPassword([]byte(paste.Password), []byte(providedPassword))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			log.Debug().Uint64("id", id).Msg("Password verification failed: incorrect password")
			return false, nil
		}
		log.Error().Err(err).Uint64("id", id).Msg("Error during password verification")
		return false, err
	}

	log.Debug().Uint64("id", id).Msg("Password successfully verified")
	return true, nil
}
