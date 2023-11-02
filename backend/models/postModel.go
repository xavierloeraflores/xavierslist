package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model    `gorm:"embedded"`
	Title         string
	Content       string
	UserId        uint
	SubcategoryId uint
	Public        bool `gorm:"default:false"`
	Location      uint
}
