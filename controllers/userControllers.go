package controllers

import(
    "log"
    "encoding/json"
    "net/http"
    "../data"
    "../models"
    "../common"
)

func Register(w http.ResponseWriter, r *http.Request){
    var dataResource UserResource
    err := json.NewDecoder(r.Body).Decode(&dataResource)
    if err != nil{
        log.Println(err)
		panic(err)
    }
    user := &dataResource.Data
    context := NewContext()
    defer context.Close()
    col := context.DbCollection("users")
    repo := &data.UserRepository{C: col}
    err = repo.CreateUser(user)
    if err != nil{
        log.Println(err)
		panic(err)
    }
    user.HashPassword = nil
    j, err := json.Marshal(UserResource{Data: *user})
	if err != nil {
		log.Println(err)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func Login(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Access-Control-Allow-Origin", "*")
    var dataResource LoginResource
    err := json.NewDecoder(r.Body).Decode(&dataResource)
    if err != nil{
        log.Println(err)
        panic(err)
    }
    LoginModel := dataResource.Data
    loginUser := models.User{
        Email: LoginModel.Email,
        Password: LoginModel.Password,
    }
    context := NewContext()
    defer context.Close()
    col := context.DbCollection("users")
    repo := &data.UserRepository{C: col}
    usuario, err := repo.Login(loginUser)
    if err != nil{
        log.Println(err)
        panic(err)
    }
    log.Println("-------------------->usuario :", usuario)
    token, err := common.GenerateJWT(usuario.Email,"member")
    if err != nil{
        log.Println("error en token : ",err)
        panic(err)
    }

    usuario.HashPassword = nil

    authUser := AuthUserModel{
        User: usuario,
        Token: token,
    }

    j, err := json.Marshal(AuthUserResource{Data: authUser})
    if err != nil{
        log.Println(err)
        panic(err)
    }
    
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}


func GetProjectsOfUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("users")
	repo := data.UserRepository{C: col}
	projectos := repo.GetProjects()
	/*if err != nil{
		log.Println("Error en consulta lookup : ",err)
		panic(err)
	}*/
	j,err := json.Marshal(UserProjectResource{Data: projectos})
	if err != nil{
		log.Println("Error en marshal projectos : ",err)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
