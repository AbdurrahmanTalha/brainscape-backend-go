package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Courses          []EnrolledCourse `json:"courses"`
	RecommendCourses []string         `json:"recommendCourses"`
	DailyStreak      []DailyStreak    `json:"dailyStreak"`
}

type EnrolledCourse struct {
	CourseID string   `json:"courseId"`
	Progress []string `json:"progress"`
}

type DailyStreak struct {
	Day    string `json:"day"`
	Date   int64  `json:"date"`
	Active bool   `json:"active"`
}
