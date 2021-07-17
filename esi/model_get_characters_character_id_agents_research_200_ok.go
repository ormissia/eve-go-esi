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
type GetCharactersCharacterIdAgentsResearch200Ok struct {
	// agent_id integer
	AgentId int32 `json:"agent_id"`
	// points_per_day number
	PointsPerDay float32 `json:"points_per_day"`
	// remainder_points number
	RemainderPoints float32 `json:"remainder_points"`
	// skill_type_id integer
	SkillTypeId int32 `json:"skill_type_id"`
	// started_at string
	StartedAt time.Time `json:"started_at"`
}