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
 * AdvantEDGE Bandwidth Management API
 * Bandwidth Management Sercice is AdvantEDGE's implementation of [ETSI MEC ISG MEC015 Traffic Management APIs](https://www.etsi.org/deliver/etsi_gs/MEC/001_099/015/02.02.01_60/gs_MEC015v020201p.pdf) <p>[Copyright (c) ETSI 2017](https://forge.etsi.org/etsi-forge-copyright-notice.txt) <p>**Micro-service**<br>[meep-tm](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-tm/server/bwm) <p>**Type & Usage**<br>Edge Service used by edge applications that want to get information about BWM Info and Session(s) in the network <p>**Note**<br>AdvantEDGE supports all Bandwidth Management API endpoints.
 *
 * OpenAPI spec version: 2.2.1
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
    define(['ApiClient', 'model/LinkType'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./LinkType'));
  } else {
    // Browser globals (root is window)
    if (!root.AdvantEdgeBandwidthManagementApi) {
      root.AdvantEdgeBandwidthManagementApi = {};
    }
    root.AdvantEdgeBandwidthManagementApi.AppTerminationNotificationLinks = factory(root.AdvantEdgeBandwidthManagementApi.ApiClient, root.AdvantEdgeBandwidthManagementApi.LinkType);
  }
}(this, function(ApiClient, LinkType) {
  'use strict';

  /**
   * The AppTerminationNotificationLinks model module.
   * @module model/AppTerminationNotificationLinks
   * @version 2.2.1
   */

  /**
   * Constructs a new <code>AppTerminationNotificationLinks</code>.
   * Object containing hyperlinks related to the resource.
   * @alias module:model/AppTerminationNotificationLinks
   * @class
   * @param subscription {module:model/LinkType} 
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
        obj.subscription = LinkType.constructFromObject(data['subscription']);
      if (data.hasOwnProperty('confirmTermination'))
        obj.confirmTermination = LinkType.constructFromObject(data['confirmTermination']);
    }
    return obj;
  }

  /**
   * @member {module:model/LinkType} subscription
   */
  exports.prototype.subscription = undefined;

  /**
   * @member {module:model/LinkType} confirmTermination
   */
  exports.prototype.confirmTermination = undefined;

  return exports;

}));
