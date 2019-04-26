/*
 * MEEP Mobility Group Manager Model
 *
 * Copyright (c) 2019 InterDigital Communications, Inc. All rights reserved. The information provided herein is the proprietary and confidential information of InterDigital Communications, Inc. 
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Event object
type MobilityGroupEvent struct {

	// Mobility Group event name
	Name string `json:"name,omitempty"`

	// Mobility Group event type
	Type_ string `json:"type,omitempty"`

	// Mobility Group UE identifier
	UeId string `json:"ueId,omitempty"`

	AppState *MobilityGroupAppState `json:"appState,omitempty"`
}
