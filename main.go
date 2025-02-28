package main

import (
	"log"
	"time"

	"ticket-booking-backend/config"
	"ticket-booking-backend/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Koneksi ke database, migrasi, dan seed data admin
	config.ConnectDB()
	config.MigrateDB()
	config.SeedAdmin()

	// Inisialisasi router Gin
	router := gin.Default()

	// Tambahkan middleware CORS agar frontend (misalnya dari http://localhost:3000) dapat mengakses backend
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Ubah sesuai origin frontend Anda
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Endpoint root
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to Ticket Booking API"})
	})

	// Daftarkan route
	routes.AdminAuthRoutes(router)    // Untuk endpoint /admin/login (tanpa proteksi)
	routes.AdminRoutes(router)        // Endpoint admin (dengan proteksi JWT)
	routes.UserRoutes(router)         // Endpoint register & login user
	routes.EventRoutes(router)        // Endpoint event
	routes.OrderRoutes(router)        // Endpoint order
	routes.TicketRoutes(router)       // Endpoint ticket
	routes.AdminProfileRoutes(router) // Endpoint admin profile

	log.Println("Server running on port 8080")
	router.Run(":8080")
}
