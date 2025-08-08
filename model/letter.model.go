package model

import "time"

type Letter struct {
	ID         string      `json:"id" gorm:"primaryKey"`
	Type       LetterType  `json:"type"`
	Content    []byte      `json:"content"`
	CreatedAt  time.Time   `json:"created_at" gorm:"default:now()"`
	UpdatedAt  time.Time   `json:"updated_at" gorm:"default:now()"`
	ConsumerId *string     `json:"consumer_id"`
	SenderId   *string     `json:"sender_id"`
	OwnerId    string      `json:"owner_id"`
}

/*
设置表名(默认为 model 的复数形式)
*/
func (l *Letter) TableName() string {
	return "letter"
}

// LetterType 邮件类型枚举
type LetterType string

const (
	LetterTypeAttack        LetterType = "ATTACK"
	LetterTypeDefense       LetterType = "DEFENSE"
	LetterTypeReinforce     LetterType = "REINFORCE"
	LetterTypeScout         LetterType = "SCOUT"
	LetterTypeInMarch       LetterType = "IN_MARCH"
	LetterTypeReturnCity    LetterType = "RETURN_CITY"
	LetterTypeOccupy        LetterType = "OCCUPY"
	LetterTypeChat          LetterType = "CHAT"
	LetterTypeCollect       LetterType = "COLLECT"
	LetterTypeFriendShip    LetterType = "FRIEND_SHIP"
	LetterTypeAllianceRequest LetterType = "ALLIANCE_REQUEST"
	LetterTypeAllianceKick  LetterType = "ALLIANCE_KICK"
	LetterTypeAllianceLeave LetterType = "ALLIANCE_LEAVE"
) 