package main

import (
	"log"
	"os"

	"klz-server/internal/handlers"
	"klz-server/internal/models"
	"klz-server/pkg/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	db, err := database.NewConnection()
	if err != nil {
		log.Fatal("Could not initialize database connection")
	}

	// Auto Migrate database
	db.AutoMigrate(&models.User{})

	app := fiber.New(fiber.Config{
		AppName: "KLZ API v1.0",
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Routes
	api := app.Group("/api")

	// Health check
	api.Get("/health", handlers.HealthCheck)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
