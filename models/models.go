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

	Proyecto struct{
		Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
		UserId bson.ObjectId `bson:",omitempty" json:"userid,omitempty"`
		Codigo string `json:"codigo"`
		Nombre string `json:"nombre"`
		Descripcion string `json:"descripcion"`
		Status int `json:"status"`
		DateStart time.Time `json:"datestart"`
		DateEnd time.Time `json:"dateend"`
		DateEndFake time.Time `json:"dateendfake"`
	}

	UserProjectModel struct{
		Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Email string `json:"email"`
		Password string `json:"password,omitempty"`
		HashPassword []byte `json:"hashpassword,omitempty"`
		DateCreated time.Time `json:"datecreated"`
		Projects []Proyecto `json:"projects"`
	}
	
)
