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
 * AdvantEDGE Device Application Interface
 *
 * Device application interface is AdvantEDGE's implementation of [ETSI MEC ISG MEC016 Device application interface API](http://www.etsi.org/deliver/etsi_gs/MEC/001_099/021/02.02.01_60/gs_MEC016v020201p.pdf) <p>[Copyright (c) ETSI 2017](https://forge.etsi.org/etsi-forge-copyright-notice.txt) <p>**Micro-service**<br>[meep-dai](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-dai) <p>**Type & Usage**<br>Edge Service used by edge applications that want to get information about application mobility in the network <p>**Note**<br>AdvantEDGE supports a selected subset of Device application interface API endpoints (see below).
 *
 * API version: 2.2.1
 * Contact: AdvantEDGE@InterDigital.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type AppContextAppInfo struct {
	// Identifier of this MEC application descriptor. This attribute shall be globally unique. It is equivalent to the appDId defined in clause 6.2.1.2 of ETSI GS MEC 0102 [1]. It shall be present if the application is one in the ApplicationList.
	AppDId string `json:"appDId,omitempty"`
	// Identifies the version of the application descriptor. It is equivalent to the appDVersion defined in clause 6.2.1.2 of ETSI GS MEC 0102 [1].
	AppDVersion string `json:"appDVersion"`
	// Human readable description of the MEC application. The length of the value shall not exceed 128 characters.
	AppDescription string `json:"appDescription,omitempty"`
	// Name of the MEC application. The length of the value shall not exceed 32 characters.
	AppName string `json:"appName"`
	// Provider of the MEC application. The length of the value shall not exceed 32 characters.
	AppProvider string `json:"appProvider"`
	// Software version of the MEC application. The length of the value shall not exceed 32 characters.
	AppSoftVersion string `json:"appSoftVersion,omitempty"`
	// URI of the application package. Included in the request if the application is not one in the ApplicationList. appPackageSource enables on-boarding of the application package into the MEC system. The application package shall comply with the definitions in clause 6.2.1.2 of ETSI GS MEC 0102 [1].
	AppPackageSource string `json:"appPackageSource,omitempty"`
	// List of user application instance information.
	UserAppInstanceInfo []AppContextAppInfoUserAppInstanceInfo `json:"userAppInstanceInfo"`
}
