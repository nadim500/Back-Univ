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
