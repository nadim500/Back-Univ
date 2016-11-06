package controllers

import(
	"log"
    "encoding/json"
    "net/http"
	"../data"
	"github.com/gorilla/mux"
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

func UpdateTrabajador(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Access-Control-Allow-Origin", "*")
    vars := mux.Vars(r)
    id := vars["id"]
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
    if trabajador.Password == ""{
        log.Println("nil password")
        err = repo.UpdateNoPwd(id,trabajador)
        if err != nil{
            log.Println("Error en editar trabajador sin pwd: ",err)
            panic(err)
        }
        j,err := json.Marshal(TrabajadorResource{Data: *trabajador})
        if err != nil{
            log.Println("Error en marshal trabajador : ",err)
            panic(err)
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(j)
    }else{
        err = repo.Update(id,trabajador)
        if err != nil{
            log.Println("Error en editar trabajador sin pwd: ",err)
            panic(err)
        }
        trabajador.HashPassword = nil
        j,err := json.Marshal(TrabajadorResource{Data: *trabajador})
        if err != nil{
            log.Println("Error en marshal trabajador : ",err)
            panic(err)
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(j)
    }
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

func GetTrabajadoresEntity(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
    id := vars["id"]
    context := NewContext()
    defer context.Close()
    col := context.DbCollection("trabajadores")
    repo := &data.TrabajadorRepository{C: col}
    trabajadores := repo.GetAllForEntity(id)
    j,err := json.Marshal(TrabajadoresResource{Data: trabajadores})
	if err != nil{
		log.Println("Error en marshal trabajadores : ",err)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)	
}
