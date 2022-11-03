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
 * AdvantEDGE Mobility Group Service REST API
 *
 * Mobility Group Service allows to form groups formed multiple edge application instances and share user states automatically withing the group <p>**Micro-service**<br>[meep-mg-manager](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-mg-manager) <p>**Type & Usage**<br>Edge Service used by edge applications to share user state between the  Mobility Group members <p>**Details**<br>API details available at _your-AdvantEDGE-ip-address/api_
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
		handler = met.MetricsHandler(handler, mgm.sandboxName, serviceName)
		handler = httpLog.LogRx(handler)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	// Path prefix router order is important
	// Service Api files
	handler = http.StripPrefix("/mgm/v1/api/", http.FileServer(http.Dir("./api/")))
	router.
		PathPrefix("/mgm/v1/api/").
		Name("Api").
		Handler(handler)
	// User supplied service API files
	handler = http.StripPrefix("/mgm/v1/user-api/", http.FileServer(http.Dir("./user-api/")))
	router.
		PathPrefix("/mgm/v1/user-api/").
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
		"/mgm/v1/",
		Index,
	},

	Route{
		"CreateMobilityGroup",
		strings.ToUpper("Post"),
		"/mgm/v1/mg/{mgName}",
		CreateMobilityGroup,
	},

	Route{
		"CreateMobilityGroupApp",
		strings.ToUpper("Post"),
		"/mgm/v1/mg/{mgName}/app/{appId}",
		CreateMobilityGroupApp,
	},

	Route{
		"CreateMobilityGroupUe",
		strings.ToUpper("Post"),
		"/mgm/v1/mg/{mgName}/app/{appId}/ue",
		CreateMobilityGroupUe,
	},

	Route{
		"DeleteMobilityGroup",
		strings.ToUpper("Delete"),
		"/mgm/v1/mg/{mgName}",
		DeleteMobilityGroup,
	},

	Route{
		"DeleteMobilityGroupApp",
		strings.ToUpper("Delete"),
		"/mgm/v1/mg/{mgName}/app/{appId}",
		DeleteMobilityGroupApp,
	},

	Route{
		"GetMobilityGroup",
		strings.ToUpper("Get"),
		"/mgm/v1/mg/{mgName}",
		GetMobilityGroup,
	},

	Route{
		"GetMobilityGroupApp",
		strings.ToUpper("Get"),
		"/mgm/v1/mg/{mgName}/app/{appId}",
		GetMobilityGroupApp,
	},

	Route{
		"GetMobilityGroupAppList",
		strings.ToUpper("Get"),
		"/mgm/v1/mg/{mgName}/app",
		GetMobilityGroupAppList,
	},

	Route{
		"GetMobilityGroupList",
		strings.ToUpper("Get"),
		"/mgm/v1/mg",
		GetMobilityGroupList,
	},

	Route{
		"SetMobilityGroup",
		strings.ToUpper("Put"),
		"/mgm/v1/mg/{mgName}",
		SetMobilityGroup,
	},

	Route{
		"SetMobilityGroupApp",
		strings.ToUpper("Put"),
		"/mgm/v1/mg/{mgName}/app/{appId}",
		SetMobilityGroupApp,
	},

	Route{
		"TransferAppState",
		strings.ToUpper("Post"),
		"/mgm/v1/mg/{mgName}/app/{appId}/state",
		TransferAppState,
	},
}
