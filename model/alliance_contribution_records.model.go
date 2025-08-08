package model

import "time"

type AllianceContributionRecords struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	AllianceId string    `json:"alliance_id"`
	MemberId   string    `json:"member_id"`
	Wood       int       `json:"wood" gorm:"default:0"`
	Stone      int       `json:"stone" gorm:"default:0"`
	Iron       int       `json:"iron" gorm:"default:0"`
	Food       int       `json:"food" gorm:"default:0"`
	Gold       int       `json:"gold" gorm:"default:0"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"default:now()"`
}

/*
设置表名(默认为 model 的复数形式)
*/
func (acr *AllianceContributionRecords) TableName() string {
	return "alliance_contribution_records"
} 