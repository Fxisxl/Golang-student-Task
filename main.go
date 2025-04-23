package main

import (
	"log"
	// "student-tracker/controllers"
	"student-tracker/database"
	"student-tracker/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv" // Add this import for godotenv
)






func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()
	database.SeedCourses()
	r := gin.Default()

	routes.AuthRoutes(r)
	routes.CourseRoutes(r)

	r.Run() // default on :8080
}