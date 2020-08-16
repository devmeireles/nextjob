package models

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User struct
type User struct {
	gorm.Model
	Username string  `json:"username" gorm:"unique;not null"`
	Email    string  `json:"email" gorm:"unique;not null"`
	Password string  `json:"password" gorm:"not null"`
	Language string  `json:"language" gorm:"default:'en'"`
	UserType int     `json:"user_type" gorm:"default:'1'"`
	Address  Address `json:"-"`
}

// UserLogin struct
type UserLogin struct {
	Username string
	Password string
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14) // 14 is the cost for hashing the password.
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return errors.New("password incorrect")
	}
	return nil
}

func (u *User) BeforeSave() (err error) {
	password := strings.TrimSpace(u.Password)
	hashedpassword, err := HashPassword(password)
	if err != nil {
		return err
	}
	u.Password = string(hashedpassword)
	return nil

	// hashedPassword, err := HashPassword(u.Password)
	// if err != nil {
	// 	return err
	// }
	// u.Password = string(hashedPassword)
	// return nil
}
