package handlers

import (
	"github.com/gin-gonic/gin"
)

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
func CreatePaste(c *gin.Context) {}

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
func GetPaste(c *gin.Context) {}

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
func UpdatePaste(c *gin.Context) {}

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
func DeletePaste(c *gin.Context) {}

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
func ListPastes(c *gin.Context) {
}
