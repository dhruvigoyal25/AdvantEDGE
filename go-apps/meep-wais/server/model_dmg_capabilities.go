/*
 * ETSI GS MEC 028 - WLAN Access Information API
 *
 * The ETSI MEC ISG MEC028 WLAN Access Information API described using OpenAPI
 *
 * API version: 2.1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package server

type DmgCapabilities struct {
	// Extended SC MCS capabilities as defined in draft IEEE P802.11/D4.0 [i.11]
	ExtScMcsCap int32 `json:"ExtScMcsCap"`
	// DMG AP or PCP capabilities information as defined in draft IEEE P802.11/D4.0 [i.11]
	DmgApOrPcpCapInfo int32 `json:"dmgApOrPcpCapInfo"`
	// DMG station beam tracking time limit as defined in draft IEEE P802.11/D4.0 [i.11]
	DmgStaBeamTrackTimeLimit int32 `json:"dmgStaBeamTrackTimeLimit"`
	// DMG station capabilities information as defined in draft IEEE P802.11/D4.0 [i.11]
	DmgStaCapInfo int64 `json:"dmgStaCapInfo"`
	// Number of basic A-MSDU subframes in A-MSDU as defined in draft IEEE P802.11/D4.0 [i.11]
	MaxNrBasicAmsduSubframes int32 `json:"maxNrBasicAmsduSubframes"`
	// Number of short A-MSDU subframes in A-MSDU as defined in draft IEEE P802.11/D4.0 [i.11]
	MaxNrShortAmsduSubframes int32 `json:"maxNrShortAmsduSubframes"`
	// SAR capabilities as defined in draft IEEE P802.11/D4.0 [i.11]
	SarCap int32 `json:"sarCap"`
	// TDD capabilities as defined in draft IEEE P802.11/D4.0 [i.11]
	TddCap int32 `json:"tddCap"`
}
