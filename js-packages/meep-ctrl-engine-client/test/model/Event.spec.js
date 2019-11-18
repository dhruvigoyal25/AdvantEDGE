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
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.AdvantEdgePlatformControllerRestApi);
  }
}(this, function(expect, AdvantEdgePlatformControllerRestApi) {
  'use strict';

  var instance;

  describe('(package)', function() {
    describe('Event', function() {
      beforeEach(function() {
        instance = new AdvantEdgePlatformControllerRestApi.Event();
      });

      it('should create an instance of Event', function() {
        // TODO: update the code to test Event
        expect(instance).to.be.a(AdvantEdgePlatformControllerRestApi.Event);
      });

      it('should have the property name (base name: "name")', function() {
        // TODO: update the code to test the property name
        expect(instance).to.have.property('name');
        // expect(instance.name).to.be(expectedValueLiteral);
      });

      it('should have the property type (base name: "type")', function() {
        // TODO: update the code to test the property type
        expect(instance).to.have.property('type');
        // expect(instance.type).to.be(expectedValueLiteral);
      });

      it('should have the property eventNetworkCharacteristicsUpdate (base name: "eventNetworkCharacteristicsUpdate")', function() {
        // TODO: update the code to test the property eventNetworkCharacteristicsUpdate
        expect(instance).to.have.property('eventNetworkCharacteristicsUpdate');
        // expect(instance.eventNetworkCharacteristicsUpdate).to.be(expectedValueLiteral);
      });

      it('should have the property eventMobility (base name: "eventMobility")', function() {
        // TODO: update the code to test the property eventMobility
        expect(instance).to.have.property('eventMobility');
        // expect(instance.eventMobility).to.be(expectedValueLiteral);
      });

      it('should have the property eventPoasInRange (base name: "eventPoasInRange")', function() {
        // TODO: update the code to test the property eventPoasInRange
        expect(instance).to.have.property('eventPoasInRange');
        // expect(instance.eventPoasInRange).to.be(expectedValueLiteral);
      });

      it('should have the property eventOther (base name: "eventOther")', function() {
        // TODO: update the code to test the property eventOther
        expect(instance).to.have.property('eventOther');
        // expect(instance.eventOther).to.be(expectedValueLiteral);
      });

    });
  });

}));
