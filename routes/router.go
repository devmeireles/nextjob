package routes

import (
	"github.com/devmeireles/jobfinder/controller"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func (server *Server) SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/skills", controller.GetAllSkills).Methods("GET")
	r.HandleFunc("/skill/{id}", controller.GetSkill).Methods("GET")

	server.Router = r

	return r
}
