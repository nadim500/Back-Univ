package data

import(
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
    "../models"
)

type TareaRepository struct{
	C *mgo.Collection
}

func(d *TareaRepository) Create(tarea *models.Tarea) error{
	obj_id := bson.NewObjectId()
	tarea.Id = obj_id
	err := d.C.Insert(&tarea)
	return err
}

func(d *TareaRepository) GetAll() []models.Tarea{
	var tareas []models.Tarea
	iter := d.C.Find(nil).Iter()
	result := models.Tarea{}
	for iter.Next(&result){
		tareas = append(tareas, result)
	}
	return tareas
}
func (d *TareaRepository) GetAllForProject(id string) []models.TareaTrabajador{
	var tareas []models.TareaTrabajador
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
            "$lookup": bson.M{
                "from": "tareas",
                "localField": "_id",
                "foreignField": "personalid",
                "as": "tareas",
            },
        },
		bson.M{
			"$unwind": bson.M{
				"path":"$tareas",
				"preserveNullAndEmptyArrays":true,
			},
		},
		bson.M{
			"$project": bson.M{
				"_id": 1,
				"trabajadorid": 1,
				"proyectoid": 1,
				"tareas.trabajador": "$trabajador.nombre",
				"tareas.nombre": 1,
				"tareas.datestart": 1,
				"tareas.dateend": 1,
				"tareas.dateregistro": 1,
				"tareas.daterecordatorio": 1,
				"tareas._id": 1,
				"tareas.personalid": 1,
			},
		},
		bson.M{
			"$group": bson.M{
				"_id": "$_id",
				"tareas": bson.M{
                    "$push": "$tareas",
                },
			},
		},
	}).Iter()
	result := models.PersonalTarea{}
	for iter.Next(&result){
		for _,element := range result.Tareas{
			tareas = append(tareas,element)
		}
	}
	return tareas
}
