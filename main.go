package main

import (
	"log"
	"net/http"
	"os"

	"github.com/devmeireles/jobfinder/routes"
	"github.com/devmeireles/jobfinder/utils"

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
		os.Getenv("DB_USER_DEV"),
		os.Getenv("DB_PASSWORD_DEV"),
		os.Getenv("DB_PORT_DEV"),
		os.Getenv("DB_HOST_DEV"),
		os.Getenv("DB_NAME_DEV"),
	)

	r := server.SetupRoutes()

	port := os.Getenv("PORT")

	http.ListenAndServe(port, r)
}
