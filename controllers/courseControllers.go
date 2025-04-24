package controllers

import (
	"net/http"
	"student-tracker/models"
	"student-tracker/database"
	"github.com/gin-gonic/gin"
)

func GetCourses(c *gin.Context) {
	var courses []models.Course
	if err := database.DB.Find(&courses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
		return
	}

	c.JSON(http.StatusOK, courses)
}

func EnrollCourse(c *gin.Context) {
	var input struct {
		CourseID uint `json:"course_id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//  Get student_id from JWT
	studentIDVal, exists := c.Get("student_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	studentID := studentIDVal.(uint)

	var course models.Course
	if err := database.DB.First(&course, input.CourseID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	// Prevent duplicate enrollment
	var existing models.Enrollment
	if err := database.DB.Where("student_id = ? AND course_id = ?", studentID, input.CourseID).First(&existing).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Already enrolled in this course"})
		return
	}

	enrollment := models.Enrollment{StudentID: studentID, CourseID: input.CourseID}
	if err := database.DB.Create(&enrollment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enroll"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enrolled successfully"})
}


func RateCourse(c *gin.Context) {
	var input struct {
		CourseID uint `json:"course_id"`
		Rating   int  `json:"rating"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//  Get student_id from JWT
	studentIDVal, exists := c.Get("student_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	studentID := studentIDVal.(uint)

	var enrollment models.Enrollment
	if err := database.DB.Where("student_id = ? AND course_id = ?", studentID, input.CourseID).First(&enrollment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Enrollment not found"})
		return
	}

	enrollment.Rating = &input.Rating
	if err := database.DB.Save(&enrollment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to rate course"})
		return
	}

	var totalRating int
	var count int64
	database.DB.Model(&models.Enrollment{}).Where("course_id = ? AND rating IS NOT NULL", input.CourseID).Select("SUM(rating)").Scan(&totalRating)
	database.DB.Model(&models.Enrollment{}).Where("course_id = ? AND rating IS NOT NULL", input.CourseID).Count(&count)

	var course models.Course
	if err := database.DB.First(&course, input.CourseID).Error; err == nil && count > 0 {
		course.Rating = float64(totalRating) / float64(count)
		database.DB.Save(&course)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Rated successfully", "new_rating": course.Rating})
}
