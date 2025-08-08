package model

import "time"

type Relation struct {
	ID         string       `json:"id" gorm:"primaryKey"`
	FromUserId string       `json:"from_user_id"`
	ToUserId   string       `json:"to_user_id"`
	Relation   RelationType `json:"relation"`
	CreatedAt  time.Time    `json:"created_at" gorm:"default:now()"`
	UpdatedAt  time.Time    `json:"updated_at" gorm:"default:now()"`
}

/*
设置表名(默认为 model 的复数形式)
*/
func (r *Relation) TableName() string {
	return "relation"
}

// RelationType 好友关系类型枚举
type RelationType string

const (
	RelationTypeFriend  RelationType = "FRIEND"
	RelationTypeEnemy   RelationType = "ENEMY"
	RelationTypeAlly    RelationType = "ALLY"
	RelationTypeNeutral RelationType = "NEUTRAL"
) 