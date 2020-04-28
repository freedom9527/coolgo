package models

import "github.com/jinzhu/gorm"

type Admin struct {
	gorm.Model
	Name         string  `gorm:"size:255"`
}

func (Admin) TableName() string {
	return "admin"
}