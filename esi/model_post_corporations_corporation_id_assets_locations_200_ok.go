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
type PostCorporationsCorporationIdAssetsLocations200Ok struct {
	// item_id integer
	ItemId   int64                                                 `json:"item_id"`
	Position *PostCorporationsCorporationIdAssetsLocationsPosition `json:"position"`
}
