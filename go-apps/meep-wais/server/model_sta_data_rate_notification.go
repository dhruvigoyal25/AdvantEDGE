/*
 * ETSI GS MEC 028 - WLAN Access Information API
 *
 * The ETSI MEC ISG MEC028 WLAN Access Information API described using OpenAPI
 *
 * API version: 2.1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package server

type StaDataRateNotification struct {
	// Shall be set to \"StaDataRateNotification\".
	NotificationType string `json:"notificationType"`
	// Data rates of a client station.
	StaDataRate []StaDataRate `json:"staDataRate,omitempty"`

	TimeStamp *TimeStamp `json:"timeStamp,omitempty"`
}
