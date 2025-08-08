package model

import "time"

type BuildingCommercial struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	Prosperity int       `json:"prosperity" gorm:"default:20"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"default:now()"`
	BuildingId string    `json:"building_id" gorm:"unique"`
	UserId     string    `json:"user_id"`
	CityId     *string   `json:"city_id" gorm:"unique"`
}

/*
设置表名(默认为 model 的复数形式)
*/
func (bc *BuildingCommercial) TableName() string {
	return "building_commercial"
} 