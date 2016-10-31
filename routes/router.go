package routes

import(
    "github.com/gorilla/mux"
)

func InitRoutes() *mux.Router{
    router := mux.NewRouter().StrictSlash(false)
    router = SetUserRoutes(router)
	router = SetProyectoRoutes(router)
    router = SetDocumentosRoutes(router)
    router = SetPersonalsRoutes(router)
    router = SetTrabajadoresRoutes(router)
    router = SetCategoriasRoutes(router)
    router = SetTareasRoutes(router)
    router = SetPartidaRoutes(router)
    router = SetOtrosRoutes(router)
    return router
}
