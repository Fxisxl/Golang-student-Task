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

	// Get student from token (for simplicity, assuming token was sent in header)
	studentID := uint(1) // Hardcoded student for now; replace with JWT verification later

	var course models.Course
	if err := database.DB.First(&course, input.CourseID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	// Enroll student in course
	enrollment := models.Enrollment{StudentID: studentID, CourseID: input.CourseID}
	if err := database.DB.Create(&enrollment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enroll in course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enrolled in course successfully"})
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

	// Get student from token (for simplicity, hardcoded)
	studentID := uint(1)

	// Find enrollment
	var enrollment models.Enrollment
	if err := database.DB.Where("student_id = ? AND course_id = ?", studentID, input.CourseID).First(&enrollment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Enrollment not found"})
		return
	}

	// Update rating
	enrollment.Rating = &input.Rating
	if err := database.DB.Save(&enrollment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to rate course"})
		return
	}

	// Recalculate course rating
	var course models.Course
	if err := database.DB.First(&course, input.CourseID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find course"})
		return
	}

	// Initialize count as int64
	var totalRating int
	var count int64 // Change the type to int64
	database.DB.Model(&models.Enrollment{}).Where("course_id = ?", input.CourseID).Select("rating").Scan(&totalRating)
	database.DB.Model(&models.Enrollment{}).Where("course_id = ?", input.CourseID).Count(&count)

	// Calculate average rating if there are enrollments
	if count > 0 {
		course.Rating = float64(totalRating) / float64(count)
	}

	database.DB.Save(&course)

	c.JSON(http.StatusOK, gin.H{"message": "Course rated successfully", "new_rating": course.Rating})
}