package routes

import (
	"ticket-booking-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/register", handlers.RegisterUser)
	router.POST("/login", handlers.LoginUser)
}
