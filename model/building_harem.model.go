package model

import "time"

type BuildingHarem struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	Capacity   int       `json:"capacity" gorm:"default:10"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"default:now()"`
	BuildingId string    `json:"building_id" gorm:"unique"`
	UserId     string    `json:"user_id"`
	CityId     *string   `json:"city_id" gorm:"unique"`
}

/*
设置表名(默认为 model 的复数形式)
*/
func (bh *BuildingHarem) TableName() string {
	return "building_harem"
}
