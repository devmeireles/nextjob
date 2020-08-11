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
