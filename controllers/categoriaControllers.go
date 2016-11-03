package controllers

import(
	"log"
    "encoding/json"
    "net/http"
	"../data"
)

func CreateCategoria(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
    var dataResource CategoriaResource
    err := json.NewDecoder(r.Body).Decode(&dataResource)
    if err != nil{
        log.Println("Error en decode nuevo categoria: ",err)
        panic(err)
    }
    categoria := &dataResource.Data
    context := NewContext()
    defer context.Close()
    col := context.DbCollection("categorias")
    repo := &data.CategoriaRepository{C: col}
    err = repo.Create(categoria)
    if err != nil{
        log.Println("Error en crear categoria: ",err)
        panic(err)
    }
	j,err := json.Marshal(CategoriaResource{Data: *categoria})
	if err != nil{
		log.Println("Error en marshal categoria : ",err)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func GetCategorias(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
    context := NewContext()
    defer context.Close()
    col := context.DbCollection("categorias")
    repo := &data.CategoriaRepository{C: col}
    categorias := repo.GetAll()
    j,err := json.Marshal(CategoriasResource{Data: categorias})
	if err != nil{
		log.Println("Error en marshal categorias : ",err)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	
}
