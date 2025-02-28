package routes

import (
	"ticket-booking-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func EventRoutes(router *gin.Engine) {
	router.GET("/events", handlers.GetEvents)
	router.POST("/events", handlers.CreateEvent)
}
