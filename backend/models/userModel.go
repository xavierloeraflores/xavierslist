package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	username string `gorm:"unique"`
	password string
	email string `gorm:"unique"`
}


