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
 * AdvantEDGE Bandwidth Management API
 * Bandwidth Management Sercice is AdvantEDGE's implementation of [ETSI MEC ISG MEC015 Traffic Management APIs](https://www.etsi.org/deliver/etsi_gs/MEC/001_099/015/02.02.01_60/gs_MEC015v020201p.pdf) <p>[Copyright (c) ETSI 2017](https://forge.etsi.org/etsi-forge-copyright-notice.txt) <p>**Micro-service**<br>[meep-tm](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-tm/server/bwm) <p>**Type & Usage**<br>Edge Service used by edge applications that want to get information about BWM Info and Session(s) in the network <p>**Note**<br>AdvantEDGE supports all Bandwidth Management API endpoints.
 *
 * OpenAPI spec version: 2.2.1
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
    factory(root.expect, root.AdvantEdgeBandwidthManagementApi);
  }
}(this, function(expect, AdvantEdgeBandwidthManagementApi) {
  'use strict';

  var instance;

  describe('(package)', function() {
    describe('OperationActionType', function() {
      beforeEach(function() {
        instance = AdvantEdgeBandwidthManagementApi.OperationActionType;
      });

      it('should create an instance of OperationActionType', function() {
        // TODO: update the code to test OperationActionType
        expect(instance).to.be.a('object');
      });

      it('should have the property STOPPING', function() {
        expect(instance).to.have.property('STOPPING');
        expect(instance.STOPPING).to.be(&quot;STOPPING&quot;);
      });

      it('should have the property TERMINATING', function() {
        expect(instance).to.have.property('TERMINATING');
        expect(instance.TERMINATING).to.be(&quot;TERMINATING&quot;);
      });

    });
  });

}));
