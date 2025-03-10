package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"memoria-backend/models"
	"memoria-backend/services"
	"memoria-backend/utils"
	"strconv"
	"time"
)

type PasteHandler struct {
	pasteService services.PasteService
}

func NewPasteHandler(pasteService services.PasteService) *PasteHandler {
	return &PasteHandler{pasteService: pasteService}
}

func IsPasteExpired(paste *models.Paste) error {
	// If the paste has no expiration time, it never expires
	if paste.ExpiresAt.IsZero() {
		return nil
	}

	// Check if the current time is past the expiration time
	if time.Now().After(paste.ExpiresAt) {
		return fmt.Errorf("paste has expired")
	}

	return nil
}

// CreatePaste godoc
// @Summary Create paste
// @Description Creates a new paste
// @Tags pastes
// @Accept json
// @Produce json
// @Param paste body models.CreatePasteRequest true "Paste data"
// @Success 200 {object} models.APIResponse[models.PasteData] "Success response with paste data"
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /paste [post]
func (h *PasteHandler) CreatePaste(c *gin.Context) {
	ctx := c.Request.Context()
	log := utils.LoggerFromContext(ctx)

	var req models.CreatePasteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind JSON for create paste request")
		utils.RespondBadRequest(c, err, "Invalid paste data format")
		return
	}

	// if req.EditorType == "" {
	// 	req.EditorType = "code"
	// }
	//

	log.Info().
		Str("title", req.Title).
		Str("contentPreview", utils.Truncate(req.Content, 50)).
		Msg("Creating new paste")

	paste, err := h.pasteService.Create(ctx, &req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create paste")
		utils.RespondInternalError(c, err, "Failed to create paste")
		return
	}

	log.Info().Uint64("pasteId", paste.ID).Msg("Successfully created paste")

	pasteData := models.PasteData{Paste: paste}
	utils.RespondCreated(c, pasteData, "Paste created successfully")
}

// GetPaste godoc
// @Summary Gets a specific paste
// @Description Retrieve a paste by ID
// @Tags pastes
// @Param id path uint64 true "Paste ID"
// @Param pw query string false "Password for protected pastes"
// @Produce json
// @Success 200 {object} models.APIResponse[models.PasteData] "Success response with paste data"
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse "Pssword required or invalid password"
// @Failure 404 {object} models.ErrorResponse "Paste not found"
// @Failure 500 {object} models.ErrorResponse
// @Router /paste/{id} [get]
func (h *PasteHandler) GetPaste(c *gin.Context) {
	ctx := c.Request.Context()
	log := utils.LoggerFromContext(ctx)

	idStr := c.Param("id")
	log.Info().Str("idStr", idStr).Msg("ParamID")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Error().Err(err).Str("idStr", idStr).Msg("Failed to parse ID for paste")
		utils.RespondBadRequest(c, err, "Invalid paste ID format")
		return
	}

	log.Info().Uint64("pasteId", id).Msg("Retrieving paste")

	paste, err := h.pasteService.GetByID(ctx, id)
	if err != nil {
		log.Error().Err(err).Uint64("pasteId", id).Msg("Failed to retrieve paste")
		utils.RespondNotFound(c, err, "Paste not found")
		return
	}
	// Check if paste is expired
	if err := IsPasteExpired(paste); err != nil {
		log.Info().Err(err).Uint64("pasteId", id).Msg("Attempted to access expired paste")
		utils.RespondNotFound(c, err, "This paste has expired and is no longer available")
		return
	}

	// Check if paste is private - if so, don't allow access through this endpoint
	if paste.Privacy == "private" {
		log.Info().Uint64("pasteId", id).Msg("Attempted to access private paste without access ID")
		utils.RespondForbidden(c, nil, "This is a private paste. Please use the private access ID to view it.")
		return
	}

	if paste.Password != "" {
		providedPassword := c.Query("pw")

		if providedPassword == "" {
			log.Info().Uint64("pasteId", id).Msg("Attempted to access password-protected paste without password")
			utils.RespondUnauthorized(c, err, "Error verifying password")
			return
		}

		passwordValid, err := h.pasteService.VerifyPassword(ctx, id, providedPassword)
		if err != nil {
			log.Error().Err(err).Uint64("pasteId", id).Msg("Error verifying password")
			utils.RespondInternalError(c, err, "Error verifying password")
			return
		}

		if !passwordValid {
			log.Info().Uint64("pasteId", id).Msg("Invalid Password provided for password protected paste")
			utils.RespondUnauthorized(c, err, "Invalid password")
			return
		}

	}

	log.Info().Uint64("pasteId", id).Msg("Successfully retrieved paste")

	pasteData := models.PasteData{Paste: paste}
	utils.RespondOK(c, pasteData, "Paste retrieved successfully")
}

// UpdatePaste godoc
// @Summary Update paste
// @Description Update a pastes value
// @Tags pastes
// @Accept json
// @Produce json
// @Param paste body models.UpdatePasteRequest true "Updated paste data"
// @Success 200 {object} models.APIResponse[models.PasteData] "Success response with paste data"
// @Failure 404 {object} models.ErrorResponse "Paste not found"
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /paste [put]
func (h *PasteHandler) UpdatePaste(c *gin.Context) {
	ctx := c.Request.Context()
	log := utils.LoggerFromContext(ctx)

	var req models.UpdatePasteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind JSON for update paste request")
		utils.RespondBadRequest(c, err, "Invalid paste data format")
		return
	}

	log.Info().
		Uint64("pasteId", req.ID).
		Str("title", req.Title).
		Str("contentPreview", utils.Truncate(req.Content, 50)).
		Msg("Updating paste")

	paste, err := h.pasteService.Update(ctx, &req)
	if err != nil {
		log.Error().Err(err).Uint64("pasteId", req.ID).Msg("Failed to update paste")
		utils.RespondInternalError(c, err, "Failed to update paste")
		return
	}

	log.Info().Uint64("pasteId", req.ID).Msg("Successfully updated paste")

	pasteData := models.PasteData{Paste: paste}
	utils.RespondOK(c, pasteData, "Paste updated successfully")
}

// DeletePaste godoc
// @Summary Deletes paste by ID
// @Description delete a paste by ID
// @Tags pastes
// @Accept json
// @Param id path uint64 true "Paste ID"
// @Produce json
// @Success 200 {object} models.APIResponse[uint64]
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /paste [delete]
func (h *PasteHandler) DeletePaste(c *gin.Context) {
	ctx := c.Request.Context()
	log := utils.LoggerFromContext(ctx)

	idStr := c.Param("id")
	log.Info().Str("idStr", idStr).Msg("ParamID")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Error().Err(err).Str("idStr", idStr).Msg("Failed to parse ID for paste")
		utils.RespondBadRequest(c, err, "Invalid paste ID format")
		return
	}

	log.Info().Uint64("pasteId", id).Msg("Deleting paste")

	deletedID, err := h.pasteService.Delete(ctx, id)
	if err != nil {
		log.Error().Err(err).Uint64("pasteId", id).Msg("Failed to delete paste")
		utils.RespondNotFound(c, err, "Paste not found or could not be deleted")
		return
	}

	log.Info().Uint64("pasteId", id).Msg("Successfully deleted paste")

	// For delete operations, you can return an empty data struct or the deleted paste
	utils.RespondOK(c, deletedID, "Paste deleted successfully")
}

// ListPastes godoc
// @Summary Lists out all the pastes
// @Description Returns a list of all the pastes. # TODO: pagination
// @Tags pastes
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} models.APIResponse[models.PasteListData] "Success response with paste list data"
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /paste/all [get]
func (h *PasteHandler) ListPastes(c *gin.Context) {
	ctx := c.Request.Context()
	log := utils.LoggerFromContext(ctx)

	// Parse pagination parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	log.Info().Int("page", page).Int("limit", limit).Msg("Retrieving pastes")

	pastes, err := h.pasteService.GetAll(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve all pastes")
		utils.RespondInternalError(c, err, "Failed to retrieve pastes")
		return
	}

	log.Info().Int("count", len(pastes)).Msg("Successfully retrieved pastes")

	pasteListData := models.PasteListData{
		Pastes: pastes,
		Count:  len(pastes),
	}
	utils.RespondOK(c, pasteListData, "Pastes retrieved successfully")
}

// GetPasteByPrivateAccessID godoc
// @Summary Gets a specific private paste using its private access ID
// @Description Retrieve a private paste by its private access ID
// @Tags pastes
// @Param accessId path string true "Private Access ID"
// @Param pw query string false "Password for protected pastes"
// @Produce json
// @Success 200 {object} models.APIResponse[models.PasteData] "Success response with paste data"
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse "Paste not found"
// @Failure 500 {object} models.ErrorResponse
// @Router /paste/private/{accessId} [get]
func (h *PasteHandler) GetPasteByPrivateAccessID(c *gin.Context) {
	ctx := c.Request.Context()
	log := utils.LoggerFromContext(ctx)

	accessID := c.Param("accessId")
	log.Info().Str("privateAccessId", accessID).Msg("Retrieving paste by private access ID")

	paste, err := h.pasteService.GetByPrivateAccessID(ctx, accessID)
	if err != nil {
		log.Error().Err(err).Str("privateAccessId", accessID).Msg("Failed to retrieve paste")
		utils.RespondNotFound(c, err, "Paste not found")
		return
	}
	// Check if paste is expired

	if err := IsPasteExpired(paste); err != nil {
		log.Info().Err(err).Uint64("pasteId", paste.ID).Msg("Attempted to access expired paste")
		utils.RespondNotFound(c, err, "This paste has expired and is no longer available")
		return
	}

	if paste.Password != "" {
		providedPassword := c.Query("pw")

		if providedPassword == "" {
			log.Info().Uint64("id", paste.ID).Msg("Attempted to access password-protected paste without password")
			utils.RespondUnauthorized(c, err, "Error verifying password")
			return
		}

		passwordValid, err := h.pasteService.VerifyPassword(ctx, paste.ID, providedPassword)
		if err != nil {
			log.Error().Err(err).Uint64("pasteId", paste.ID).Msg("Error verifying password")
			utils.RespondInternalError(c, err, "Error verifying password")
			return
		}

		if !passwordValid {
			log.Info().Uint64("pasteId", paste.ID).Msg("Invalid Password provided for password protected paste")
			utils.RespondUnauthorized(c, err, "Invalid password")
			return
		}

	}

	log.Info().Str("privateAccessId", accessID).Msg("Successfully retrieved paste")

	pasteData := models.PasteData{Paste: paste}
	utils.RespondOK(c, pasteData, "Paste retrieved successfully")
}
