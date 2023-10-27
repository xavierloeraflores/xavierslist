package models

import (
	"gorm.io/gorm"
)


type Category struct {
	gorm.Model `gorm:"embedded"`
	name string
	subcategories []Subcategory `gorm:"foreignKey:categoryId"`
}