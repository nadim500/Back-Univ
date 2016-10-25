package controllers

import(
	"log"
    "encoding/json"
    "net/http"
	"../data"
    "github.com/gorilla/mux"
)

func CreateProject(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var dataResource ProyectoResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil{
		log.Println("Error en decode nuevo projecto : ",err)
		panic(err)
	}
	proyecto := &dataResource.Data
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("proyectos")
	repo := data.ProyectoRepository{C: col}
	err = repo.Create(proyecto)
	if err != nil{
		log.Println("Error en crear projecto : ",err)
		panic(err)
	}
	j,err := json.Marshal(ProyectoResource{Data: *proyecto})
	if err != nil{
		log.Println("Error en marshal projecto : ",err)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func GetProjectById(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Access-Control-Allow-Origin", "*");
    vars := mux.Vars(r)
    id := vars["id"]
    context := NewContext()
    defer context.Close()
    col := context.DbCollection("proyectos")
    repo := &data.ProyectoRepository{C: col}
    project, err := repo.GetById(id)
    if err != nil{
        log.Println("Error en encontrar 1 projecto: ",err)
        panic(err)
    }
    j,err := json.Marshal(ProyectoResource{Data: project})
	if err != nil{
		log.Println("Error en marshal 1 projecto : ",err)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
    
    
}

func GetProjects(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("proyectos")
	repo := data.ProyectoRepository{C: col}
	projects := repo.GetAll()
	j,err := json.Marshal(ProyectosResource{Data: projects})
	if err != nil{
		log.Println("Error en marshal projectos : ",err)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
