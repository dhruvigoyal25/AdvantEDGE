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

package client

type StaStatisticsConfig struct {
	// As per Table 9-92 of IEEE 802.11-2016 [8].
	GroupIdentity int32 `json:"groupIdentity"`
	// Valid if triggeredReport = true.   Specifies the number of MAC service data units or protocol data units to determine if the trigger conditions are met.
	MeasurementCount int32                       `json:"measurementCount,omitempty"`
	TriggerCondition *StaCounterTriggerCondition `json:"triggerCondition,omitempty"`
	// Valid if triggeredReport = true.   The Trigger Timeout field contains a value in units of 100 time-units of 1024 µs during which a measuring STA does not generate further triggered STA Statistics Reports after a trigger condition has been met.
	TriggerTimeout int32 `json:"triggerTimeout,omitempty"`
	// True = triggered reporting, otherwise duration.
	TriggeredReport bool `json:"triggeredReport"`
}
