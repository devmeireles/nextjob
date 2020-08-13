package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/devmeireles/jobfinder/models"

	"github.com/devmeireles/jobfinder/services"
	"github.com/devmeireles/jobfinder/utils"
	"github.com/gorilla/mux"
)

// GetAllSkills godoc
// @Summary Get details of all skills
// @Description Get details of all skills
// @Tags skills
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Skill
// @Router /skills [get]
func GetAllSkills(w http.ResponseWriter, r *http.Request) {
	skills, err := services.GetAllSkills()

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	utils.ResSuc(w, skills)
}

// GetSkill godoc
// @Summary Get details for a given skill id
// @Description Get details of skill corresponding to the input id
// @Tags skills
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the skill"
// @Success 200 {object} models.Skill
// @Router /skill/{id} [get]
func GetSkill(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	skillReceived, err := services.GetSkill(id)

	if err != nil {
		utils.ResErr(w, err, http.StatusNotFound)
		return
	}
	utils.ResSuc(w, skillReceived)
}

// CreateSkill godoc
// @Summary Create a new skill
// @Description Create a new skill with the input paylod
// @Tags skills
// @Accept  json
// @Produce  json
// @Param skill body models.Skill true "Create skill"
// @Success 200 {object} models.Skill
// @Router /skill [post]
func CreateSkill(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	skill := models.Skill{}
	err = json.Unmarshal(body, &skill)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	newSkill, err := services.CreateSkill(&skill)

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}
	utils.ResSuc(w, newSkill)

}

// UpdateSkill godoc
// @Summary Update skill identified by the given id
// @Description Update the skill corresponding to the input id
// @Tags skills
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the skill to be updated"
// @Param skill body models.Skill true "Create skill"
// @Success 200 {object} models.Skill
// @Router /skill/{id} [put]
func UpdateSkill(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	skill := models.Skill{}
	err = json.Unmarshal(body, &skill)

	updatedSkill, err := services.UpdateSkill(&skill, id)

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	utils.ResSuc(w, updatedSkill)
}
