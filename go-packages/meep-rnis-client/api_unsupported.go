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
 * ETSI GS MEC 012 - Radio Network Information API
 *
 * Radio Network Information Service is AdvantEDGE's implementation of [ETSI MEC ISG MEC012 RNI API](https://www.etsi.org/deliver/etsi_gs/MEC/001_099/012/02.02.01_60/gs_MEC012v020201p.pdf) <p>[Copyright (c) ETSI 2017](https://forge.etsi.org/etsi-forge-copyright-notice.txt) <p>**Micro-service**<br>[meep-rnis](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-rnis) <p>**Type & Usage**<br>Edge Service used by edge applications that want to get information about radio conditions in the network <p>**Note**<br>AdvantEDGE supports a selected subset of RNI API endpoints (see below) and a subset of subscription types. <p>Supported subscriptions: <p> - CellChangeSubscription <p> - RabEstSubscription <p> - RabRelSubscription <p> - MeasRepUeSubscription <p> - NrMeasRepUeSubscription
 *
 * API version: 2.2.1
 * Contact: AdvantEDGE@InterDigital.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type UnsupportedApiService service

/*
UnsupportedApiService Retrieve S1-U bearer information related to specific UE(s)
Queries information about the S1 bearer(s)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *S1BearerInfoGETOpts - Optional Parameters:
     * @param "TempUeId" (optional.Interface of []string) -  Comma separated list of temporary identifiers allocated for the specific UE as defined in   ETSI TS 136 413
     * @param "UeIpv4Address" (optional.Interface of []string) -  Comma separated list of IE IPv4 addresses as defined for the type for AssociateId
     * @param "UeIpv6Address" (optional.Interface of []string) -  Comma separated list of IE IPv6 addresses as defined for the type for AssociateId
     * @param "NatedIpAddress" (optional.Interface of []string) -  Comma separated list of IE NATed IP addresses as defined for the type for AssociateId
     * @param "GtpTeid" (optional.Interface of []string) -  Comma separated list of GTP TEID addresses as defined for the type for AssociateId
     * @param "CellId" (optional.Interface of []string) -  Comma separated list of E-UTRAN Cell Identities
     * @param "ErabId" (optional.Interface of []int32) -  Comma separated list of E-RAB identifiers

@return S1BearerInfo
*/

type S1BearerInfoGETOpts struct {
	TempUeId       optional.Interface
	UeIpv4Address  optional.Interface
	UeIpv6Address  optional.Interface
	NatedIpAddress optional.Interface
	GtpTeid        optional.Interface
	CellId         optional.Interface
	ErabId         optional.Interface
}

func (a *UnsupportedApiService) S1BearerInfoGET(ctx context.Context, localVarOptionals *S1BearerInfoGETOpts) (S1BearerInfo, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue S1BearerInfo
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/queries/s1_bearer_info"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.TempUeId.IsSet() {
		localVarQueryParams.Add("temp_ue_id", parameterToString(localVarOptionals.TempUeId.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.UeIpv4Address.IsSet() {
		localVarQueryParams.Add("ue_ipv4_address", parameterToString(localVarOptionals.UeIpv4Address.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.UeIpv6Address.IsSet() {
		localVarQueryParams.Add("ue_ipv6_address", parameterToString(localVarOptionals.UeIpv6Address.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.NatedIpAddress.IsSet() {
		localVarQueryParams.Add("nated_ip_address", parameterToString(localVarOptionals.NatedIpAddress.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.GtpTeid.IsSet() {
		localVarQueryParams.Add("gtp_teid", parameterToString(localVarOptionals.GtpTeid.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.CellId.IsSet() {
		localVarQueryParams.Add("cell_id", parameterToString(localVarOptionals.CellId.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.ErabId.IsSet() {
		localVarQueryParams.Add("erab_id", parameterToString(localVarOptionals.ErabId.Value(), "multi"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v S1BearerInfo
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 401 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 403 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 406 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 429 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
