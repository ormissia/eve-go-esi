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
type GetCharactersCharacterIdContacts200Ok struct {
	// contact_id integer
	ContactId int32 `json:"contact_id"`
	// contact_type string
	ContactType string `json:"contact_type"`
	// Whether this contact is in the blocked list. Note a missing value denotes unknown, not true or false
	IsBlocked bool `json:"is_blocked,omitempty"`
	// Whether this contact is being watched
	IsWatched bool `json:"is_watched,omitempty"`
	// label_ids array
	LabelIds []int64 `json:"label_ids,omitempty"`
	// Standing of the contact
	Standing float32 `json:"standing"`
}