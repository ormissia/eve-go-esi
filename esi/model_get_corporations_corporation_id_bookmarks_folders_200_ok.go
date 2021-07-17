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
type GetCorporationsCorporationIdBookmarksFolders200Ok struct {
	// creator_id integer
	CreatorId int32 `json:"creator_id,omitempty"`
	// folder_id integer
	FolderId int32 `json:"folder_id"`
	// name string
	Name string `json:"name"`
}
