package services

import (
	"github.com/devmeireles/jobfinder/models"
	"github.com/devmeireles/jobfinder/utils"
)

func CreateUser(user *models.User) (*models.User, error) {
	var db = utils.DBConn()

	if err := db.Model(&models.User{}).Create(&user).Error; err != nil {
		return &models.User{}, err
	}

	return user, nil
}

func GetAllUsers() (*[]models.User, error) {
	var db = utils.DBConn()
	var err error
	users := []models.User{}

	err = db.Model(&models.User{}).Find(&users).Error

	if err != nil {
		return &[]models.User{}, err
	}

	return &users, nil
}

func GetUser(id int) (*models.User, error) {
	var db = utils.DBConn()
	var user models.User
	var err error

	err = db.Find(&user, id).Error

	if err != nil {
		return &models.User{}, err
	}

	return &user, nil
}
