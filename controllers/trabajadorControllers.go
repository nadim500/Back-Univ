package controllers

import(
	"log"
    "encoding/json"
    "net/http"
	"../data"
)

func CreateTrabajador(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
    var dataResource TrabajadorResource
    err := json.NewDecoder(r.Body).Decode(&dataResource)
    if err != nil{
        log.Println("Error en decode nuevo trabajador: ",err)
        panic(err)
    }
    trabajador := &dataResource.Data
    context := NewContext()
    defer context.Close()
    col := context.DbCollection("trabajadores")
    repo := &data.TrabajadorRepository{C: col}
    err = repo.Create(trabajador)
    if err != nil{
        log.Println("Error en crear trabajador: ",err)
        panic(err)
    }
	j,err := json.Marshal(TrabajadorResource{Data: *trabajador})
	if err != nil{
		log.Println("Error en marshal trabajador : ",err)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func GetTrabajadores(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
    context := NewContext()
    defer context.Close()
    col := context.DbCollection("trabajadores")
    repo := &data.TrabajadorRepository{C: col}
    trabajadores := repo.GetAll()
    j,err := json.Marshal(TrabajadoresResource{Data: trabajadores})
	if err != nil{
		log.Println("Error en marshal trabajadores : ",err)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	
}
