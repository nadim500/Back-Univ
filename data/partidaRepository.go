package data

import(
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
    "../models"
)

type PartidaRepository struct{
	C *mgo.Collection
}

func(d *PartidaRepository) Create(partida *models.Partida) error{
	obj_id := bson.NewObjectId()
	partida.Id = obj_id
	err := d.C.Insert(&partida)
	return err
}

func(d *PartidaRepository) GetAll() []models.Partida{
	var partidas []models.Partida
	iter := d.C.Find(nil).Iter()
	result := models.Partida{}
	for iter.Next(&result){
		partidas = append(partidas, result)
	}
	return partidas
}

func(d *PartidaRepository) GetAllForProject(id string) []models.Partida{
    var partidas []models.Partida
	iter := d.C.Pipe([]bson.M{
		bson.M{
			"$match": bson.M{
				"proyectoid": bson.ObjectIdHex(id),
			},
		},
	}).Iter()
	result := models.Partida{}
	for iter.Next(&result){
		partidas = append(partidas, result)
	}
	return partidas
}
