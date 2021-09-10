/*
 * Copyright (c) 2020  InterDigital Communications, Inc
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
 * AdvantEDGE Location Service REST API
 *
 * Location Service is AdvantEDGE's implementation of [ETSI MEC ISG MEC013 Location API](http://www.etsi.org/deliver/etsi_gs/MEC/001_099/013/02.01.01_60/gs_mec013v020101p.pdf) <p>The API is based on the Open Mobile Alliance's specification RESTful Network API for Zonal Presence <p>[Copyright (c) ETSI 2017](https://forge.etsi.org/etsi-forge-copyright-notice.txt) <p>**Micro-service**<br>[meep-loc-serv](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-loc-serv) <p>**Type & Usage**<br>Edge Service used by edge applications that want to get information about Users (UE) and Zone locations <p>**Note**<br>AdvantEDGE supports all of Location API endpoints (see below).
 *
 * API version: 2.1.1
 * Contact: AdvantEDGE@InterDigital.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package server

type LocationInfo struct {
	// Horizontal accuracy / (semi-major) uncertainty of location provided in meters, as defined in [14]. Present only if \"shape\" equals 4, 5 or 6
	Accuracy int32 `json:"accuracy,omitempty"`
	// Altitude accuracy / uncertainty of location provided in meters, as defined in [14]. Present only if \"shape\" equals 3 or 4
	AccuracyAltitude int32 `json:"accuracyAltitude,omitempty"`
	// Horizontal accuracy / (semi-major) uncertainty of location provided in meters, as defined in [14]. Present only if \"shape\" equals 4, 5 or 6
	AccuracySemiMinor int32 `json:"accuracySemiMinor,omitempty"`
	// Location altitude relative to the WGS84 ellipsoid surface.
	Altitude float32 `json:"altitude,omitempty"`
	// Confidence by which the position of a target entity is known to be within the shape description, expressed as a percentage and defined in [14]. Present only if \"shape\" equals 1, 4 or 6
	Confidence int32 `json:"confidence,omitempty"`
	// Present only if \"shape\" equals 6
	IncludedAngle int32 `json:"includedAngle,omitempty"`
	// Present only if \"shape\" equals 6
	InnerRadius int32 `json:"innerRadius,omitempty"`
	// Location latitude, expressed in the range -90° to +90°. Cardinality greater than one only if \"shape\" equals 7.
	Latitude []float32 `json:"latitude"`
	// Location longitude, expressed in the range -180° to +180°. Cardinality greater than one only if \"shape\" equals 7.
	Longitude []float32 `json:"longitude"`
	// Present only if \"shape\" equals 6
	OffsetAngle int32 `json:"offsetAngle,omitempty"`
	// Angle of orientation of the major axis, expressed in the range 0° to 180°, as defined in [14]. Present only if \"shape\" equals 4 or 6
	OrientationMajorAxis int32 `json:"orientationMajorAxis,omitempty"`
	// Shape information, as detailed in [14], associated with the reported location coordinate: <p>1 = ELLIPSOID_ARC <p>2 = ELLIPSOID_POINT <p>3 = ELLIPSOID_POINT_ALTITUDE <p>4 = ELLIPSOID_POINT_ALTITUDE_UNCERT_ELLIPSOID <p>5 = ELLIPSOID_POINT_UNCERT_CIRCLE <p>6 = ELLIPSOID_POINT_UNCERT_ELLIPSE <p>7 = POLYGON
	Shape int32 `json:"shape"`

	Timestamp *TimeStamp `json:"timestamp,omitempty"`
	// Present only if \"shape\" equals 6
	UncertaintyRadius int32 `json:"uncertaintyRadius,omitempty"`

	Velocity *LocationInfoVelocity `json:"velocity,omitempty"`
}
