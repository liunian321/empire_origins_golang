package model

import "time"

type MapElement struct {
	ID               string    `json:"id" gorm:"primaryKey"`
	Category         string    `json:"category"`
	Level            int       `json:"level" gorm:"default:1"`
	Experience       int       `json:"experience" gorm:"default:0"`
	X                int       `json:"x"`
	Y                int       `json:"y"`
	OwnerId          *string   `json:"owner_id"`
	Resource         *int      `json:"resource"`
	CreatedAt        time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"default:now()"`
	CityId           *string   `json:"city_id"`
	AdministrationId *string   `json:"administration_id"`
}

/*
设置表名(默认为 model 的复数形式)
*/
func (m *MapElement) TableName() string {
	return "map_element"
}
