/*
 * ETSI GS MEC 028 - WLAN Access Information API
 *
 * The ETSI MEC ISG MEC028 WLAN Access Information API described using OpenAPI
 *
 * API version: 2.1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package server

type WanMetrics struct {
	// 1-octet positive integer representing the current percentage loading of the downlink WAN connection, scaled linearly with 255 representing 100 %, as measured over an interval the duration of which is reported in Load Measurement Duration. In cases where the downlink load is unknown to the AP, the value is set to zero.
	DownlinkLoad int32 `json:"downlinkLoad"`
	// 4-octet positive integer whose value is an estimate of the WAN Backhaul link current downlink speed in kilobits per second.
	DownlinkSpeed int32 `json:"downlinkSpeed"`
	// The LMD (Load Measurement Duration) field is a 2-octet positive integer representing the duration over which the Downlink Load and Uplink Load have been measured, in tenths of a second. When the actual load measurement duration is greater than the maximum value, the maximum value will be reported. The value of the LMD field is set to 0 when neither the uplink nor downlink load can be computed. When the uplink and downlink loads are computed over different intervals, the maximum interval is reported.
	Lmd int32 `json:"lmd"`
	// 1-octet positive integer representing the current percentage loading of the uplink WAN connection, scaled linearly with 255 representing 100 %, as measured over an interval, the duration of which is reported in Load Measurement Duration. In cases where the uplink load is unknown to the AP, the value is set to zero.
	UplinkLoad int32 `json:"uplinkLoad"`
	// 4-octet positive integer whose value is an estimate of the WAN Backhaul link's current uplink speed in kilobits per second.
	UplinkSpeed int32 `json:"uplinkSpeed"`
	// Info about WAN link status, link symmetricity and capacity currently used.
	WanInfo int32 `json:"wanInfo"`
}
