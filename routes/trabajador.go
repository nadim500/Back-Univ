package routes

import(
    "github.com/gorilla/mux"
    "../controllers"
)

func SetTrabajadoresRoutes(router *mux.Router) *mux.Router{
    router.HandleFunc("/trabajadores",controllers.CreateTrabajador).Methods("POST")
    router.HandleFunc("/trabajadores",controllers.GetTrabajadores).Methods("GET")
    router.HandleFunc("/trabajadores/{id}", controllers.UpdateTrabajador).Methods("POST")
	router.HandleFunc("/trabajadores/entity/{id}",controllers.GetTrabajadoresEntity).Methods("GET")
    return router
}
