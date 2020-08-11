package models

import "github.com/jinzhu/gorm"

// Skill struct
type Skill struct {
	gorm.Model
	Title  string `json:"title" gorm:"unique;not null"`
	Status int    `json:"status" sql:"DEFAULT:1"`
}
