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
 * AdvantEDGE WLAN Access Information API
 *
 * WLAN Access Information Service is AdvantEDGE's implementation of [ETSI MEC ISG MEC028 WAI API](http://www.etsi.org/deliver/etsi_gs/MEC/001_099/028/02.02.01_60/gs_MEC028v020201p.pdf) <p>[Copyright (c) ETSI 2020](https://forge.etsi.org/etsi-forge-copyright-notice.txt) <p>**Micro-service**<br>[meep-wais](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-wais) <p>**Type & Usage**<br>Edge Service used by edge applications that want to get information about WLAN access information in the network <p>**Note**<br>AdvantEDGE supports a selected subset of WAI API subscription types. <p>Supported subscriptions: <p> - AssocStaSubscription <p> - StaDataRateSubscription
 *
 * API version: 2.2.1
 * Contact: AdvantEDGE@InterDigital.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type BssidInfo struct {
	// The apReachability field indicates whether the AP identified by this BSSID is reachable by the STA that requested the neighbor report. Valid values: 0 = reserved 1 = not reachable 2 = unknown 3 = reachable.
	ApReachability int32            `json:"apReachability"`
	Capabilities   *BssCapabilities `json:"capabilities"`
	// True indicates the AP represented by this BSSID is an AP that has set the Fine Timing Measurement Responder field of the Extended Capabilities element to 1.  False indicates either that the reporting AP has dot11FineTimingMsmtRespActivated equal to false, or the reported AP has not set the Fine Timing Measurement Responder field of the Extended Capabilities element to 1 or that the Fine Timing Measurement Responder field of the reported AP is not available to the reporting AP at this time.
	Ftm bool `json:"ftm"`
	// True indicates that the AP represented by this BSSID is an HT AP including the HT Capabilities element in its Beacons, and that the contents of that HT Capabilities element are identical to the HT Capabilities element advertised by the AP sending the report.
	HighThroughput bool `json:"highThroughput"`
	// True indicates the AP represented by this BSSID is including an MDE in its Beacon frames and that the contents of that MDE are identical to the MDE advertised by the AP sending the report.
	MobilityDomain bool `json:"mobilityDomain"`
	// True indicates the AP identified by this BSSID supports the same security provisioning as used by the STA in its current association.  False indicates either that the AP does not support the same security provisioning or that the security information is not available at this time.
	Security bool `json:"security"`
	// True indicates that the AP represented by this BSSID is a VHT AP and that the VHT Capabilities element, if included as a subelement in the report, is identical in content to the VHT Capabilities element included in the AP's Beacon.
	VeryHighThroughput bool `json:"veryHighThroughput"`
}
