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
 * AdvantEDGE GIS Engine REST API
 *
 * This API allows to control geo-spatial behavior and simulation. <p>**Micro-service**<br>[meep-gis-engine](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-gis-engine) <p>**Type & Usage**<br>Platform runtime interface to control geo-spatial behavior and simulation <p>**Details**<br>API details available at _your-AdvantEDGE-ip-address/api_
 *
 * API version: 1.0.0
 * Contact: AdvantEDGE@InterDigital.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Coordinates with their power values.
type CoordinatePower struct {
	// Latitude of a second element for query purpose.
	Latitude float32 `json:"latitude"`
	// Longitude of a second element for query purpose.
	Longitude float32 `json:"longitude"`
	// Reference Signal Received Quality as defined in ETSI TS 136 214.
	Rsrq int32 `json:"rsrq"`
	// Reference Signal Received Power as defined in ETSI TS 136 214.
	Rsrp int32 `json:"rsrp"`
	// Name of the POA for which RSRP/RSRQ values are calculated.
	PoaName string `json:"poaName"`
}
