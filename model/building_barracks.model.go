package model

import "time"

type BuildingBarracks struct {
	ID                      string    `json:"id" gorm:"primaryKey"`
	CanRecruitInfantry      bool      `json:"can_recruit_infantry" gorm:"default:true"`
	CanRecruitLightCavalry  bool      `json:"can_recruit_light_cavalry" gorm:"default:false"`
	CanRecruitHeavyCavalry  bool      `json:"can_recruit_heavy_cavalry" gorm:"default:false"`
	CanRecruitArcher        bool      `json:"can_recruit_archer" gorm:"default:false"`
	CanRecruitCrossbowmen   bool      `json:"can_recruit_crossbowmen" gorm:"default:false"`
	CanRecruitScout         bool      `json:"can_recruit_scout" gorm:"default:false"`
	CanRecruitCatapult      bool      `json:"can_recruit_catapult" gorm:"default:false"`
	CreatedAt               time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt               time.Time `json:"updated_at" gorm:"default:now()"`
	BuildingId              string    `json:"building_id" gorm:"unique"`
	UserId                  string    `json:"user_id"`
	CityId                  *string   `json:"city_id" gorm:"unique"`
}

/*
设置表名(默认为 model 的复数形式)
*/
func (bb *BuildingBarracks) TableName() string {
	return "building_barracks"
} 