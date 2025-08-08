package model

import "time"

type SoldierType struct {
	ID           int           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         SoldierTypeEnum `json:"name" gorm:"unique"`
	Attack       int           `json:"attack" gorm:"default:10"`
	Defense      int           `json:"defense" gorm:"default:10"`
	Health       int           `json:"health" gorm:"default:100"`
	Scout        int           `json:"scout" gorm:"default:5"`
	Speed        int           `json:"speed" gorm:"default:10"`
	Range        int           `json:"range" gorm:"default:1"`
	Capacity     int           `json:"capacity" gorm:"default:10"`
	Food         int           `json:"food" gorm:"default:50"`
	Wood         int           `json:"wood" gorm:"default:100"`
	Iron         int           `json:"iron" gorm:"default:50"`
	Stone        int           `json:"stone" gorm:"default:50"`
	Gold         int           `json:"gold" gorm:"default:0"`
	FoodConsume  int           `json:"food_consume" gorm:"default:1"`
	RecruitTime  int           `json:"recruit_time" gorm:"default:60"`
	CreatedAt    time.Time     `json:"created_at" gorm:"default:now()"`
	UpdatedAt    time.Time     `json:"updated_at" gorm:"default:now()"`
}

/*
设置表名(默认为 model 的复数形式)
*/
func (st *SoldierType) TableName() string {
	return "soldier_type"
}

// SoldierTypeEnum 士兵类型枚举
type SoldierTypeEnum string

const (
	SoldierTypeInfantry     SoldierTypeEnum = "INFANTRY"
	SoldierTypeLightCavalry SoldierTypeEnum = "LIGHT_CAVALRY"
	SoldierTypeHeavyCavalry SoldierTypeEnum = "HEAVY_CAVALRY"
	SoldierTypeArcher       SoldierTypeEnum = "ARCHER"
	SoldierTypeCrossbowman  SoldierTypeEnum = "CROSSBOWMAN"
	SoldierTypeScout        SoldierTypeEnum = "SCOUT"
	SoldierTypeCatapult     SoldierTypeEnum = "CATAPULT"
) 