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
		EntityId bson.ObjectId `bson:",omitempty" json:"entityid,omitempty"`
		TrabajadorId bson.ObjectId `bson:",omitempty" json:"trabajadorid,omitempty"`
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

	DocumentoResponsable struct{
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
		Categoria string `json:"categoria"`
		Responsable string `json:"responsable"`
    }

    Categoria struct{
        Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
        Nombre string `json:"nombre"`
    }

    Partida struct{
        Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
        ProyectoId bson.ObjectId `bson:",omitempty" json:"proyectoid,omitempty"`
        Nombre string `json:"nombre"`
    }

	PartidaOtro struct{
        Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
        ProyectoId bson.ObjectId `bson:",omitempty" json:"proyectoid,omitempty"`
        Nombre string `json:"nombre"`
		Otros []OtroAll `json:"otros"`
    }

    Otro struct{
        Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
        PartidaId bson.ObjectId `bson:",omitempty" json:"partidaid,omitempty"`
        PersonalId bson.ObjectId `bson:",omitempty" json:"personalid,omitempty"`
        DateRegistro time.Time `json:"dateregistro"`
		DateRecordatorio time.Time `json:"daterecordatorio"`
        Descripcion string `json:"descripcion"`
    }

    OtroAll struct{
        Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
        PartidaId bson.ObjectId `bson:",omitempty" json:"partidaid,omitempty"`
        PersonalId bson.ObjectId `bson:",omitempty" json:"personalid,omitempty"`
        DateRegistro time.Time `json:"dateregistro"`
		DateRecordatorio time.Time `json:"daterecordatorio"`
        Descripcion string `json:"descripcion"`
        Responsable string `json:"responsable"`
        Partida string `json:"partida"`
    }

    Trabajador struct{
        Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
		EntityId bson.ObjectId `bson:",omitempty" json:"entityid,omitempty"`
		Email string `json:"email"`
		Password string `json:"password,omitempty"`
		HashPassword []byte `json:"hashpassword,omitempty"`
		Type string `json:"type"`
        Nombre string `json:"nombre"`
    }

	Entity struct{
		Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Nombre string `json:"nombre"`
	}

    Personal struct{
        Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
        ProyectoId bson.ObjectId `bson:",omitempty" json:"proyectoid,omitempty"`
        TrabajadorId bson.ObjectId `bson:",omitempty" json:"trabajadorid,omitempty"`
    }

	PersonalTarea struct{
        Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
        ProyectoId bson.ObjectId `bson:",omitempty" json:"proyectoid,omitempty"`
        TrabajadorId bson.ObjectId `bson:",omitempty" json:"trabajadorid,omitempty"`
		Tareas []TareaTrabajador `json:"tareas"`
    }

    Tarea struct{
        Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
        PersonalId bson.ObjectId `bson:",omitempty" json:"personalid,omitempty"`
        Nombre string `json:"nombre"`
        DateStart time.Time `json:"datestart"`
        DateEnd time.Time `json:"dateend"`
		DateRegistro time.Time `json:"dateregistro"`
		DateRecordatorio time.Time `json:"daterecordatorio"`
    }

	CheckTarea struct{
		ProyectoId bson.ObjectId `bson:",omitempty" json:"proyectoid,omitempty"`
        TrabajadorId bson.ObjectId `bson:",omitempty" json:"trabajadorid,omitempty"`
        Nombre string `json:"nombre"`
        DateStart time.Time `json:"datestart"`
        DateEnd time.Time `json:"dateend"`
		DateRegistro time.Time `json:"dateregistro"`
		DateRecordatorio time.Time `json:"daterecordatorio"`
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
		DateRegistro time.Time `json:"dateregistro"`
		DateRecordatorio time.Time `json:"daterecordatorio"`
    }

    ProyectoWithAll struct{
        Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
		EntityId bson.ObjectId `bson:",omitempty" json:"entityid,omitempty"`
		TrabajadorId bson.ObjectId `bson:",omitempty" json:"trabajadorid,omitempty"`
		Codigo string `json:"codigo"`
		Nombre string `json:"nombre"`
		Descripcion string `json:"descripcion"`
		Status int `json:"status"`
		DateStart time.Time `json:"datestart"`
		DateEnd time.Time `json:"dateend"`
		DateEndFake time.Time `json:"dateendfake"`
        Documents []DocumentoCategoryPersonal `json:"documents"`
        Tareas []TareaTrabajador `json:"tareas"`
        Otros []OtroAll `json:"otros"`
        //Partidas Partida `json:"partidas"`
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
