package model

import "time"

type AllianceTech struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	TechType    TechType  `json:"tech_type"`
	Level       int       `json:"level" gorm:"default:1"`
	Bonus       int       `json:"bonus" gorm:"default:0"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:now()"`
	UpgradeWood int       `json:"upgrade_wood" gorm:"default:500"`
	UpgradeStone int      `json:"upgrade_stone" gorm:"default:400"`
	UpgradeIron int       `json:"upgrade_iron" gorm:"default:300"`
	UpgradeFood int       `json:"upgrade_food" gorm:"default:1000"`
	UpgradeGold int       `json:"upgrade_gold" gorm:"default:50"`
	UpgradeTime int       `json:"upgrade_time" gorm:"default:1500"`
	AllianceId  string    `json:"alliance_id"`
}

/*
设置表名(默认为 model 的复数形式)
*/
func (at *AllianceTech) TableName() string {
	return "alliance_tech"
}

// TechType 科技类型枚举
type TechType string

const (
	TechTypeAttackTech   TechType = "ATTACK_TECH"
	TechTypeDefenseTech  TechType = "DEFENSE_TECH"
	TechTypeHealthTech   TechType = "HEALTH_TECH"
	TechTypeScoutTech    TechType = "SCOUT_TECH"
	TechTypeSpeedTech    TechType = "SPEED_TECH"
	TechTypeCollectTech  TechType = "COLLECT_TECH"
	TechTypeCarriageTech TechType = "CARRIAGE_TECH"
) 