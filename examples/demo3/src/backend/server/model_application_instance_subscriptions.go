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
 * MEC Demo 3 API
 *
 * This section describes use-case for demo 3 that the user can accomplish using the MEC Sandbox APIs from a MEC application
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package server

// List of ams subscription unique to device and their id
type ApplicationInstanceSubscriptions struct {
	AppTerminationSubscription *ApplicationInstanceSubscriptionsAppTerminationSubscription `json:"AppTerminationSubscription,omitempty"`

	SerAvailabilitySubscription *ApplicationInstanceSubscriptionsSerAvailabilitySubscription `json:"SerAvailabilitySubscription,omitempty"`

	AmsLinkListSubscription *ApplicationInstanceSubscriptionsAmsLinkListSubscription `json:"AmsLinkListSubscription,omitempty"`
}
