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
package server

// NTP server detail.
type TimingCapsNtpServers struct {
	NtpServerAddrType *TimingCapsNtpServersNtpServerAddrType `json:"ntpServerAddrType"`
	// NTP server address
	NtpServerAddr string `json:"ntpServerAddr"`
	// Minimum poll interval for NTP messages, in seconds as a power of two. Range 3 to 17
	MinPollingInterval int32 `json:"minPollingInterval"`
	// Maximum poll interval for NTP messages, in seconds as a power of two. Range 3 to 17
	MaxPollingInterval int32 `json:"maxPollingInterval"`
	// NTP server local priority
	LocalPriority int32 `json:"localPriority"`

	AuthenticationOption *TimingCapsNtpServersAuthenticationOption `json:"authenticationOption"`
	// Authentication key number
	AuthenticationKeyNum int32 `json:"authenticationKeyNum"`
}
