package main

import (
	"fmt"
	"log"
	"os"
	"task-manager-api/config"
	"task-manager-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env globally
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	// Connect to MySQL
	config.ConnectDB()

	// Initialize Fiber
	app := fiber.New()

	routes.SetupRoutes(app)

	// Get PORT from environment, default to 3000 if not set
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Start server
	fmt.Println("🚀 Server running on port", port)
	log.Fatal(app.Listen(":" + port))
}
