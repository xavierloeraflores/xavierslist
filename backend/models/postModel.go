package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model `gorm:"embedded"`
	title string
	content string
	userId uint

	
}