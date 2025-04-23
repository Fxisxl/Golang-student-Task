package controllers

// import (
// 	"net/http"
// 	"student-tracker/database"
// 	"student-tracker/models"

// 	"github.com/gin-gonic/gin"
// )

// // func Register(c *gin.Context) {
// // 	var input models.Student
// // 	if err := c.ShouldBindJSON(&input); err != nil {
// // 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// // 		return
// // 	}

// // 	if err := database.DB.Create(&input).Error; err != nil {
// // 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// // 		return
// // 	}

// // 	c.JSON(http.StatusOK, input)
// // }
