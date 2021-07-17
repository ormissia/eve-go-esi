# {{classname}}

All URIs are relative to *https://esi.evetech.net/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDiffVersionAVersionB**](WebUIApi.md#GetDiffVersionAVersionB) | **Get** /diff/{version_a}/{version_b}/ | Diff route
[**GetUi**](WebUIApi.md#GetUi) | **Get** /ui/ | SwaggerUI route (v3)

# **GetDiffVersionAVersionB**
> string GetDiffVersionAVersionB(ctx, versionA, versionB, optional)
Diff route

Diff two ESI specs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **versionA** | **string**| Swagger spec version to compare with | 
  **versionB** | **string**| Swagger spec version to compare against | 
 **optional** | ***WebUIApiGetDiffVersionAVersionBOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebUIApiGetDiffVersionAVersionBOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userAgent** | **optional.String**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **optional.String**| Client identifier, takes precedence over User-Agent | 
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/html; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUi**
> string GetUi(ctx, optional)
SwaggerUI route (v3)

ESI web UI. This is an open source project. If you find ESI web UI specific bugs, please report them to https://github.com/esi/esi-swagger-ui/issues

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WebUIApiGetUiOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebUIApiGetUiOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAgent** | **optional.String**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **optional.String**| Client identifier, takes precedence over User-Agent | 
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **version** | **optional.String**| The Swagger spec version to display | [default to latest]

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/html; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

