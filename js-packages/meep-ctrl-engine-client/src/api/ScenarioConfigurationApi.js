/*
 * Copyright (c) 2019  InterDigital Communications, Inc
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
 * AdvantEDGE Platform Controller REST API
 * This API is the main platform API and mainly used by the AdvantEDGE frontend to interact with scenarios <p>**Micro-service**<br>[meep-ctrl-engine](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-ctrl-engine) <p>**Type & Usage**<br>Platform main interface used by controller software that want to interact with the AdvantEDGE platform <p>**Details**<br>API details available at _your-AdvantEDGE-ip-address:30000/api_ <p>**Default Port**<br>`30000` 
 *
 * OpenAPI spec version: 1.0.0
 * Contact: AdvantEDGE@InterDigital.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.9
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/Scenario', 'model/ScenarioList'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('../model/Scenario'), require('../model/ScenarioList'));
  } else {
    // Browser globals (root is window)
    if (!root.AdvantEdgePlatformControllerRestApi) {
      root.AdvantEdgePlatformControllerRestApi = {};
    }
    root.AdvantEdgePlatformControllerRestApi.ScenarioConfigurationApi = factory(root.AdvantEdgePlatformControllerRestApi.ApiClient, root.AdvantEdgePlatformControllerRestApi.Scenario, root.AdvantEdgePlatformControllerRestApi.ScenarioList);
  }
}(this, function(ApiClient, Scenario, ScenarioList) {
  'use strict';

  /**
   * ScenarioConfiguration service.
   * @module api/ScenarioConfigurationApi
   * @version 1.0.0
   */

  /**
   * Constructs a new ScenarioConfigurationApi. 
   * @alias module:api/ScenarioConfigurationApi
   * @class
   * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
   * default to {@link module:ApiClient#instance} if unspecified.
   */
  var exports = function(apiClient) {
    this.apiClient = apiClient || ApiClient.instance;


    /**
     * Callback function to receive the result of the createScenario operation.
     * @callback module:api/ScenarioConfigurationApi~createScenarioCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Add a scenario
     * Add a scenario to the platform scenario store
     * @param {String} name Scenario name
     * @param {module:model/Scenario} scenario Scenario to add to MEEP store
     * @param {module:api/ScenarioConfigurationApi~createScenarioCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.createScenario = function(name, scenario, callback) {
      var postBody = scenario;

      // verify the required parameter 'name' is set
      if (name === undefined || name === null) {
        throw new Error("Missing the required parameter 'name' when calling createScenario");
      }

      // verify the required parameter 'scenario' is set
      if (scenario === undefined || scenario === null) {
        throw new Error("Missing the required parameter 'scenario' when calling createScenario");
      }


      var pathParams = {
        'name': name
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
      var returnType = null;

      return this.apiClient.callApi(
        '/scenarios/{name}', 'POST',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the deleteScenario operation.
     * @callback module:api/ScenarioConfigurationApi~deleteScenarioCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete a scenario
     * Delete a scenario by name from the platform scenario store
     * @param {String} name Scenario name
     * @param {module:api/ScenarioConfigurationApi~deleteScenarioCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.deleteScenario = function(name, callback) {
      var postBody = null;

      // verify the required parameter 'name' is set
      if (name === undefined || name === null) {
        throw new Error("Missing the required parameter 'name' when calling deleteScenario");
      }


      var pathParams = {
        'name': name
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
      var returnType = null;

      return this.apiClient.callApi(
        '/scenarios/{name}', 'DELETE',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the deleteScenarioList operation.
     * @callback module:api/ScenarioConfigurationApi~deleteScenarioListCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete all scenarios
     * Delete all scenarios present in the platform scenario store
     * @param {module:api/ScenarioConfigurationApi~deleteScenarioListCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.deleteScenarioList = function(callback) {
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
      var returnType = null;

      return this.apiClient.callApi(
        '/scenarios', 'DELETE',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the getScenario operation.
     * @callback module:api/ScenarioConfigurationApi~getScenarioCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Scenario} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get a specific scenario
     * Get a scenario by name from the platform scenario store
     * @param {String} name Scenario name
     * @param {module:api/ScenarioConfigurationApi~getScenarioCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Scenario}
     */
    this.getScenario = function(name, callback) {
      var postBody = null;

      // verify the required parameter 'name' is set
      if (name === undefined || name === null) {
        throw new Error("Missing the required parameter 'name' when calling getScenario");
      }


      var pathParams = {
        'name': name
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
      var returnType = Scenario;

      return this.apiClient.callApi(
        '/scenarios/{name}', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the getScenarioList operation.
     * @callback module:api/ScenarioConfigurationApi~getScenarioListCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ScenarioList} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get all scenarios
     * Returns all scenarios from the platform scenario store
     * @param {module:api/ScenarioConfigurationApi~getScenarioListCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/ScenarioList}
     */
    this.getScenarioList = function(callback) {
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
      var returnType = ScenarioList;

      return this.apiClient.callApi(
        '/scenarios', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the setScenario operation.
     * @callback module:api/ScenarioConfigurationApi~setScenarioCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update a scenario
     * Update a scenario by name in the platform scenario store
     * @param {String} name Scenario name
     * @param {module:model/Scenario} scenario Scenario to add to MEEP store
     * @param {module:api/ScenarioConfigurationApi~setScenarioCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.setScenario = function(name, scenario, callback) {
      var postBody = scenario;

      // verify the required parameter 'name' is set
      if (name === undefined || name === null) {
        throw new Error("Missing the required parameter 'name' when calling setScenario");
      }

      // verify the required parameter 'scenario' is set
      if (scenario === undefined || scenario === null) {
        throw new Error("Missing the required parameter 'scenario' when calling setScenario");
      }


      var pathParams = {
        'name': name
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
      var returnType = null;

      return this.apiClient.callApi(
        '/scenarios/{name}', 'PUT',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
  };

  return exports;
}));
