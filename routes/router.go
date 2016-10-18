package routes

import(
    "github.com/gorilla/mux"
    "net/http"
)

func InitRoutes() *mux.Router{
    router := mux.NewRouter().StrictSlash(false)
    fs := http.StripPrefix("/media/", http.FileServer(http.Dir("./media")))
	router.PathPrefix("/media/").Handler(fs)
    router = SetUserRoutes(router)
    router = SetAlbumRoutes(router)
    router = SetMusicRoutes(router)
    return router
}
