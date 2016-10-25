package data

import(
    "golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
    "log"
    "../models"
    "time"
)

type UserRepository struct{
    C *mgo.Collection
}

func (r *UserRepository) CreateUser(user *models.User) error{
    obj_id := bson.NewObjectId()
    user.Id = obj_id
    user.DateCreated = time.Now()
    hpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil{
        log.Fatalf("[hpass]: %s\n", err)
        panic(err)
    }
    user.HashPassword = hpass
    user.Password = ""
    err = r.C.Insert(&user)
    return err
}

func (r *UserRepository) Login(user models.User)(u models.User, err error){
    err = r.C.Find(bson.M{"email": user.Email}).One(&u)
	if err != nil {
        u = models.User{}
		return u, err
	}

	err = bcrypt.CompareHashAndPassword(u.HashPassword, []byte(user.Password))
	if err != nil {
		u = models.User{}
        return u, err
	}
	return u, nil
}

func (r *UserRepository) GetProjects() []models.UserProjectModel{
	var users []models.UserProjectModel
	iter := r.C.Pipe([]bson.M{
		bson.M{"$lookup": bson.M{"from": "proyectos", "localField": "_id", "foreignField": "userid", "as": "projects"}},
	}).Iter()
	result := models.UserProjectModel{}
	for iter.Next(&result){
		result.HashPassword = nil
		users = append(users, result)
	}
	return users
}
