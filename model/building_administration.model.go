package model

import "time"

type BuildingAdministration struct {
	ID                      string    `json:"id" gorm:"primaryKey"`
	ResourcePointsCapacity  int       `json:"resource_points_capacity" gorm:"default:1"`
	CreatedAt               time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt               time.Time `json:"updated_at" gorm:"default:now()"`
	BuildingId              string    `json:"building_id" gorm:"unique"`
	UserId                  string    `json:"user_id"`
	MapElementId            string    `json:"map_element_id" gorm:"unique"`
	CityId                  *string   `json:"city_id" gorm:"unique"`
}

/*
设置表名(默认为 model 的复数形式)
*/
func (ba *BuildingAdministration) TableName() string {
	return "building_administration"
} 