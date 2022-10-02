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
 * MEC Application Support API
 *
 * The ETSI MEC ISG MEC011 MEC Application Support API described using OpenAPI
 *
 * API version: 2.2.1
 * Contact: AdvantEDGE@InterDigital.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// This type represents the tunnel information.
type TunnelInfo struct {
	TunnelType *TunnelInfoTunnelType `json:"tunnelType"`
	// Destination address of the tunnel
	TunnelDstAddress string `json:"tunnelDstAddress,omitempty"`
	// Source address of the tunnel
	TunnelSrcAddress string `json:"tunnelSrcAddress,omitempty"`
}
