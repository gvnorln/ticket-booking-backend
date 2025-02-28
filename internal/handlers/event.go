package handlers

import (
	"net/http"
	"ticket-booking-backend/config"
	"ticket-booking-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {
	var events []models.Event
	config.DB.Find(&events)
	c.JSON(http.StatusOK, events)
}

func CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&event)
	c.JSON(http.StatusCreated, event)
}
