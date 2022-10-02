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
 * ETSI GS MEC 013 - Location API
 *
 * Location Service is AdvantEDGE's implementation of [ETSI MEC ISG MEC013 Location API](https://www.etsi.org/deliver/etsi_gs/MEC/001_099/013/02.02.01_60/gs_mec013v020201p.pdf) <p>The API is based on the Open Mobile Alliance's specification RESTful Network API for Zonal Presence <p>[Copyright (c) ETSI 2017](https://forge.etsi.org/etsi-forge-copyright-notice.txt) <p>**Micro-service**<br>[meep-loc-serv](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-loc-serv) <p>**Type & Usage**<br>Edge Service used by edge applications that want to get information about Users (UE) and Zone locations <p>**Note**<br>AdvantEDGE supports all of Location API endpoints (see below).
 *
 * API version: 2.2.1
 * Contact: AdvantEDGE@InterDigital.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// A type containing zonal traffic subscription
type ZonalTrafficSubscription struct {
	CallbackReference *CallbackReference `json:"callbackReference"`
	// A correlator that the client can use to tag this particular resource representation during a request to create a resource on the server.
	ClientCorrelator string `json:"clientCorrelator,omitempty"`
	// Period (in seconds) of time notifications are provided for. If set to \"0\" (zero), a default duration time, which is specified by the service policy, will be used. If the parameter is omitted, the notifications will continue until the maximum duration time, which is specified by the service policy, unless the notifications are stopped by deletion of subscription for notifications. This element MAY be given by the client during resource creation in order to signal the desired lifetime of the subscription. The server MUST return in this element the   period of time for which the subscription will still be valid.
	Duration int32 `json:"duration,omitempty"`
	// Interest realm of access point (e.g. geographical area, a type of industry etc.).
	InterestRealm []string `json:"interestRealm,omitempty"`
	// Self referring URL
	ResourceURL string `json:"resourceURL,omitempty"`
	// List of user event values to generate notifications for (these apply to zone identifier or all interest realms within zone identifier specified). If this element is missing, a notification is requested to be generated for any change in user event.
	UserEventCriteria []UserEventType `json:"userEventCriteria,omitempty"`
	// Identifier of zone
	ZoneId string `json:"zoneId"`
}
