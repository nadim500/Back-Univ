package data

import(
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
    "../models"
)

type OtroRepository struct{
	C *mgo.Collection
}

func(d *OtroRepository) Create(otro *models.Otro) error{
	obj_id := bson.NewObjectId()
	otro.Id = obj_id
	err := d.C.Insert(&otro)
	return err
}

func(d *OtroRepository) GetAll() []models.Otro{
	var otros []models.Otro
	iter := d.C.Find(nil).Iter()
	result := models.Otro{}
	for iter.Next(&result){
		otros = append(otros, result)
	}
	return otros
}

func(d * OtroRepository) GetAllForProject(id string) []models.OtroAll{
	//var otros []models.PartidaOtro
	var otros []models.OtroAll
	iter := d.C.Pipe([]bson.M{
		bson.M{
			"$match": bson.M{
				"proyectoid": bson.ObjectIdHex(id),
			},
		},
		bson.M{
			"$lookup": bson.M{
                "from": "otros",
                "localField": "_id",
                "foreignField": "partidaid",
                "as": "otros",
            },
		},
		bson.M{
			"$unwind": "$otros",
		},
		bson.M{
			"$lookup": bson.M{
				"from": "personals",
                "localField": "otros.personalid",
                "foreignField": "_id",
                "as": "otros.personals",
			},
		},
		bson.M{
			"$unwind": bson.M{
				"path":"$otros.personals",
				"preserveNullAndEmptyArrays":true,
			},
        },
		bson.M{
            "$lookup": bson.M{
                "from": "trabajadores",
                "localField": "otros.personals.trabajadorid",
                "foreignField": "_id",
                "as": "otros.personals.trabajador",
            },
        },
		bson.M{
			"$unwind": bson.M{
				"path":"$otros.personals.trabajador",
				"preserveNullAndEmptyArrays":true,
			},
        },
		bson.M{
			"$project": bson.M{
				"_id": 1,
				"otros.responsable": "$otros.personals.trabajador.nombre",
                "otros.partida": "$nombre",
                "otros._id": "$otros._id",
                "otros.partidaid": "$otros.partidaid",
                "otros.personalid": "$otros.personalid",
                "otros.dateregistro": "$otros.dateregistro",
                "otros.daterecordatorio": "$otros.daterecordatorio",
                "otros.descripcion": "$otros.descripcion",
			},
		},
		bson.M{
			"$group": bson.M{
				"_id":"$_id",
				"nombre":bson.M{
					"$first": "$nombre",
				},
				"otros": bson.M{
                    "$push": "$otros",
                },
			},
		},
	}).Iter()
	result := models.PartidaOtro{}
	for iter.Next(&result){
		for _,element := range result.Otros{
			otros = append(otros,element)
		}
		//otros = append(otros, result)
	}
	return otros
}
