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

func(d *DocumentoRepository) GetAllForProject(id string) []models.DocumentoResponsable{
	var documentos []models.DocumentoResponsable
	iter := d.C.Pipe([]bson.M{
		bson.M{
			"$match": bson.M{
				"proyectoid": bson.ObjectIdHex(id),
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":"personals",
				"localField":"personalid",
				"foreignField":"_id",
				"as":"personal",
			},
		},
		bson.M{
			"$unwind": "$personal",
		},
		bson.M{
			"$lookup": bson.M{
				"from":"categorias",
				"localField":"categoriaid",
				"foreignField":"_id",
				"as":"categoria",
			},
		},
		bson.M{
			"$unwind": "$categoria",
		},
		bson.M{
			"$lookup": bson.M{
				"from":"trabajadores",
				"localField":"personal.trabajadorid",
				"foreignField":"_id",
				"as":"personal.trabajador",
			},
		},
		bson.M{
			"$unwind": "$personal.trabajador",
		},
		bson.M{
			"$project": bson.M{
				"_id": 1,
				"categoriaid": 1,
				"proyectoid": 1,
				"personalid": 1,
				"nombre": 1,
				"version": 1,
				"datecreated": 1,
				"daterecordatorio": 1,
				"comentario": 1,
				"urldocument": 1,
				"categoria": "$categoria.nombre",
				"responsable": "$personal.trabajador.nombre",
			},
		},
	}).Iter()
	result := models.DocumentoResponsable{}
	for iter.Next(&result){
		documentos = append(documentos,result)
	}
	return documentos
}
