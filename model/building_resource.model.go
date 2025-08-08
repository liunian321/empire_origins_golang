package model

import "time"

type BuildingResource struct {
	ID          string        `json:"id" gorm:"primaryKey"`
	Type        MaterialType  `json:"type"`
	Level       int           `json:"level" gorm:"default:1"`
	Yield       int           `json:"yield" gorm:"default:500"`
	Scale       int           `json:"scale"`
	Wood        int           `json:"wood"`
	Iron        int           `json:"iron"`
	Stone       int           `json:"stone"`
	Food        int           `json:"food"`
	Gold        int           `json:"gold" gorm:"default:0"`
	UpgradeTime int           `json:"upgrade_time"`
	CreatedAt   time.Time     `json:"created_at" gorm:"default:now()"`
	UpdatedAt   time.Time     `json:"updated_at" gorm:"default:now()"`
	UserId      string        `json:"user_id"`
	CityId      string        `json:"city_id"`
}

/*
设置表名(默认为 model 的复数形式)
*/
func (br *BuildingResource) TableName() string {
	return "building_resource"
}

// MaterialType 资源类型枚举
type MaterialType string

const (
	MaterialTypeWood  MaterialType = "WOOD"
	MaterialTypeStone MaterialType = "STONE"
	MaterialTypeIron  MaterialType = "IRON"
	MaterialTypeFood  MaterialType = "FOOD"
	MaterialTypeGold  MaterialType = "GOLD"
) 