package main

import (
	"fmt"
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
		fmt.Println(err)
	}

	dbDriver := os.Getenv("DB_DRIVER")
	if dbDriver == "" {
		log.Fatal("$DB_DRIVER must be set")
	}

	dbUser := os.Getenv("DB_USER_TEST")
	if dbUser == "" {
		log.Fatal("$DB_USER_TEST must be set")
	}

	dbPassword := os.Getenv("DB_PASSWORD_TEST")
	if dbPassword == "" {
		log.Fatal("$DB_PASSWORD_TEST must be set")
	}

	dbPort := os.Getenv("DB_PORT_TEST")
	if dbPort == "" {
		log.Fatal("$DB_PORT_TEST must be set")
	}

	dbHost := os.Getenv("DB_HOST_TEST")
	if dbHost == "" {
		log.Fatal("$DB_HOST_TEST must be set")
	}

	dbName := os.Getenv("DB_NAME_TEST")
	if dbName == "" {
		log.Fatal("$DB_DRIVER must be set")
	}

	utils.InitDatabase(
		dbDriver,
		dbUser,
		dbPassword,
		dbPort,
		dbHost,
		dbName,
	)

	r := server.SetupRoutes()

	port := os.Getenv("PORT")

	if port == "" {
		port = "3333"
	}

	fmt.Println(port)

	log.Fatal(http.ListenAndServe(":"+port, r))
}
