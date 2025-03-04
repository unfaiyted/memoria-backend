// router/user_routes.go
package router

import (
	"memoria-backend/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	users := rg.Group("/users")
	{
		users.POST("", handlers.CreateUser(db))
		users.GET("", handlers.GetUsers(db))
		users.GET("/:id", handlers.GetUser(db))
		users.PUT("/:id", handlers.UpdateUser(db))
		users.DELETE("/:id", handlers.DeleteUser(db))
	}
}
