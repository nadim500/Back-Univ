package data

import(
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
    "../models"
)

type DocumentoRepository struct{
	C *mgo.Collection
}

func(d *DocumentoRepository) Create(documento *models.Documento) error{
	obj_id := bson.NewObjectId()
	documento.Id = obj_id
	err := d.C.Insert(&documento)
	return err
}

func(d *DocumentoRepository) GetAll() []models.Documento{
	var documents []models.Documento
	iter := d.C.Find(nil).Iter()
	result := models.Documento{}
	for iter.Next(&result){
		documents = append(documents, result)
	}
	return documents
}
