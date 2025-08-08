package model

import (
	"time"
)

type User struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique"`
	Password  string    `json:"password"`
	Email     string    `json:"email" gorm:"unique"`
	Order     uint      `json:"order"`
	IsEnabled bool      `json:"is_enabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

/*
设置表名(默认为 model 的复数形式)
*/
func (u *User) TableName() string {
	return "user"
}

type LoginUser struct {
	Username string `json:"username" binding:"required,min=1,max=16"`
	Password string `json:"password" binding:"required,min=8,max=32"`
}

type AddUserResponse struct {
	UserId       string `json:"user_id"`
	AccessToken  string `json:"access_token"`
	MapElementId string `json:"map_element_id"`
}