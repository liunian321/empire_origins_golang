package dto

import "github.com/golang-jwt/jwt/v5"

type LoginUser struct {
	Username string `json:"username" binding:"required,min=1,max=16"`
	Password string `json:"password" binding:"required,min=8,max=32"`
}

type AddUserResponse struct {
	UserId       string `json:"user_id"`
	AccessToken  string `json:"access_token"`
	MapElementId string `json:"map_element_id"`
}

type UserInfo struct {
	ID           string `json:"id"`
	MapElementId string `json:"map_element_id"`
	jwt.RegisteredClaims
}
