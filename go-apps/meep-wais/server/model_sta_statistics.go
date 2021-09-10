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
 * AdvantEDGE WLAN Access Information API
 *
 * WLAN Access Information Service is AdvantEDGE's implementation of [ETSI MEC ISG MEC028 WAI API](http://www.etsi.org/deliver/etsi_gs/MEC/001_099/028/02.02.01_60/gs_MEC028v020201p.pdf) <p>[Copyright (c) ETSI 2020](https://forge.etsi.org/etsi-forge-copyright-notice.txt) <p>**Micro-service**<br>[meep-wais](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-wais) <p>**Type & Usage**<br>Edge Service used by edge applications that want to get information about WLAN access information in the network <p>**Note**<br>AdvantEDGE supports a selected subset of WAI API subscription types. <p>Supported subscriptions: <p> - AssocStaSubscription <p> - StaDataRateSubscription
 *
 * API version: 2.2.1
 * Contact: AdvantEDGE@InterDigital.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package server

type StaStatistics struct {
	Group2to9Data *StaStatisticsGroup2to9Data `json:"group2to9Data,omitempty"`
	// Indicates the requested statistics group describing the Statistics Group Data according to Table 9-114 of IEEE 802.11-2016 [8]. Depending on group identity, one and only one of the STA Statistics Group Data will be present.
	GroupIdentity int32 `json:"groupIdentity"`

	GroupOneData *StaStatisticsGroupOneData `json:"groupOneData,omitempty"`

	GroupZeroData *StaStatisticsGroupZeroData `json:"groupZeroData,omitempty"`
	// Duration over which the Statistics Group Data was measured in time units of 1024 µs.   Duration equal to zero indicates a report of current values.
	MeasurementDuration int32 `json:"measurementDuration"`
	// Measurement ID of the Measurement configuration applied to this STA Statistics Report.
	MeasurementId string `json:"measurementId"`

	StaId *StaIdentity `json:"staId,omitempty"`
}
