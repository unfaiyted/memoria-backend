// In handlers/health.go
package handlers

import (
	"memoria-backend/models"
	"memoria-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	service services.HealthService
}

func NewHealthHandler(service services.HealthService) *HealthHandler {
	return &HealthHandler{
		service: service,
	}
}

// CheckHealth godoc
// @Summary checks app and database health
// @Description returns JSON object with health statuses.
// @Tags health
// @Produce json
// @Success 200 {object} models.HealthResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /health [get]
func (h *HealthHandler) CheckHealth(c *gin.Context) {
	appStatus := h.service.CheckApplicationStatus()
	dbStatus := h.service.CheckDatabaseConnection()

	// Set overall status based on individual component statuses
	overallStatus := "up"
	httpStatus := http.StatusOK

	// Check if any component is down
	if !appStatus || !dbStatus {
		overallStatus = "down"
		httpStatus = http.StatusInternalServerError

		// Create error response
		errorResponse := models.ErrorResponse{
			Error: "Health check failed",
			Details: map[string]interface{}{
				"application": appStatus,
				"database":    dbStatus,
			},
		}

		c.JSON(httpStatus, errorResponse)
		return
	}

	// All components are healthy
	response := models.HealthResponse{
		Status:      overallStatus,
		Application: appStatus,
		Database:    dbStatus,
	}

	c.JSON(httpStatus, response)
}
