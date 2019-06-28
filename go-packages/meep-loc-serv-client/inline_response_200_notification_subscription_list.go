/*
 * Location API
 *
 * The ETSI MEC ISG MEC012 Location API described using OpenAPI. The API is based on the Open Mobile Alliance's specification RESTful Network API for Zonal Presence
 *
 * API version: 1.1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type InlineResponse200NotificationSubscriptionList struct {
	ZonalTrafficSubscription []ZonalTrafficSubscription `json:"zonalTrafficSubscription,omitempty"`

	// Self referring URL.
	ResourceURL string `json:"resourceURL,omitempty"`
}
