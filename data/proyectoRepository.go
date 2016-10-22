package data

import(
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
    "../models"
)

type ProyectoRepository struct{
	C *mgo.Collection
}

func(p *ProyectoRepository) Create(proyecto *models.Proyecto) error{
	obj_id := bson.NewObjectId()
	proyecto.Id = obj_id
	err := p.C.Insert(&proyecto)
	return err
}

func(p *ProyectoRepository) GetAll() []models.Proyecto{
	var projects []models.Proyecto
	iter := p.C.Find(nil).Iter()
	result := models.Proyecto{}
	for iter.Next(&result){
		projects = append(projects, result)
	}
	return projects
}
