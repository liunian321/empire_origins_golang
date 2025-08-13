package dto

import "empire_origins_golang/model"

type CityInfoArgs struct {
	Point *Point `json:"point"`
}

type Point struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
}

type CityInfoResponse struct {
	City              model.City `json:"city"`
	SoldierCount      int32      `json:"soldier_count"`
	WarehouseCapacity int32      `json:"warehouse_capacity"`
}
