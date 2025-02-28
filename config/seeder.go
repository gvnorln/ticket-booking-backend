package config

import (
	"fmt"
	"ticket-booking-backend/internal/models"
)

func SeedAdmin() {
	var count int64
	DB.Model(&models.Admin{}).Count(&count)

	if count == 0 {
		admin := models.Admin{
			Username: "admin",
			Email:    "admin@example.com",
			Password: "admin123",
			Role:     "superadmin",
		}

		if err := admin.HashPassword(); err != nil {
			fmt.Println("Failed to hash admin password:", err)
			return
		}

		if err := DB.Create(&admin).Error; err != nil {
			fmt.Println("Gagal membuat akun admin:", err)
		} else {
			fmt.Println("Akun admin berhasil dibuat! Email: admin@example.com, Password: admin123")
		}
	}
}
