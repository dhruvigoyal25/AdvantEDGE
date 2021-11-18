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
 * AdvantEDGE MEC Application Support API
 * MEC Application Support Service is AdvantEDGE's implementation of [ETSI MEC ISG MEC011 Application Enablement API](https://www.etsi.org/deliver/etsi_gs/MEC/001_099/011/02.01.01_60/gs_MEC011v020101p.pdf) <p>[Copyright (c) ETSI 2017](https://forge.etsi.org/etsi-forge-copyright-notice.txt) <p>**Micro-service**<br>[meep-app-enablement](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-app-enablement/server/app-support) <p>**Type & Usage**<br>Edge Service used by edge applications that want to get information about applications in the network <p>**Note**<br>AdvantEDGE supports a selected subset of Application Support API endpoints (see below).
 *
 * OpenAPI spec version: 2.1.1
 * Contact: AdvantEDGE@InterDigital.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 3.0.29
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/DestinationInterface', 'model/TrafficFilter'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./DestinationInterface'), require('./TrafficFilter'));
  } else {
    // Browser globals (root is window)
    if (!root.AdvantEdgeMecApplicationSupportApi) {
      root.AdvantEdgeMecApplicationSupportApi = {};
    }
    root.AdvantEdgeMecApplicationSupportApi.TrafficRule = factory(root.AdvantEdgeMecApplicationSupportApi.ApiClient, root.AdvantEdgeMecApplicationSupportApi.DestinationInterface, root.AdvantEdgeMecApplicationSupportApi.TrafficFilter);
  }
}(this, function(ApiClient, DestinationInterface, TrafficFilter) {
  'use strict';

  /**
   * The TrafficRule model module.
   * @module model/TrafficRule
   * @version 2.1.1
   */

  /**
   * Constructs a new <code>TrafficRule</code>.
   * This type represents the general information of a traffic rule.
   * @alias module:model/TrafficRule
   * @class
   * @param trafficRuleId {String} Identify the traffic rule.
   * @param filterType {module:model/TrafficRule.FilterTypeEnum} Definition of filter per FLOW or PACKET. If flow the filter match UE->EPC packet and the reverse packet is handled in the same context
   * @param priority {Number} Priority of this traffic rule. If traffic rule conflicts, the one with higher priority take precedence
   * @param trafficFilter {Array.<module:model/TrafficFilter>} 
   * @param action {module:model/TrafficRule.ActionEnum} The action of the MEC host data plane when a packet matches the trafficFilter
   * @param state {module:model/TrafficRule.StateEnum} Contains the traffic rule state. This attribute may be updated using HTTP PUT   method
   */
  var exports = function(trafficRuleId, filterType, priority, trafficFilter, action, state) {
    this.trafficRuleId = trafficRuleId;
    this.filterType = filterType;
    this.priority = priority;
    this.trafficFilter = trafficFilter;
    this.action = action;
    this.state = state;
  };

  /**
   * Constructs a <code>TrafficRule</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/TrafficRule} obj Optional instance to populate.
   * @return {module:model/TrafficRule} The populated <code>TrafficRule</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      if (data.hasOwnProperty('trafficRuleId'))
        obj.trafficRuleId = ApiClient.convertToType(data['trafficRuleId'], 'String');
      if (data.hasOwnProperty('filterType'))
        obj.filterType = ApiClient.convertToType(data['filterType'], 'String');
      if (data.hasOwnProperty('priority'))
        obj.priority = ApiClient.convertToType(data['priority'], 'Number');
      if (data.hasOwnProperty('trafficFilter'))
        obj.trafficFilter = ApiClient.convertToType(data['trafficFilter'], [TrafficFilter]);
      if (data.hasOwnProperty('action'))
        obj.action = ApiClient.convertToType(data['action'], 'String');
      if (data.hasOwnProperty('dstInterface'))
        obj.dstInterface = DestinationInterface.constructFromObject(data['dstInterface']);
      if (data.hasOwnProperty('state'))
        obj.state = ApiClient.convertToType(data['state'], 'String');
    }
    return obj;
  }

  /**
   * Identify the traffic rule.
   * @member {String} trafficRuleId
   */
  exports.prototype.trafficRuleId = undefined;

  /**
   * Definition of filter per FLOW or PACKET. If flow the filter match UE->EPC packet and the reverse packet is handled in the same context
   * @member {module:model/TrafficRule.FilterTypeEnum} filterType
   */
  exports.prototype.filterType = undefined;

  /**
   * Priority of this traffic rule. If traffic rule conflicts, the one with higher priority take precedence
   * @member {Number} priority
   */
  exports.prototype.priority = undefined;

  /**
   * @member {Array.<module:model/TrafficFilter>} trafficFilter
   */
  exports.prototype.trafficFilter = undefined;

  /**
   * The action of the MEC host data plane when a packet matches the trafficFilter
   * @member {module:model/TrafficRule.ActionEnum} action
   */
  exports.prototype.action = undefined;

  /**
   * @member {module:model/DestinationInterface} dstInterface
   */
  exports.prototype.dstInterface = undefined;

  /**
   * Contains the traffic rule state. This attribute may be updated using HTTP PUT   method
   * @member {module:model/TrafficRule.StateEnum} state
   */
  exports.prototype.state = undefined;


  /**
   * Allowed values for the <code>filterType</code> property.
   * @enum {String}
   * @readonly
   */
  exports.FilterTypeEnum = {
    /**
     * value: "FLOW"
     * @const
     */
    FLOW: "FLOW",

    /**
     * value: "PACKET"
     * @const
     */
    PACKET: "PACKET"
  };


  /**
   * Allowed values for the <code>action</code> property.
   * @enum {String}
   * @readonly
   */
  exports.ActionEnum = {
    /**
     * value: "DROP"
     * @const
     */
    DROP: "DROP",

    /**
     * value: "FORWARD_DECAPSULATED"
     * @const
     */
    FORWARD_DECAPSULATED: "FORWARD_DECAPSULATED",

    /**
     * value: "FORWARD_ENCAPSULATED"
     * @const
     */
    FORWARD_ENCAPSULATED: "FORWARD_ENCAPSULATED",

    /**
     * value: "PASSTHROUGH"
     * @const
     */
    PASSTHROUGH: "PASSTHROUGH",

    /**
     * value: "DUPLICATE_DECAPSULATED"
     * @const
     */
    DUPLICATE_DECAPSULATED: "DUPLICATE_DECAPSULATED",

    /**
     * value: "DUPLICATE_ENCAPSULATED"
     * @const
     */
    DUPLICATE_ENCAPSULATED: "DUPLICATE_ENCAPSULATED"
  };


  /**
   * Allowed values for the <code>state</code> property.
   * @enum {String}
   * @readonly
   */
  exports.StateEnum = {
    /**
     * value: "ACTIVE"
     * @const
     */
    ACTIVE: "ACTIVE",

    /**
     * value: "INACTIVE"
     * @const
     */
    INACTIVE: "INACTIVE"
  };

  return exports;

}));
