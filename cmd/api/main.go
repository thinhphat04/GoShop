package main

import (
	"API/config"
	"API/routes"
	"log"
)

// @title           API Documentation
// @version         1.0
// @description     API for managing users, products, orders, and carts.
// @host            localhost:8080
// @BasePath        /api
func main() {
	// Load environment variables
	config.Load()

	// Set up routes
	router := routes.SetupRoutes()

	// Start the server
	log.Println("Server is running on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
