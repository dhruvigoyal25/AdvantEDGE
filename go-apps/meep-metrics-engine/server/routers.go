/*
 * Copyright (c) 2019  InterDigital Communications, Inc
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
 * AdvantEDGE Metrics Service REST API
 *
 * Metrics Service provides metrics about the active scenario <p>**Micro-service**<br>[meep-metrics-engine](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-metrics-engine) <p>**Type & Usage**<br>Platform Service used by control/monitoring software and possibly by edge applications that require metrics <p>**Details**<br>API details available at _your-AdvantEDGE-ip-address:30000/api_ <p>**Default Port**<br>`30005`
 *
 * API version: 1.0.0
 * Contact: AdvantEDGE@InterDigital.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	v2 "github.com/InterDigitalInc/AdvantEDGE/go-apps/meep-metrics-engine/server/v2"
)

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

func IndexV2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! on v2")
}

func Init() (err error) {
	err = v2.Init()
	if err != nil {
		return err
	}
	return nil
}

var routes = Routes{
	Route{
		"IndexV2",
		"GET",
		"/metrics/v2/",
		IndexV2,
	},

	Route{
		"PostEventQuery",
		strings.ToUpper("Post"),
		"/metrics/v2/metrics/query/event",
		v2.PostEventQuery,
	},

	Route{
		"PostNetworkQuery",
		strings.ToUpper("Post"),
		"/metrics/v2/metrics/query/network",
		v2.PostNetworkQuery,
	},

	Route{
		"CreateEventSubscription",
		strings.ToUpper("Post"),
		"/metrics/v2/metrics/subscriptions/event",
		v2.CreateEventSubscription,
	},

	Route{
		"CreateNetworkSubscription",
		strings.ToUpper("Post"),
		"/metrics/v2/metrics/subscriptions/network",
		v2.CreateNetworkSubscription,
	},

	Route{
		"DeleteEventSubscriptionById",
		strings.ToUpper("Delete"),
		"/metrics/v2/metrics/subscriptions/event/{subscriptionId}",
		v2.DeleteEventSubscriptionById,
	},

	Route{
		"DeleteNetworkSubscriptionById",
		strings.ToUpper("Delete"),
		"/metrics/v2/metrics/subscriptions/network/{subscriptionId}",
		v2.DeleteNetworkSubscriptionById,
	},

	Route{
		"GetEventSubscription",
		strings.ToUpper("Get"),
		"/metrics/v2/metrics/subscriptions/event",
		v2.GetEventSubscription,
	},

	Route{
		"GetEventSubscriptionById",
		strings.ToUpper("Get"),
		"/metrics/v2/metrics/subscriptions/event/{subscriptionId}",
		v2.GetEventSubscriptionById,
	},

	Route{
		"GetNetworkSubscription",
		strings.ToUpper("Get"),
		"/metrics/v2/metrics/subscriptions/network",
		v2.GetNetworkSubscription,
	},

	Route{
		"GetNetworkSubscriptionById",
		strings.ToUpper("Get"),
		"/metrics/v2/metrics/subscriptions/network/{subscriptionId}",
		v2.GetNetworkSubscriptionById,
	},
}
