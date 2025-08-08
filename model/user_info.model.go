package model

import "github.com/golang-jwt/jwt/v5"

type UserInfo struct {
	ID           string `json:"id"`
	MapElementId string `json:"map_element_id"`
	jwt.RegisteredClaims
}
