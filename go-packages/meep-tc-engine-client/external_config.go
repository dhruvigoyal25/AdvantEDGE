/*
 * MEEP TC controller API
 *
 * Copyright (c) 2019 InterDigital Communications, Inc. All rights reserved. The information provided herein is the proprietary and confidential information of InterDigital Communications, Inc. 
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// External Process configuration. NOTE: Only valid if 'isExternal' is set.
type ExternalConfig struct {

	IngressServiceMap []ServiceMap `json:"ingressServiceMap,omitempty"`

	EgressServiceMap []ServiceMap `json:"egressServiceMap,omitempty"`
}
