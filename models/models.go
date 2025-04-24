package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Enrollments []Enrollment
}

type Course struct {
	gorm.Model
	Name     string  `json:"name"`
	Rating   float64 `json:"rating"`
	Enrollments []Enrollment
}

type Enrollment struct {
	gorm.Model
	StudentID uint
	CourseID  uint
	Course    Course `gorm:"foreignKey:CourseID"` //to fetch course details
	Rating    *int   `json:"rating"`
}

