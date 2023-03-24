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
 * AdvantEDGE Application Mobility API
 *
 * Device application interface is AdvantEDGE's implementation of [ETSI MEC ISG MEC016 Device application interface API](http://www.etsi.org/deliver/etsi_gs/MEC/001_099/021/02.02.01_60/gs_MEC016v020201p.pdf) <p>[Copyright (c) ETSI 2017](https://forge.etsi.org/etsi-forge-copyright-notice.txt) <p>**Micro-service**<br>[meep-dai](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-dai) <p>**Type & Usage**<br>Edge Service used by edge applications that want to get information about application mobility in the network <p>**Note**<br>AdvantEDGE supports a selected subset of Device application interface API endpoints (see below).
 *
 * API version: 2.2.1
 * Contact: AdvantEDGE@InterDigital.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type AppContext struct {
	// Provides indication to the MEC system that instantiation of the requested application is desired should a requested appLocation become available that was not at the time of the request.
	AppAutoInstantiation bool               `json:"appAutoInstantiation,omitempty"`
	AppInfo              *AppContextAppInfo `json:"appInfo"`
	// Used by the device application to request to receive notifications at the callbackReference URI relating to location availability for user application instantiation.
	AppLocationUpdates bool `json:"appLocationUpdates,omitempty"`
	// Uniquely identifies the device application. The length of the value shall not exceed 32 characters.
	AssociateDevAppId string `json:"associateDevAppId"`
	// URI assigned by the device application to receive application lifecycle related notifications. Inclusion in the request implies the client supports the pub/sub mechanism and is capable of receiving notifications. This endpoint shall be maintained for the lifetime of the application context.
	CallbackReference string `json:"callbackReference,omitempty"`
	// Uniquely identifies the application context in the MEC system. Assigned by the MEC system and shall be present other than in a create request. The length of the value shall not exceed 32 characters.
	ContextId string `json:"contextId,omitempty"`
}
