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
    define(['ApiClient', 'model/AppTerminationNotification', 'model/ApplicationContextState', 'model/InlineNotification', 'model/ProblemDetails', 'model/ServiceAvailabilityNotification'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('../model/AppTerminationNotification'), require('../model/ApplicationContextState'), require('../model/InlineNotification'), require('../model/ProblemDetails'), require('../model/ServiceAvailabilityNotification'));
  } else {
    // Browser globals (root is window)
    if (!root.MecDemo3Api) {
      root.MecDemo3Api = {};
    }
    root.MecDemo3Api.NotificationApi = factory(root.MecDemo3Api.ApiClient, root.MecDemo3Api.AppTerminationNotification, root.MecDemo3Api.ApplicationContextState, root.MecDemo3Api.InlineNotification, root.MecDemo3Api.ProblemDetails, root.MecDemo3Api.ServiceAvailabilityNotification);
  }
}(this, function(ApiClient, AppTerminationNotification, ApplicationContextState, InlineNotification, ProblemDetails, ServiceAvailabilityNotification) {
  'use strict';

  /**
   * Notification service.
   * @module api/NotificationApi
   * @version 0.0.1
   */

  /**
   * Constructs a new NotificationApi. 
   * @alias module:api/NotificationApi
   * @class
   * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
   * default to {@link module:ApiClient#instance} if unspecified.
   */
  var exports = function(apiClient) {
    this.apiClient = apiClient || ApiClient.instance;


    /**
     * Callback function to receive the result of the amsNotificationCallback operation.
     * @callback module:api/NotificationApi~amsNotificationCallbackCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Callback endpoint for AMS Notifications
     * Callback endpoint for AMS Notifications
     * @param {module:model/InlineNotification} body 
     * @param {module:api/NotificationApi~amsNotificationCallbackCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.amsNotificationCallback = function(body, callback) {
      var postBody = body;

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling amsNotificationCallback");
      }


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
        '/services/callback/amsevent', 'POST',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the appTerminationNotificationCallback operation.
     * @callback module:api/NotificationApi~appTerminationNotificationCallbackCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Represents the information that the MEP notifies the subscribed application instance about the corresponding application instance termination/stop&#39;
     * @param {Object} opts Optional parameters
     * @param {module:model/AppTerminationNotification} opts.body 
     * @param {module:api/NotificationApi~appTerminationNotificationCallbackCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.appTerminationNotificationCallback = function(opts, callback) {
      opts = opts || {};
      var postBody = opts['body'];


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
      var accepts = [];
      var returnType = null;

      return this.apiClient.callApi(
        '/application/termination', 'POST',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the contextTransferNotificationCallback operation.
     * @callback module:api/NotificationApi~contextTransferNotificationCallbackCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Callback endpoint for MEC021 context-state transfer notification
     * Callback endpoint for MEC021 context-state transfer notification
     * @param {module:model/ApplicationContextState} body app termination notification details
     * @param {module:api/NotificationApi~contextTransferNotificationCallbackCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.contextTransferNotificationCallback = function(body, callback) {
      var postBody = body;

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling contextTransferNotificationCallback");
      }


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
      var accepts = [];
      var returnType = null;

      return this.apiClient.callApi(
        '/application/transfer', 'POST',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the serviceAvailNotificationCallback operation.
     * @callback module:api/NotificationApi~serviceAvailNotificationCallbackCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Callback endpoint for MEC011 Notifications
     * Callback endpoint for MEC011 Notifications
     * @param {Object} opts Optional parameters
     * @param {module:model/ServiceAvailabilityNotification} opts.body 
     * @param {module:api/NotificationApi~serviceAvailNotificationCallbackCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.serviceAvailNotificationCallback = function(opts, callback) {
      opts = opts || {};
      var postBody = opts['body'];


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
      var accepts = [];
      var returnType = null;

      return this.apiClient.callApi(
        '/services/callback/service-availability', 'POST',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
  };

  return exports;
}));
