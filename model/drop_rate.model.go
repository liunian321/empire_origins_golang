package model

import "time"

type DropRate struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Business  string    `json:"business"`
	Category  string    `json:"category"`
	Name      string    `json:"name"`
	DropRate  int       `json:"drop_rate"`
	CreatedAt time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:now()"`
}

/*
设置表名(默认为 model 的复数形式)
*/
func (dr *DropRate) TableName() string {
	return "drop_rate"
} 