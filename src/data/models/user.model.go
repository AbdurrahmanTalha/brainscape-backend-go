package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint
	FullName string
	Email    string
	Password []byte `gorm:"type:bytea"`
	Role     string
}
