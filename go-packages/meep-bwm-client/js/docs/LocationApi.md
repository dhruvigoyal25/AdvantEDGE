# AdvantEdgeBandwidthManagementApi.LocationApi

All URIs are relative to *https://localhost/sandboxname/bwm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**mec011AppTerminationPOST**](LocationApi.md#mec011AppTerminationPOST) | **POST** /notifications/mec011/appTermination | MEC011 Application Termination notification for self termination


<a name="mec011AppTerminationPOST"></a>
# **mec011AppTerminationPOST**
> mec011AppTerminationPOST(body)

MEC011 Application Termination notification for self termination

Terminates itself.

### Example
```javascript
var AdvantEdgeBandwidthManagementApi = require('advant_edge_bandwidth_management_api');

var apiInstance = new AdvantEdgeBandwidthManagementApi.LocationApi();

var body = new AdvantEdgeBandwidthManagementApi.AppTerminationNotification(); // AppTerminationNotification | Termination notification details


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
apiInstance.mec011AppTerminationPOST(body, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AppTerminationNotification**](AppTerminationNotification.md)| Termination notification details | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

