package routes

import (
	"ticket-booking-backend/internal/handlers"
	middlewares "ticket-booking-backend/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine) {
	admin := router.Group("/admin")
	// Proteksi route dengan middleware JWT
	admin.Use(middlewares.AdminAuthMiddleware())
	admin.GET("/users", handlers.GetAllUsers)
	admin.GET("/movies", handlers.GetAllMovies)
	admin.POST("/movies", handlers.CreateMovie)
	admin.PUT("/movies/:id", handlers.UpdateMovie)
	admin.DELETE("/movies/:id", handlers.DeleteMovie)
}
