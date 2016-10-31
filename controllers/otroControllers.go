package controllers

import(
	"log"
    "encoding/json"
    "net/http"
	"../data"
)

func CreateOtro(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
    var dataResource OtroResource
    err := json.NewDecoder(r.Body).Decode(&dataResource)
    if err != nil{
        log.Println("Error en decode nuevo otro: ",err)
        panic(err)
    }
    otro := &dataResource.Data
    context := NewContext()
    defer context.Close()
    col := context.DbCollection("otros")
    repo := &data.OtroRepository{C: col}
    err = repo.Create(otro)
    if err != nil{
        log.Println("Error en crear otro: ",err)
        panic(err)
    }
	j,err := json.Marshal(OtroResource{Data: *otro})
	if err != nil{
		log.Println("Error en marshal otro : ",err)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func GetOtros(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
    context := NewContext()
    defer context.Close()
    col := context.DbCollection("otros")
    repo := &data.OtroRepository{C: col}
    otros := repo.GetAll()
    j,err := json.Marshal(OtrosResource{Data: otros})
	if err != nil{
		log.Println("Error en marshal otros : ",err)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	
}
