package models

import (
	"gorm.io/gorm"
)

type Site struct {
	gorm.Model `gorm:"embedded"`
	name string
	url string
	locations []Location `gorm:"foreignKey:siteId"`
}