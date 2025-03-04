package handlers

import (
	"memoria-backend/models"
	"memoria-backend/services"

	"github.com/gin-gonic/gin"
)

type PasteHandler struct {
	pasteService services.PasteService
}

func NewPasteHandler(pasteService services.PasteService) *PasteHandler {
	return &PasteHandler{pasteService: pasteService}
}

// CreatePaste godoc
// @Summary create paste
// @Description creates a new paste
// @Tags pastes
// @Accept json
// @Produce json
// @Success 200 {object} models.PasteResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /pastes [put]
func (h *PasteHandler) CreatePaste(c *gin.Context) {
	var req models.CreatePasteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	paste, err := h.pasteService.Create(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"paste": paste})
}

// GetPaste godoc
// @Summary Gets a specific paste
// @Description Retrieve a paste by ID
// @Tags pastes
// @Accept json
// @Produce json
// @Success 200 {object} models.PasteResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /pastes [get]
func (h *PasteHandler) GetPaste(c *gin.Context) {
	var req models.PasteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	paste, err := h.pasteService.GetByID(req.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"paste": paste})
}

// UpdatePaste godoc
// @Summary Update paste
// @Description Update a pastes value
// @Tags pastes
// @Accept json
// @Produce json
// @Success 200 {object} models.PasteResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /pastes [put]
func (h *PasteHandler) UpdatePaste(c *gin.Context) {
	var req models.UpdatePasteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	paste, err := h.pasteService.Update(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"paste": paste})
}

// DeletePaste godoc
// @Summary Deletes paste by ID
// @Description delete a paste by ID
// @Tags pastes
// @Accept json
// @Produce json
// @Success 200 {object} models.PasteResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /pastes [delete]
func (h *PasteHandler) DeletePaste(c *gin.Context) {
	var req models.PasteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	paste, err := h.pasteService.Delete(req.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"paste": paste})
}

// ListPastes godoc
// @Summary deletes pastes by ID
// @Description Deletes a paste by ID
// @Tags pastes
// @Accept json
// @Produce json
// @Success 200 {object} models.PasteListResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /pastes [get]
func (h *PasteHandler) ListPastes(c *gin.Context) {
	var req models.PasteListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// TODO: Implement pagination, per page list request format
	pastes, err := h.pasteService.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"pastes": pastes})
}
