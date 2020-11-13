/*
 * ETSI GS MEC 028 - WLAN Access Information API
 *
 * The ETSI MEC ISG MEC028 WLAN Access Information API described using OpenAPI
 *
 * API version: 2.1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package server

type ApLocation struct {
	CivicLocation *CivicLocation `json:"civicLocation,omitempty"`

	Geolocation *GeoLocation `json:"geolocation,omitempty"`
}
