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
 * AdvantEDGE Sandbox Controller REST API
 * This API is the main Sandbox Controller API for scenario deployment & event injection <p>**Micro-service**<br>[meep-sandbox-ctrl](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-sandbox-ctrl) <p>**Type & Usage**<br>Platform runtime interface to manage active scenarios and inject events in AdvantEDGE platform <p>**Details**<br>API details available at _your-AdvantEDGE-ip-address/api_
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
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.AdvantEdgeSandboxControllerRestApi);
  }
}(this, function(expect, AdvantEdgeSandboxControllerRestApi) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new AdvantEdgeSandboxControllerRestApi.EventsApi();
  });

  describe('(package)', function() {
    describe('EventsApi', function() {
      describe('sendEvent', function() {
        it('should call sendEvent successfully', function(done) {
          // TODO: uncomment, update parameter values for sendEvent call
          /*
          var type = "type_example";
          var event = new AdvantEdgeSandboxControllerRestApi.Event();
          event.name = "";
          event.type = "MOBILITY";
          event.eventMobility = new AdvantEdgeSandboxControllerRestApi.EventMobility();
          event.eventMobility.elementName = "";
          event.eventMobility.dest = "";
          event.eventNetworkCharacteristicsUpdate = new AdvantEdgeSandboxControllerRestApi.EventNetworkCharacteristicsUpdate();
          event.eventNetworkCharacteristicsUpdate.elementName = "";
          event.eventNetworkCharacteristicsUpdate.elementType = "SCENARIO";
          event.eventNetworkCharacteristicsUpdate.latency = 0;
          event.eventNetworkCharacteristicsUpdate.latencyVariation = 0;
          event.eventNetworkCharacteristicsUpdate.throughput = 0;
          event.eventNetworkCharacteristicsUpdate.packetLoss = 0.0;
          event.eventPoasInRange = new AdvantEdgeSandboxControllerRestApi.EventPoasInRange();
          event.eventPoasInRange.ue = "";
          event.eventPoasInRange.poasInRange = [""];
          event.eventScenarioUpdate = new AdvantEdgeSandboxControllerRestApi.EventScenarioUpdate();
          event.eventScenarioUpdate.action = "ADD";
          event.eventScenarioUpdate.nodes = [new AdvantEdgeSandboxControllerRestApi.ScenarioNode()];
          event.eventScenarioUpdate.nodes[0].type = "UE";
          event.eventScenarioUpdate.nodes[0].nodeDataUnion = new AdvantEdgeSandboxControllerRestApi.NodeDataUnion();
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation = new AdvantEdgeSandboxControllerRestApi.PhysicalLocation();
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.id = "";
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.name = "";
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.type = "UE";
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.isExternal = false;
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.geoData = new AdvantEdgeSandboxControllerRestApi.GeoData();
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.geoData.location = new AdvantEdgeSandboxControllerRestApi.Point();
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.geoData.location.type = "Point";
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.geoData.location.coordinates = [];
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.geoData.radius = ;
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.geoData.path = new AdvantEdgeSandboxControllerRestApi.LineString();
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.geoData.path.type = "LineString";
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.geoData.path.coordinates = [[]];
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.networkLocationsInRange = [""];
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.meta = {key: ""};
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.userMeta = {key: ""};
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes = [new AdvantEdgeSandboxControllerRestApi.Process()];
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].id = "";
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].name = "";
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].type = "UE-APP";
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].isExternal = false;
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].image = "";
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].environment = "";
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].commandArguments = "";
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].commandExe = "";
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].serviceConfig = new AdvantEdgeSandboxControllerRestApi.ServiceConfig();
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].serviceConfig.name = "";
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].serviceConfig.meSvcName = "";
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].serviceConfig.ports = [new AdvantEdgeSandboxControllerRestApi.ServicePort()];
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].serviceConfig.ports[0].protocol = "";
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].serviceConfig.ports[0].port = 0;
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].serviceConfig.ports[0].externalPort = 0;
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].gpuConfig = new AdvantEdgeSandboxControllerRestApi.GpuConfig();
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].gpuConfig.type = "";
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].gpuConfig.count = 0;
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].externalConfig = new AdvantEdgeSandboxControllerRestApi.ExternalConfig();
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].externalConfig.ingressServiceMap = [new AdvantEdgeSandboxControllerRestApi.IngressService()];
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].externalConfig.ingressServiceMap[0].name = "";
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].externalConfig.ingressServiceMap[0].port = 0;
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].externalConfig.ingressServiceMap[0].externalPort = 0;
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].externalConfig.ingressServiceMap[0].protocol = "";
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].externalConfig.egressServiceMap = [new AdvantEdgeSandboxControllerRestApi.EgressService()];
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].externalConfig.egressServiceMap[0].name = "";
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].externalConfig.egressServiceMap[0].meSvcName = "";
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].externalConfig.egressServiceMap[0].ip = "";
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].externalConfig.egressServiceMap[0].port = 0;
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].externalConfig.egressServiceMap[0].protocol = "";
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].status = "";
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].userChartLocation = "";
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].userChartAlternateValues = "";
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].userChartGroup = "";
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].meta = {key: ""};
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].userMeta = {key: ""};
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].appLatency = 0;
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].appLatencyVariation = 0;
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].appThroughput = 0;
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].appPacketLoss = 0.0;
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.processes[0].placementId = "";
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.linkLatency = 0;
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.linkLatencyVariation = 0;
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.linkThroughput = 0;
          event.eventScenarioUpdate.nodes[0].nodeDataUnion.physicalLocation.linkPacketLoss = 0.0;
          event.eventScenarioUpdate.nodes[0].parent = "";
          event.eventScenarioUpdate.nodes[0].children = [""];
          event.eventOther = new AdvantEdgeSandboxControllerRestApi.EventOther();
          event.eventOther.event = "";

          instance.sendEvent(type, event, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
    });
  });

}));
