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
type GetCharactersCharacterIdBookmarks200Ok struct {
	// bookmark_id integer
	BookmarkId  int32                                         `json:"bookmark_id"`
	Coordinates *GetCharactersCharacterIdBookmarksCoordinates `json:"coordinates,omitempty"`
	// created string
	Created time.Time `json:"created"`
	// creator_id integer
	CreatorId int32 `json:"creator_id"`
	// folder_id integer
	FolderId int32                                  `json:"folder_id,omitempty"`
	Item     *GetCharactersCharacterIdBookmarksItem `json:"item,omitempty"`
	// label string
	Label string `json:"label"`
	// location_id integer
	LocationId int32 `json:"location_id"`
	// notes string
	Notes string `json:"notes"`
}
