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
 * ETSI GS MEC 028 - WLAN Access Information API
 *
 * The ETSI MEC ISG MEC028 WLAN Access Information API described using OpenAPI
 *
 * API version: 2.1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type AssociatedStations struct {
	// Unique number which identifies a particular association between an Access Point and a station.
	AssocId string `json:"assocId,omitempty"`
	// IPv4 or IPv6 address allocated for the station associated with the Access Point.
	IpAddress []string `json:"ipAddress,omitempty"`
	// Unique identifier assigned to a station (as network interface controller) for communications at the data link layer of a network segment.
	MacId string `json:"macId"`
}
