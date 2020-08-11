package tests

import (
	"os"
	"testing"

	"github.com/devmeireles/jobfinder/routes"
	"github.com/devmeireles/jobfinder/utils"
)

var server = routes.Server{}

func TestMain(m *testing.M) {
	utils.LoadEnv()

	utils.InitDatabase(
		os.Getenv("DB_DRIVER_TEST"),
		os.Getenv("DB_USER_TEST"),
		os.Getenv("DB_PASSWORD_TEST"),
		os.Getenv("DB_PORT_TEST"),
		os.Getenv("DB_HOST_TEST"),
		os.Getenv("DB_NAME_TEST"),
	)

	server.SetupRoutes()

	os.Exit(m.Run())
}
