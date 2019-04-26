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
    define(['ApiClient', 'model/NetworkLocation'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./NetworkLocation'));
  } else {
    // Browser globals (root is window)
    if (!root.MeepControllerRestApi) {
      root.MeepControllerRestApi = {};
    }
    root.MeepControllerRestApi.Zone = factory(root.MeepControllerRestApi.ApiClient, root.MeepControllerRestApi.NetworkLocation);
  }
}(this, function(ApiClient, NetworkLocation) {
  'use strict';




  /**
   * The Zone model module.
   * @module model/Zone
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>Zone</code>.
   * Logical zone (MEC network) object
   * @alias module:model/Zone
   * @class
   */
  var exports = function() {
    var _this = this;

















  };

  /**
   * Constructs a <code>Zone</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/Zone} obj Optional instance to populate.
   * @return {module:model/Zone} The populated <code>Zone</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('id')) {
        obj['id'] = ApiClient.convertToType(data['id'], 'String');
      }
      if (data.hasOwnProperty('name')) {
        obj['name'] = ApiClient.convertToType(data['name'], 'String');
      }
      if (data.hasOwnProperty('type')) {
        obj['type'] = ApiClient.convertToType(data['type'], 'String');
      }
      if (data.hasOwnProperty('interFogLatency')) {
        obj['interFogLatency'] = ApiClient.convertToType(data['interFogLatency'], 'Number');
      }
      if (data.hasOwnProperty('interFogLatencyVariation')) {
        obj['interFogLatencyVariation'] = ApiClient.convertToType(data['interFogLatencyVariation'], 'Number');
      }
      if (data.hasOwnProperty('interFogThroughput')) {
        obj['interFogThroughput'] = ApiClient.convertToType(data['interFogThroughput'], 'Number');
      }
      if (data.hasOwnProperty('interFogPacketLoss')) {
        obj['interFogPacketLoss'] = ApiClient.convertToType(data['interFogPacketLoss'], 'Number');
      }
      if (data.hasOwnProperty('interEdgeLatency')) {
        obj['interEdgeLatency'] = ApiClient.convertToType(data['interEdgeLatency'], 'Number');
      }
      if (data.hasOwnProperty('interEdgeLatencyVariation')) {
        obj['interEdgeLatencyVariation'] = ApiClient.convertToType(data['interEdgeLatencyVariation'], 'Number');
      }
      if (data.hasOwnProperty('interEdgeThroughput')) {
        obj['interEdgeThroughput'] = ApiClient.convertToType(data['interEdgeThroughput'], 'Number');
      }
      if (data.hasOwnProperty('interEdgePacketLoss')) {
        obj['interEdgePacketLoss'] = ApiClient.convertToType(data['interEdgePacketLoss'], 'Number');
      }
      if (data.hasOwnProperty('edgeFogLatency')) {
        obj['edgeFogLatency'] = ApiClient.convertToType(data['edgeFogLatency'], 'Number');
      }
      if (data.hasOwnProperty('edgeFogLatencyVariation')) {
        obj['edgeFogLatencyVariation'] = ApiClient.convertToType(data['edgeFogLatencyVariation'], 'Number');
      }
      if (data.hasOwnProperty('edgeFogThroughput')) {
        obj['edgeFogThroughput'] = ApiClient.convertToType(data['edgeFogThroughput'], 'Number');
      }
      if (data.hasOwnProperty('edgeFogPacketLoss')) {
        obj['edgeFogPacketLoss'] = ApiClient.convertToType(data['edgeFogPacketLoss'], 'Number');
      }
      if (data.hasOwnProperty('networkLocations')) {
        obj['networkLocations'] = ApiClient.convertToType(data['networkLocations'], [NetworkLocation]);
      }
    }
    return obj;
  }

  /**
   * Unique zone ID
   * @member {String} id
   */
  exports.prototype['id'] = undefined;
  /**
   * Zone name
   * @member {String} name
   */
  exports.prototype['name'] = undefined;
  /**
   * Zone type
   * @member {module:model/Zone.TypeEnum} type
   */
  exports.prototype['type'] = undefined;
  /**
   * Latency in ms between fog nodes (or PoAs) within zone
   * @member {Number} interFogLatency
   */
  exports.prototype['interFogLatency'] = undefined;
  /**
   * Latency variation in ms between fog nodes (or PoAs) within zone
   * @member {Number} interFogLatencyVariation
   */
  exports.prototype['interFogLatencyVariation'] = undefined;
  /**
   * The limit of the traffic supported between fog nodes (or PoAs) within the zone
   * @member {Number} interFogThroughput
   */
  exports.prototype['interFogThroughput'] = undefined;
  /**
   * Packet lost (in terms of percentage) between fog nodes (or PoAs) within the zone
   * @member {Number} interFogPacketLoss
   */
  exports.prototype['interFogPacketLoss'] = undefined;
  /**
   * Latency in ms between edge nodes within zone
   * @member {Number} interEdgeLatency
   */
  exports.prototype['interEdgeLatency'] = undefined;
  /**
   * Latency variation in ms between edge nodes within zone
   * @member {Number} interEdgeLatencyVariation
   */
  exports.prototype['interEdgeLatencyVariation'] = undefined;
  /**
   * The limit of the traffic supported between edge nodes within the zone
   * @member {Number} interEdgeThroughput
   */
  exports.prototype['interEdgeThroughput'] = undefined;
  /**
   * Packet lost (in terms of percentage) between edge nodes within the zone
   * @member {Number} interEdgePacketLoss
   */
  exports.prototype['interEdgePacketLoss'] = undefined;
  /**
   * Latency in ms between fog nodes (or PoAs) and edge nodes within zone
   * @member {Number} edgeFogLatency
   */
  exports.prototype['edgeFogLatency'] = undefined;
  /**
   * Latency variation in ms between fog nodes (or PoAs) and edge nodes within zone
   * @member {Number} edgeFogLatencyVariation
   */
  exports.prototype['edgeFogLatencyVariation'] = undefined;
  /**
   * The limit of the traffic supported between fog nodes (or PoAs) and edge nodes within the zone
   * @member {Number} edgeFogThroughput
   */
  exports.prototype['edgeFogThroughput'] = undefined;
  /**
   * Packet lost (in terms of percentage) between fog nodes (or PoAs) and edge nodes within the zone
   * @member {Number} edgeFogPacketLoss
   */
  exports.prototype['edgeFogPacketLoss'] = undefined;
  /**
   * @member {Array.<module:model/NetworkLocation>} networkLocations
   */
  exports.prototype['networkLocations'] = undefined;


  /**
   * Allowed values for the <code>type</code> property.
   * @enum {String}
   * @readonly
   */
  exports.TypeEnum = {
    /**
     * value: "ZONE"
     * @const
     */
    "ZONE": "ZONE",
    /**
     * value: "COMMON"
     * @const
     */
    "COMMON": "COMMON"  };


  return exports;
}));


