package routes

import(
    "github.com/gorilla/mux"
    "../controllers"
)

func SetDocumentosRoutes(router *mux.Router) *mux.Router{
    router.HandleFunc("/documentos",controllers.CreateDocument).Methods("POST")
    router.HandleFunc("/documentos",controllers.GetDocumentos).Methods("GET")
    return router
}
