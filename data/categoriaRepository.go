package data

import(
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
    "../models"
)

type CategoriaRepository struct{
	C *mgo.Collection
}

func(d *CategoriaRepository) Create(categoria *models.Categoria) error{
	obj_id := bson.NewObjectId()
	categoria.Id = obj_id
	err := d.C.Insert(&categoria)
	return err
}

func(d *CategoriaRepository) GetAll() []models.Categoria{
	var categorias []models.Categoria
	iter := d.C.Find(nil).Iter()
	result := models.Categoria{}
	for iter.Next(&result){
		categorias = append(categorias, result)
	}
	return categorias
}
