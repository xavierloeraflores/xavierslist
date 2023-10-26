package models

import (
	"backend/utils"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	username string `gorm:"unique"`
	password []byte 
	email string `gorm:"unique"`
}


func (user *User) HashPassword(password string) {
	hashedPassword := utils.HashPassword(password)
	user.password = hashedPassword
}