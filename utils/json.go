package utils

import (
	"encoding/json"
	"net/http"

	"github.com/devmeireles/jobfinder/models"
)

type UnstructuredJSON = map[string]interface{}

type Model struct{}

func ResErr(res http.ResponseWriter, err error, statusCode int) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)

	response := models.ResponseMsg{
		Success: false,
		Message: err.Error(),
	}

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(res, err.Error(), statusCode)
		return
	}

	res.Write(json)
	return
}

func ResSuc(res http.ResponseWriter, data interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	response := models.ResponseMsg{
		Success: true,
		Data:    data,
	}

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Write(json)
	return
}
