package routes

import(
    "github.com/gorilla/mux"
    "../controllers"
)

func SetOtrosRoutes(router *mux.Router) *mux.Router{
    router.HandleFunc("/otros",controllers.CreateOtro).Methods("POST")
    router.HandleFunc("/otros",controllers.GetOtros).Methods("GET")
    return router
}
