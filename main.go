// main.go
package main

import (
	"log"
	"memoria-backend/database"
	"memoria-backend/models"
	"memoria-backend/repository"
	"memoria-backend/router"
	"memoria-backend/services"

	_ "memoria-backend/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	configRepo := repository.NewConfigRepository()
	configService := services.NewConfigService(configRepo)

	if err := configService.InitConfig(); err != nil {
		log.Fatalf("Failed to initialize config: %v", err)
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
		log.Fatal("Failed to connect to database:", err)
	}
	// Auto Migrate the schema
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Paste{})

	r := router.Setup(db, configService)

	// Swagger Docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server
	r.Run(":8080")
}
