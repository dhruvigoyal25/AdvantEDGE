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
 * AdvantEDGE WLAN Access Information API
 *
 * WLAN Access Information Service is AdvantEDGE's implementation of [ETSI MEC ISG MEC028 WAI API](http://www.etsi.org/deliver/etsi_gs/MEC/001_099/028/02.02.01_60/gs_MEC028v020201p.pdf) <p>[Copyright (c) ETSI 2020](https://forge.etsi.org/etsi-forge-copyright-notice.txt) <p>**Micro-service**<br>[meep-wais](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-wais) <p>**Type & Usage**<br>Edge Service used by edge applications that want to get information about WLAN access information in the network <p>**Note**<br>AdvantEDGE supports a selected subset of WAI API subscription types. <p>Supported subscriptions: <p> - AssocStaSubscription <p> - StaDataRateSubscription
 *
 * API version: 2.2.1
 * Contact: AdvantEDGE@InterDigital.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package server

type StaDataRateSubscription struct {
	Links *AssocStaSubscriptionLinks `json:"_links,omitempty"`

	CallbackReference string `json:"callbackReference,omitempty"`

	ExpiryDeadline *TimeStamp `json:"expiryDeadline,omitempty"`

	NotificationEvent *StaDataRateSubscriptionNotificationEvent `json:"notificationEvent,omitempty"`
	// Set for periodic notification reporting. Value indicates the notification period in seconds.
	NotificationPeriod int32 `json:"notificationPeriod,omitempty"`
	// Set to TRUE by the service consumer to request a test notification on the callbackReference URI to determine if it is reachable by the WAIS for notifications.
	RequestTestNotification bool `json:"requestTestNotification,omitempty"`
	// Identifier(s) to uniquely specify the target client station(s) for the subscription.
	StaId []StaIdentity `json:"staId"`
	// Shall be set to \"StaDataRateSubscription\".
	SubscriptionType string `json:"subscriptionType"`

	WebsockNotifConfig *WebsockNotifConfig `json:"websockNotifConfig,omitempty"`
}
