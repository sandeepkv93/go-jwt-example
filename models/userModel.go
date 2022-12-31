package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"Unique"`
	Password string
}
