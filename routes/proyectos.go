package routes

import(
	"github.com/gorilla/mux"
	"../controllers"
)

func SetProyectoRoutes(router *mux.Router) *mux.Router{
	router.HandleFunc("/proyectos",controllers.CreateProject).Methods("POST")
	router.HandleFunc("/proyectos",controllers.GetProjects).Methods("GET")
	return router
}
