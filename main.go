// main.go
package main

import (
	"context"
	"memoria-backend/database"
	"memoria-backend/middleware"
	"memoria-backend/repository"
	"memoria-backend/router"
	"memoria-backend/services"
	logger "memoria-backend/utils"

	_ "memoria-backend/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/rs/zerolog/log"
)

//	@title			Memoria API
//	@version		1.0
//	@description	API Server for Memoria application
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
// @BasePath	/api/v1
// @schemes	http
// @openapi	3.0.0
func main() {
	logger.Initialize()

	ctx := context.Background()

	configRepo := repository.NewConfigRepository()
	configService := services.NewConfigService(configRepo)

	if err := configService.InitConfig(ctx); err != nil {
		log.Fatal().Err(err).Msg("Failed to init config")
	}

	appConfig := configService.GetConfig()

	dbConfig := database.Config{
		Host:     appConfig.Db.Host,
		User:     appConfig.Db.User,
		Password: appConfig.Db.Password,
		Name:     appConfig.Db.Name,
		Port:     appConfig.Db.Port,
	}

	db, err := database.Initialize(dbConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database:")
	}

	r := router.Setup(ctx, db, configService)

	r.Use(middleware.LoggerMiddleware())

	// Swagger Docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server
	r.Run(":8080")
}
