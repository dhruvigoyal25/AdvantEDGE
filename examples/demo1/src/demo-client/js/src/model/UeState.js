/**
 * MEEP Demo App API
 * This is the MEEP Demo App API
 *
 * OpenAPI spec version: 0.0.1
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.5
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
    if (!root.MeepDemoAppApi) {
      root.MeepDemoAppApi = {};
    }
    root.MeepDemoAppApi.UeState = factory(root.MeepDemoAppApi.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';




  /**
   * The UeState model module.
   * @module model/UeState
   * @version 0.0.1
   */

  /**
   * Constructs a new <code>UeState</code>.
   * Ue state basic information object
   * @alias module:model/UeState
   * @class
   */
  var exports = function() {
    var _this = this;



  };

  /**
   * Constructs a <code>UeState</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/UeState} obj Optional instance to populate.
   * @return {module:model/UeState} The populated <code>UeState</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('duration')) {
        obj['duration'] = ApiClient.convertToType(data['duration'], 'Number');
      }
      if (data.hasOwnProperty('trafficBw')) {
        obj['trafficBw'] = ApiClient.convertToType(data['trafficBw'], 'Number');
      }
    }
    return obj;
  }

  /**
   * Duration since the game stated
   * @member {Number} duration
   */
  exports.prototype['duration'] = undefined;
  /**
   * Traffic info for the registered Ue
   * @member {Number} trafficBw
   */
  exports.prototype['trafficBw'] = undefined;



  return exports;
}));


