package routes

import(
    "github.com/gorilla/mux"
    "../controllers"
)

func SetMusicRoutes(router *mux.Router) *mux.Router{
    router.HandleFunc("/musics",controllers.CreateMusic).Methods("POST")
    //router.HandleFunc("/musics",controllers.GetMusics).Methods("GET")
    //router.HandleFunc("/musics/{id}",controllers.UpdateMusic).Methods("PUT")
    //router.HandleFunc("/musics/{id}",controllers.GetMusicById).Methods("GET")
    //router.HandleFunc("/musics/{id}",controller.DeleteById).Methods("DELETE")
    router.HandleFunc("/musics/uploadFile",controllers.UploadFile).Methods("POST")
    return router
}
