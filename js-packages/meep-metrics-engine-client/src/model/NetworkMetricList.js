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
 * AdvantEDGE Metrics Service REST API
 * Metrics Service provides metrics about the active scenario <p>**Micro-service**<br>[meep-metrics-engine](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-metrics-engine) <p>**Type & Usage**<br>Platform Service used by control/monitoring software and possibly by edge applications that require metrics <p>**Details**<br>API details available at _your-AdvantEDGE-ip-address/api_
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
    define(['ApiClient', 'model/NetworkMetric'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./NetworkMetric'));
  } else {
    // Browser globals (root is window)
    if (!root.AdvantEdgeMetricsServiceRestApi) {
      root.AdvantEdgeMetricsServiceRestApi = {};
    }
    root.AdvantEdgeMetricsServiceRestApi.NetworkMetricList = factory(root.AdvantEdgeMetricsServiceRestApi.ApiClient, root.AdvantEdgeMetricsServiceRestApi.NetworkMetric);
  }
}(this, function(ApiClient, NetworkMetric) {
  'use strict';

  /**
   * The NetworkMetricList model module.
   * @module model/NetworkMetricList
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>NetworkMetricList</code>.
   * Network metrics query response
   * @alias module:model/NetworkMetricList
   * @class
   */
  var exports = function() {
  };

  /**
   * Constructs a <code>NetworkMetricList</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/NetworkMetricList} obj Optional instance to populate.
   * @return {module:model/NetworkMetricList} The populated <code>NetworkMetricList</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      if (data.hasOwnProperty('name'))
        obj.name = ApiClient.convertToType(data['name'], 'String');
      if (data.hasOwnProperty('columns'))
        obj.columns = ApiClient.convertToType(data['columns'], ['String']);
      if (data.hasOwnProperty('values'))
        obj.values = ApiClient.convertToType(data['values'], [NetworkMetric]);
    }
    return obj;
  }

  /**
   * Response name
   * @member {String} name
   */
  exports.prototype.name = undefined;

  /**
   * columns included in response based on queried values
   * @member {Array.<String>} columns
   */
  exports.prototype.columns = undefined;

  /**
   * @member {Array.<module:model/NetworkMetric>} values
   */
  exports.prototype.values = undefined;

  return exports;

}));
