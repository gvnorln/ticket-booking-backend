package routes

import (
	"ticket-booking-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func TicketRoutes(router *gin.Engine) {
	router.GET("/tickets", handlers.GetTickets)
}
