package data

import(
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
    "../models"
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
		"type": "trabajador",
	}).Iter()
	result := models.Trabajador{}
	for iter.Next(&result){
		trabajadores = append(trabajadores, result)
	}
	return trabajadores
}
