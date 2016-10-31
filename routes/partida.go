package routes

import(
    "github.com/gorilla/mux"
    "../controllers"
)

func SetPartidaRoutes(router *mux.Router) *mux.Router{
    router.HandleFunc("/partidas",controllers.CreatePartida).Methods("POST")
    router.HandleFunc("/partidas",controllers.GetPartidas).Methods("GET")
    router.HandleFunc("/partidas/project/{id}",controllers.GetPartidasProject).Methods("GET")
    return router
}
