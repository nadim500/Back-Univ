package common

import(
    "encoding/json"
    "log"
    "os"
)

type(
    configuration struct{
        Server, MongoDBHost, DBUser, DBPwd, Database string
    }
)

var AppConfig configuration

func initConfig(){
    loadAppConfig()
}

func loadAppConfig(){
    file, err := os.Open("./common/config.json")
    defer file.Close()
    if err != nil{
        log.Fatalf("[loadConfig]: %s\n", err)
        panic(err)
    }
    decoder := json.NewDecoder(file)
    AppConfig = configuration{}
    err = decoder.Decode(&AppConfig)
    if err != nil{
        log.Fatalf("[loadAppConfig]: %s\n", err)
        panic(err)
    }
}
