package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/devmeireles/nextjob/models"
	"github.com/devmeireles/nextjob/utils"
	"github.com/gorilla/mux"
	"github.com/steinfletcher/apitest"
)

func TestUser(t *testing.T) {
	utils.LoadEnv()

	utils.InitDatabase(
		os.Getenv("DB_DRIVER_TEST"),
		os.Getenv("DB_USER_TEST"),
		os.Getenv("DB_PASSWORD_TEST"),
		os.Getenv("DB_PORT_TEST"),
		os.Getenv("DB_HOST_TEST"),
		os.Getenv("DB_NAME_TEST"),
	)

	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", GetAllUsers).Methods("GET")
	r.HandleFunc("/user", CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")
	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("Create a user", func(t *testing.T) {
		var user = models.User{
			Email:    "mail@example.com.br",
			Password: "723878g37242732243121",
			Username: "au12sernamexxx222",
			Language: "es",
		}
		userSave, _ := json.Marshal(user)

		apitest.New().
			Handler(r).
			Post("/user").
			JSON(userSave).
			Expect(t).
			Status(http.StatusOK).
			End()
	})

	t.Run("get all users", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Get("/users").
			Expect(t).
			Status(http.StatusOK).
			End()
	})

	t.Run("found", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Get("/user/40").
			Expect(t).
			Status(http.StatusOK).
			End()
	})

	t.Run("found", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Get("/user/9045904").
			Expect(t).
			Status(http.StatusNotFound).
			End()
	})

	t.Run("Update a user", func(t *testing.T) {
		var user = models.User{
			Language: "es",
			Username: "elgabo",
		}
		userSave, _ := json.Marshal(user)

		apitest.New().
			Handler(r).
			Method(http.MethodPut).
			URL("/user/40").
			JSON(userSave).
			Expect(t).
			Status(http.StatusOK).
			End()
	})

	t.Run("delete a user", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Method(http.MethodDelete).
			URL("/user/1").
			Expect(t).
			Status(http.StatusOK).
			End()
	})
}
