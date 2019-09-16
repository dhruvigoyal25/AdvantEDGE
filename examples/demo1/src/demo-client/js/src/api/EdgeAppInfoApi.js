/**
 * MEEP Demo App API
 * Copyright (c) 2019  InterDigital Communications, Inc Licensed under the Apache License, Version 2.0 (the \"License\"); you may not use this file except in compliance with the License. You may obtain a copy of the License at      http://www.apache.org/licenses/LICENSE-2.0  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an \"AS IS\" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License. 
 *
 * OpenAPI spec version: 0.0.1
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
    define(['ApiClient', 'model/EdgeInfo'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('../model/EdgeInfo'));
  } else {
    // Browser globals (root is window)
    if (!root.MeepDemoAppApi) {
      root.MeepDemoAppApi = {};
    }
    root.MeepDemoAppApi.EdgeAppInfoApi = factory(root.MeepDemoAppApi.ApiClient, root.MeepDemoAppApi.EdgeInfo);
  }
}(this, function(ApiClient, EdgeInfo) {
  'use strict';

  /**
   * EdgeAppInfo service.
   * @module api/EdgeAppInfoApi
   * @version 0.0.1
   */

  /**
   * Constructs a new EdgeAppInfoApi. 
   * @alias module:api/EdgeAppInfoApi
   * @class
   * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
   * default to {@link module:ApiClient#instance} if unspecified.
   */
  var exports = function(apiClient) {
    this.apiClient = apiClient || ApiClient.instance;


    /**
     * Callback function to receive the result of the getEdgeInfo operation.
     * @callback module:api/EdgeAppInfoApi~getEdgeInfoCallback
     * @param {String} error Error message, if any.
     * @param {module:model/EdgeInfo} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Retrieve edge add info
     * 
     * @param {module:api/EdgeAppInfoApi~getEdgeInfoCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/EdgeInfo}
     */
    this.getEdgeInfo = function(callback) {
      var postBody = null;


      var pathParams = {
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = [];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = EdgeInfo;

      return this.apiClient.callApi(
        '/edge-app', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
  };

  return exports;
}));
