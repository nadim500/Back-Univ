package controllers

import(
    "log"
    "net/http"
    "encoding/json"
    "../data"
)

func CreateAlbum(w http.ResponseWriter, r *http.Request){

    var dataResource AlbumResource
    err := json.NewDecoder(r.Body).Decode(&dataResource)
    if err != nil{
        log.Println("error en Decode : ",err)
        panic(err)
    }

    album := &dataResource.Data
    context := NewContext()
    defer context.Close()
    col := context.DbCollection("albums")
    repo := data.AlbumRepository{C: col}
    err = repo.CreateAlbum(album)
    if err != nil{
        log.Println("error en crear album : ",err)
        panic(err)
    }

    j,err := json.Marshal(AlbumResource{Data: *album})
    if err != nil{
        log.Println("error en Marshal : ",err)
        panic(err)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write(j)
    
}
