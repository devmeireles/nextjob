package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/devmeireles/nextjob/models"
	"github.com/devmeireles/nextjob/services"
	"github.com/devmeireles/nextjob/utils"
)

// Login godoc
// @Summary Login with user and password to generate a JWT
// @Description Generate a JWT to logged user
// @Tags auth
// @Accept  json
// @Produce  json
// @Param auth body models.UserLogin true "Generate a JWT"
// @Success 200 {object} models.UserLogin
// @Router /auth/login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	userLogin := models.UserLogin{}
	err = json.Unmarshal(body, &userLogin)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	userReceived, err := services.GetUsername(userLogin.Username)
	if err != nil {
		utils.ResErr(w, err, http.StatusNotFound)
		return
	}

	err = models.CheckPasswordHash(userLogin.Password, userReceived.Password)
	if err != nil {
		utils.ResErr(w, err, http.StatusNotFound)
		return
	}

	token, err := utils.EncodeAuthToken(userReceived.ID)
	if err != nil {
		utils.ResErr(w, err, http.StatusNotFound)
		return
	}

	resp["token"] = token

	utils.ResSuc(w, resp)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input paylod
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body models.User true "Create user"
// @Success 200 {object} models.User
// @Router /auth/register [post]
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
