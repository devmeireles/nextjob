package controller

import (
	"net/http"
	"strconv"

	"github.com/devmeireles/jobfinder/services"
	"github.com/devmeireles/jobfinder/utils"
	"github.com/gorilla/mux"
)

func GetAllSkills(w http.ResponseWriter, r *http.Request) {
	skills, err := services.GetAllSkills()

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	utils.ResSuc(w, skills)
}

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
