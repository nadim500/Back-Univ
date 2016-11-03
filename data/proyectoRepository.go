package data

import(
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
    "../models"
	"log"
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

func(p *ProyectoRepository) GetById(id string) (project models.Proyecto, err error){
    err = p.C.FindId(bson.ObjectIdHex(id)).One(&project)
    return
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

func(p *ProyectoRepository) GetAllForEntity(id string) []models.Proyecto{
	var projects []models.Proyecto
	iter := p.C.Find(bson.M{
		"entityid" : bson.ObjectIdHex(id),
	}).Iter()
	result := models.Proyecto{}
	for iter.Next(&result){
		projects = append(projects, result)
	}
	return projects
}

func(p *ProyectoRepository) GetAllTodo(id string) []models.ProyectoWithAll{
	var projects []models.ProyectoWithAll
	log.Println("VINO AQUI")
	iter := p.C.Pipe([]bson.M{
		bson.M{
			"$match": bson.M{
				"_id": bson.ObjectIdHex(id),
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":"documentos",
				"localField":"_id",
				"foreignField":"proyectoid",
				"as":"documents",
			},
		},
		bson.M{
			"$unwind": bson.M{
				"path":"$documents",
				"preserveNullAndEmptyArrays":true,
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":"personals",
				"localField":"documents.personalid",
				"foreignField":"_id",
				"as":"documents.personal",
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":"categorias",
				"localField":"documents.categoriaid",
				"foreignField":"_id",
				"as":"documents.categoria",
			},
		},
		bson.M{
			"$unwind": bson.M{
				"path":"$documents.personal",
				"preserveNullAndEmptyArrays":true,
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":"trabajadores",
				"localField":"documents.personal.trabajadorid",
				"foreignField":"_id",
				"as":"documents.personal.trabajador",
			},
		},
		bson.M{
			"$unwind": bson.M{
				"path":"$documents.personal.trabajador",
				"preserveNullAndEmptyArrays":true,
			},
		},
		bson.M{
			"$unwind": bson.M{
				"path":"$documents.categoria",
				"preserveNullAndEmptyArrays":true,
			},
		},
		bson.M{
			"$group": bson.M{
				"_id":"$_id",
                "userid":bson.M{
                    "$first": "$userid",
                },
				"codigo":bson.M{
					"$first": "$codigo",
				},
                "entityid":bson.M{
                    "$first": "$entityid",
                },
                "trabajadorid":bson.M{
                    "$first": "$trabajadorid",
                },
				"nombre":bson.M{
					"$first": "$nombre",
				},
				"descripcion":bson.M{
					"$first": "$descripcion",
				},
				"datestart":bson.M{
					"$first": "$datestart",
				},
				"dateend":bson.M{
					"$first": "$dateend",
				},
				"dateendfake":bson.M{
					"$first": "$dateendfake",
				},
				"status":bson.M{
					"$first": "$status",
				},
				"documents":bson.M{
					"$push": "$documents",
				},
			},
		},

        bson.M{
			"$lookup": bson.M{
				"from":"personals",
				"localField":"_id",
				"foreignField":"proyectoid",
				"as":"personals",
			},
		},
        
        bson.M{
			"$unwind": bson.M{
				"path":"$personals",
				"preserveNullAndEmptyArrays":true,
			},
		},
        
		bson.M{
			"$lookup": bson.M{
				"from": "trabajadores",
				"localField": "personals.trabajadorid",
				"foreignField": "_id",
				"as": "personals.trabajador",
			},
		},
        bson.M{
            "$lookup": bson.M{
                "from": "tareas",
                "localField": "personals._id",
                "foreignField": "personalid",
                "as": "tareas",
            },
        },
        bson.M{
			"$unwind": bson.M{
				"path":"$personals.trabajador",
				"preserveNullAndEmptyArrays":true,
			},
		},

        bson.M{
			"$unwind": bson.M{
				"path":"$tareas",
				"preserveNullAndEmptyArrays":true,
			},
		},

        bson.M{
            "$project": bson.M{
                "_id": 1,
                "codigo": 1,
                "nombre": 1,
                "entityid": 1,
                "trabajadorid": 1,
                "descripcion": 1,
                "status": 1,
                "datestart": 1,
                "dateend": 1,
                "dateendfake": 1,
                "documents": 1,
                "personals": 1,
                "tareas.trabajador": "$personals.trabajador.nombre",
                "tareas.nombre": 1,
                "tareas.datestart": 1,
                "tareas.dateend": 1,
                "tareas._id": 1,
                "tareas.personalid": 1,
            },
        },
        
        bson.M{
			"$group": bson.M{
				"_id":"$_id",
                "userid":bson.M{
                    "$first": "$userid",
                },
                "entityid":bson.M{
                    "$first": "$entityid",
                },
                "trabajadorid":bson.M{
                    "$first": "$trabajadorid",
                },
				"codigo":bson.M{
					"$first": "$codigo",
				},
				"nombre":bson.M{
					"$first": "$nombre",
				},
				"descripcion":bson.M{
					"$first": "$descripcion",
				},
				"datestart":bson.M{
					"$first": "$datestart",
				},
				"dateend":bson.M{
					"$first": "$dateend",
				},
				"dateendfake":bson.M{
					"$first": "$dateendfake",
				},
				"status":bson.M{
					"$first": "$status",
				},
				"documents":bson.M{
					"$last": "$documents",
				},
                "tareas": bson.M{
                    "$push": "$tareas",
                },
			},
		},

        bson.M{
            "$lookup" : bson.M{
                "from": "partidas",
                "localField": "_id",
                "foreignField": "proyectoid",
                "as": "partidas",
            },
        },

        bson.M{
			"$unwind": bson.M{
				"path":"$partidas",
				"preserveNullAndEmptyArrays":true,
			},
        },

        /*bson.M{
            "$unwind": "$partidas",
        },*/

        bson.M{
            "$lookup": bson.M{
                "from": "otros",
                "localField": "partidas._id",
                "foreignField": "partidaid",
                "as": "partidas.otros",
            },
        },

        bson.M{
			"$unwind": bson.M{
				"path":"$partidas.otros",
				"preserveNullAndEmptyArrays":true,
			},
        },

        /*bson.M{
            "$unwind": "$partidas.otros",
        },*/

        bson.M{
            "$lookup": bson.M{
                "from": "personals",
                "localField": "partidas.otros.personalid",
                "foreignField": "_id",
                "as": "partidas.otros.personals",
            },
        },

        bson.M{
			"$unwind": bson.M{
				"path":"$partidas.otros.personals",
				"preserveNullAndEmptyArrays":true,
			},
        },

        bson.M{
            "$lookup": bson.M{
                "from": "trabajadores",
                "localField": "partidas.otros.personals.trabajadorid",
                "foreignField": "_id",
                "as": "partidas.otros.personals.trabajador",
            },
        },

        bson.M{
			"$unwind": bson.M{
				"path":"$partidas.otros.personals.trabajador",
				"preserveNullAndEmptyArrays":true,
			},
        },

        bson.M{
            "$project": bson.M{
                "_id": 1,
                "codigo": 1,
                "nombre": 1,
                "descripcion": 1,
                "status": 1,
                "datestart": 1,
                "dateend": 1,
                "dateendfake": 1,
                "documents": 1,
                "personals": 1,
                "tareas": 1,
                "entityid": 1,
                "trabajadorid": 1,
                "otros.responsable": "$partidas.otros.personals.trabajador.nombre",
                "otros.partida": "$partidas.nombre",
                "otros._id": "$partidas.otros._id",
                "otros.partidaid": "$partidas.otros.partidaid",
                "otros.personalid": "$partidas.otros.personalid",
                "otros.dateregistro": "$partidas.otros.dateregistro",
                "otros.daterecordatorio": "$partidas.otros.daterecordatorio",
                "otros.descripcion": "$partidas.otros.descripcion",
            },
        },

		bson.M{
			"$group": bson.M{
				"_id":"$_id",
                "entityid":bson.M{
                    "$first": "$entityid",
                },
                "trabajadorid":bson.M{
                    "$first": "$trabajadorid",
                },
				"codigo":bson.M{
					"$first": "$codigo",
				},
				"nombre":bson.M{
					"$first": "$nombre",
				},
				"descripcion":bson.M{
					"$first": "$descripcion",
				},
				"datestart":bson.M{
					"$first": "$datestart",
				},
				"dateend":bson.M{
					"$first": "$dateend",
				},
				"dateendfake":bson.M{
					"$first": "$dateendfake",
				},
				"status":bson.M{
					"$first": "$status",
				},
				"documents":bson.M{
					"$last": "$documents",
				},
                "tareas": bson.M{
                    "$last": "$tareas",
                },
                "otros": bson.M{
                    "$push": "$otros",
                },
			},
		},

        
	}).Iter()
	//log.Println("iter : ",iter)
	result := models.ProyectoWithAll{}
	for iter.Next(&result){
		//log.Println("result: ",result)
		projects = append(projects, result)
	}
	return projects
	
}
