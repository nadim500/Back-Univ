package routes

import(
    "github.com/gorilla/mux"
    "../controllers"
)

func SetTrabajadoresRoutes(router *mux.Router) *mux.Router{
    router.HandleFunc("/trabajadores",controllers.CreateTrabajador).Methods("POST")
    router.HandleFunc("/trabajadores",controllers.GetTrabajadores).Methods("GET")
    return router
}
