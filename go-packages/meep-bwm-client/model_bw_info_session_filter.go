/*
 * Copyright (c) 2022  The AdvantEDGE Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * AdvantEDGE Bandwidth Management API
 *
 * Bandwidth Management Sercice is AdvantEDGE's implementation of [ETSI MEC ISG MEC015 Traffic Management APIs](https://www.etsi.org/deliver/etsi_gs/MEC/001_099/015/02.02.01_60/gs_MEC015v020201p.pdf) <p>[Copyright (c) ETSI 2017](https://forge.etsi.org/etsi-forge-copyright-notice.txt) <p>**Micro-service**<br>[meep-tm](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-tm/server/bwm) <p>**Type & Usage**<br>Edge Service used by edge applications that want to get information about BWM Info and Session(s) in the network <p>**Note**<br>AdvantEDGE supports all Bandwidth Management API endpoints.
 *
 * API version: 2.2.1
 * Contact: AdvantEDGE@InterDigital.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type BwInfoSessionFilter struct {
	// Destination address identity of session. The string for a IPv4 address shall be formatted in the \"dotted decimal\" notation as defined in IETF RFC 1166 [10]. The string for a IPv6 address shall be formatted according to clause 4 of IETF RFC 5952 [11], with in CIDR notation [12] used to provide the routing prefix.
	DstAddress string `json:"dstAddress,omitempty"`
	// Destination port identity of session
	DstPort string `json:"dstPort,omitempty"`
	// Protocol number
	Protocol string `json:"protocol,omitempty"`
	// Source address identity of session. The string for a IPv4 address shall be formatted in the \"dotted decimal\" notation as defined in IETF RFC 1166 [10]. The string for a IPv6 address shall be formatted according to clause 4 of IETF RFC 5952 [11], with in CIDR notation [12] used to provide the routing prefix.
	SourceIp string `json:"sourceIp,omitempty"`
	// Source port identity of session
	SourcePort string `json:"sourcePort,omitempty"`
}
