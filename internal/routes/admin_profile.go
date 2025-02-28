package routes

import (
	"ticket-booking-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func AdminProfileRoutes(router *gin.Engine) {
	// Endpoint untuk mendapatkan profil admin
	router.GET("/admin/profile", handlers.GetAdminProfile)
	// Endpoint untuk mengupdate profil admin
	router.PUT("/admin/profile", handlers.UpdateAdminProfile)
}
