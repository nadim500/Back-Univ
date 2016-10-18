package routes

import(
  "github.com/gorilla/mux"
  "../controllers"
)

func SetUserRoutes(router *mux.Router) *mux.Router{
  router.HandleFunc("/user/register",controllers.Register).Methods("POST")
  router.HandleFunc("/user/login",controllers.Login).Methods("POST")
  return router
}
