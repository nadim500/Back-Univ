package data

import(
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
    "../models"
)

type PersonalRepository struct{
	C *mgo.Collection
}

func(d *PersonalRepository) Create(personal *models.Personal) error{
	obj_id := bson.NewObjectId()
	personal.Id = obj_id
	err := d.C.Insert(&personal)
	return err
}

func(d *PersonalRepository) GetAll() []models.Personal{
	var personals []models.Personal
	iter := d.C.Find(nil).Iter()
	result := models.Personal{}
	for iter.Next(&result){
		personals = append(personals, result)
	}
	return personals
}

func(d *PersonalRepository) Check(personal models.CheckTarea) []models.Personal{
	var person []models.Personal
	idTrabajador:= personal.TrabajadorId
	idProyecto := personal.ProyectoId
	iter := d.C.Find(bson.M{
		"$and": []bson.M{
			bson.M{
				"trabajadorid": idTrabajador,
			},
			bson.M{
				"proyectoid": idProyecto,
			},
		},
	}).Iter()
	result := models.Personal{}
	for iter.Next(&result){
		person = append(person, result)
	}
	return person
}

func(d *PersonalRepository) GetAllForProject(id string) []models.Trabajador{
	var personals []models.Trabajador
	iter := d.C.Pipe([]bson.M{
		bson.M{
			"$match": bson.M{
				"proyectoid": bson.ObjectIdHex(id),
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from": "trabajadores",
				"localField": "trabajadorid",
				"foreignField": "_id",
				"as": "trabajador",
			},
		},
		bson.M{
			"$unwind": "$trabajador",
		},
		bson.M{
			"$project": bson.M{
				"_id": 1,
				"nombre": "$trabajador.nombre",
			},
		},
	}).Iter()
	result := models.Trabajador{}
	for iter.Next(&result){
		personals = append(personals, result)
	}
	return personals
}
