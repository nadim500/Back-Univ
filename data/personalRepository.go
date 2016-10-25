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
