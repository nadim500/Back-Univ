package routes

import(
    "github.com/gorilla/mux"
    "../controllers"
)

func SetAlbumRoutes(router *mux.Router) *mux.Router{
    router.HandleFunc("/albums",controllers.CreateAlbum).Methods("POST")
    return router
}
