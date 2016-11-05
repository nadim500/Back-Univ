package controllers

import(
	"log"
    "encoding/json"
    "net/http"
	"../data"
	"github.com/gorilla/mux"
)

func CreateDocument(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
    var dataResource DocumentoResource
    err := json.NewDecoder(r.Body).Decode(&dataResource)
    if err != nil{
        log.Println("Error en decode nuevo documento: ",err)
        panic(err)
    }
    documento := &dataResource.Data
    context := NewContext()
    defer context.Close()
    col := context.DbCollection("documentos")
    repo := &data.DocumentoRepository{C: col}
    err = repo.Create(documento)
    if err != nil{
        log.Println("Error en crear documento: ",err)
        panic(err)
    }
	j,err := json.Marshal(DocumentoResource{Data: *documento})
	if err != nil{
		log.Println("Error en marshal documento : ",err)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func GetDocumentos(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
    context := NewContext()
    defer context.Close()
    col := context.DbCollection("documentos")
    repo := &data.DocumentoRepository{C: col}
    documents := repo.GetAll()
    log.Println("documentos : ",documents)
    j,err := json.Marshal(DocumentosResource{Data: documents})
	if err != nil{
		log.Println("Error en marshal documentos : ",err)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	
}

func GetDocumentosProject(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin","*")
	vars := mux.Vars(r)
    id := vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("documentos")
    repo := &data.DocumentoRepository{C: col}
	documentos := repo.GetAllForProject(id)
	j,err := json.Marshal(DocumentosResponsableResource{Data: documentos})
	if err != nil{
		log.Println("Error en marshal documentos for project: ",err)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

