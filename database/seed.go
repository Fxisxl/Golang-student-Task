package database

import (
	"log"
	"student-tracker/models"
)

// SeedCourses seeds the database with sample courses
func SeedCourses() {
	courses := []models.Course{
		{Name: "Go Programming", Rating: 70},
		{Name: "Web Development with React", Rating: 70},
		{Name: "Machine Learning Basics", Rating: 70},
		{Name: "Databases and SQL", Rating: 70},
		{Name: "Introduction to AI", Rating: 70},
	}

	for _, course := range courses {
		if err := DB.Create(&course).Error; err != nil {
			log.Println("Error seeding course:", err)
		}
	}
}
