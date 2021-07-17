/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.8.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// recipient object
type GetCharactersCharacterIdMailRecipient struct {
	// recipient_id integer
	RecipientId int32 `json:"recipient_id"`
	// recipient_type string
	RecipientType string `json:"recipient_type"`
}
