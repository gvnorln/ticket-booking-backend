package routes

import (
	"ticket-booking-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine) {
	router.GET("/orders", handlers.GetOrders)
	router.POST("/orders", handlers.CreateOrder)
}
