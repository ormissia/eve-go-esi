/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.8.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// 200 ok object
type GetCorporationsCorporationIdFacilities200Ok struct {
	// facility_id integer
	FacilityId int64 `json:"facility_id"`
	// system_id integer
	SystemId int32 `json:"system_id"`
	// type_id integer
	TypeId int32 `json:"type_id"`
}