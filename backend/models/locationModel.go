package models

import (
	"gorm.io/gorm"
)


type Location struct {
	gorm.Model `gorm:"embedded"`
	name string
	siteId uint
}
