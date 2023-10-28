package models

import (
	"gorm.io/gorm"
)


type Location struct {
	gorm.Model `gorm:"embedded"`
	Name string
	SiteId uint
}
