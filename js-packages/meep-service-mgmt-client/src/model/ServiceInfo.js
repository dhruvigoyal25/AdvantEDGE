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
    define(['ApiClient', 'model/CategoryRef', 'model/LocalityType', 'model/SerInstanceId', 'model/SerName', 'model/SerializerType', 'model/ServiceState', 'model/TransportInfo'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./CategoryRef'), require('./LocalityType'), require('./SerInstanceId'), require('./SerName'), require('./SerializerType'), require('./ServiceState'), require('./TransportInfo'));
  } else {
    // Browser globals (root is window)
    if (!root.AdvantEdgeMecServiceManagementApi) {
      root.AdvantEdgeMecServiceManagementApi = {};
    }
    root.AdvantEdgeMecServiceManagementApi.ServiceInfo = factory(root.AdvantEdgeMecServiceManagementApi.ApiClient, root.AdvantEdgeMecServiceManagementApi.CategoryRef, root.AdvantEdgeMecServiceManagementApi.LocalityType, root.AdvantEdgeMecServiceManagementApi.SerInstanceId, root.AdvantEdgeMecServiceManagementApi.SerName, root.AdvantEdgeMecServiceManagementApi.SerializerType, root.AdvantEdgeMecServiceManagementApi.ServiceState, root.AdvantEdgeMecServiceManagementApi.TransportInfo);
  }
}(this, function(ApiClient, CategoryRef, LocalityType, SerInstanceId, SerName, SerializerType, ServiceState, TransportInfo) {
  'use strict';

  /**
   * The ServiceInfo model module.
   * @module model/ServiceInfo
   * @version 2.1.1
   */

  /**
   * Constructs a new <code>ServiceInfo</code>.
   * This type represents the general information of a MEC service.
   * @alias module:model/ServiceInfo
   * @class
   * @param serName {module:model/SerName} 
   * @param version {String} Service version
   * @param state {module:model/ServiceState} 
   * @param transportInfo {module:model/TransportInfo} 
   * @param serializer {module:model/SerializerType} 
   */
  var exports = function(serName, version, state, transportInfo, serializer) {
    this.serName = serName;
    this.version = version;
    this.state = state;
    this.transportInfo = transportInfo;
    this.serializer = serializer;
  };

  /**
   * Constructs a <code>ServiceInfo</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/ServiceInfo} obj Optional instance to populate.
   * @return {module:model/ServiceInfo} The populated <code>ServiceInfo</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      if (data.hasOwnProperty('serInstanceId'))
        obj.serInstanceId = SerInstanceId.constructFromObject(data['serInstanceId']);
      if (data.hasOwnProperty('serName'))
        obj.serName = SerName.constructFromObject(data['serName']);
      if (data.hasOwnProperty('serCategory'))
        obj.serCategory = CategoryRef.constructFromObject(data['serCategory']);
      if (data.hasOwnProperty('version'))
        obj.version = ApiClient.convertToType(data['version'], 'String');
      if (data.hasOwnProperty('state'))
        obj.state = ServiceState.constructFromObject(data['state']);
      if (data.hasOwnProperty('transportInfo'))
        obj.transportInfo = TransportInfo.constructFromObject(data['transportInfo']);
      if (data.hasOwnProperty('serializer'))
        obj.serializer = SerializerType.constructFromObject(data['serializer']);
      if (data.hasOwnProperty('scopeOfLocality'))
        obj.scopeOfLocality = LocalityType.constructFromObject(data['scopeOfLocality']);
      if (data.hasOwnProperty('consumedLocalOnly'))
        obj.consumedLocalOnly = ApiClient.convertToType(data['consumedLocalOnly'], 'Boolean');
      if (data.hasOwnProperty('isLocal'))
        obj.isLocal = ApiClient.convertToType(data['isLocal'], 'Boolean');
    }
    return obj;
  }

  /**
   * @member {module:model/SerInstanceId} serInstanceId
   */
  exports.prototype.serInstanceId = undefined;

  /**
   * @member {module:model/SerName} serName
   */
  exports.prototype.serName = undefined;

  /**
   * @member {module:model/CategoryRef} serCategory
   */
  exports.prototype.serCategory = undefined;

  /**
   * Service version
   * @member {String} version
   */
  exports.prototype.version = undefined;

  /**
   * @member {module:model/ServiceState} state
   */
  exports.prototype.state = undefined;

  /**
   * @member {module:model/TransportInfo} transportInfo
   */
  exports.prototype.transportInfo = undefined;

  /**
   * @member {module:model/SerializerType} serializer
   */
  exports.prototype.serializer = undefined;

  /**
   * @member {module:model/LocalityType} scopeOfLocality
   */
  exports.prototype.scopeOfLocality = undefined;

  /**
   * Indicate whether the service can only be consumed by the MEC applications located in the same locality (as defined by scopeOfLocality) as this  service instance.
   * @member {Boolean} consumedLocalOnly
   */
  exports.prototype.consumedLocalOnly = undefined;

  /**
   * Indicate whether the service is located in the same locality (as defined by scopeOfLocality) as the consuming MEC application.
   * @member {Boolean} isLocal
   */
  exports.prototype.isLocal = undefined;

  return exports;

}));
