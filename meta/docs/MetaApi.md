# {{classname}}

All URIs are relative to *https://esi.evetech.net/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHeaders**](MetaApi.md#GetHeaders) | **Get** /headers/ | Debug request headers
[**GetPing**](MetaApi.md#GetPing) | **Get** /ping | Ping route
[**GetStatus**](MetaApi.md#GetStatus) | **Get** /status.json | ESI health status
[**GetVerify**](MetaApi.md#GetVerify) | **Get** /verify/ | Verify access token
[**GetVersions**](MetaApi.md#GetVersions) | **Get** /versions/ | List versions

# **GetHeaders**
> map[string]string GetHeaders(ctx, )
Debug request headers

Echo the request headers for debugging purposes. Note that the 'Connection' header and any 'X-' headers are not included

### Required Parameters
This endpoint does not need any parameter.

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPing**
> string GetPing(ctx, )
Ping route

Ping the ESI routers

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStatus**
> []GetStatusItem GetStatus(ctx, optional)
ESI health status

Provides a general health indicator per route and method

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MetaApiGetStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MetaApiGetStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **optional.String**| The version of metrics to request. Note that alternate versions are grouped together | [default to latest]

### Return type

[**[]GetStatusItem**](get_status_item.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVerify**
> GetVerifyOk GetVerify(ctx, optional)
Verify access token

Verify authorization tokens in ESI's auth cache

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MetaApiGetVerifyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MetaApiGetVerifyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAgent** | **optional.String**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **optional.String**| Client identifier, takes precedence over User-Agent | 
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **token** | **optional.String**| Access token to use if unable to set a header | 
 **authorization** | **optional.String**| Access token, in the format of \&quot;Bearer &lt;access token&gt;\&quot; | 

### Return type

[**GetVerifyOk**](get_verify_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVersions**
> []string GetVersions(ctx, optional)
List versions

List all endpoint versions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MetaApiGetVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MetaApiGetVersionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAgent** | **optional.String**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **optional.String**| Client identifier, takes precedence over User-Agent | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

