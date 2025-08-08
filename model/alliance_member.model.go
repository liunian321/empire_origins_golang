package model

import "time"

type AllianceMember struct {
	ID        string            `json:"id" gorm:"primaryKey"`
	UserId    string            `json:"user_id" gorm:"unique"`
	AllianceId string           `json:"alliance_id"`
	Position  AlliancePosition  `json:"position" gorm:"default:MEMBER"`
	JoinedAt  time.Time         `json:"joined_at" gorm:"default:now()"`
	CreatedAt time.Time         `json:"created_at" gorm:"default:now()"`
	UpdatedAt time.Time         `json:"updated_at" gorm:"default:now()"`
}

/*
设置表名(默认为 model 的复数形式)
*/
func (am *AllianceMember) TableName() string {
	return "alliance_member"
}

// AlliancePosition 联盟职位枚举
type AlliancePosition string

const (
	AlliancePositionLeader     AlliancePosition = "LEADER"
	AlliancePositionViceLeader AlliancePosition = "VICE_LEADER"
	AlliancePositionMember     AlliancePosition = "MEMBER"
) 