package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	CourseCreated []Course `gorm:"many2many:course_teachers;" json:"createdCourses"`
}
