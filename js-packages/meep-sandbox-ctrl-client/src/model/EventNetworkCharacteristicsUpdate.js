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
 * AdvantEDGE Sandbox Controller REST API
 * This API is the main Sandbox Controller API for scenario deployment & event injection <p>**Micro-service**<br>[meep-sandbox-ctrl](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-sandbox-ctrl) <p>**Type & Usage**<br>Platform runtime interface to manage active scenarios and inject events in AdvantEDGE platform <p>**Details**<br>API details available at _your-AdvantEDGE-ip-address/api_
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
    define(['ApiClient', 'model/NetworkCharacteristics'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./NetworkCharacteristics'));
  } else {
    // Browser globals (root is window)
    if (!root.AdvantEdgeSandboxControllerRestApi) {
      root.AdvantEdgeSandboxControllerRestApi = {};
    }
    root.AdvantEdgeSandboxControllerRestApi.EventNetworkCharacteristicsUpdate = factory(root.AdvantEdgeSandboxControllerRestApi.ApiClient, root.AdvantEdgeSandboxControllerRestApi.NetworkCharacteristics);
  }
}(this, function(ApiClient, NetworkCharacteristics) {
  'use strict';

  /**
   * The EventNetworkCharacteristicsUpdate model module.
   * @module model/EventNetworkCharacteristicsUpdate
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>EventNetworkCharacteristicsUpdate</code>.
   * Network Characteristics update Event object
   * @alias module:model/EventNetworkCharacteristicsUpdate
   * @class
   */
  var exports = function() {
  };

  /**
   * Constructs a <code>EventNetworkCharacteristicsUpdate</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/EventNetworkCharacteristicsUpdate} obj Optional instance to populate.
   * @return {module:model/EventNetworkCharacteristicsUpdate} The populated <code>EventNetworkCharacteristicsUpdate</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      if (data.hasOwnProperty('elementName'))
        obj.elementName = ApiClient.convertToType(data['elementName'], 'String');
      if (data.hasOwnProperty('elementType'))
        obj.elementType = ApiClient.convertToType(data['elementType'], 'String');
      if (data.hasOwnProperty('netChar'))
        obj.netChar = NetworkCharacteristics.constructFromObject(data['netChar']);
    }
    return obj;
  }

  /**
   * Name of the network element to be updated
   * @member {String} elementName
   */
  exports.prototype.elementName = undefined;

  /**
   * Type of the network element to be updated
   * @member {module:model/EventNetworkCharacteristicsUpdate.ElementTypeEnum} elementType
   */
  exports.prototype.elementType = undefined;

  /**
   * @member {module:model/NetworkCharacteristics} netChar
   */
  exports.prototype.netChar = undefined;


  /**
   * Allowed values for the <code>elementType</code> property.
   * @enum {String}
   * @readonly
   */
  exports.ElementTypeEnum = {
    /**
     * value: "SCENARIO"
     * @const
     */
    SCENARIO: "SCENARIO",

    /**
     * value: "OPERATOR"
     * @const
     */
    OPERATOR: "OPERATOR",

    /**
     * value: "OPERATOR-CELLULAR"
     * @const
     */
    OPERATOR_CELLULAR: "OPERATOR-CELLULAR",

    /**
     * value: "ZONE"
     * @const
     */
    ZONE: "ZONE",

    /**
     * value: "POA"
     * @const
     */
    POA: "POA",

    /**
     * value: "POA-4G"
     * @const
     */
    pOA4G: "POA-4G",

    /**
     * value: "POA-5G"
     * @const
     */
    pOA5G: "POA-5G",

    /**
     * value: "POA-WIFI"
     * @const
     */
    POA_WIFI: "POA-WIFI",

    /**
     * value: "EDGE"
     * @const
     */
    EDGE: "EDGE",

    /**
     * value: "FOG"
     * @const
     */
    FOG: "FOG",

    /**
     * value: "UE"
     * @const
     */
    UE: "UE",

    /**
     * value: "DC"
     * @const
     */
    DC: "DC",

    /**
     * value: "UE-APP"
     * @const
     */
    UE_APP: "UE-APP",

    /**
     * value: "EDGE-APP"
     * @const
     */
    EDGE_APP: "EDGE-APP",

    /**
     * value: "CLOUD-APP"
     * @const
     */
    CLOUD_APP: "CLOUD-APP"
  };

  return exports;

}));
