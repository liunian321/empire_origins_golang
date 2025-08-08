package model

import "time"

type City struct {
	ID string `json:"id" gorm:"primaryKey"`
	// 城市名称
	Name string `json:"name"`
	// 人口数量
	Population   int       `json:"population" gorm:"default:0"`
	OwnerId      string    `json:"owner_id"`
	MapElementId string    `json:"map_element_id" gorm:"unique"`
	CreatedAt    time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"default:now()"`
	// 粮食
	Food int `json:"food" gorm:"default:0"`
	// 木材
	Wood int `json:"wood" gorm:"default:0"`
	// 石材
	Stone int `json:"stone" gorm:"default:0"`
	// 黄金
	Gold int `json:"gold" gorm:"default:0"`
	// 铁矿
	Iron int `json:"iron" gorm:"default:0"`
	// 城市规模
	Scale int `json:"scale" gorm:"default:0"`
	// 人口更新时间
	PopulationUpdateAt *time.Time `json:"population_update_at"`
	// 黄金更新时间
	GoldUpdateAt *time.Time `json:"gold_update_at"`
	// 资源更新时间
	ResourceUpdateAt *time.Time `json:"resource_update_at"`
}

/*
设置表名(默认为 model 的复数形式)
*/
func (c *City) TableName() string {
	return "city"
}
