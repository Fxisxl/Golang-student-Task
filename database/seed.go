package database

import (
	"log"
	"student-tracker/models"
)

func SeedCourses() {
	var count int64
	if err := DB.Model(&models.Course{}).Count(&count).Error; err != nil {
		log.Println("Failed to count courses:", err)
		return
	}

	if count > 0 {
		log.Println("Courses already exist, skipping seeding process")
		return
	}

	courses := []models.Course{
		{Name: "Go Programming", Rating: 70},
		{Name: "Web Development with React", Rating: 70},
		{Name: "Machine Learning Basics", Rating: 70},
		{Name: "Databases and SQL", Rating: 70},
		{Name: "Introduction to AI", Rating: 70},
	}

	if err := DB.Create(&courses).Error; err != nil {
		log.Println("Failed to seed courses:", err)
	} else {
		log.Println("Courses seeded successfully")
	}
}