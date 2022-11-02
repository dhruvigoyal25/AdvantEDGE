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
 * MEC Demo 3 API
 * Demo 3 is an edge application that can be used with AdvantEDGE or ETSI MEC Sandbox to demonstrate MEC011 and MEC021 usage
 *
 * OpenAPI spec version: 0.0.1
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
    define(['ApiClient', 'model/CommunicationInterface'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./CommunicationInterface'));
  } else {
    // Browser globals (root is window)
    if (!root.MecDemo3Api) {
      root.MecDemo3Api = {};
    }
    root.MecDemo3Api.AdjacentAppInfoNotificationAdjacentAppInfo = factory(root.MecDemo3Api.ApiClient, root.MecDemo3Api.CommunicationInterface);
  }
}(this, function(ApiClient, CommunicationInterface) {
  'use strict';

  /**
   * The AdjacentAppInfoNotificationAdjacentAppInfo model module.
   * @module model/AdjacentAppInfoNotificationAdjacentAppInfo
   * @version 0.0.1
   */

  /**
   * Constructs a new <code>AdjacentAppInfoNotificationAdjacentAppInfo</code>.
   * @alias module:model/AdjacentAppInfoNotificationAdjacentAppInfo
   * @class
   * @param appInstanceId {String} Identifier of the adjacent application instance.
   * @param commInterface {Array.<module:model/CommunicationInterface>} If present, it represents the communication interface(s) information of the application instance.
   */
  var exports = function(appInstanceId, commInterface) {
    this.appInstanceId = appInstanceId;
    this.commInterface = commInterface;
  };

  /**
   * Constructs a <code>AdjacentAppInfoNotificationAdjacentAppInfo</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/AdjacentAppInfoNotificationAdjacentAppInfo} obj Optional instance to populate.
   * @return {module:model/AdjacentAppInfoNotificationAdjacentAppInfo} The populated <code>AdjacentAppInfoNotificationAdjacentAppInfo</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      if (data.hasOwnProperty('appInstanceId'))
        obj.appInstanceId = ApiClient.convertToType(data['appInstanceId'], 'String');
      if (data.hasOwnProperty('commInterface'))
        obj.commInterface = ApiClient.convertToType(data['commInterface'], [CommunicationInterface]);
    }
    return obj;
  }

  /**
   * Identifier of the adjacent application instance.
   * @member {String} appInstanceId
   */
  exports.prototype.appInstanceId = undefined;

  /**
   * If present, it represents the communication interface(s) information of the application instance.
   * @member {Array.<module:model/CommunicationInterface>} commInterface
   */
  exports.prototype.commInterface = undefined;

  return exports;

}));
