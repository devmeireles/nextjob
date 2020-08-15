package routes

import (
	"github.com/devmeireles/nextjob/controller"

	"github.com/gorilla/mux"

	_ "github.com/devmeireles/nextjob/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Server struct {
	Router *mux.Router
}

func (server *Server) SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Skill routes
	r.HandleFunc("/skills", controller.GetAllSkills).Methods("GET")
	r.HandleFunc("/skill/{id}", controller.GetSkill).Methods("GET")
	r.HandleFunc("/skill", controller.CreateSkill).Methods("POST")
	r.HandleFunc("/skill/{id}", controller.UpdateSkill).Methods("PUT")
	r.HandleFunc("/skill/{id}", controller.DeleteSkill).Methods("DELETE")

	// User routes
	r.HandleFunc("/user/address", controller.CreateUserAddress).Methods("POST")
	r.HandleFunc("/user/address", controller.UpdateUserAddress).Methods("PUT")

	r.HandleFunc("/user", controller.CreateUser).Methods("POST")
	r.HandleFunc("/users", controller.GetAllUsers).Methods("GET")
	r.HandleFunc("/user/{id}", controller.GetUser).Methods("GET")
	r.HandleFunc("/user/{id}", controller.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", controller.DeleteUser).Methods("DELETE")

	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	server.Router = r

	return r
}
