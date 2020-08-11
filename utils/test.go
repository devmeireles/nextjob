package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"regexp"

	"github.com/devmeireles/jobfinder/models"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

const projectDirName = "jobfinder"

// LoadEnv loads env vars from .env
func LoadEnv() {
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		log.WithFields(log.Fields{
			"cause": err,
			"cwd":   cwd,
		}).Fatal("Problem loading .env file")

		os.Exit(-1)
	}
}

func ParseBody(content *httptest.ResponseRecorder) models.ResponseMsg {
	parsedRes := models.ResponseMsg{}
	body, _ := ioutil.ReadAll(content.Body)
	json.Unmarshal(body, &parsedRes)

	return parsedRes

}
