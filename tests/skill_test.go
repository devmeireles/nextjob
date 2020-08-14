package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/devmeireles/nextjob/models"
	"github.com/stretchr/testify/assert"
)

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	server.Router.ServeHTTP(rr, req)

	return rr
}

func parseBody(content *httptest.ResponseRecorder) models.ResponseMsg {
	parsedRes := models.ResponseMsg{}
	body, _ := ioutil.ReadAll(content.Body)
	json.Unmarshal(body, &parsedRes)

	return parsedRes

}

func TestGetSkill(t *testing.T) {
	t.Run("Get all skills", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/skills", nil)
		response := executeRequest(req)
		parsedBody := parseBody(response)

		assert.True(t, parsedBody.Success)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("Get a skill", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/skill/1", nil)
		response := executeRequest(req)
		parsedBody := parseBody(response)

		assert.True(t, parsedBody.Success)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("Get a nonexistent skill", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/skill/190347", nil)
		response := executeRequest(req)
		parsedBody := parseBody(response)

		assert.False(t, parsedBody.Success)
		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Try to get a wrong parsed url skill", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/skill/12x", nil)
		response := executeRequest(req)
		parsedBody := parseBody(response)

		assert.False(t, parsedBody.Success)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("Create a skill", func(t *testing.T) {
		var skill = models.Skill{
			Status: 1,
			Title:  "Redis",
		}

		skillSave, _ := json.Marshal(skill)

		req, _ := http.NewRequest("POST", "/skill", bytes.NewBuffer(skillSave))

		response := executeRequest(req)
		parsedBody := parseBody(response)

		assert.True(t, parsedBody.Success)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("Try to save existent a skill", func(t *testing.T) {
		var skill = models.Skill{
			Status: 1,
			Title:  "MySQL",
		}

		skillSave, _ := json.Marshal(skill)

		req, _ := http.NewRequest("POST", "/skill", bytes.NewBuffer(skillSave))

		response := executeRequest(req)
		parsedBody := parseBody(response)

		assert.False(t, parsedBody.Success)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}