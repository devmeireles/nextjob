package routes

import (
	"github.com/devmeireles/nextjob/controller"
	"github.com/devmeireles/nextjob/middlewares"

	"github.com/gorilla/mux"

	_ "github.com/devmeireles/nextjob/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Server struct {
	Router *mux.Router
}

func (server *Server) SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Auth Routes
	r.HandleFunc("/auth/login", controller.Login).Methods("POST")

	// Swagger Routes
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	r.Use(middlewares.AuthJwtVerify)

	// Skill routes
	r.HandleFunc("/skills", controller.GetAllSkills).Methods("GET")
	r.HandleFunc("/skill/{id:[0-9]+}", controller.GetSkill).Methods("GET")
	r.HandleFunc("/skill", controller.CreateSkill).Methods("POST")
	r.HandleFunc("/skill/{id:[0-9]+}", controller.UpdateSkill).Methods("PUT")
	r.HandleFunc("/skill/{id:[0-9]+}", controller.DeleteSkill).Methods("DELETE")

	// User routes
	r.HandleFunc("/user/address", controller.CreateUserAddress).Methods("POST")
	r.HandleFunc("/user/address", controller.UpdateUserAddress).Methods("PUT")
	r.HandleFunc("/user", controller.CreateUser).Methods("POST")
	r.HandleFunc("/users", controller.GetAllUsers).Methods("GET")
	r.HandleFunc("/user/{id:[0-9]+}", controller.GetUser).Methods("GET")
	r.HandleFunc("/user/{id:[0-9]+}", controller.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id:[0-9]+}", controller.DeleteUser).Methods("DELETE")

	server.Router = r

	return r
}
