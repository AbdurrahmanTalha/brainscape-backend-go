package models

import "gorm.io/gorm"

type Status string

const (
	Public  Status = "public"
	Private Status = "private"
	Review  Status = "review"
)

type Course struct {
	gorm.Model
	ID              uint   `json:"id"`
	Category        string `json:"category"`
	Description     string `json:"description"`
	Title           string `json:"title"`
	Image           string `json:"image"`
	Status          Status `json:"status"`
	CourseCreatedBy []Teacher `gorm:"many2many:course_teachers;" json:"teachers"`
}
