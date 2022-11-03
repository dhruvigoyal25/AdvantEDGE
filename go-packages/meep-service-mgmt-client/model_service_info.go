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
 * AdvantEDGE Service Management API
 *
 * MEC Service Management Service is AdvantEDGE's implementation of [ETSI MEC ISG MEC011 Application Enablement API](https://www.etsi.org/deliver/etsi_gs/MEC/001_099/011/02.02.01_60/gs_MEC011v020201p.pdf) <p>[Copyright (c) ETSI 2017](https://forge.etsi.org/etsi-forge-copyright-notice.txt) <p>**Micro-service**<br>[meep-app-enablement](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-app-enablement/server/service-mgmt) <p>**Type & Usage**<br>Edge Service used by edge applications that want to get information about services in the network <p>**Note**<br>AdvantEDGE supports all of Service Management API endpoints (see below).
 *
 * API version: 2.2.1
 * Contact: AdvantEDGE@InterDigital.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// This type represents the general information of a MEC service.
type ServiceInfo struct {
	SerInstanceId string `json:"serInstanceId,omitempty"`

	SerName string `json:"serName"`

	SerCategory *CategoryRef `json:"serCategory,omitempty"`
	// Service version
	Version string `json:"version"`

	State *ServiceState `json:"state"`
	// Identifier of the platform-provided transport to be used by the service. Valid identifiers may be obtained using the \"Transport information query\" procedure. May be present in POST requests to signal the use of a platform-provided transport for the service, and shall be absent otherwise.
	TransportId string `json:"transportId,omitempty"`

	TransportInfo *TransportInfo `json:"transportInfo,omitempty"`

	Serializer *SerializerType `json:"serializer"`

	ScopeOfLocality *LocalityType `json:"scopeOfLocality,omitempty"`
	// Indicate whether the service can only be consumed by the MEC applications located in the same locality (as defined by scopeOfLocality) as this  service instance.
	// manually removed the omitempty
	ConsumedLocalOnly bool `json:"consumedLocalOnly"`
	// Indicate whether the service is located in the same locality (as defined by scopeOfLocality) as the consuming MEC application.
	// manually removed the omitempty
	IsLocal bool `json:"isLocal"`

	LivenessInterval int32 `json:"livenessInterval,omitempty"`

	Links *ServiceInfoLinks `json:"_links"`
}
