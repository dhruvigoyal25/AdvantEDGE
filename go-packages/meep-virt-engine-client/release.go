/*
 * MEEP Virtualization Engine REST API
 *
 * Copyright (c) 2019 InterDigital Communications, Inc. All rights reserved. The information provided herein is the proprietary and confidential information of InterDigital Communications, Inc. 
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type Release struct {

	// Release name
	Name string `json:"name,omitempty"`

	// Current release state
	State string `json:"state,omitempty"`
}
