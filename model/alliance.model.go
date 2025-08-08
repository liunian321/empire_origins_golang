package model

import "time"

type Alliance struct {
	ID                   string    `json:"id" gorm:"primaryKey"`
	Name                 string    `json:"name" gorm:"unique"`
	Level                int       `json:"level" gorm:"default:1"`
	MemberLimit          int       `json:"member_limit" gorm:"default:50"`
	CurrentMembers       int       `json:"current_members" gorm:"default:1"`
	Description          *string   `json:"description"`
	CreatedAt            time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt            time.Time `json:"updated_at" gorm:"default:now()"`
	TroopsRequest        int       `json:"troops_request" gorm:"default:0"`
	BlockRequest         bool      `json:"block_request" gorm:"default:false"`
	BlockRequestTime     *int      `json:"block_request_time"`
	AllianceResourceId   *string   `json:"alliance_resource_id"`
	UpgradeWood          int       `json:"upgrade_wood" gorm:"default:1000"`
	UpgradeStone         int       `json:"upgrade_stone" gorm:"default:800"`
	UpgradeIron          int       `json:"upgrade_iron" gorm:"default:500"`
	UpgradeFood          int       `json:"upgrade_food" gorm:"default:2000"`
	UpgradeGold          int       `json:"upgrade_gold" gorm:"default:100"`
	UpgradeTime          int       `json:"upgrade_time" gorm:"default:7200"`
}

/*
设置表名(默认为 model 的复数形式)
*/
func (a *Alliance) TableName() string {
	return "alliance"
} 