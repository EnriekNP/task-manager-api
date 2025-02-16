package main

import (
	"log"
	"task-manager-api/config"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	// Connect to MySQL
	config.ConnectDB()

	// Run migration
	config.MigrateDB()
}
