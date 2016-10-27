package controllers

import(
	"log"
    "encoding/json"
    "net/http"
	"../data"
)

func CreateTarea(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
    var dataResource TareaResource
    err := json.NewDecoder(r.Body).Decode(&dataResource)
    if err != nil{
        log.Println("Error en decode nuevo tarea: ",err)
        panic(err)
    }
    tarea := &dataResource.Data
    context := NewContext()
    defer context.Close()
    col := context.DbCollection("tareas")
    repo := &data.TareaRepository{C: col}
    err = repo.Create(tarea)
    if err != nil{
        log.Println("Error en crear tarea: ",err)
        panic(err)
    }
	j,err := json.Marshal(TareaResource{Data: *tarea})
	if err != nil{
		log.Println("Error en marshal tarea : ",err)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func GetTareas(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
    context := NewContext()
    defer context.Close()
    col := context.DbCollection("tareas")
    repo := &data.TareaRepository{C: col}
    tareas := repo.GetAll()
    log.Println("tareas : ",tareas)
    j,err := json.Marshal(TareasResource{Data: tareas})
	if err != nil{
		log.Println("Error en marshal tareas : ",err)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	
}
