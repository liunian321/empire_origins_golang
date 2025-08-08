package model

import "time"

type MilitaryAction struct {
	ID                  string      `json:"id" gorm:"primaryKey"`
	Type                LetterType  `json:"type"`
	RemainTime          int         `json:"remain_time"`
	CreatedAt           time.Time   `json:"created_at" gorm:"default:now()"`
	UpdatedAt           time.Time   `json:"updated_at" gorm:"default:now()"`
	OccupyTime          *time.Time  `json:"occupy_time"`
	AttackerElementId   string      `json:"attacker_element_id"`
	DefenderElementId   string      `json:"defender_element_id"`
	AttackerId          string      `json:"attacker_id"`
	DefenderId          *string     `json:"defender_id"`
}

/*
设置表名(默认为 model 的复数形式)
*/
func (ma *MilitaryAction) TableName() string {
	return "military_action"
} 