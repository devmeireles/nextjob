package models

import "github.com/jinzhu/gorm"

// Address struct
type Address struct {
	gorm.Model
	Line string `json:"line" gorm:"unique;not null"`
	City string `json:"city" gorm:"not null"`
}
