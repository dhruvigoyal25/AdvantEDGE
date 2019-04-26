/**
 * MEEP Controller REST API
 * Copyright (c) 2019 InterDigital Communications, Inc. All rights reserved. The information provided herein is the proprietary and confidential information of InterDigital Communications, Inc. 
 *
 * OpenAPI spec version: 1.0.0
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.3.1
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
    factory(root.expect, root.MeepControllerRestApi);
  }
}(this, function(expect, MeepControllerRestApi) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new MeepControllerRestApi.EventNetworkCharacteristicsUpdate();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('EventNetworkCharacteristicsUpdate', function() {
    it('should create an instance of EventNetworkCharacteristicsUpdate', function() {
      // uncomment below and update the code to test EventNetworkCharacteristicsUpdate
      //var instane = new MeepControllerRestApi.EventNetworkCharacteristicsUpdate();
      //expect(instance).to.be.a(MeepControllerRestApi.EventNetworkCharacteristicsUpdate);
    });

    it('should have the property elementName (base name: "elementName")', function() {
      // uncomment below and update the code to test the property elementName
      //var instane = new MeepControllerRestApi.EventNetworkCharacteristicsUpdate();
      //expect(instance).to.be();
    });

    it('should have the property elementType (base name: "elementType")', function() {
      // uncomment below and update the code to test the property elementType
      //var instane = new MeepControllerRestApi.EventNetworkCharacteristicsUpdate();
      //expect(instance).to.be();
    });

    it('should have the property latency (base name: "latency")', function() {
      // uncomment below and update the code to test the property latency
      //var instane = new MeepControllerRestApi.EventNetworkCharacteristicsUpdate();
      //expect(instance).to.be();
    });

    it('should have the property latencyVariation (base name: "latencyVariation")', function() {
      // uncomment below and update the code to test the property latencyVariation
      //var instane = new MeepControllerRestApi.EventNetworkCharacteristicsUpdate();
      //expect(instance).to.be();
    });

    it('should have the property throughput (base name: "throughput")', function() {
      // uncomment below and update the code to test the property throughput
      //var instane = new MeepControllerRestApi.EventNetworkCharacteristicsUpdate();
      //expect(instance).to.be();
    });

    it('should have the property packetLoss (base name: "packetLoss")', function() {
      // uncomment below and update the code to test the property packetLoss
      //var instane = new MeepControllerRestApi.EventNetworkCharacteristicsUpdate();
      //expect(instance).to.be();
    });

  });

}));
