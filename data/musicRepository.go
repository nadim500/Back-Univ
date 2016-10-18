package data

import(
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
    "../models"
)

type MusicRepository struct{
    C *mgo.Collection
}

func (m *MusicRepository) CreateMusic(music *models.Music) error{
    obj_id := bson.NewObjectId()
    music.Id = obj_id
    err := m.C.Insert(&music)
    return err
}
