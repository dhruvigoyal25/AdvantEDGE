# AdvantEdgeBandwidthManagementApi.BwInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**allocationId** | **String** | Bandwidth allocation instance identifier | [optional] 
**allocationDirection** | **String** | The direction of the requested BW allocation: 00 &#x3D; Downlink (towards the UE) 01 &#x3D; Uplink (towards the application/session) 10 &#x3D; Symmetrical | 
**appInsId** | **String** | Application instance identifier | 
**appName** | **String** | Name of the application | [optional] 
**fixedAllocation** | **String** | Size of requested fixed BW allocation in [bps] | 
**fixedBWPriority** | **String** | Indicates the allocation priority when dealing with several applications or sessions in parallel. Values are not defined in the present document | [optional] 
**requestType** | **Number** | Numeric value (0 - 255) corresponding to specific type of consumer as following: 0 &#x3D; APPLICATION_SPECIFIC_BW_ALLOCATION 1 &#x3D; SESSION_SPECIFIC_BW_ALLOCATION | 
**sessionFilter** | [**[BwInfoSessionFilter]**](BwInfoSessionFilter.md) | Session filtering criteria, applicable when requestType is set as SESSION_SPECIFIC_BW_ALLOCATION. Any filtering criteria shall define a single session only. In case multiple sessions match sessionFilter the request shall be rejected | [optional] 
**timeStamp** | [**BwInfoTimeStamp**](BwInfoTimeStamp.md) |  | [optional] 


<a name="FixedBWPriorityEnum"></a>
## Enum: FixedBWPriorityEnum


* `SEE_DESCRIPTION` (value: `"SEE_DESCRIPTION"`)




<a name="RequestTypeEnum"></a>
## Enum: RequestTypeEnum


* `_0` (value: `0`)

* `_1` (value: `1`)




