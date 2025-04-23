package routes

import (
	"github.com/gin-gonic/gin"
	"student-tracker/controllers"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
}

func CourseRoutes(r *gin.Engine) {
	r.GET("/courses", controllers.GetCourses)
	r.POST("/enroll", controllers.EnrollCourse)
	r.POST("/rate", controllers.RateCourse)
}
