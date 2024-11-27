package main

import (
	"project/config"
	"project/models"
	"project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	config.ConnectDatabase()

	// Migrate User model to the database
	config.DB.AutoMigrate(&models.User{})

	// Setup Gin router and routes
	router := gin.Default()
	router.LoadHTMLGlob("templates/*") // Load HTML templates
	routes.SetupRoutes(router)

	// Start the server
	router.Run(":8080")
}
