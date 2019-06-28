/*
 * Location API
 *
 * The ETSI MEC ISG MEC012 Location API described using OpenAPI. The API is based on the Open Mobile Alliance's specification RESTful Network API for Zonal Presence
 *
 * API version: 1.1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// exampleAPI coming from the spec
const basepath = "/etsi-013"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		basepath + "/location/v1/",
		Index,
	},

	Route{
		"UserTrackingSubDelById",
		strings.ToUpper("Delete"),
		basepath + "/location/v1/subscriptions/userTracking/{subscriptionId}",
		UserTrackingSubDelById,
	},

	Route{
		"UserTrackingSubGet",
		strings.ToUpper("Get"),
		basepath + "/location/v1/subscriptions/userTracking",
		UserTrackingSubGet,
	},

	Route{
		"UserTrackingSubGetById",
		strings.ToUpper("Get"),
		basepath + "/location/v1/subscriptions/userTracking/{subscriptionId}",
		UserTrackingSubGetById,
	},

	Route{
		"UserTrackingSubPost",
		strings.ToUpper("Post"),
		basepath + "/location/v1/subscriptions/userTracking",
		UserTrackingSubPost,
	},

	Route{
		"UserTrackingSubPutById",
		strings.ToUpper("Put"),
		basepath + "/location/v1/subscriptions/userTracking/{subscriptionId}",
		UserTrackingSubPutById,
	},

	Route{
		"ZonalTrafficSubDelById",
		strings.ToUpper("Delete"),
		basepath + "/location/v1/subscriptions/zonalTraffic/{subscriptionId}",
		ZonalTrafficSubDelById,
	},

	Route{
		"ZonalTrafficSubGet",
		strings.ToUpper("Get"),
		basepath + "/location/v1/subscriptions/zonalTraffic",
		ZonalTrafficSubGet,
	},

	Route{
		"ZonalTrafficSubGetById",
		strings.ToUpper("Get"),
		basepath + "/location/v1/subscriptions/zonalTraffic/{subscriptionId}",
		ZonalTrafficSubGetById,
	},

	Route{
		"ZonalTrafficSubPost",
		strings.ToUpper("Post"),
		basepath + "/location/v1/subscriptions/zonalTraffic",
		ZonalTrafficSubPost,
	},

	Route{
		"ZonalTrafficSubPutById",
		strings.ToUpper("Put"),
		basepath + "/location/v1/subscriptions/zonalTraffic/{subscriptionId}",
		ZonalTrafficSubPutById,
	},

	Route{
		"ZoneStatusDelById",
		strings.ToUpper("Delete"),
		basepath + "/location/v1/subscriptions/zoneStatus/{subscriptionId}",
		ZoneStatusDelById,
	},

	Route{
		"ZoneStatusGet",
		strings.ToUpper("Get"),
		basepath + "/location/v1/subscriptions/zonalStatus",
		ZoneStatusGet,
	},

	Route{
		"ZoneStatusGetById",
		strings.ToUpper("Get"),
		basepath + "/location/v1/subscriptions/zoneStatus/{subscriptionId}",
		ZoneStatusGetById,
	},

	Route{
		"ZoneStatusPost",
		strings.ToUpper("Post"),
		basepath + "/location/v1/subscriptions/zonalStatus",
		ZoneStatusPost,
	},

	Route{
		"ZoneStatusPutById",
		strings.ToUpper("Put"),
		basepath + "/location/v1/subscriptions/zoneStatus/{subscriptionId}",
		ZoneStatusPutById,
	},

	Route{
		"UsersGet",
		strings.ToUpper("Get"),
		basepath + "/location/v1/users",
		UsersGet,
	},

	Route{
		"UsersGetById",
		strings.ToUpper("Get"),
		basepath + "/location/v1/users/{userId}",
		UsersGetById,
	},

	Route{
		"ZonesByIdGetAps",
		strings.ToUpper("Get"),
		basepath + "/location/v1/zones/{zoneId}/accessPoints",
		ZonesByIdGetAps,
	},

	Route{
		"ZonesByIdGetApsById",
		strings.ToUpper("Get"),
		basepath + "/location/v1/zones/{zoneId}/accessPoints/{accessPointId}",
		ZonesByIdGetApsById,
	},

	Route{
		"ZonesGet",
		strings.ToUpper("Get"),
		basepath + "/location/v1/zones",
		ZonesGet,
	},

	Route{
		"ZonesGetById",
		strings.ToUpper("Get"),
		basepath + "/location/v1/zones/{zoneId}",
		ZonesGetById,
	},
}
