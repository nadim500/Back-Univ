package routes

import(
    "github.com/gorilla/mux"
    "../controllers"
)

func SetPersonalsRoutes(router *mux.Router) *mux.Router{
    router.HandleFunc("/personals",controllers.CreatePersonal).Methods("POST")
    router.HandleFunc("/personals",controllers.GetPersonals).Methods("GET")
	router.HandleFunc("/personals/project/{id}",controllers.GetPersonalsProject).Methods("GET")
    return router
}
