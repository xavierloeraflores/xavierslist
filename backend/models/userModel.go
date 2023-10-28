package models

import (
	"backend/utils"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model 	`gorm:"embedded"`
	Username string `gorm:"unique"`
	Password []byte 
	Email string `gorm:"unique"`
	Posts []Post `gorm:"foreignKey:userId"`
}


func (user *User) HashPassword(password string) {
	hashedPassword := utils.HashPassword(password)
	user.Password = hashedPassword
}