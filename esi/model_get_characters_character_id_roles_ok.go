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
type GetCharactersCharacterIdRolesOk struct {
	// roles array
	Roles []string `json:"roles,omitempty"`
	// roles_at_base array
	RolesAtBase []string `json:"roles_at_base,omitempty"`
	// roles_at_hq array
	RolesAtHq []string `json:"roles_at_hq,omitempty"`
	// roles_at_other array
	RolesAtOther []string `json:"roles_at_other,omitempty"`
}