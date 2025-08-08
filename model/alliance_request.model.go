package model

import "time"

type AllianceRequest struct {
	ID         string                    `json:"id" gorm:"primaryKey"`
	UserId     string                    `json:"user_id"`
	AllianceId string                    `json:"alliance_id"`
	CreatedAt  time.Time                 `json:"created_at" gorm:"default:now()"`
	UpdatedAt  time.Time                 `json:"updated_at" gorm:"default:now()"`
	Status     AllianceRequestStatus     `json:"status" gorm:"default:PENDING"`
	HandlerId  *string                   `json:"handler_id"`
}

/*
设置表名(默认为 model 的复数形式)
*/
func (ar *AllianceRequest) TableName() string {
	return "alliance_request"
}

// AllianceRequestStatus 联盟请求状态枚举
type AllianceRequestStatus string

const (
	AllianceRequestStatusPending  AllianceRequestStatus = "PENDING"
	AllianceRequestStatusAccepted AllianceRequestStatus = "ACCEPTED"
	AllianceRequestStatusRejected AllianceRequestStatus = "REJECTED"
) 