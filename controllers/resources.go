package controllers

import(
    "../models"
)

type(
    
    UserResource struct {
        Data models.User `json:"data"`
    }

    MusicResource struct{
        Data models.Music `json:"data"`
    }

    AlbumResource struct{
        Data models.Album `json:"data"`
    }

    MessageUploadResource struct{
        Data MessageUploadModel `json:"data"`
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

    MessageUploadModel struct{
        Message string `json:"message"`
        Url string `json:"url"`
    }
    
    AuthUserModel struct {
        User models.User `json:"user"`
        Token string `json:"token"`
        
    }
)
