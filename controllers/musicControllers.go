package controllers

import(
    "log"
    "time"
    "net/http"
    "io"
    "os"
    "strconv"
    "encoding/json"
    "../data"
)

func UploadFile(w http.ResponseWriter, r *http.Request){

    file, _, err := r.FormFile("musicToUpload")
    if err != nil{
        log.Println("error en FormFile :", err)
        panic(err)
    }
    defer file.Close()

    hoy := time.Now();
    hora, min, sec := hoy.Clock()
    year, month, day := hoy.Date()
    nanosecond := hoy.Nanosecond()
    nombre := strconv.Itoa(year) + month.String()+ strconv.Itoa(day) + strconv.Itoa(hora) + strconv.Itoa(min) + strconv.Itoa(sec) + strconv.Itoa(nanosecond) + ".mp3"
	url := "localhost:8080/media/"+nombre

    out, err := os.Create("./media/"+nombre)
    if err != nil{
        log.Println("error en Crear el archivo : ",err)
        panic(err)
    }
    defer out.Close()
    
    _, err = io.Copy(out,file)
    if err != nil{
        log.Println("error en Copiar el archivo: ", err)
        panic(err)
    }

    message := MessageUploadModel{
        Message: "Musica Upload con exito",
        Url: url,
    }

    j, err := json.Marshal(MessageUploadResource{Data: message})
    if err != nil{
        log.Println("error en generar la respuesta : ",err)
        panic(err)
    }

    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(j)
    
}

func CreateMusic(w http.ResponseWriter, r *http.Request){

    var dataResource MusicResource
    err := json.NewDecoder(r.Body).Decode(&dataResource)
    if err != nil{
        log.Println("error en Decoder : ",err)
        panic(err)
    }
    
    music := &dataResource.Data
    context := NewContext()
    defer context.Close()
    col := context.DbCollection("musics")
    repo := data.MusicRepository{C: col}
    err = repo.CreateMusic(music)
    if err != nil{
        log.Println("error en crear musica : ",err)
        panic(err)
    }

    j,err := json.Marshal(MusicResource{Data: *music})
    if err != nil{
        log.Println("error en Marshal : ",err)
        panic(err)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write(j)
    
}
