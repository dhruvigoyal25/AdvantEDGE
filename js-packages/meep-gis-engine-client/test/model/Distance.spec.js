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
 * AdvantEDGE GIS Engine REST API
 * This API allows to control geo-spatial behavior and simulation. <p>**Micro-service**<br>[meep-gis-engine](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-gis-engine) <p>**Type & Usage**<br>Platform runtime interface to control geo-spatial behavior and simulation <p>**Details**<br>API details available at _your-AdvantEDGE-ip-address/api_
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
    factory(root.expect, root.AdvantEdgeGisEngineRestApi);
  }
}(this, function(expect, AdvantEdgeGisEngineRestApi) {
  'use strict';

  var instance;

  describe('(package)', function() {
    describe('Distance', function() {
      beforeEach(function() {
        instance = new AdvantEdgeGisEngineRestApi.Distance();
      });

      it('should create an instance of Distance', function() {
        // TODO: update the code to test Distance
        expect(instance).to.be.a(AdvantEdgeGisEngineRestApi.Distance);
      });

      it('should have the property distance (base name: "distance")', function() {
        // TODO: update the code to test the property distance
        expect(instance).to.have.property('distance');
        // expect(instance.distance).to.be(expectedValueLiteral);
      });

      it('should have the property srcLatitude (base name: "srcLatitude")', function() {
        // TODO: update the code to test the property srcLatitude
        expect(instance).to.have.property('srcLatitude');
        // expect(instance.srcLatitude).to.be(expectedValueLiteral);
      });

      it('should have the property srcLongitude (base name: "srcLongitude")', function() {
        // TODO: update the code to test the property srcLongitude
        expect(instance).to.have.property('srcLongitude');
        // expect(instance.srcLongitude).to.be(expectedValueLiteral);
      });

      it('should have the property dstLatitude (base name: "dstLatitude")', function() {
        // TODO: update the code to test the property dstLatitude
        expect(instance).to.have.property('dstLatitude');
        // expect(instance.dstLatitude).to.be(expectedValueLiteral);
      });

      it('should have the property dstLongitude (base name: "dstLongitude")', function() {
        // TODO: update the code to test the property dstLongitude
        expect(instance).to.have.property('dstLongitude');
        // expect(instance.dstLongitude).to.be(expectedValueLiteral);
      });

    });
  });

}));
