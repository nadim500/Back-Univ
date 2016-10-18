package main

import(
    "log"
    "net/http"
    "./common"
    "./routes"
)

func main(){
    common.StartUp()
    router := routes.InitRoutes()
    server := &http.Server{
        Addr : common.AppConfig.Server,
        Handler : router,
    }
    log.Println("Listening...")
    server.ListenAndServe()
//    log.Fatal(http.ListenAndServe(":8080", router))
}
