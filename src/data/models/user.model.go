package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint `json:"id"`
	FullName string `json:"fullName"`
	Email    string	`json:"email"`
	Password []byte `gorm:"type:bytea" json:"password"`
	Role     string `json:"role"`
}
