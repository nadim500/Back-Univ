package controllers

import(
	"log"
    "encoding/json"
    "net/http"
	"../data"
	"github.com/gorilla/mux"
)

func CreatePersonal(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
    var dataResource PersonalResource
    err := json.NewDecoder(r.Body).Decode(&dataResource)
    if err != nil{
        log.Println("Error en decode nuevo personal: ",err)
        panic(err)
    }
    personal := &dataResource.Data
    context := NewContext()
    defer context.Close()
    col := context.DbCollection("personals")
    repo := &data.PersonalRepository{C: col}
    err = repo.Create(personal)
    if err != nil{
        log.Println("Error en crear personal: ",err)
        panic(err)
    }
	j,err := json.Marshal(PersonalResource{Data: *personal})
	if err != nil{
		log.Println("Error en marshal personal : ",err)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func GetPersonals(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
    context := NewContext()
    defer context.Close()
    col := context.DbCollection("personals")
    repo := &data.PersonalRepository{C: col}
    personals := repo.GetAll()
    j,err := json.Marshal(PersonalsResource{Data: personals})
	if err != nil{
		log.Println("Error en marshal personals : ",err)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	
}

func GetPersonalsProject(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
    id := vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("personals")
	repo := data.PersonalRepository{C: col}
	personals := repo.GetAllForProject(id)
	j,err := json.Marshal(TrabajadoresResource{Data: personals})
	if err != nil{
		log.Println("Error en marshal personals for project: ",err)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
