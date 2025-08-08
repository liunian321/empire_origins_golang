package model

import "time"

type Woman struct {
	ID             string         `json:"id" gorm:"primaryKey"`
	Name           string         `json:"name"`
	Star           int            `json:"star" gorm:"default:0"`
	Charm          int            `json:"charm"`
	Figure         int            `json:"figure"`
	Height         int            `json:"height"`
	Intelligence   int            `json:"intelligence"`
	Appearance     int            `json:"appearance"`
	Morality       int            `json:"morality"`
	Affection      int            `json:"affection" gorm:"default:0"`
	Personality    Personality    `json:"personality"`
	Position       *WomanPosition `json:"position"`
	AvatarUrl      *string        `json:"avatar_url"`
	CreatedAt      time.Time      `json:"created_at" gorm:"default:now()"`
	UpdatedAt      time.Time      `json:"updated_at" gorm:"default:now()"`
	IsOwned        bool           `json:"is_owned" gorm:"default:false"`
	BuildingHaremId string        `json:"building_harem_id"`
	OwnerId        string         `json:"owner_id"`
}

/*
设置表名(默认为 model 的复数形式)
*/
func (w *Woman) TableName() string {
	return "woman"
}

// Personality 性格枚举
type Personality string

const (
	PersonalitySweet  Personality = "SWEET"
	PersonalitySexy   Personality = "SEXY"
	PersonalityCute   Personality = "CUTE"
	PersonalityCold   Personality = "COLD"
	PersonalityStrong Personality = "STRONG"
	PersonalityGentle Personality = "GENTLE"
)

// WomanPosition 女人职位枚举
type WomanPosition string

const (
	WomanPositionOfficer WomanPosition = "OFFICER"
) 