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
 * AdvantEDGE MEC Service Management API
 * MEC Service Management Service is AdvantEDGE's implementation of [ETSI MEC ISG MEC011 Application Enablement API](https://www.etsi.org/deliver/etsi_gs/MEC/001_099/011/02.01.01_60/gs_MEC011v020101p.pdf) <p>[Copyright (c) ETSI 2017](https://forge.etsi.org/etsi-forge-copyright-notice.txt) <p>**Micro-service**<br>[meep-app-enablement](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-app-enablement/server/service-mgmt) <p>**Type & Usage**<br>Edge Service used by edge applications that want to get information about services in the network <p>**Note**<br>AdvantEDGE supports all of Service Management API endpoints (see below).
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
    define(['ApiClient', 'model/ServiceState'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./ServiceState'));
  } else {
    // Browser globals (root is window)
    if (!root.AdvantEdgeMecServiceManagementApi) {
      root.AdvantEdgeMecServiceManagementApi = {};
    }
    root.AdvantEdgeMecServiceManagementApi.ServiceStates = factory(root.AdvantEdgeMecServiceManagementApi.ApiClient, root.AdvantEdgeMecServiceManagementApi.ServiceState);
  }
}(this, function(ApiClient, ServiceState) {
  'use strict';

  /**
   * The ServiceStates model module.
   * @module model/ServiceStates
   * @version 2.1.1
   */

  /**
   * Constructs a new <code>ServiceStates</code>.
   * States of the services about which to report events. If the event is  a state change, this filter represents the state after the change.
   * @alias module:model/ServiceStates
   * @class
   * @extends Array
   */
  var exports = function() {
  };

  /**
   * Constructs a <code>ServiceStates</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/ServiceStates} obj Optional instance to populate.
   * @return {module:model/ServiceStates} The populated <code>ServiceStates</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      ApiClient.constructFromObject(data, obj, 'ServiceState');
    }
    return obj;
  }

  Object.setPrototypeOf(exports.prototype, new Array());
  return exports;

}));
