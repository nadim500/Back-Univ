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
