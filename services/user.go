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
