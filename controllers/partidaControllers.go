package controllers

import(
	"log"
    "encoding/json"
    "net/http"
	"../data"
    "github.com/gorilla/mux"
)

func CreatePartida(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
    var dataResource PartidaResource
    err := json.NewDecoder(r.Body).Decode(&dataResource)
    if err != nil{
        log.Println("Error en decode nuevo partida: ",err)
        panic(err)
    }
    partida := &dataResource.Data
    context := NewContext()
    defer context.Close()
    col := context.DbCollection("partidas")
    repo := &data.PartidaRepository{C: col}
    err = repo.Create(partida)
    if err != nil{
        log.Println("Error en crear partida: ",err)
        panic(err)
    }
	j,err := json.Marshal(PartidaResource{Data: *partida})
	if err != nil{
		log.Println("Error en marshal partida : ",err)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func GetPartidas(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
    context := NewContext()
    defer context.Close()
    col := context.DbCollection("partidas")
    repo := &data.PartidaRepository{C: col}
    partidas := repo.GetAll()
    j,err := json.Marshal(PartidasResource{Data: partidas})
	if err != nil{
		log.Println("Error en marshal partidas : ",err)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	
}

func GetPartidasProject(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
    id := vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("partidas")
	repo := data.PartidaRepository{C: col}
	partidas := repo.GetAllForProject(id)
	j,err := json.Marshal(PartidasResource{Data: partidas})
	if err != nil{
		log.Println("Error en marshal partidas for project: ",err)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
