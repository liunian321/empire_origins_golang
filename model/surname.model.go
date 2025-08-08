package model

import "time"

type Surname struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"unique"`
	CreatedAt time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:now()"`
}

/*
设置表名(默认为 model 的复数形式)
*/
func (s *Surname) TableName() string {
	return "surname"
} 