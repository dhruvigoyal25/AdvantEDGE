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
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.MecDemo3Api);
  }
}(this, function(expect, MecDemo3Api) {
  'use strict';

  var instance;

  describe('(package)', function() {
    describe('Link', function() {
      beforeEach(function() {
        instance = new MecDemo3Api.Link();
      });

      it('should create an instance of Link', function() {
        // TODO: update the code to test Link
        expect(instance).to.be.a(MecDemo3Api.Link);
      });

      it('should have the property subscription (base name: "subscription")', function() {
        // TODO: update the code to test the property subscription
        expect(instance).to.have.property('subscription');
        // expect(instance.subscription).to.be(expectedValueLiteral);
      });

    });
  });

}));
