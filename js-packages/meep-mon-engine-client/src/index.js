/*
 * Copyright (c) 2022  The AdvantEDGE Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * AdvantEDGE Monitoring Engine REST API
 * This API provides AdvantEDGE microservice & scenario deployment status information collected in the Monitoring Engine. <p>**Micro-service**<br>[meep-mon-engine](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-mon-engine) <p>**Type & Usage**<br>Platform interface to retrieve AdvantEDGE microservice & scenario deployment status information <p>**Details**<br>API details available at _your-AdvantEDGE-ip-address/api_
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

(function(factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/PodStatus', 'model/PodsStatus', 'api/PodStatesApi'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('./ApiClient'), require('./model/PodStatus'), require('./model/PodsStatus'), require('./api/PodStatesApi'));
  }
}(function(ApiClient, PodStatus, PodsStatus, PodStatesApi) {
  'use strict';

  /**
   * This_API_provides_AdvantEDGE_microservice__scenario_deployment_status_information_collected_in_the_Monitoring_Engine__pMicro_servicebr_meep_mon_engine_httpsgithub_comInterDigitalIncAdvantEDGEtreemastergo_appsmeep_mon_engine_pType__UsagebrPlatform_interface_to_retrieve_AdvantEDGE_microservice__scenario_deployment_status_information_pDetailsbrAPI_details_available_at__your_AdvantEDGE_ip_addressapi_.<br>
   * The <code>index</code> module provides access to constructors for all the classes which comprise the public API.
   * <p>
   * An AMD (recommended!) or CommonJS application will generally do something equivalent to the following:
   * <pre>
   * var AdvantEdgeMonitoringEngineRestApi = require('index'); // See note below*.
   * var xxxSvc = new AdvantEdgeMonitoringEngineRestApi.XxxApi(); // Allocate the API class we're going to use.
   * var yyyModel = new AdvantEdgeMonitoringEngineRestApi.Yyy(); // Construct a model instance.
   * yyyModel.someProperty = 'someValue';
   * ...
   * var zzz = xxxSvc.doSomething(yyyModel); // Invoke the service.
   * ...
   * </pre>
   * <em>*NOTE: For a top-level AMD script, use require(['index'], function(){...})
   * and put the application logic within the callback function.</em>
   * </p>
   * <p>
   * A non-AMD browser application (discouraged) might do something like this:
   * <pre>
   * var xxxSvc = new AdvantEdgeMonitoringEngineRestApi.XxxApi(); // Allocate the API class we're going to use.
   * var yyy = new AdvantEdgeMonitoringEngineRestApi.Yyy(); // Construct a model instance.
   * yyyModel.someProperty = 'someValue';
   * ...
   * var zzz = xxxSvc.doSomething(yyyModel); // Invoke the service.
   * ...
   * </pre>
   * </p>
   * @module index
   * @version 1.0.0
   */
  var exports = {
    /**
     * The ApiClient constructor.
     * @property {module:ApiClient}
     */
    ApiClient: ApiClient,
    /**
     * The PodStatus model constructor.
     * @property {module:model/PodStatus}
     */
    PodStatus: PodStatus,
    /**
     * The PodsStatus model constructor.
     * @property {module:model/PodsStatus}
     */
    PodsStatus: PodsStatus,
    /**
     * The PodStatesApi service constructor.
     * @property {module:api/PodStatesApi}
     */
    PodStatesApi: PodStatesApi
  };

  return exports;
}));
