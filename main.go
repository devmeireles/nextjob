package main

import (
	"log"
	"net/http"
	"os"

	"github.com/devmeireles/nextjob/routes"
	"github.com/devmeireles/nextjob/utils"

	"github.com/joho/godotenv"
)

var server = routes.Server{}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	utils.InitDatabase(
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER_TEST"),
		os.Getenv("DB_PASSWORD_TEST"),
		os.Getenv("DB_PORT_TEST"),
		os.Getenv("DB_HOST_TEST"),
		os.Getenv("DB_NAME_TEST"),
	)

	r := server.SetupRoutes()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT must be set")
	}

	log.Fatal(http.ListenAndServe(":3333", r))
}
