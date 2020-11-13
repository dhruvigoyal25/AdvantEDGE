/*
 * ETSI GS MEC 028 - WLAN Access Information API
 *
 * The ETSI MEC ISG MEC028 WLAN Access Information API described using OpenAPI
 *
 * API version: 2.1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package server

type BssLoad struct {
	// Available Admission Capacity that specifies the remaining amount of medium time available via explicit admission control, in units of 32 s/s.
	AvailAdmCap int32 `json:"availAdmCap"`
	// The percentage of time, linearly scaled with 255 representing 100 %, that the AP sensed the medium was busy, as indicated by either the physical or virtual Carrier Sense (CS) mechanism.
	ChannelUtilization int32 `json:"channelUtilization"`
	// An unsigned integer that indicates the total number of STAs currently associated with this BSS.
	StaCount int32 `json:"staCount"`
}
