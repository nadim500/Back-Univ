package models

import(
	"time"
	"gopkg.in/mgo.v2/bson"
)

type(
	User struct{
		Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Email string `json:"email"`
		Password string `json:"password,omitempty"`
		HashPassword []byte `json:"hashpassword,omitempty"`
		DateCreated time.Time `json:"datecreated"`
	}

	Proyecto struct{
		Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
		UserId bson.ObjectId `bson:",omitempty" json:"userid,omitempty"`
		Codigo string `json:"codigo"`
		Nombre string `json:"nombre"`
		Descripcion string `json:"descripcion"`
		Status int `json:"status"`
		DateStart time.Time `json:"datestart"`
		DateEnd time.Time `json:"dateend"`
		DateEndFake time.Time `json:"dateendfake"`
	}

    Documento struct{
        Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
        CategoriaId bson.ObjectId `bson:",omitempty" json:"categoriaid,omitempty"`
        ProyectoId bson.ObjectId `bson:",omitempty" json:"proyectoid,omitempty"`
        PersonalId bson.ObjectId `bson:",omitempty" json:"personalid,omitempty"`
        Nombre string `json:"nombre"`
        Version string `json:"version"`
        DateCreated time.Time `json:"datecreated"`
		DateRecordatorio time.Time `json:"daterecordatorio"`
        Comentario string `json:"comentario"`
        UrlDocument string `json:"urldocument"`
    }

    Categoria struct{
        Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
        Nombre string `json:"nombre"`
    }

    Trabajador struct{
        Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
        UserId bson.ObjectId `bson:",omitempty" json:"userid,omitempty"`
        Nombre string `json:"nombre"`
    }

    Personal struct{
        Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
        ProyectoId bson.ObjectId `bson:",omitempty" json:"proyectoid,omitempty"`
        TrabajadorId bson.ObjectId `bson:",omitempty" json:"trabajadorid,omitempty"`
        DocumentoId bson.ObjectId `bson:",omitempty" json:"documentoid,omitempty"`
    }

    Tarea struct{
        Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
        PersonalId bson.ObjectId `bson:",omitempty" json:"personalid,omitempty"`
        Nombre string `json:"nombre"`
        DateStart time.Time `json:"datestart"`
        DateEnd time.Time `json:"dateend"`
    }

    DocumentoCategory struct{
        Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
        CategoriaId bson.ObjectId `bson:",omitempty" json:"categoriaid,omitempty"`
        ProyectoId bson.ObjectId `bson:",omitempty" json:"proyectoid,omitempty"`
        PersonalId bson.ObjectId `bson:",omitempty" json:"personalid,omitempty"`
        Nombre string `json:"nombre"`
        Version string `json:"version"`
        DateCreated time.Time `json:"datecreated"`
        Comentario string `json:"comentario"`
        UrlDocument string `json:"urldocument"`
        Category Categoria `json:"category"`
    }

    PersonalTrabajador struct{
        Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
        ProyectoId bson.ObjectId `bson:",omitempty" json:"proyectoid,omitempty"`
        TrabajadorId bson.ObjectId `bson:",omitempty" json:"trabajadorid,omitempty"`
        DocumentoId bson.ObjectId `bson:",omitempty" json:"documentoid,omitempty"`
        Trabajador Trabajador `json:"trabajador"`
    }

    DocumentoCategoryPersonal struct{
        Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
        CategoriaId bson.ObjectId `bson:",omitempty" json:"categoriaid,omitempty"`
        ProyectoId bson.ObjectId `bson:",omitempty" json:"proyectoid,omitempty"`
        PersonalId bson.ObjectId `bson:",omitempty" json:"personalid,omitempty"`
        Nombre string `json:"nombre"`
        Version string `json:"version"`
        DateCreated time.Time `json:"datecreated"`
        Comentario string `json:"comentario"`
        UrlDocument string `json:"urldocument"`
        Categoria Categoria `json:"categoria"`
        Personal PersonalTrabajador `json:"personal"`
    }

    TareaTrabajador struct{
        Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
        PersonalId bson.ObjectId `bson:",omitempty" json:"personalid,omitempty"`
        Nombre string `json:"nombre"`
        DateStart time.Time `json:"datestart"`
        DateEnd time.Time `json:"dateend"`
        Trabajador string `json:"trabajador"`
    }

    PersonalTarea struct{
        Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
        ProyectoId bson.ObjectId `bson:",omitempty" json:"proyectoid,omitempty"`
        TrabajadorId bson.ObjectId `bson:",omitempty" json:"trabajadorid,omitempty"`
        DocumentoId bson.ObjectId `bson:",omitempty" json:"documentoid,omitempty"`
        Trabajador Trabajador `json:"trabajador"`
        //Tareas []Tarea `json:"tareas"`
    }

    ProyectoWithAll struct{
        Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
		UserId bson.ObjectId `bson:",omitempty" json:"userid,omitempty"`
		Codigo string `json:"codigo"`
		Nombre string `json:"nombre"`
		Descripcion string `json:"descripcion"`
		Status int `json:"status"`
		DateStart time.Time `json:"datestart"`
		DateEnd time.Time `json:"dateend"`
		DateEndFake time.Time `json:"dateendfake"`
        Documents []DocumentoCategoryPersonal `json:"documents"`
        Tareas []TareaTrabajador `json:"tareas"`
    }

	UserProjectModel struct{
		Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Email string `json:"email"`
		Password string `json:"password,omitempty"`
		HashPassword []byte `json:"hashpassword,omitempty"`
		DateCreated time.Time `json:"datecreated"`
		Projects []Proyecto `json:"projects"`
	}
	
)
