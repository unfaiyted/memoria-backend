package handlers

import (
	"memoria-backend/models"
	"memoria-backend/services"
	"memoria-backend/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PasteHandler struct {
	pasteService services.PasteService
}

func NewPasteHandler(pasteService services.PasteService) *PasteHandler {
	return &PasteHandler{pasteService: pasteService}
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

	log.Info().
		Str("title", req.Title).
		Str("content_preview", utils.Truncate(req.Content, 50)).
		Msg("Creating new paste")

	paste, err := h.pasteService.Create(ctx, &req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create paste")
		utils.RespondInternalError(c, err, "Failed to create paste")
		return
	}

	log.Info().Uint64("paste_id", paste.ID).Msg("Successfully created paste")

	pasteData := models.PasteData{Paste: paste}
	utils.RespondCreated(c, pasteData, "Paste created successfully")
}

// GetPaste godoc
// @Summary Gets a specific paste
// @Description Retrieve a paste by ID
// @Tags pastes
// @Param id path uint64 true "Paste ID"
// @Produce json
// @Success 200 {object} models.APIResponse[models.PasteData] "Success response with paste data"
// @Failure 400 {object} models.ErrorResponse
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

	log.Info().Uint64("paste_id", id).Msg("Retrieving paste")

	paste, err := h.pasteService.GetByID(ctx, id)
	if err != nil {
		log.Error().Err(err).Uint64("paste_id", id).Msg("Failed to retrieve paste")
		utils.RespondNotFound(c, err, "Paste not found")
		return
	}

	log.Info().Uint64("paste_id", id).Msg("Successfully retrieved paste")

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
		Uint64("paste_id", req.ID).
		Str("title", req.Title).
		Str("content_preview", utils.Truncate(req.Content, 50)).
		Msg("Updating paste")

	paste, err := h.pasteService.Update(ctx, &req)
	if err != nil {
		log.Error().Err(err).Uint64("paste_id", req.ID).Msg("Failed to update paste")
		utils.RespondInternalError(c, err, "Failed to update paste")
		return
	}

	log.Info().Uint64("paste_id", req.ID).Msg("Successfully updated paste")

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

	log.Info().Uint64("paste_id", id).Msg("Deleting paste")

	deletedID, err := h.pasteService.Delete(ctx, id)
	if err != nil {
		log.Error().Err(err).Uint64("paste_id", id).Msg("Failed to delete paste")
		utils.RespondNotFound(c, err, "Paste not found or could not be deleted")
		return
	}

	log.Info().Uint64("paste_id", id).Msg("Successfully deleted paste")

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
