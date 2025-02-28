package handlers

import (
	"net/http"
	"ticket-booking-backend/config"
	"ticket-booking-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&order)
	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully!"})
}

func GetOrders(c *gin.Context) {
	var orders []models.Order
	config.DB.Find(&orders)
	c.JSON(http.StatusOK, orders)
}
