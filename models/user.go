package models

import "github.com/jinzhu/gorm"

// User struct
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Language string `json:"language" gorm:"default:'en'"`
	UserType int    `json:"user_type" gorm:"default:'1'"`
}
