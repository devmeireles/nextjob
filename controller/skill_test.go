package controller

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/devmeireles/jobfinder/utils"
	"github.com/gorilla/mux"
	"github.com/steinfletcher/apitest"
)

func TestSkills(t *testing.T) {
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
	r.HandleFunc("/skill/{id}", GetSkill).Methods("GET")
	r.HandleFunc("/skills", GetAllSkills).Methods("GET")
	r.HandleFunc("/skill", CreateSkill).Methods("POST")
	r.HandleFunc("/skill/{id}", UpdateSkill).Methods("PUT")
	r.HandleFunc("/skill/{id}", DeleteSkill).Methods("DELETE")
	ts := httptest.NewServer(r)
	defer ts.Close()

	// t.Run("get all skills", func(t *testing.T) {
	// 	apitest.New().
	// 		Handler(r).
	// 		Get("/skills").
	// 		Expect(t).
	// 		Status(http.StatusOK).
	// 		End()
	// })

	// t.Run("found", func(t *testing.T) {
	// 	apitest.New().
	// 		Handler(r).
	// 		Get("/skill/1").
	// 		Expect(t).
	// 		Status(http.StatusOK).
	// 		End()
	// })

	// t.Run("not found", func(t *testing.T) {
	// 	apitest.New().
	// 		Handler(r).
	// 		Get("/skill/12390482390").
	// 		Expect(t).
	// 		Status(http.StatusNotFound).
	// 		End()
	// })

	// t.Run("Try to get a wrong parsed url skill", func(t *testing.T) {
	// 	apitest.New().
	// 		Handler(r).
	// 		Get("/skill/1x").
	// 		Expect(t).
	// 		Status(http.StatusInternalServerError).
	// 		End()
	// })

	// t.Run("Create a skill", func(t *testing.T) {
	// 	var skill = models.Skill{
	// 		Status: 1,
	// 		Title:  "Google Cloud+10xx",
	// 	}
	// 	skillSave, _ := json.Marshal(skill)

	// 	apitest.New().
	// 		Handler(r).
	// 		Post("/skill").
	// 		JSON(skillSave).
	// 		Expect(t).
	// 		Status(http.StatusOK).
	// 		End()
	// })

	// t.Run("Update a skill", func(t *testing.T) {
	// 	var skill = models.Skill{
	// 		Title: "Google Cloud+20xx",
	// 	}
	// 	skillSave, _ := json.Marshal(skill)

	// 	apitest.New().
	// 		Handler(r).
	// 		Post("/skill").
	// 		JSON(skillSave).
	// 		Expect(t).
	// 		Status(http.StatusOK).
	// 		End()
	// })

	// t.Run("Update a skill but with a existent title", func(t *testing.T) {
	// 	var skill = models.Skill{
	// 		Title: "SQLxx",
	// 	}
	// 	skillUpdate, _ := json.Marshal(skill)

	// 	apitest.New().
	// 		Handler(r).
	// 		Method(http.MethodPut).
	// 		URL("/skill/1").
	// 		JSON(skillUpdate).
	// 		Expect(t).
	// 		Status(http.StatusInternalServerError).
	// 		End()
	// })

	t.Run("delete a skill", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Method(http.MethodDelete).
			URL("/skill/3").
			Expect(t).
			Status(http.StatusOK).
			End()
	})
}
