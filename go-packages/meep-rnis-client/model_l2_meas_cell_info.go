/*
 * Copyright (c) 2022  InterDigital Communications, Inc
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
 * AdvantEDGE Radio Network Information API
 *
 * Radio Network Information Service is AdvantEDGE's implementation of [ETSI MEC ISG MEC012 RNI API](https://www.etsi.org/deliver/etsi_gs/MEC/001_099/012/02.02.01_60/gs_MEC012v020201p.pdf) <p>[Copyright (c) ETSI 2017](https://forge.etsi.org/etsi-forge-copyright-notice.txt) <p>**Micro-service**<br>[meep-rnis](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-rnis) <p>**Type & Usage**<br>Edge Service used by edge applications that want to get information about radio conditions in the network <p>**Note**<br>AdvantEDGE supports a selected subset of RNI API endpoints (see below) and a subset of subscription types. <p>Supported subscriptions: <p> - CellChangeSubscription <p> - RabEstSubscription <p> - RabRelSubscription <p> - MeasRepUeSubscription <p> - NrMeasRepUeSubscription
 *
 * API version: 2.2.1
 * Contact: AdvantEDGE@InterDigital.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type L2MeasCellInfo struct {
	// It indicates the packet discard rate in percentage of the downlink GBR traffic in a cell, as defined in ETSI TS 136 314 [i.11].
	DlGbrPdrCell int32 `json:"dl_gbr_pdr_cell,omitempty"`
	// It indicates the PRB usage for downlink GBR traffic, as defined in ETSI TS 136 314 [i.11] and ETSI TS 136 423 [i.12].
	DlGbrPrbUsageCell int32 `json:"dl_gbr_prb_usage_cell,omitempty"`
	// It indicates the packet discard rate in percentage of the downlink non-GBR traffic in a cell, as defined in ETSI TS 136 314 [i.11].
	DlNongbrPdrCell int32 `json:"dl_nongbr_pdr_cell,omitempty"`
	// It indicates (in percentage) the PRB usage for downlink non-GBR traffic, as defined in ETSI TS 136 314 [i.11] and ETSI TS 136 423 [i.12].
	DlNongbrPrbUsageCell int32 `json:"dl_nongbr_prb_usage_cell,omitempty"`
	// It indicates (in percentage) the PRB usage for total downlink traffic, as defined in ETSI TS 136 314 [i.11] and ETSI TS 136 423 [i.12].
	DlTotalPrbUsageCell int32 `json:"dl_total_prb_usage_cell,omitempty"`
	Ecgi                *Ecgi `json:"ecgi,omitempty"`
	// It indicates the number of active UEs with downlink GBR traffic, as defined in ETSI TS 136 314 [i.11].
	NumberOfActiveUeDlGbrCell int32 `json:"number_of_active_ue_dl_gbr_cell,omitempty"`
	// It indicates the number of active UEs with downlink non-GBR traffic, as defined in ETSI TS 136 314 [i.11].
	NumberOfActiveUeDlNongbrCell int32 `json:"number_of_active_ue_dl_nongbr_cell,omitempty"`
	// It indicates the number of active UEs with uplink GBR traffic, as defined in ETSI TS 136 314 [i.11].
	NumberOfActiveUeUlGbrCell int32 `json:"number_of_active_ue_ul_gbr_cell,omitempty"`
	// It indicates the number of active UEs with uplink non-GBR traffic, as defined in ETSI TS 136 314 [i.11].
	NumberOfActiveUeUlNongbrCell int32 `json:"number_of_active_ue_ul_nongbr_cell,omitempty"`
	// It indicates (in percentage) the received dedicated preamples, as defined in ETSI TS 136 314 [i.11].
	ReceivedDedicatedPreamblesCell int32 `json:"received_dedicated_preambles_cell,omitempty"`
	// It indicates (in percentage) the received randomly selected preambles in the high range, as defined in ETSI TS 136 314 [i.11].
	ReceivedRandomlySelectedPreamblesHighRangeCell int32 `json:"received_randomly_selected_preambles_high_range_cell,omitempty"`
	// It indicates (in percentage) the received randomly selected preambles in the low range, as defined in ETSI TS 136 314 [i.11].
	ReceivedRandomlySelectedPreamblesLowRangeCell int32 `json:"received_randomly_selected_preambles_low_range_cell,omitempty"`
	// It indicates the packet discard rate in percentage of the uplink GBR traffic in a cell, as defined in ETSI TS 136 314 [i.11].
	UlGbrPdrCell int32 `json:"ul_gbr_pdr_cell,omitempty"`
	// It indicates (in percentage) the PRB usage for uplink GBR traffic, as defined in ETSI TS 136 314 [i.11] and ETSI TS 136 423 [i.12].
	UlGbrPrbUsageCell int32 `json:"ul_gbr_prb_usage_cell,omitempty"`
	// It indicates the packet discard rate in percentage of the uplink non-GBR traffic in a cell, as defined in ETSI TS 136 314 [i.11].
	UlNongbrPdrCell int32 `json:"ul_nongbr_pdr_cell,omitempty"`
	// It indicates (in percentage) the PRB usage for uplink non-GBR traffic, as defined in ETSI TS 136 314 [i.11] and ETSI TS 136 423 [i.12].
	UlNongbrPrbUsageCell int32 `json:"ul_nongbr_prb_usage_cell,omitempty"`
	// It indicates (in percentage) the PRB usage for total uplink traffic, as defined in ETSI TS 136 314 [i.11] and ETSI TS 136 423 [i.12].
	UlTotalPrbUsageCell int32 `json:"ul_total_prb_usage_cell,omitempty"`
}
