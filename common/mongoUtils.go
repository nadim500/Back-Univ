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
    autorIndex := mgo.Index{
        Key: []string{"datecreated"},
        Unique: false,
        Background: true,
        Sparse: true,
    }
    albumIndex := mgo.Index{
        Key: []string{"autorid"},
        Unique: false,
        Background: true,
        Sparse: true,
    }
    musicIndex := mgo.Index{
        Key: []string{"albumid"},
        Unique: false,
        Background: true,
        Sparse: true,
    }
    session := GetSession().Copy()
    defer session.Close()
    userCol := session.DB(AppConfig.Database).C("users")
    autorCol := session.DB(AppConfig.Database).C("autors")
    albumCol := session.DB(AppConfig.Database).C("albums")
    musicCol := session.DB(AppConfig.Database).C("musics")
    err = userCol.EnsureIndex(userIndex)
    if err != nil{
        log.Fatalf("[addIndexes]: %s\n", err)
    }
    err = autorCol.EnsureIndex(autorIndex)
    if err != nil{
        log.Fatalf("[addIndexes]: %s\n", err)
    }
    err = albumCol.EnsureIndex(albumIndex)
    if err != nil{
        log.Fatalf("[addIndexes]: %s\n", err)
    }
    err = musicCol.EnsureIndex(musicIndex)
    if err != nil{
        log.Fatalf("[addIndexes]: %s\n", err)
    }
    
}
