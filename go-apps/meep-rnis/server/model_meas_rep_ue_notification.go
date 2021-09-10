/*
 * Copyright (c) 2020  InterDigital Communications, Inc
 *
 * Licensed under the Apache License, Version 2.0 (the \"License\");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an \"AS IS\" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * AdvantEDGE Radio Network Information Service REST API
 *
 * Radio Network Information Service is AdvantEDGE's implementation of [ETSI MEC ISG MEC012 RNI API](http://www.etsi.org/deliver/etsi_gs/MEC/001_099/012/02.01.01_60/gs_MEC012v020101p.pdf) <p>[Copyright (c) ETSI 2017](https://forge.etsi.org/etsi-forge-copyright-notice.txt) <p>**Micro-service**<br>[meep-rnis](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-rnis) <p>**Type & Usage**<br>Edge Service used by edge applications that want to get information about radio conditions in the network <p>**Note**<br>AdvantEDGE supports a selected subset of RNI API endpoints (see below) and a subset of subscription types. <p>Supported subscriptions: <p> - CellChangeSubscription <p> - RabEstSubscription <p> - RabRelSubscription <p> - MeasRepUeSubscription <p> - NrMeasRepUeSubscription
 *
 * API version: 2.1.1
 * Contact: AdvantEDGE@InterDigital.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package server

type MeasRepUeNotification struct {
	// 0 to N identifiers to associate the event for a specific UE or flow.
	AssociateId []AssociateId `json:"associateId,omitempty"`
	// This parameter can be repeated to contain information of all the carriers assign for Carrier Aggregation up to M.
	CarrierAggregationMeasInfo []MeasRepUeNotificationCarrierAggregationMeasInfo `json:"carrierAggregationMeasInfo,omitempty"`

	Ecgi *Ecgi `json:"ecgi"`
	// This parameter can be repeated to contain information of all the neighbouring cells up to N.
	EutranNeighbourCellMeasInfo []MeasRepUeNotificationEutranNeighbourCellMeasInfo `json:"eutranNeighbourCellMeasInfo,omitempty"`
	// Indicates height of the UE in meters relative to the sea level as defined in ETSI TS 136.331 [i.7].
	HeightUe int32 `json:"heightUe,omitempty"`
	// 5G New Radio secondary serving cells measurement information.
	NewRadioMeasInfo []MeasRepUeNotificationNewRadioMeasInfo `json:"newRadioMeasInfo,omitempty"`
	// Measurement quantities concerning the 5G NR neighbours.
	NewRadioMeasNeiInfo []MeasRepUeNotificationNewRadioMeasNeiInfo `json:"newRadioMeasNeiInfo,omitempty"`
	// Shall be set to \"MeasRepUeNotification\".
	NotificationType string `json:"notificationType"`
	// Reference Signal Received Power as defined in ETSI TS 136 214 [i.5].
	Rsrp int32 `json:"rsrp"`
	// Extended Reference Signal Received Power, with value mapping defined in ETSI TS 136 133 [i.16].
	RsrpEx int32 `json:"rsrpEx,omitempty"`
	// Reference Signal Received Quality as defined in ETSI TS 136 214 [i.5].
	Rsrq int32 `json:"rsrq"`
	// Extended Reference Signal Received Quality, with value mapping defined in ETSI TS 136 133 [i.16].
	RsrqEx int32 `json:"rsrqEx,omitempty"`
	// Reference Signal \"Signal to Interference plus Noise Ratio\", with value mapping defined in ETSI TS 136 133 [i.16].
	Sinr int32 `json:"sinr,omitempty"`

	TimeStamp *TimeStamp `json:"timeStamp,omitempty"`

	Trigger *Trigger `json:"trigger"`
}
