package routes

import (
	"github.com/devmeireles/jobfinder/controller"

	"github.com/gorilla/mux"

	_ "github.com/devmeireles/jobfinder/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Server struct {
	Router *mux.Router
}

// @title NextJob API
// @version 1.0
// @description This is a sample serice for managing orders
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email dev.meireles@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3333
// @BasePath /
func (server *Server) SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/skills", controller.GetAllSkills).Methods("GET")
	r.HandleFunc("/skill/{id}", controller.GetSkill).Methods("GET")
	r.HandleFunc("/skill", controller.CreateSkill).Methods("POST")
	r.HandleFunc("/skill/{id}", controller.UpdateSkill).Methods("PUT")

	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	server.Router = r

	return r
}
