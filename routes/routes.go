package routes

import (
	"github.com/gin-gonic/gin"
	"student-tracker/controllers"
	"student-tracker/middleware"
)

func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/")

	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)
	auth.GET("/courses", controllers.GetCourses)
}

func CourseRoutes(r *gin.Engine) {
	courses := r.Group("/courses")
	courses.Use(middleware.AuthMiddleware()) //  Apply JWT middleware


	
	courses.POST("/enroll", controllers.EnrollCourse)
	courses.POST("/rate", controllers.RateCourse)
	courses.GET("/my", controllers.GetMyCourses)
}



