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

// GetAllUsers godoc
// @Summary Get details of all users
// @Description Get details of all users
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} models.User
// @Router /api/users [get]
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
// @Security ApiKeyAuth
// @Param id path int true "ID of the user"
// @Success 200 {object} models.User
// @Router /api/user/{id} [get]
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
// @Security ApiKeyAuth
// @Param id path int true "ID of the user to be updated"
// @Param user body models.User true "Update user"
// @Success 200 {object} models.User
// @Router /api/user/{id} [put]
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
// @Security ApiKeyAuth
// @Param id path int true "ID of the user to be deleted"
// @Success 204 "No Content"
// @Router /api/user/{id} [delete]
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

// CreateUserAddress godoc
// @Summary Create a new user address
// @Description Create a new user address with the input paylod
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user body models.Address true "Create a user address"
// @Success 200 {object} models.Address
// @Router /api/user/address [post]
func CreateUserAddress(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	userAddress := models.Address{}
	err = json.Unmarshal(body, &userAddress)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	newUserAddress, err := services.CreateUserAddress(&userAddress)

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}
	utils.ResSuc(w, newUserAddress)
}

// UpdateUserAddress godoc
// @Summary Update user address identified by the given id
// @Description Update the user address corresponding to the input id
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user body models.Address true "Update user address"
// @Success 200 {object} models.Address
// @Router /api/user/address [put]
func UpdateUserAddress(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	userAddress := models.Address{}
	err = json.Unmarshal(body, &userAddress)

	updatedUserAddress, err := services.UpdateUserAddress(&userAddress)

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	utils.ResSuc(w, updatedUserAddress)
}
