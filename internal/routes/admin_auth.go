package routes

import (
	"ticket-booking-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

// AdminAuthRoutes menyediakan endpoint untuk login admin (tanpa middleware)
func AdminAuthRoutes(router *gin.Engine) {
	router.POST("/admin/login", handlers.LoginAdmin)
}
