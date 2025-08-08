package model

import "time"

type Officer struct {
	ID              string    `json:"id" gorm:"primaryKey"`
	Level           int       `json:"level" gorm:"default:1"`
	Experience      int       `json:"experience" gorm:"default:0"`
	Defense         int       `json:"defense" gorm:"default:1"`
	Attack          int       `json:"attack" gorm:"default:1"`
	Health          int       `json:"health" gorm:"default:1"`
	Scout           int       `json:"scout" gorm:"default:1"`
	Speed           int       `json:"speed" gorm:"default:1"`
	Salary          int       `json:"salary" gorm:"default:100"`
	IsEnlisted      bool      `json:"is_enlisted" gorm:"default:false"`
	CreatedAt       time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"default:now()"`
	WomanId         string    `json:"woman_id" gorm:"unique"`
	MapElementId    string    `json:"map_element_id"`
	MilitaryActionId *string  `json:"military_action_id"`
}

/*
设置表名(默认为 model 的复数形式)
*/
func (o *Officer) TableName() string {
	return "officer"
} 