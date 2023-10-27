package models

import (
	"gorm.io/gorm"
)

type Subcategory struct {
	gorm.Model `gorm:"embedded"`
	name string
	categoryId uint
}
