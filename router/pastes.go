package router

import (
	"memoria-backend/handlers"
	"memoria-backend/repository"
	"memoria-backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterPasteRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	pasteRepo := repository.NewPasteRepository(db)
	pasteService := services.NewPasteService(pasteRepo)
	pasteHandlers := handlers.NewPasteHandler(pasteService)

	pastes := rg.Group("/paste")
	{
		pastes.POST("", pasteHandlers.CreatePaste)
		pastes.GET("/all", pasteHandlers.ListPastes)
		pastes.GET("/:id", pasteHandlers.GetPaste)
		pastes.PUT("/:id", pasteHandlers.UpdatePaste)
		pastes.DELETE("/:id", pasteHandlers.DeletePaste)
	}
}
