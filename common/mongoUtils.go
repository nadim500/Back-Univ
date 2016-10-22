package common

import(
    "log"
    "time"
    "gopkg.in/mgo.v2"
)

var session *mgo.Session

func GetSession() *mgo.Session{
    if session == nil{
        var err error
        session, err = mgo.DialWithInfo(&mgo.DialInfo{
            Addrs: []string{AppConfig.MongoDBHost},
            Username: AppConfig.DBUser,
            Password: AppConfig.DBPwd,
            Timeout: 60 * time.Second,
        })
        if err != nil{
            log.Fatalf("[GetSession]: %s\n", err)
        }
    }
    return session
}

func createDbSession(){
    var err error
    session, err = mgo.DialWithInfo(&mgo.DialInfo{
        Addrs: []string{AppConfig.MongoDBHost},
        Username: AppConfig.DBUser,
        Password: AppConfig.DBPwd,
        Timeout: 60 * time.Second,
    })
    if err != nil{
        log.Fatalf("[GetSession]: %s\n", err)
    }
}

func addIndexes(){
    var err error
    userIndex := mgo.Index{
        Key: []string{"email"},
        Unique: true,
        Background: true,
        Sparse: true,
    }
	proyectoIndex := mgo.Index{
		Key: []string{"datestart"},
        Unique: false,
        Background: true,
        Sparse: true,
	}
    session := GetSession().Copy()
    defer session.Close()
    userCol := session.DB(AppConfig.Database).C("users")
    err = userCol.EnsureIndex(userIndex)
    if err != nil{
        log.Fatalf("[addIndexes]: %s\n", err)
    }
	proyectoCol := session.DB(AppConfig.Database).C("proyectos")
	err = proyectoCol.EnsureIndex(proyectoIndex)
	if err != nil{
		log.Fatalf("[addIndexes]: %s\n", err)
	}
}
