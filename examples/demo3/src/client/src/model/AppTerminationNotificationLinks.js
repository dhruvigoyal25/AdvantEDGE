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
    define(['ApiClient', 'model/LinkType2', 'model/LinkTypeConfirmTermination'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./LinkType2'), require('./LinkTypeConfirmTermination'));
  } else {
    // Browser globals (root is window)
    if (!root.MecDemo3Api) {
      root.MecDemo3Api = {};
    }
    root.MecDemo3Api.AppTerminationNotificationLinks = factory(root.MecDemo3Api.ApiClient, root.MecDemo3Api.LinkType2, root.MecDemo3Api.LinkTypeConfirmTermination);
  }
}(this, function(ApiClient, LinkType2, LinkTypeConfirmTermination) {
  'use strict';

  /**
   * The AppTerminationNotificationLinks model module.
   * @module model/AppTerminationNotificationLinks
   * @version 0.0.1
   */

  /**
   * Constructs a new <code>AppTerminationNotificationLinks</code>.
   * Object containing hyperlinks related to the resource.
   * @alias module:model/AppTerminationNotificationLinks
   * @class
   * @param subscription {module:model/LinkType2} 
   */
  var exports = function(subscription) {
    this.subscription = subscription;
  };

  /**
   * Constructs a <code>AppTerminationNotificationLinks</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/AppTerminationNotificationLinks} obj Optional instance to populate.
   * @return {module:model/AppTerminationNotificationLinks} The populated <code>AppTerminationNotificationLinks</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      if (data.hasOwnProperty('subscription'))
        obj.subscription = LinkType2.constructFromObject(data['subscription']);
      if (data.hasOwnProperty('confirmTermination'))
        obj.confirmTermination = LinkTypeConfirmTermination.constructFromObject(data['confirmTermination']);
    }
    return obj;
  }

  /**
   * @member {module:model/LinkType2} subscription
   */
  exports.prototype.subscription = undefined;

  /**
   * @member {module:model/LinkTypeConfirmTermination} confirmTermination
   */
  exports.prototype.confirmTermination = undefined;

  return exports;

}));
