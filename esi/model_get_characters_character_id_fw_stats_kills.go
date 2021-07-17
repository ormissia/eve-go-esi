/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.8.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Summary of kills done by the given character against enemy factions
type GetCharactersCharacterIdFwStatsKills struct {
	// Last week's total number of kills by a given character against enemy factions
	LastWeek int32 `json:"last_week"`
	// Total number of kills by a given character against enemy factions since the character enlisted
	Total int32 `json:"total"`
	// Yesterday's total number of kills by a given character against enemy factions
	Yesterday int32 `json:"yesterday"`
}
