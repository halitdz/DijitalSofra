package models

import "gorm.io/gorm"

var DB *gorm.DB

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Balance  float64 `gorm:"default:1000"`
}
