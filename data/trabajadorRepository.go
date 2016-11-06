package data

import(
    "golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
    "../models"
    "log"
)

type TrabajadorRepository struct{
	C *mgo.Collection
}

func(d *TrabajadorRepository) Create(trabajador *models.Trabajador) error{
	obj_id := bson.NewObjectId()
	trabajador.Id = obj_id
	trabajador.Type = "trabajador"
	err := d.C.Insert(&trabajador)
	return err
}

func(d *TrabajadorRepository) Update(id string,trabajador *models.Trabajador) error{
    trabajador.Id = bson.ObjectIdHex(id)
    hpass, err := bcrypt.GenerateFromPassword([]byte(trabajador.Password), bcrypt.DefaultCost)
    if err != nil{
        log.Fatalf("[hpass]: %s\n", err)
        panic(err)
    }
    trabajador.HashPassword = hpass
    trabajador.Password = ""
    err = d.C.Update(bson.M{
        "_id": trabajador.Id,
    },bson.M{
        "$set": bson.M{
            "email": trabajador.Email,
            "nombre": trabajador.Nombre,
            "hashpassword": trabajador.HashPassword,
            "type": "responsable",
        },
    })
    return err
}

func(d *TrabajadorRepository) UpdateNoPwd(id string,trabajador *models.Trabajador) error{
	err := d.C.Update(bson.M{
        "_id": bson.ObjectIdHex(id),
    },bson.M{
        "$set": bson.M{
            "email": trabajador.Email,
            "nombre": trabajador.Nombre,
        },
    })
    trabajador.Id = bson.ObjectIdHex(id)
	return err
}

func(d *TrabajadorRepository) GetAll() []models.Trabajador{
	var trabajadores []models.Trabajador
	iter := d.C.Find(nil).Iter()
	result := models.Trabajador{}
	for iter.Next(&result){
		trabajadores = append(trabajadores, result)
	}
	return trabajadores
}

func(d *TrabajadorRepository) GetAllForEntity(id string) []models.Trabajador{
	var trabajadores []models.Trabajador
	iter := d.C.Find(bson.M{
		"entityid": bson.ObjectIdHex(id),
        "$or": []bson.M{
            bson.M{
                "type": "trabajador",
            },
            bson.M{
                "type": "responsable",
            },
        },
	}).Iter()
	result := models.Trabajador{}
	for iter.Next(&result){
		trabajadores = append(trabajadores, result)
	}
	return trabajadores
}
