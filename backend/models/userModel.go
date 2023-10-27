package models

import (
	"backend/utils"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model 	`gorm:"embedded"`
	username string `gorm:"unique"`
	password []byte 
	email string `gorm:"unique"`
	posts []Post `gorm:"foreignKey:userId"`
}


func (user *User) HashPassword(password string) {
	hashedPassword := utils.HashPassword(password)
	user.password = hashedPassword
}