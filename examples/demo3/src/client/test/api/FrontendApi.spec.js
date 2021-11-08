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
 * This section describes use-case for demo 3 that the user can accomplish using the MEC Sandbox APIs from a MEC application
 *
 * OpenAPI spec version: 0.0.1
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
    factory(root.expect, root.MecDemo3Api);
  }
}(this, function(expect, MecDemo3Api) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new MecDemo3Api.FrontendApi();
  });

  describe('(package)', function() {
    describe('FrontendApi', function() {
      describe('infoAmsLogsGet', function() {
        it('should call infoAmsLogsGet successfully', function(done) {
          // TODO: uncomment, update parameter values for infoAmsLogsGet call
          /*
          var numLogs = null;

          instance.infoAmsLogsGet(numLogs, function(error, data, response) {
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
      describe('infoApplicationMecPlatformDeleteDelete', function() {
        it('should call infoApplicationMecPlatformDeleteDelete successfully', function(done) {
          // TODO: uncomment, update parameter values for infoApplicationMecPlatformDeleteDelete call
          /*
          var mecPlatform = null;

          instance.infoApplicationMecPlatformDeleteDelete(mecPlatform, function(error, data, response) {
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
      describe('infoApplicationMecPlatformGet', function() {
        it('should call infoApplicationMecPlatformGet successfully', function(done) {
          // TODO: uncomment, update parameter values for infoApplicationMecPlatformGet call
          /*
          var mecPlatform = null;

          instance.infoApplicationMecPlatformGet(mecPlatform, function(error, data, response) {
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
      describe('infoLogsGet', function() {
        it('should call infoLogsGet successfully', function(done) {
          // TODO: uncomment, update parameter values for infoLogsGet call
          /*
          var numLogs = null;

          instance.infoLogsGet(numLogs, function(error, data, response) {
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
      describe('registerAppMecPlatformPost', function() {
        it('should call registerAppMecPlatformPost successfully', function(done) {
          // TODO: uncomment, update parameter values for registerAppMecPlatformPost call
          /*
          var mecPlatform = null;

          instance.registerAppMecPlatformPost(mecPlatform, function(error, data, response) {
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
      describe('serviceAmsDeleteDeviceDelete', function() {
        it('should call serviceAmsDeleteDeviceDelete successfully', function(done) {
          // TODO: uncomment, update parameter values for serviceAmsDeleteDeviceDelete call
          /*
          var device = null;

          instance.serviceAmsDeleteDeviceDelete(device, function(error, data, response) {
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
      describe('serviceAmsUpdateDevicePut', function() {
        it('should call serviceAmsUpdateDevicePut successfully', function(done) {
          // TODO: uncomment, update parameter values for serviceAmsUpdateDevicePut call
          /*
          var device = null;

          instance.serviceAmsUpdateDevicePut(device, function(error, data, response) {
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
