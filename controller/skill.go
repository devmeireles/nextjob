package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/devmeireles/nextjob/models"

	"github.com/devmeireles/nextjob/services"
	"github.com/devmeireles/nextjob/utils"
	"github.com/gorilla/mux"
)

// GetAllSkills godoc
// @Summary Get details of all skills
// @Description Get details of all skills
// @Tags skills
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} models.Skill
// @Router /api/skills [get]
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
// @Security ApiKeyAuth
// @Param id path int true "ID of the skill"
// @Success 200 {object} models.Skill
// @Router /api/skill/{id} [get]
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
// @Description Create a new skill with the input
// @Tags skills
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param skill body models.Skill true "Create a skill"
// @Success 200 {object} models.Skill
// @Router /api/skill [post]
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
// @Security ApiKeyAuth
// @Param id path int true "ID of the skill to be updated"
// @Param skill body models.Skill true "Create skill"
// @Success 200 {object} models.Skill
// @Router /api/skill/{id} [put]
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

// DeleteSkill godoc
// @Summary Delete skill identified by the given id
// @Description Delete the skill corresponding to the input id
// @Tags skills
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "ID of the skill to be deleted"
// @Success 204 "No Content"
// @Router /api/skill/{id} [delete]
func DeleteSkill(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	_, err = services.DeleteSkill(id)

	if err != nil {
		utils.ResErr(w, err, http.StatusNotFound)
		return
	}

	utils.ResTrue(w)
}
