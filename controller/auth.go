package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/devmeireles/nextjob/models"
	"github.com/devmeireles/nextjob/services"
	"github.com/devmeireles/nextjob/utils"
)

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
