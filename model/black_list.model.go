package model

import "time"

type BlackList struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	IP        string    `json:"ip"`
	UserId    string    `json:"user_id" gorm:"unique"`
	Reason    string    `json:"reason"`
	CreatedAt time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:now()"`
}

/*
设置表名(默认为 model 的复数形式)
*/
func (bl *BlackList) TableName() string {
	return "black_list"
} 