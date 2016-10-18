package models

import(
	"time"
	"gopkg.in/mgo.v2/bson"
)

type(
	User struct{
		Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
        Email string `json:"email"`
        Password string `json:"password,omitempty"`
        HashPassword []byte `json:"hashpassword,omitempty"`
        DateCreated time.Time `json:"datecreated"`
	}
    Autor struct{
        Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
        Name string `json:"name"`
        Country string `json:"country"`
        DateCreated time.Time `json:"datecreated"`
    }
    Album struct{
        Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
        AutorId bson.ObjectId `bson:",omitempty" json:"autorid,omitempty"`
        Name string `json:"name"`
        Type string `json:"type"`
        DateRelease time.Time `json:"daterelease,omitempty"`
        Description string `json:"description"`
    }
    Music struct{
        Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
        AutorId bson.ObjectId `bson:",omitempty" json:"autorid,omitempty"`
        AlbumId bson.ObjectId `bson:",omitempty" json:"albumid,omitempty"`
        Name string `json:"name"`
        Duration int `json:"duration"`
        Url string `json:"url"`
    }
)
