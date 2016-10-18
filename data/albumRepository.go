package data

import(
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
    "../models"
    "time"
)

type AlbumRepository struct{
    C *mgo.Collection
}

func (a *AlbumRepository) CreateAlbum(album *models.Album) error{
    obj_id := bson.NewObjectId()
    album.Id = obj_id
    album.DateRelease = time.Now()
    err := a.C.Insert(&album)
    return err
}
