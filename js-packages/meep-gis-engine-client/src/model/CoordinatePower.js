/*
 * Copyright (c) 2022  The AdvantEDGE Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * AdvantEDGE GIS Engine REST API
 * This API allows to control geo-spatial behavior and simulation. <p>**Micro-service**<br>[meep-gis-engine](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-gis-engine) <p>**Type & Usage**<br>Platform runtime interface to control geo-spatial behavior and simulation <p>**Details**<br>API details available at _your-AdvantEDGE-ip-address/api_
 *
 * OpenAPI spec version: 1.0.0
 * Contact: AdvantEDGE@InterDigital.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.9
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'));
  } else {
    // Browser globals (root is window)
    if (!root.AdvantEdgeGisEngineRestApi) {
      root.AdvantEdgeGisEngineRestApi = {};
    }
    root.AdvantEdgeGisEngineRestApi.CoordinatePower = factory(root.AdvantEdgeGisEngineRestApi.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';

  /**
   * The CoordinatePower model module.
   * @module model/CoordinatePower
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>CoordinatePower</code>.
   * Coordinates with their power values.
   * @alias module:model/CoordinatePower
   * @class
   * @param latitude {Number} Latitude of a second element for query purpose.
   * @param longitude {Number} Longitude of a second element for query purpose.
   * @param rsrq {Number} Reference Signal Received Quality as defined in ETSI TS 136 214.
   * @param rsrp {Number} Reference Signal Received Power as defined in ETSI TS 136 214.
   * @param poaName {String} Name of the POA for which RSRP/RSRQ values are calculated.
   */
  var exports = function(latitude, longitude, rsrq, rsrp, poaName) {
    this.latitude = latitude;
    this.longitude = longitude;
    this.rsrq = rsrq;
    this.rsrp = rsrp;
    this.poaName = poaName;
  };

  /**
   * Constructs a <code>CoordinatePower</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/CoordinatePower} obj Optional instance to populate.
   * @return {module:model/CoordinatePower} The populated <code>CoordinatePower</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      if (data.hasOwnProperty('latitude'))
        obj.latitude = ApiClient.convertToType(data['latitude'], 'Number');
      if (data.hasOwnProperty('longitude'))
        obj.longitude = ApiClient.convertToType(data['longitude'], 'Number');
      if (data.hasOwnProperty('rsrq'))
        obj.rsrq = ApiClient.convertToType(data['rsrq'], 'Number');
      if (data.hasOwnProperty('rsrp'))
        obj.rsrp = ApiClient.convertToType(data['rsrp'], 'Number');
      if (data.hasOwnProperty('poaName'))
        obj.poaName = ApiClient.convertToType(data['poaName'], 'String');
    }
    return obj;
  }

  /**
   * Latitude of a second element for query purpose.
   * @member {Number} latitude
   */
  exports.prototype.latitude = undefined;

  /**
   * Longitude of a second element for query purpose.
   * @member {Number} longitude
   */
  exports.prototype.longitude = undefined;

  /**
   * Reference Signal Received Quality as defined in ETSI TS 136 214.
   * @member {Number} rsrq
   */
  exports.prototype.rsrq = undefined;

  /**
   * Reference Signal Received Power as defined in ETSI TS 136 214.
   * @member {Number} rsrp
   */
  exports.prototype.rsrp = undefined;

  /**
   * Name of the POA for which RSRP/RSRQ values are calculated.
   * @member {String} poaName
   */
  exports.prototype.poaName = undefined;

  return exports;

}));
