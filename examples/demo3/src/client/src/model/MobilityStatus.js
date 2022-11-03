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
    define(['ApiClient'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'));
  } else {
    // Browser globals (root is window)
    if (!root.MecDemo3Api) {
      root.MecDemo3Api = {};
    }
    root.MecDemo3Api.MobilityStatus = factory(root.MecDemo3Api.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';

  /**
   * Enum class MobilityStatus.
   * @enum {String}
   * @readonly
   */
  var exports = {
    /**
     * value: "INTERHOST_MOVEOUT_TRIGGERED"
     * @const
     */
    TRIGGERED: "INTERHOST_MOVEOUT_TRIGGERED",

    /**
     * value: "INTERHOST_MOVEOUT_COMPLETED"
     * @const
     */
    COMPLETED: "INTERHOST_MOVEOUT_COMPLETED",

    /**
     * value: "INTERHOST_MOVEOUT_FAILED"
     * @const
     */
    FAILED: "INTERHOST_MOVEOUT_FAILED"
  };

  /**
   * Returns a <code>MobilityStatus</code> enum value from a JavaScript object name.
   * @param {Object} data The plain JavaScript object containing the name of the enum value.
   * @return {module:model/MobilityStatus} The enum <code>MobilityStatus</code> value.
   */
  exports.constructFromObject = function(object) {
    return object;
  }

  return exports;
}));
