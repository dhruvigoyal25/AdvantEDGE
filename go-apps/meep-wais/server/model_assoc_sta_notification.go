/*
 * ETSI GS MEC 028 - WLAN Access Information API
 *
 * The ETSI MEC ISG MEC028 WLAN Access Information API described using OpenAPI
 *
 * API version: 2.1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package server

type AssocStaNotification struct {
	ApId *ApIdentity `json:"apId"`
	// Shall be set to \"AssocStaNotification\".
	NotificationType string `json:"notificationType"`
	// Identifier(s) to uniquely specify the client station(s) associated.
	StaId []StaIdentity `json:"staId,omitempty"`

	TimeStamp *TimeStamp `json:"timeStamp,omitempty"`
}
