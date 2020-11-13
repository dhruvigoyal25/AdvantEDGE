/*
 * ETSI GS MEC 028 - WLAN Access Information API
 *
 * The ETSI MEC ISG MEC028 WLAN Access Information API described using OpenAPI
 *
 * API version: 2.1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package server

type HeCapabilities struct {
	// MAC capabilities of an Access Point.
	HeMacCapInfo int32 `json:"heMacCapInfo"`
	// PHY capabilities of an Access Point.
	HePhyCapinfo int32 `json:"hePhyCapinfo"`
	// PPE Threshold determines the nominal packet padding value for a HE PPDU.
	PpeThresholds int32 `json:"ppeThresholds,omitempty"`
	// Supported MCS and NSS Set.
	SupportedHeMcsNssSet int32 `json:"supportedHeMcsNssSet"`
}
