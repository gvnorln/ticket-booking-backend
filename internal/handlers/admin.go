package handlers

import (
	"net/http"
	"ticket-booking-backend/config"
	"ticket-booking-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetAllMovies(c *gin.Context) {
	var movies []models.Movie
	if err := config.DB.Find(&movies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"movies": movies})
}

func CreateMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&movie)
	c.JSON(http.StatusCreated, gin.H{"message": "Movie created successfully"})
}

func UpdateMovie(c *gin.Context) {
	var movie models.Movie
	id := c.Param("id")
	if err := config.DB.First(&movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&movie)
	c.JSON(http.StatusOK, gin.H{"message": "Movie updated successfully"})
}

func DeleteMovie(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Movie{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete movie"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}
