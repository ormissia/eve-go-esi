/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.8.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Top 4 rankings of factions by victory points from yesterday, last week and in total
type GetFwLeaderboardsVictoryPoints struct {
	// Top 4 ranking of factions active in faction warfare by total victory points. A faction is considered \"active\" if they have participated in faction warfare in the past 14 days
	ActiveTotal []GetFwLeaderboardsActiveTotalActiveTotal1 `json:"active_total"`
	// Top 4 ranking of factions by victory points in the past week
	LastWeek []GetFwLeaderboardsLastWeekLastWeek1 `json:"last_week"`
	// Top 4 ranking of factions by victory points in the past day
	Yesterday []GetFwLeaderboardsYesterdayYesterday1 `json:"yesterday"`
}
