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
type GetStatusOk struct {
	// Current online player count
	Players int32 `json:"players"`
	// Running version as string
	ServerVersion string `json:"server_version"`
	// Server start timestamp
	StartTime time.Time `json:"start_time"`
	// If the server is in VIP mode
	Vip bool `json:"vip,omitempty"`
}
