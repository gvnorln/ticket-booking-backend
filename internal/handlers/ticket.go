package handlers

import (
	"net/http"
	"ticket-booking-backend/config"
	"ticket-booking-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func GetTickets(c *gin.Context) {
	var tickets []models.Ticket
	config.DB.Find(&tickets)
	c.JSON(http.StatusOK, tickets)
}
