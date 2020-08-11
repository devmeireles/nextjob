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
	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("get all skills", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Get("/skills").
			Expect(t).
			Status(http.StatusOK).
			End()
	})

	t.Run("found", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Get("/skill/1").
			Expect(t).
			Status(http.StatusOK).
			End()
	})

	t.Run("not found", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Get("/skill/12390482390").
			Expect(t).
			Status(http.StatusNotFound).
			End()
	})

	t.Run("Try to get a wrong parsed url skill", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Get("/skill/1x").
			Expect(t).
			Status(http.StatusInternalServerError).
			End()
	})

	// t.Run("Get all skills", func(t *testing.T) {
	// 	req, err := http.NewRequest("GET", "/skills", nil)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	rr := httptest.NewRecorder()
	// 	handler := http.HandlerFunc(GetAllSkills)

	// 	handler.ServeHTTP(rr, req)

	// 	parsedBody := utils.ParseBody(rr)

	// 	assert.True(t, parsedBody.Success)
	// 	assert.Equal(t, http.StatusOK, rr.Code)
	// })

	// t.Run("Get a skill", func(t *testing.T) {
	// 	req, err := http.NewRequest("GET", "/skill/1", nil)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	rr := httptest.NewRecorder()
	// 	handler := http.HandlerFunc(GetSkill)

	// 	handler.ServeHTTP(rr, req)

	// 	parsedBody := utils.ParseBody(rr)

	// 	assert.True(t, parsedBody.Success)
	// 	assert.Equal(t, http.StatusOK, rr.Code)
	// })
}
