# AdvantEdgeBandwidthManagementApi.BwmApi

All URIs are relative to *https://localhost/sandboxname/bwm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**bandwidthAllocationDELETE**](BwmApi.md#bandwidthAllocationDELETE) | **DELETE** /bw_allocations/{allocationId} | Remove a specific bandwidthAllocation
[**bandwidthAllocationGET**](BwmApi.md#bandwidthAllocationGET) | **GET** /bw_allocations/{allocationId} | Retrieve information about a specific bandwidthAllocation
[**bandwidthAllocationListGET**](BwmApi.md#bandwidthAllocationListGET) | **GET** /bw_allocations | Retrieve information about a list of bandwidthAllocation resources
[**bandwidthAllocationPATCH**](BwmApi.md#bandwidthAllocationPATCH) | **PATCH** /bw_allocations/{allocationId} | Modify the information about a specific existing bandwidthAllocation by sending updates on the data structure
[**bandwidthAllocationPOST**](BwmApi.md#bandwidthAllocationPOST) | **POST** /bw_allocations | Create a bandwidthAllocation resource
[**bandwidthAllocationPUT**](BwmApi.md#bandwidthAllocationPUT) | **PUT** /bw_allocations/{allocationId} | Update the information about a specific bandwidthAllocation


<a name="bandwidthAllocationDELETE"></a>
# **bandwidthAllocationDELETE**
> bandwidthAllocationDELETE(allocationId)

Remove a specific bandwidthAllocation

Used in &#39;Unregister from Bandwidth Management Service&#39; procedure as described in clause 6.2.3.

### Example
```javascript
var AdvantEdgeBandwidthManagementApi = require('advant_edge_bandwidth_management_api');

var apiInstance = new AdvantEdgeBandwidthManagementApi.BwmApi();

var allocationId = "allocationId_example"; // String | Represents a bandwidth allocation instance


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
apiInstance.bandwidthAllocationDELETE(allocationId, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allocationId** | **String**| Represents a bandwidth allocation instance | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/problem+json

<a name="bandwidthAllocationGET"></a>
# **bandwidthAllocationGET**
> BwInfo bandwidthAllocationGET(allocationId)

Retrieve information about a specific bandwidthAllocation

Retrieves information about a bandwidthAllocation resource. Typically used in &#39;Get configured bandwidth allocation from Bandwidth Management Service&#39; procedure as described in clause 6.2.5.

### Example
```javascript
var AdvantEdgeBandwidthManagementApi = require('advant_edge_bandwidth_management_api');

var apiInstance = new AdvantEdgeBandwidthManagementApi.BwmApi();

var allocationId = "allocationId_example"; // String | Represents a bandwidth allocation instance


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.bandwidthAllocationGET(allocationId, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allocationId** | **String**| Represents a bandwidth allocation instance | 

### Return type

[**BwInfo**](BwInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

<a name="bandwidthAllocationListGET"></a>
# **bandwidthAllocationListGET**
> [BwInfo] bandwidthAllocationListGET(opts)

Retrieve information about a list of bandwidthAllocation resources

Retrieves information about a list of bandwidthAllocation resources. Typically used in &#39;Get configured bandwidth allocation from Bandwidth Management Service&#39; procedure as described in clause 6.2.5.

### Example
```javascript
var AdvantEdgeBandwidthManagementApi = require('advant_edge_bandwidth_management_api');

var apiInstance = new AdvantEdgeBandwidthManagementApi.BwmApi();

var opts = { 
  'appInstanceId': ["appInstanceId_example"], // [String] | A MEC application instance may use multiple app_instance_ids as an input parameter to query the bandwidth allocation of a list of MEC application instances. app_instance_id corresponds to appInsId defined in table 7.2.2-1. See note.
  'appName': ["appName_example"], // [String] | A MEC application instance may use multiple app_names as an input parameter to query the bandwidth allocation of a list of MEC application instances. app_name corresponds to appName defined in table 7.2.2-1. See note.
  'sessionId': ["sessionId_example"] // [String] | A MEC application instance may use session_id as an input parameter to query the bandwidth allocation of a list of sessions. session_id corresponds to allocationId defined in table 7.2.2-1. See note.
};

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.bandwidthAllocationListGET(opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appInstanceId** | [**[String]**](String.md)| A MEC application instance may use multiple app_instance_ids as an input parameter to query the bandwidth allocation of a list of MEC application instances. app_instance_id corresponds to appInsId defined in table 7.2.2-1. See note. | [optional] 
 **appName** | [**[String]**](String.md)| A MEC application instance may use multiple app_names as an input parameter to query the bandwidth allocation of a list of MEC application instances. app_name corresponds to appName defined in table 7.2.2-1. See note. | [optional] 
 **sessionId** | [**[String]**](String.md)| A MEC application instance may use session_id as an input parameter to query the bandwidth allocation of a list of sessions. session_id corresponds to allocationId defined in table 7.2.2-1. See note. | [optional] 

### Return type

[**[BwInfo]**](BwInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

<a name="bandwidthAllocationPATCH"></a>
# **bandwidthAllocationPATCH**
> BwInfo bandwidthAllocationPATCH(body, allocationId)

Modify the information about a specific existing bandwidthAllocation by sending updates on the data structure

Updates the information about a bandwidthAllocation resource. As specified in ETSI GS MEC 009 [6], the PATCH HTTP method updates a resource on top of the existing resource state by just including the changes (&#39;deltas&#39;) in the request body.

### Example
```javascript
var AdvantEdgeBandwidthManagementApi = require('advant_edge_bandwidth_management_api');

var apiInstance = new AdvantEdgeBandwidthManagementApi.BwmApi();

var body = new AdvantEdgeBandwidthManagementApi.BwInfoDeltas(); // BwInfoDeltas | Description of the changes to instruct the server how to modify the resource representation.

var allocationId = "allocationId_example"; // String | Represents a bandwidth allocation instance


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.bandwidthAllocationPATCH(body, allocationId, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**BwInfoDeltas**](BwInfoDeltas.md)| Description of the changes to instruct the server how to modify the resource representation. | 
 **allocationId** | **String**| Represents a bandwidth allocation instance | 

### Return type

[**BwInfo**](BwInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

<a name="bandwidthAllocationPOST"></a>
# **bandwidthAllocationPOST**
> BwInfo bandwidthAllocationPOST(body)

Create a bandwidthAllocation resource

Used to create a bandwidthAllocation resource. Typically used in &#39;Register to Bandwidth Management Service&#39; procedure as described in clause 6.2.1.

### Example
```javascript
var AdvantEdgeBandwidthManagementApi = require('advant_edge_bandwidth_management_api');

var apiInstance = new AdvantEdgeBandwidthManagementApi.BwmApi();

var body = new AdvantEdgeBandwidthManagementApi.BwInfo(); // BwInfo | Entity body in the request contains BwInfo to be created.


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.bandwidthAllocationPOST(body, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**BwInfo**](BwInfo.md)| Entity body in the request contains BwInfo to be created. | 

### Return type

[**BwInfo**](BwInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

<a name="bandwidthAllocationPUT"></a>
# **bandwidthAllocationPUT**
> BwInfo bandwidthAllocationPUT(body, allocationId)

Update the information about a specific bandwidthAllocation

Updates the information about a bandwidthAllocation resource. As specified in ETSI GS MEC 009 [6], the PUT HTTP method has &#39;replace&#39; semantics.

### Example
```javascript
var AdvantEdgeBandwidthManagementApi = require('advant_edge_bandwidth_management_api');

var apiInstance = new AdvantEdgeBandwidthManagementApi.BwmApi();

var body = new AdvantEdgeBandwidthManagementApi.BwInfo(); // BwInfo | BwInfo with updated information is included as entity body of the request.

var allocationId = "allocationId_example"; // String | Represents a bandwidth allocation instance


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.bandwidthAllocationPUT(body, allocationId, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**BwInfo**](BwInfo.md)| BwInfo with updated information is included as entity body of the request. | 
 **allocationId** | **String**| Represents a bandwidth allocation instance | 

### Return type

[**BwInfo**](BwInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

