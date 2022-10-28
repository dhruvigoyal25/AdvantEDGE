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
    root.MecDemo3Api.TimeStamp1 = factory(root.MecDemo3Api.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';

  /**
   * The TimeStamp1 model module.
   * @module model/TimeStamp1
   * @version 0.0.1
   */

  /**
   * Constructs a new <code>TimeStamp1</code>.
   * &#39;This data type represents the time stamp as Unix-time since January 1, 1970, 00:00:00 UTC&#39;
   * @alias module:model/TimeStamp1
   * @class
   * @param seconds {Number} 'The seconds part of the Time. Time is defined as Unix-time since January 1, 1970, 00:00:00 UTC.'
   * @param nanoSeconds {Number} 'The nanoseconds part of the Time. Time is defined as Unix-time since January 1, 1970, 00:00:00 UTC.'
   */
  var exports = function(seconds, nanoSeconds) {
    this.seconds = seconds;
    this.nanoSeconds = nanoSeconds;
  };

  /**
   * Constructs a <code>TimeStamp1</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/TimeStamp1} obj Optional instance to populate.
   * @return {module:model/TimeStamp1} The populated <code>TimeStamp1</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      if (data.hasOwnProperty('seconds'))
        obj.seconds = ApiClient.convertToType(data['seconds'], 'Number');
      if (data.hasOwnProperty('nanoSeconds'))
        obj.nanoSeconds = ApiClient.convertToType(data['nanoSeconds'], 'Number');
    }
    return obj;
  }

  /**
   * 'The seconds part of the Time. Time is defined as Unix-time since January 1, 1970, 00:00:00 UTC.'
   * @member {Number} seconds
   */
  exports.prototype.seconds = undefined;

  /**
   * 'The nanoseconds part of the Time. Time is defined as Unix-time since January 1, 1970, 00:00:00 UTC.'
   * @member {Number} nanoSeconds
   */
  exports.prototype.nanoSeconds = undefined;

  return exports;

}));
