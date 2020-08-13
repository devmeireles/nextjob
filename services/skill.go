package services

import (
	"github.com/devmeireles/jobfinder/models"
	"github.com/devmeireles/jobfinder/utils"
)

func GetAllSkills() (*[]models.Skill, error) {
	var err error
	var db = utils.DBConn()

	skills := []models.Skill{}
	err = db.Model(&models.Skill{}).Find(&skills).Error

	if err != nil {
		return &[]models.Skill{}, err
	}

	return &skills, nil
}

func GetSkill(id int) (*models.Skill, error) {
	var skill models.Skill
	var err error
	var db = utils.DBConn()

	err = db.Find(&skill, id).Error

	if err != nil {
		return &models.Skill{}, err
	}

	return &skill, nil
}

func CreateSkill(skill *models.Skill) (*models.Skill, error) {
	var db = utils.DBConn()

	if err := db.Model(&models.Skill{}).Create(&skill).Error; err != nil {
		return &models.Skill{}, err
	}

	return skill, nil
}

func UpdateSkill(skill *models.Skill, id int) (*models.Skill, error) {
	var db = utils.DBConn()

	if err := db.Model(&models.Skill{}).Where("id = ?", id).Update(&skill).Error; err != nil {
		return &models.Skill{}, err
	}

	return skill, nil
}

func DeleteSkill(id int) (*models.Skill, error) {
	var err error
	var skill models.Skill
	var db = utils.DBConn()

	err = db.Find(&skill, id).Error

	if err != nil {
		return &models.Skill{}, err
	}

	err = db.Delete(&skill, id).Error

	if err != nil {
		return &models.Skill{}, err
	}

	return &skill, nil
}
