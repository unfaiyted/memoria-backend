package router

import (
	"memoria-backend/handlers"
	"memoria-backend/services"

	"github.com/gin-gonic/gin"
)

func RegisterHealthRoutes(rg *gin.RouterGroup, service services.HealthService) {
	healthHandlers := handlers.NewHealthHandler(service)

	// Create a health endpoint
	rg.GET("/health", healthHandlers.CheckHealth)
}
