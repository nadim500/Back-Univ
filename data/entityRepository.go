package data

import(
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
    "../models"
)

type EntityRepository struct{
	C *mgo.Collection
}

func(d *EntityRepository) Create(entity *models.Entity) error{
	obj_id := bson.NewObjectId()
	entity.Id = obj_id
	err := d.C.Insert(&entity)
	return err
}

func(d *EntityRepository) GetAll() []models.Entity{
	var entitys []models.Entity
	iter := d.C.Find(nil).Iter()
	result := models.Entity{}
	for iter.Next(&result){
		entitys = append(entitys, result)
	}
	return entitys
}
