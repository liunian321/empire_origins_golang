package city

import (
	"empire_origins_golang/db"
	"empire_origins_golang/dto"
	"empire_origins_golang/model"
	"strings"
	"time"

	"empire_origins_golang/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CityInfo(c *gin.Context) {
	userInfo, exists := c.Get("userInfo")
	if !exists {
		c.JSON(401, gin.H{"message": "Unauthorized"})
		return
	}
	args := dto.CityInfoArgs{}
	if err := c.ShouldBindJSON(&args); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	selectFields := []string{"id", "name", "map_element.x", "map_element.y"}
	cityInterface := gorm.G[model.City](db.DB)
	// 如果传入的是坐标,则不查询过多信息
	cityInterface.Joins(clause.InnerJoin.AssociationFrom("MapElement", gorm.G[model.MapElement](db.DB).Select("x", "y")).As("map_element"), func(db gorm.JoinBuilder, joinTable, curTable clause.Table) error {
		if args.Point != nil {
			db.Where("map_element.x = ? AND map_element.y = ?", args.Point.X, args.Point.Y)
		} else {
			selectFields = append(selectFields, "population", "scale", "wood", "iron", "stone", "gold", "food")
		}
		return nil
	})

	userInfoData := userInfo.(dto.UserInfo)

	if userInfoData.MapElementId != "" {
		cityInterface.Where("map_element_id = ?", userInfoData.MapElementId)
	}
	cityInterface.Select(strings.Join(selectFields, ","))
	city, error := cityInterface.First(c)
	if error != nil {
		c.JSON(500, gin.H{"message": "Failed to get city"})
		return
	}

	// 获取士兵数量
	soldiers, _ := gorm.G[model.Soldier](db.DB).Select("quantity").Where("map_element_id = ?", userInfoData.MapElementId).Find(c)

	// 获取仓库容量
	warehouse, error := gorm.G[model.BuildingWarehouse](db.DB).Select("capacity").Where("city_id = ?", city.ID).Where("user_id = ?", userInfoData.ID).First(c)
	warehouseCapacity := int32(10000)
	if error == nil {
		warehouseCapacity = warehouse.Capacity * 1000
	}

	soldier_count := int32(0)
	for _, soldier := range soldiers {
		soldier_count += soldier.Quantity
	}

	c.JSON(200, dto.CityInfoResponse{
		City:              city,
		SoldierCount:      soldier_count,
		WarehouseCapacity: warehouseCapacity,
	})
}

func InitCity(c *gin.Context) {
	point, err := util.NewMapUtil().GetUniqueRandomCoordinates(c)
	if err != nil {
		c.JSON(500, gin.H{"message": "Failed to get unique random coordinates"})
		return
	}
	db.DB.Transaction(func(tx *gorm.DB) error {
		mapElement := model.MapElement{
			Category: "城市",
			X:        point.X,
			Y:        point.Y,
		}
		tx.Create(&mapElement)

		city := model.City{
			Name:               "test",
			MapElementID:       mapElement.ID,
			Population:         10000,
			Scale:              20,
			Food:               1000,
			Wood:               1000,
			Stone:              1000,
			Gold:               1000,
			Iron:               1000,
			ResourceUpdateAt:   time.Now(),
			PopulationUpdateAt: time.Now(),
			GoldUpdateAt:       time.Now(),
		}
		tx.Create(&city)

		// 更新地图元素的所有者
		mapElement.OwnerID = city.ID
		mapElement.CityID = city.ID
		tx.Save(&mapElement)

		// TODO:创建各种建筑

		return nil
	})
}
