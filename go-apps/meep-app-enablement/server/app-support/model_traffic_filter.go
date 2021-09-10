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
 * AdvantEDGE MEC Application Support API
 *
 * MEC Application Support Service is AdvantEDGE's implementation of [ETSI MEC ISG MEC011 Application Enablement API](http://www.etsi.org/deliver/etsi_gs/MEC/001_099/011/02.02.01_60/gs_MEC011v020201p.pdf) <p>[Copyright (c) ETSI 2017](https://forge.etsi.org/etsi-forge-copyright-notice.txt) <p>**Micro-service**<br>[meep-app-enablement](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-app-enablement/server/app-support) <p>**Type & Usage**<br>Edge Service used by edge applications that want to get information about applications in the network <p>**Note**<br>AdvantEDGE supports a selected subset of Application Support API endpoints (see below).
 *
 * API version: 2.2.1
 * Contact: AdvantEDGE@InterDigital.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package server

// This type represents the traffic filter.
type TrafficFilter struct {
	// An IP address or a range of IP address. For IPv4, the IP address could be an IP address plus mask, or an individual IP address, or a range of IP addresses. For IPv6, the IP address could be an IP prefix, or a range of IP prefixes.
	SrcAddress []string `json:"srcAddress,omitempty"`
	// An IP address or a range of IP address. For IPv4, the IP address could be an IP address plus mask, or an individual IP address, or a range of IP addresses. For IPv6, the IP address could be an IP prefix, or a range of IP prefixes.
	DstAddress []string `json:"dstAddress,omitempty"`
	// A port or a range of ports
	SrcPort []string `json:"srcPort,omitempty"`
	// A port or a range of ports
	DstPort []string `json:"dstPort,omitempty"`
	// Specify the protocol of the traffic filter
	Protocol []string `json:"protocol,omitempty"`
	// Used for token based traffic rule
	Token []string `json:"token,omitempty"`
	// Used for GTP tunnel based traffic rule
	SrcTunnelAddress []string `json:"srcTunnelAddress,omitempty"`
	// Used for GTP tunnel based traffic rule
	TgtTunnelAddress []string `json:"tgtTunnelAddress,omitempty"`
	// Used for GTP tunnel based traffic rule
	SrcTunnelPort []string `json:"srcTunnelPort,omitempty"`
	// Used for GTP tunnel based traffic rule
	DstTunnelPort []string `json:"dstTunnelPort,omitempty"`
	// Used to match all packets that have the same Quality Class Indicator (QCI).
	QCI int32 `json:"qCI,omitempty"`
	// Used to match all IPv4 packets that have the same Differentiated Services Code Point (DSCP)
	DSCP int32 `json:"dSCP,omitempty"`
	// Used to match all IPv6 packets that have the same Traffic Class.
	TC int32 `json:"tC,omitempty"`
}
