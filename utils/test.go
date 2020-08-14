package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"regexp"

	"github.com/devmeireles/nextjob/models"
	"github.com/joho/godotenv"
)

const projectDirName = "nextjob-api"

// LoadEnv loads env vars from .env
func LoadEnv() {
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		// log.WithFields(log.Fields{
		// 	"cause": err,
		// 	"cwd":   cwd,
		// }).Fatal("Problem loading .env file")

		// os.Exit(-1)

		fmt.Println(err)
	}
}

func ParseBody(content *httptest.ResponseRecorder) models.ResponseMsg {
	parsedRes := models.ResponseMsg{}
	body, _ := ioutil.ReadAll(content.Body)
	json.Unmarshal(body, &parsedRes)

	return parsedRes

}
