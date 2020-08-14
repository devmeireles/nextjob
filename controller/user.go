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
// @Success 200 {array} models.User
// @Router /users [get]
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetAllUsers()

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	utils.ResSuc(w, users)
}

// GetUser godoc
// @Summary Get details for a given user id
// @Description Get details of user corresponding to the input id
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the user"
// @Success 200 {object} models.User
// @Router /user/{id} [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	userReceived, err := services.GetUser(id)

	if err != nil {
		utils.ResErr(w, err, http.StatusNotFound)
		return
	}
	utils.ResSuc(w, userReceived)
}

// UpdateUser godoc
// @Summary Update user identified by the given id
// @Description Update the user corresponding to the input id
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the user to be updated"
// @Param user body models.User true "Update user"
// @Success 200 {object} models.User
// @Router /user/{id} [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)

	updatedUser, err := services.UpdateUser(&user, id)

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	utils.ResSuc(w, updatedUser)
}

// DeleteUser godoc
// @Summary Delete user identified by the given id
// @Description Delete the user corresponding to the input id
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the user to be deleted"
// @Success 204 "No Content"
// @Router /user/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	_, err = services.DeleteUser(id)

	if err != nil {
		utils.ResErr(w, err, http.StatusNotFound)
		return
	}

	utils.ResTrue(w)
}
