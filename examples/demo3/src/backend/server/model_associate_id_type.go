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
 * MEC Demo 3 API
 *
 * Demo 3 is an edge application that can be used with AdvantEDGE or ETSI MEC Sandbox to demonstrate MEC011 and MEC021 usage
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package server

// AssociateIdType : Numeric value (0-255) corresponding to specified type of identifier
type AssociateIdType string

// List of AssociateIdType
const (
	UE_I_PV4_ADDRESS_AssociateIdType AssociateIdType = "UE_IPv4_ADDRESS"
	UE_IPV6_ADDRESS_AssociateIdType  AssociateIdType = "UE_IPV6_ADDRESS"
	NATED_IP_ADDRESS_AssociateIdType AssociateIdType = "NATED_IP_ADDRESS"
	GTP_TEID_AssociateIdType         AssociateIdType = "GTP_TEID"
)
