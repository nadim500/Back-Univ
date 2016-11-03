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
