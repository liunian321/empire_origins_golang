package model

import "time"

type Soldier struct {
	ID              string        `json:"id" gorm:"primaryKey"`
	Level           int           `json:"level" gorm:"default:1"`
	Experience      int           `json:"experience" gorm:"default:0"`
	Quantity        int           `json:"quantity" gorm:"default:0"`
	Status          SoldierStatus `json:"status" gorm:"default:IN_CITY"`
	OwnerId         *string       `json:"owner_id"`
	CityId          *string       `json:"city_id"`
	CreatedAt       time.Time     `json:"created_at" gorm:"default:now()"`
	UpdatedAt       time.Time     `json:"updated_at" gorm:"default:now()"`
	MapElementId    *string       `json:"map_element_id"`
	SoldierTypeId   *int          `json:"soldier_type_id"`
	MilitaryActionId *string      `json:"military_action_id"`
}

/*
设置表名(默认为 model 的复数形式)
*/
func (s *Soldier) TableName() string {
	return "soldier"
}

// SoldierStatus 士兵状态枚举
type SoldierStatus string

const (
	SoldierStatusInCity     SoldierStatus = "IN_CITY"
	SoldierStatusStationed  SoldierStatus = "STATIONED"
	SoldierStatusInBattle   SoldierStatus = "IN_BATTLE"
	SoldierStatusReturnCity SoldierStatus = "RETURN_CITY"
	SoldierStatusInMarch    SoldierStatus = "IN_MARCH"
	SoldierStatusReinforce  SoldierStatus = "REINFORCE"
) 