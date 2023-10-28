package models

import (
	"gorm.io/gorm"
)

type Site struct {
	gorm.Model `gorm:"embedded"`
	Name string
	Url string
	Locations []Location `gorm:"foreignKey:siteId"`
}