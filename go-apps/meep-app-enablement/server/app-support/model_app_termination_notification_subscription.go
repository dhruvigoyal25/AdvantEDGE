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

// This type represents the information that the MEC platform notifies the subscribed application instance about  the corresponding application instance termination/stop.
type AppTerminationNotificationSubscription struct {
	// Shall be set to AppTerminationNotificationSubscription.
	SubscriptionType string `json:"subscriptionType"`
	// URI selected by the MEC application instance to receive notifications on the subscribed MEC application instance management information. This shall be included in both the request and the response.
	CallbackReference string `json:"callbackReference"`

	Links *Self `json:"_links"`
	// It is used as the filtering criterion for the subscribed events.
	AppInstanceId string `json:"appInstanceId"`
}
