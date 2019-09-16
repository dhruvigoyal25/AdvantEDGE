/*
 * MEEP Controller REST API
 *
 * Copyright (c) 2019  InterDigital Communications, Inc Licensed under the Apache License, Version 2.0 (the \"License\"); you may not use this file except in compliance with the License. You may obtain a copy of the License at      http://www.apache.org/licenses/LICENSE-2.0  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an \"AS IS\" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Logical network location object
type NetworkLocation struct {

	// Unique network location ID
	Id string `json:"id,omitempty"`

	// Network location name
	Name string `json:"name,omitempty"`

	// Network location type
	Type_ string `json:"type,omitempty"`

	// Latency in ms for all terminal links within network location
	TerminalLinkLatency int32 `json:"terminalLinkLatency,omitempty"`

	// Latency variation in ms for all terminal links within network location
	TerminalLinkLatencyVariation int32 `json:"terminalLinkLatencyVariation,omitempty"`

	// The limit of the traffic supported for all terminal links within the network location
	TerminalLinkThroughput int32 `json:"terminalLinkThroughput,omitempty"`

	// Packet lost (in terms of percentage) for all terminal links within the network location
	TerminalLinkPacketLoss float64 `json:"terminalLinkPacketLoss,omitempty"`

	// Key/Value Pair Map (string, string)
	Meta map[string]string `json:"meta,omitempty"`

	// Key/Value Pair Map (string, string)
	UserMeta map[string]string `json:"userMeta,omitempty"`

	PhysicalLocations []PhysicalLocation `json:"physicalLocations,omitempty"`
}
