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

    CategoriaResource struct{
        Data models.Categoria `json:"data"`
    }

    CategoriasResource struct{
        Data []models.Categoria `json:"data"`
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
	
    AuthUserResource struct {
        Data AuthUserModel `json:"data"`
    }
    
    LoginResource struct{
        Data LoginModel `json:"data"`
    }
    
    LoginModel struct {
        Email string `json:"email"`
        Password string `json:"password"`
    }
    
    AuthUserModel struct {
        User models.User `json:"user"`
        Token string `json:"token"`
        
    }
)
