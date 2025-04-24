package main

import (
	"log"
	"os"
	"student-tracker/database"
	"student-tracker/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file (for local development)
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Connect to the database
	database.ConnectDB()

	// Seed courses (only when needed)
	database.SeedCourses()

	// Set up the Gin router
	r := gin.Default()

	// CORS middleware configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Get frontend URL from environment variables
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Set up routes
	routes.AuthRoutes(r)
	routes.CourseRoutes(r)

	// Set the port for the application (use environment variable for Render or default to 8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	// Run the server on the specified port
	err = r.Run(":" + port)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
