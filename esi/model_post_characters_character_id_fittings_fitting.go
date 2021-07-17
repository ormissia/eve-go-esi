/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.8.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// fitting object
type PostCharactersCharacterIdFittingsFitting struct {
	// description string
	Description string `json:"description"`
	// items array
	Items []PostCharactersCharacterIdFittingsItem `json:"items"`
	// name string
	Name string `json:"name"`
	// ship_type_id integer
	ShipTypeId int32 `json:"ship_type_id"`
}
