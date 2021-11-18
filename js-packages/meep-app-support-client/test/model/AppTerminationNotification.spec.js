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
 * AdvantEDGE MEC Application Support API
 * MEC Application Support Service is AdvantEDGE's implementation of [ETSI MEC ISG MEC011 Application Enablement API](https://www.etsi.org/deliver/etsi_gs/MEC/001_099/011/02.01.01_60/gs_MEC011v020101p.pdf) <p>[Copyright (c) ETSI 2017](https://forge.etsi.org/etsi-forge-copyright-notice.txt) <p>**Micro-service**<br>[meep-app-enablement](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-app-enablement/server/app-support) <p>**Type & Usage**<br>Edge Service used by edge applications that want to get information about applications in the network <p>**Note**<br>AdvantEDGE supports a selected subset of Application Support API endpoints (see below).
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
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.AdvantEdgeMecApplicationSupportApi);
  }
}(this, function(expect, AdvantEdgeMecApplicationSupportApi) {
  'use strict';

  var instance;

  describe('(package)', function() {
    describe('AppTerminationNotification', function() {
      beforeEach(function() {
        instance = new AdvantEdgeMecApplicationSupportApi.AppTerminationNotification();
      });

      it('should create an instance of AppTerminationNotification', function() {
        // TODO: update the code to test AppTerminationNotification
        expect(instance).to.be.a(AdvantEdgeMecApplicationSupportApi.AppTerminationNotification);
      });

      it('should have the property notificationType (base name: "notificationType")', function() {
        // TODO: update the code to test the property notificationType
        expect(instance).to.have.property('notificationType');
        // expect(instance.notificationType).to.be(expectedValueLiteral);
      });

      it('should have the property operationAction (base name: "operationAction")', function() {
        // TODO: update the code to test the property operationAction
        expect(instance).to.have.property('operationAction');
        // expect(instance.operationAction).to.be(expectedValueLiteral);
      });

      it('should have the property maxGracefulTimeout (base name: "maxGracefulTimeout")', function() {
        // TODO: update the code to test the property maxGracefulTimeout
        expect(instance).to.have.property('maxGracefulTimeout');
        // expect(instance.maxGracefulTimeout).to.be(expectedValueLiteral);
      });

      it('should have the property links (base name: "_links")', function() {
        // TODO: update the code to test the property links
        expect(instance).to.have.property('links');
        // expect(instance.links).to.be(expectedValueLiteral);
      });

    });
  });

}));
