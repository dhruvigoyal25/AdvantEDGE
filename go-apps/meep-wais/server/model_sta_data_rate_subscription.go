/*
 * ETSI GS MEC 028 - WLAN Access Information API
 *
 * The ETSI MEC ISG MEC028 WLAN Access Information API described using OpenAPI
 *
 * API version: 2.1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package server

type StaDataRateSubscription struct {
	Links *AssocStaSubscriptionLinks `json:"_links,omitempty"`
	// URI selected by the service consumer to receive notifications on the subscribed WLAN Access Information Service. This shall be included both in the request and in response.
	CallbackReference string `json:"callbackReference"`

	ExpiryDeadline *TimeStamp `json:"expiryDeadline,omitempty"`
	// Identifier(s) to uniquely specify the target client station(s) for the subscription
	StaId []StaIdentity `json:"staId"`
	// Shall be set to \"StaDataRateSubscription\".
	SubscriptionType string `json:"subscriptionType"`
}
