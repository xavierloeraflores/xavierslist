package models

import (
	"gorm.io/gorm"
)

type Subcategory struct {
	gorm.Model `gorm:"embedded"`
	Name       string
	CategoryId uint
}
