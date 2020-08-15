package services

import (
	"github.com/devmeireles/nextjob/models"
	"github.com/devmeireles/nextjob/utils"
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

func UpdateUser(user *models.User, id int) (*models.User, error) {
	var db = utils.DBConn()

	if err := db.Model(&models.User{}).Where("id = ?", id).Update(&user).Error; err != nil {
		return &models.User{}, err
	}

	return user, nil
}

func DeleteUser(id int) (*models.User, error) {
	var err error
	var user models.User
	var db = utils.DBConn()

	err = db.Find(&user, id).Error

	if err != nil {
		return &models.User{}, err
	}

	err = db.Delete(&user, id).Error

	if err != nil {
		return &models.User{}, err
	}

	return &user, nil
}

func CreateUserAddress(address *models.Address) (*models.Address, error) {
	var db = utils.DBConn()

	if err := db.Model(&models.Address{}).Create(&address).Error; err != nil {
		return &models.Address{}, err
	}

	return address, nil
}

func UpdateUserAddress(address *models.Address) (*models.Address, error) {
	var db = utils.DBConn()

	if err := db.Model(&models.Address{}).Where("id = ?", address.UserID).Update(&address).Error; err != nil {
		return &models.Address{}, err
	}

	return address, nil
}
