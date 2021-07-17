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
type GetUniverseGraphicsGraphicIdOk struct {
	// collision_file string
	CollisionFile string `json:"collision_file,omitempty"`
	// graphic_file string
	GraphicFile string `json:"graphic_file,omitempty"`
	// graphic_id integer
	GraphicId int32 `json:"graphic_id"`
	// icon_folder string
	IconFolder string `json:"icon_folder,omitempty"`
	// sof_dna string
	SofDna string `json:"sof_dna,omitempty"`
	// sof_fation_name string
	SofFationName string `json:"sof_fation_name,omitempty"`
	// sof_hull_name string
	SofHullName string `json:"sof_hull_name,omitempty"`
	// sof_race_name string
	SofRaceName string `json:"sof_race_name,omitempty"`
}
