package routes

import(
    "github.com/gorilla/mux"
    "../controllers"
)

func SetCategoriasRoutes(router *mux.Router) *mux.Router{
    router.HandleFunc("/categorias",controllers.CreateCategoria).Methods("POST")
    router.HandleFunc("/categorias",controllers.GetCategorias).Methods("GET")
    return router
}
