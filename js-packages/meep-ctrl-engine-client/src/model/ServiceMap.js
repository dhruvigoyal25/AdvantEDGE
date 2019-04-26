/**
 * MEEP Controller REST API
 * Copyright (c) 2019 InterDigital Communications, Inc. All rights reserved. The information provided herein is the proprietary and confidential information of InterDigital Communications, Inc. 
 *
 * OpenAPI spec version: 1.0.0
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.3.1
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
    if (!root.MeepControllerRestApi) {
      root.MeepControllerRestApi = {};
    }
    root.MeepControllerRestApi.ServiceMap = factory(root.MeepControllerRestApi.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';




  /**
   * The ServiceMap model module.
   * @module model/ServiceMap
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>ServiceMap</code>.
   * Mapping of exposed ports to internal or external services
   * @alias module:model/ServiceMap
   * @class
   */
  var exports = function() {
    var _this = this;






  };

  /**
   * Constructs a <code>ServiceMap</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/ServiceMap} obj Optional instance to populate.
   * @return {module:model/ServiceMap} The populated <code>ServiceMap</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('name')) {
        obj['name'] = ApiClient.convertToType(data['name'], 'String');
      }
      if (data.hasOwnProperty('ip')) {
        obj['ip'] = ApiClient.convertToType(data['ip'], 'String');
      }
      if (data.hasOwnProperty('port')) {
        obj['port'] = ApiClient.convertToType(data['port'], 'Number');
      }
      if (data.hasOwnProperty('externalPort')) {
        obj['externalPort'] = ApiClient.convertToType(data['externalPort'], 'Number');
      }
      if (data.hasOwnProperty('protocol')) {
        obj['protocol'] = ApiClient.convertToType(data['protocol'], 'String');
      }
    }
    return obj;
  }

  /**
   * Service name
   * @member {String} name
   */
  exports.prototype['name'] = undefined;
  /**
   * Service IP address for external service only (egress)   <li>N/A for internal services 
   * @member {String} ip
   */
  exports.prototype['ip'] = undefined;
  /**
   * Service port number
   * @member {Number} port
   */
  exports.prototype['port'] = undefined;
  /**
   * Port used to expose internal service only (ingress)   <li>Must be unique port in range (30000 - 32767)   <li>N/A for external services 
   * @member {Number} externalPort
   */
  exports.prototype['externalPort'] = undefined;
  /**
   * Protocol that the application is using (TCP or UDP)
   * @member {String} protocol
   */
  exports.prototype['protocol'] = undefined;



  return exports;
}));


