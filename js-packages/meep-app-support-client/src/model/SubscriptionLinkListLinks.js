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
    define(['ApiClient', 'model/LinkType', 'model/SubscriptionLinkListLinksSubscriptions'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./LinkType'), require('./SubscriptionLinkListLinksSubscriptions'));
  } else {
    // Browser globals (root is window)
    if (!root.AdvantEdgeMecApplicationSupportApi) {
      root.AdvantEdgeMecApplicationSupportApi = {};
    }
    root.AdvantEdgeMecApplicationSupportApi.SubscriptionLinkListLinks = factory(root.AdvantEdgeMecApplicationSupportApi.ApiClient, root.AdvantEdgeMecApplicationSupportApi.LinkType, root.AdvantEdgeMecApplicationSupportApi.SubscriptionLinkListLinksSubscriptions);
  }
}(this, function(ApiClient, LinkType, SubscriptionLinkListLinksSubscriptions) {
  'use strict';

  /**
   * The SubscriptionLinkListLinks model module.
   * @module model/SubscriptionLinkListLinks
   * @version 2.1.1
   */

  /**
   * Constructs a new <code>SubscriptionLinkListLinks</code>.
   * Self-referring URI.
   * @alias module:model/SubscriptionLinkListLinks
   * @class
   * @param self {module:model/LinkType} 
   */
  var exports = function(self) {
    this.self = self;
  };

  /**
   * Constructs a <code>SubscriptionLinkListLinks</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/SubscriptionLinkListLinks} obj Optional instance to populate.
   * @return {module:model/SubscriptionLinkListLinks} The populated <code>SubscriptionLinkListLinks</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      if (data.hasOwnProperty('self'))
        obj.self = LinkType.constructFromObject(data['self']);
      if (data.hasOwnProperty('subscriptions'))
        obj.subscriptions = ApiClient.convertToType(data['subscriptions'], [SubscriptionLinkListLinksSubscriptions]);
    }
    return obj;
  }

  /**
   * @member {module:model/LinkType} self
   */
  exports.prototype.self = undefined;

  /**
   * The MEC application instance's subscriptions
   * @member {Array.<module:model/SubscriptionLinkListLinksSubscriptions>} subscriptions
   */
  exports.prototype.subscriptions = undefined;

  return exports;

}));
