package model

import "time"

type Building struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Type        string    `json:"type"`
	Level       int       `json:"level" gorm:"default:1"`
	CityId      string    `json:"city_id"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:now()"`
	Wood        int       `json:"wood" gorm:"default:0"`
	Stone       int       `json:"stone" gorm:"default:0"`
	Iron        int       `json:"iron" gorm:"default:0"`
	Food        int       `json:"food" gorm:"default:0"`
	Gold        int       `json:"gold" gorm:"default:0"`
	Scale       int       `json:"scale" gorm:"default:0"`
	UpgradeTime int       `json:"upgrade_time" gorm:"default:0"`
}

/*
设置表名(默认为 model 的复数形式)
*/
func (b *Building) TableName() string {
	return "building"
}

// BuildingType 建筑类型枚举
type BuildingType string

const (
	BuildingTypeBarracks      BuildingType = "BARRACKS"
	BuildingTypeStable        BuildingType = "STABLE"
	BuildingTypeFactory       BuildingType = "FACTORY"
	BuildingTypeWarehouse     BuildingType = "WAREHOUSE"
	BuildingTypeCityWall      BuildingType = "CITY_WALL"
	BuildingTypeResidential   BuildingType = "RESIDENTIAL"
	BuildingTypeCommercial    BuildingType = "COMMERCIAL"
	BuildingTypeHarem         BuildingType = "HAREM"
	BuildingTypeAdministration BuildingType = "ADMINISTRATION"
	BuildingTypePalace        BuildingType = "PALACE"
) 