package controllers

import(
	"log"
    "encoding/json"
    "net/http"
	"../models"
	"../data"
	"github.com/gorilla/mux"
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

func CreateTareaCheck(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var dataResource CheckTareaResource
    err := json.NewDecoder(r.Body).Decode(&dataResource)
    if err != nil{
        log.Println("Error en decode check tarea: ",err)
        panic(err)
    }
    checktarea := dataResource.Data
    context := NewContext()
    defer context.Close()
    col := context.DbCollection("personals")
    repo := &data.PersonalRepository{C: col}
    personals := repo.Check(checktarea)
	log.Println("peron: ", personals)
	if len(personals)== 0{
		log.Println("0")
		personal := &models.Personal{
			ProyectoId: checktarea.ProyectoId,
			TrabajadorId: checktarea.TrabajadorId,
		}
		col = context.DbCollection("personals")
		repo = &data.PersonalRepository{C: col}
		err = repo.Create(personal)
		tarea := &models.Tarea{
			PersonalId: personal.Id,
			Nombre: checktarea.Nombre,
			DateStart: checktarea.DateStart,
			DateEnd: checktarea.DateEnd,
			DateRegistro: checktarea.DateRegistro,
			DateRecordatorio: checktarea.DateRecordatorio,
		}
		col = context.DbCollection("tareas")
		repo1 := &data.TareaRepository{C: col}
		err = repo1.Create(tarea)
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
	}else{
		log.Println("existe ersonal")
		tarea := &models.Tarea{
			PersonalId: personals[0].Id,
			Nombre: checktarea.Nombre,
			DateStart: checktarea.DateStart,
			DateEnd: checktarea.DateEnd,
			DateRegistro: checktarea.DateRegistro,
			DateRecordatorio: checktarea.DateRecordatorio,
		}
		col = context.DbCollection("tareas")
		repo1 := &data.TareaRepository{C: col}
		err = repo1.Create(tarea)
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
}

func GetTareas(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
    context := NewContext()
    defer context.Close()
    col := context.DbCollection("tareas")
    repo := &data.TareaRepository{C: col}
    tareas := repo.GetAll()
    j,err := json.Marshal(TareasResource{Data: tareas})
	if err != nil{
		log.Println("Error en marshal tareas : ",err)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	
}

func GetTareasProject(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
    id := vars["id"]
    context := NewContext()
    defer context.Close()
    col := context.DbCollection("personals")
    repo := &data.TareaRepository{C: col}
    tareas := repo.GetAllForProject(id)
    j,err := json.Marshal(TareasTrabajadorResource{Data: tareas})
	if err != nil{
		log.Println("Error en marshal tareas trabajador de projecto : ",err)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	
}
