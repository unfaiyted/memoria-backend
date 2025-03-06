// router/router.go
package router

import (
	"context"
	"memoria-backend/services"
	"memoria-backend/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(ctx context.Context, db *gorm.DB, configService services.ConfigService) *gin.Engine {
	r := gin.Default()
	log := utils.LoggerFromContext(ctx)

	appConfig := configService.GetConfig()
	// CORS config
	config := cors.DefaultConfig()
	config.AllowOrigins = appConfig.Auth.AllowedOrigins

	log.Info().
		Strs("AllowedOrigins", config.AllowOrigins).
		Msg("Allowed Origins set.")

	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type"}
	r.Use(cors.New(config))

	// Setup API v1 routes
	v1 := r.Group("api/v1")

	healthService := services.NewHealthService(db)

	// Register all routes
	RegisterUserRoutes(v1, db)
	RegisterConfigRoutes(v1, configService)
	RegisterHealthRoutes(v1, healthService)
	RegisterPasteRoutes(v1, db)

	return r
}
