/*
 * ETSI GS MEC 028 - WLAN Access Information API
 *
 * The ETSI MEC ISG MEC028 WLAN Access Information API described using OpenAPI
 *
 * API version: 2.1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package server

type GeoLocation struct {
	// The altitude value of location as defined in IETF RFC 6225 [6]
	Altitude int32 `json:"altitude,omitempty"`
	// The type description for altitude information e.g. floors or meters as defined in IETF RFC 6225 [6]
	AltitudeType int32 `json:"altitudeType,omitempty"`
	// The uncertainty for altitude information as defined in IETF RFC 6225 [6]
	AltitudeUncertainty int32 `json:"altitudeUncertainty,omitempty"`
	// The datum value to express how coordinates are organized and related to real world as defined in IETF RFC 6225 [6]
	Datum int32 `json:"datum"`
	// The latitude value of location as defined in IETF RFC 6225 [6]
	Lat int64 `json:"lat"`
	// The uncertainty for Latitude information as defined in IETF RFC 6225 [6]
	LatUncertainty int32 `json:"latUncertainty"`
	// The longitude value of location as defined in IETF RFC 6225 [6]
	Long int64 `json:"long"`
	// The uncertainty for Longitude information as defined in IETF RFC 6225 [6]
	LongUncertainty int32 `json:"longUncertainty"`
}
