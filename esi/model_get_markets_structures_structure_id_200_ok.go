/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.8.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"time"
)

// 200 ok object
type GetMarketsStructuresStructureId200Ok struct {
	// duration integer
	Duration int32 `json:"duration"`
	// is_buy_order boolean
	IsBuyOrder bool `json:"is_buy_order"`
	// issued string
	Issued time.Time `json:"issued"`
	// location_id integer
	LocationId int64 `json:"location_id"`
	// min_volume integer
	MinVolume int32 `json:"min_volume"`
	// order_id integer
	OrderId int64 `json:"order_id"`
	// price number
	Price float64 `json:"price"`
	// range string
	Range_ string `json:"range"`
	// type_id integer
	TypeId int32 `json:"type_id"`
	// volume_remain integer
	VolumeRemain int32 `json:"volume_remain"`
	// volume_total integer
	VolumeTotal int32 `json:"volume_total"`
}