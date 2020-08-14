package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/devmeireles/jobfinder/models"
	"github.com/devmeireles/jobfinder/services"
	"github.com/devmeireles/jobfinder/utils"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input paylod
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "Create user"
// @Success 200 {object} models.User
// @Router /user [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	newUser, err := services.CreateUser(&user)

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}
	utils.ResSuc(w, newUser)
}

// GetAllUsers godoc
// @Summary Get details of all users
// @Description Get details of all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Skill
// @Router /users [get]
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	skills, err := services.GetAllUsers()

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	utils.ResSuc(w, skills)
}
