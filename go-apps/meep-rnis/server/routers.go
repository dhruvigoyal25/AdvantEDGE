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
 * AdvantEDGE Radio Network Information Service REST API
 *
 * Radio Network Information Service is AdvantEDGE's implementation of [ETSI MEC ISG MEC012 RNI API](https://www.etsi.org/deliver/etsi_gs/MEC/001_099/012/02.02.01_60/gs_MEC012v020201p.pdf) <p>[Copyright (c) ETSI 2017](https://forge.etsi.org/etsi-forge-copyright-notice.txt) <p>**Micro-service**<br>[meep-rnis](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-rnis) <p>**Type & Usage**<br>Edge Service used by edge applications that want to get information about radio conditions in the network <p>**Note**<br>AdvantEDGE supports a selected subset of RNI API endpoints (see below) and a subset of subscription types. <p>Supported subscriptions: <p> - CellChangeSubscription <p> - RabEstSubscription <p> - RabRelSubscription <p> - MeasRepUeSubscription <p> - NrMeasRepUeSubscription
 *
 * API version: 2.2.1
 * Contact: AdvantEDGE@InterDigital.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package server

import (
	"fmt"
	"net/http"
	"strings"

	httpLog "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-http-logger"
	met "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-metrics"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	var handler http.Handler
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		handler = met.MetricsHandler(handler, sandboxName, serviceName)
		handler = httpLog.LogRx(handler, "")
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	// Path prefix router order is important
	// Service Api files
	handler = http.StripPrefix("/rni/v2/api/", http.FileServer(http.Dir("./api/")))
	router.
		PathPrefix("/rni/v2/api/").
		Name("Api").
		Handler(handler)
	// User supplied service API files
	handler = http.StripPrefix("/rni/v2/user-api/", http.FileServer(http.Dir("./user-api/")))
	router.
		PathPrefix("/rni/v2/user-api/").
		Name("UserApi").
		Handler(handler)

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/rni/v2/",
		Index,
	},

	Route{
		"Layer2MeasInfoGET",
		strings.ToUpper("Get"),
		"/rni/v2/queries/layer2_meas",
		Layer2MeasInfoGET,
	},

	Route{
		"Mec011AppTerminationPOST",
		strings.ToUpper("Post"),
		"/rni/v2/notifications/mec011/appTermination",
		Mec011AppTerminationPOST,
	},

	Route{
		"PlmnInfoGET",
		strings.ToUpper("Get"),
		"/rni/v2/queries/plmn_info",
		PlmnInfoGET,
	},

	Route{
		"RabInfoGET",
		strings.ToUpper("Get"),
		"/rni/v2/queries/rab_info",
		RabInfoGET,
	},

	Route{
		"SubscriptionLinkListSubscriptionsGET",
		strings.ToUpper("Get"),
		"/rni/v2/subscriptions",
		SubscriptionLinkListSubscriptionsGET,
	},

	Route{
		"SubscriptionsDELETE",
		strings.ToUpper("Delete"),
		"/rni/v2/subscriptions/{subscriptionId}",
		SubscriptionsDELETE,
	},

	Route{
		"SubscriptionsGET",
		strings.ToUpper("Get"),
		"/rni/v2/subscriptions/{subscriptionId}",
		SubscriptionsGET,
	},

	Route{
		"SubscriptionsPOST",
		strings.ToUpper("Post"),
		"/rni/v2/subscriptions",
		SubscriptionsPOST,
	},

	Route{
		"SubscriptionsPUT",
		strings.ToUpper("Put"),
		"/rni/v2/subscriptions/{subscriptionId}",
		SubscriptionsPUT,
	},

	Route{
		"S1BearerInfoGET",
		strings.ToUpper("Get"),
		"/rni/v2/queries/s1_bearer_info",
		S1BearerInfoGET,
	},
}
