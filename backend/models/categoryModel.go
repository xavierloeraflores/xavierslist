package models

import (
	"gorm.io/gorm"
)


type Category struct {
	gorm.Model `gorm:"embedded"`
	Name string
	Subcategories []Subcategory `gorm:"foreignKey:CategoryId"`
}