package config

import (
	"fmt"
	"log"
	"ticket-booking-backend/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// JwtSecret dipakai untuk menandatangani token JWT
var JwtSecret = []byte("your-secret-key")

func ConnectDB() {
	dsn := "host=host.docker.internal user=admin password=glissha dbname=ticket_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("Database connected successfully")
	DB = db
}

func MigrateDB() {
	if DB == nil {
		log.Fatal("Database is not connected. Make sure to call ConnectDB() first.")
	}
	err := DB.AutoMigrate(&models.User{}, &models.Admin{}, &models.Movie{}, &models.Event{}, &models.Order{}, &models.Ticket{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	fmt.Println("Database Migrated")
}
