package handlers

import (
	"net/http"
	"ticket-booking-backend/config"
	"ticket-booking-backend/internal/models"

	"github.com/gin-gonic/gin"
)

// GetAdminProfile mengembalikan data profil admin.
// Untuk contoh, kita ambil admin dengan ID 1; sebaiknya gunakan data dari token.
func GetAdminProfile(c *gin.Context) {
	var admin models.Admin
	if err := config.DB.First(&admin, 1).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}
	c.JSON(http.StatusOK, admin)
}

// UpdateAdminProfile menerima data JSON untuk memperbarui profil admin.
func UpdateAdminProfile(c *gin.Context) {
	var admin models.Admin
	// Untuk contoh, kita ambil admin dengan ID 1 (gantilah logika ini sesuai dengan autentikasi yang Anda gunakan)
	if err := config.DB.First(&admin, 1).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}

	// Struktur input untuk update profil
	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	admin.Username = input.Username
	admin.Email = input.Email

	if err := config.DB.Save(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, admin)
}
