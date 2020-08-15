package models

import "github.com/jinzhu/gorm"

// Address struct
type Address struct {
	gorm.Model
	UserID  int    `json:"user_id" gorm:"foreignkey:ID"`
	Line    string `json:"line" gorm:"unique;not null"`
	City    string `json:"city" gorm:"not null"`
	State   string `json:"state"`
	Country string `json:"country" gorm:"not null"`
	Zipcode string `json:"zipcode"`
}
