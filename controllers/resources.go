package controllers

import(
    "../models"
)

type(
    
    UserResource struct {
        Data models.User `json:"data"`
    }

	UserProjectResource struct{
		Data []models.UserProjectModel `json:"data"`
	}

	ProyectoResource struct {
		Data models.Proyecto `json:"data"`
	}

	ProyectosResource struct{
		Data []models.Proyecto `json:"data"`
	}

	ProyectosResourceAll struct{
		Data []models.ProyectoWithAll `json:"data"`
	}

    DocumentoResource struct{
        Data models.Documento `json:"data"`
    }

    DocumentosResource struct{
        Data []models.Documento `json:"data"`
    }

	DocumentosResponsableResource struct{
		Data []models.DocumentoResponsable `json:"data"`
	}

    CategoriaResource struct{
        Data models.Categoria `json:"data"`
    }

    CategoriasResource struct{
        Data []models.Categoria `json:"data"`
    }

    PartidaResource struct{
        Data models.Partida `json:"data"`
    }

    PartidasResource struct{
        Data []models.Partida `json:"data"`
    }

	PartidasOtroResource struct{
        Data []models.PartidaOtro `json:"data"`
    }

    OtroResource struct{
        Data models.Otro `json:"data"`
    }

    OtrosResource struct{
        Data []models.Otro `json:"data"`
    }

	OtrosAllResource struct{
		Data []models.OtroAll `json:"data"`
	}

    TrabajadorResource struct{
        Data models.Trabajador `json:"data"`
    }

    TrabajadoresResource struct{
        Data []models.Trabajador `json:"data"`
    }

    PersonalResource struct{
        Data models.Personal `json:"data"`
    }

    PersonalsResource struct{
        Data []models.Personal `json:"data"`
    }

	CheckTareaResource struct{
		Data models.CheckTarea `json:"data"`
	}
	
    TareaResource struct {
		Data models.Tarea `json:"data"`
	}

	TareasResource struct{
		Data []models.Tarea `json:"data"`
	}

	TareasTrabajadorResource struct{
		Data []models.TareaTrabajador `json:"data"`
	}
	
    AuthUserResource struct {
        Data AuthUserModel `json:"data"`
    }
    
    LoginResource struct{
        Data LoginModel `json:"data"`
    }

	LoginTrabajadorResource struct{
        Data LoginTrabajador `json:"data"`
    }

	LoginTrabajador struct{
		Email string `json:"email"`
        Password string `json:"password"`
	}
	
    LoginModel struct {
        Email string `json:"email"`
        Password string `json:"password"`
		Empresa string `json:"empresa"`
		Nombre string `json:"nombre"`
    }
    
    AuthUserModel struct {
        User models.Trabajador `json:"user"`
        Token string `json:"token"`
        
    }
)
