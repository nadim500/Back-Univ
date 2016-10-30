package routes

import(
    "github.com/gorilla/mux"
    "../controllers"
)

func SetTareasRoutes(router *mux.Router) *mux.Router{
    router.HandleFunc("/tareas",controllers.CreateTarea).Methods("POST")
	router.HandleFunc("/tareas/check",controllers.CreateTareaCheck).Methods("POST")
    router.HandleFunc("/tareas",controllers.GetTareas).Methods("GET")
    return router
}
