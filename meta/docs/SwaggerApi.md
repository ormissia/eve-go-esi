# {{classname}}

All URIs are relative to *https://esi.evetech.net/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDevSwagger**](SwaggerApi.md#GetDevSwagger) | **Get** /_dev/swagger.json | Get _dev spec (versioned)
[**GetDevSwagger_0**](SwaggerApi.md#GetDevSwagger_0) | **Get** /dev/swagger.json | Get dev swagger spec
[**GetLatestSwagger**](SwaggerApi.md#GetLatestSwagger) | **Get** /_latest/swagger.json | Get _latest spec (versioned)
[**GetLatestSwagger_0**](SwaggerApi.md#GetLatestSwagger_0) | **Get** /latest/swagger.json | Get latest swagger spec
[**GetLegacySwagger**](SwaggerApi.md#GetLegacySwagger) | **Get** /_legacy/swagger.json | Get _legacy spec (versioned)
[**GetLegacySwagger_0**](SwaggerApi.md#GetLegacySwagger_0) | **Get** /legacy/swagger.json | Get legacy swagger spec
[**GetMetaSwagger**](SwaggerApi.md#GetMetaSwagger) | **Get** /swagger.json | Get meta swagger spec
[**GetV1Swagger**](SwaggerApi.md#GetV1Swagger) | **Get** /v1/swagger.json | Get v1 swagger spec
[**GetV2Swagger**](SwaggerApi.md#GetV2Swagger) | **Get** /v2/swagger.json | Get v2 swagger spec
[**GetV3Swagger**](SwaggerApi.md#GetV3Swagger) | **Get** /v3/swagger.json | Get v3 swagger spec
[**GetV4Swagger**](SwaggerApi.md#GetV4Swagger) | **Get** /v4/swagger.json | Get v4 swagger spec
[**GetV5Swagger**](SwaggerApi.md#GetV5Swagger) | **Get** /v5/swagger.json | Get v5 swagger spec
[**GetV6Swagger**](SwaggerApi.md#GetV6Swagger) | **Get** /v6/swagger.json | Get v6 swagger spec

# **GetDevSwagger**
> interface{} GetDevSwagger(ctx, optional)
Get _dev spec (versioned)

The _dev ESI swagger specification. The underscore swagger specs do not have a basePath, giving you an easy way to generate stable clients on versioned routes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SwaggerApiGetDevSwaggerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SwaggerApiGetDevSwaggerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAgent** | **optional.String**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **optional.String**| Client identifier, takes precedence over User-Agent | 
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDevSwagger_0**
> interface{} GetDevSwagger_0(ctx, optional)
Get dev swagger spec

The dev ESI swagger specification.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SwaggerApiGetDevSwagger_1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SwaggerApiGetDevSwagger_1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAgent** | **optional.String**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **optional.String**| Client identifier, takes precedence over User-Agent | 
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestSwagger**
> interface{} GetLatestSwagger(ctx, optional)
Get _latest spec (versioned)

The _latest ESI swagger specification. The underscore swagger specs do not have a basePath, giving you an easy way to generate stable clients on versioned routes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SwaggerApiGetLatestSwaggerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SwaggerApiGetLatestSwaggerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAgent** | **optional.String**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **optional.String**| Client identifier, takes precedence over User-Agent | 
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestSwagger_0**
> interface{} GetLatestSwagger_0(ctx, optional)
Get latest swagger spec

The latest ESI swagger specification.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SwaggerApiGetLatestSwagger_2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SwaggerApiGetLatestSwagger_2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAgent** | **optional.String**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **optional.String**| Client identifier, takes precedence over User-Agent | 
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLegacySwagger**
> interface{} GetLegacySwagger(ctx, optional)
Get _legacy spec (versioned)

The _legacy ESI swagger specification. The underscore swagger specs do not have a basePath, giving you an easy way to generate stable clients on versioned routes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SwaggerApiGetLegacySwaggerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SwaggerApiGetLegacySwaggerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAgent** | **optional.String**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **optional.String**| Client identifier, takes precedence over User-Agent | 
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLegacySwagger_0**
> interface{} GetLegacySwagger_0(ctx, optional)
Get legacy swagger spec

The legacy ESI swagger specification.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SwaggerApiGetLegacySwagger_3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SwaggerApiGetLegacySwagger_3Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAgent** | **optional.String**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **optional.String**| Client identifier, takes precedence over User-Agent | 
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMetaSwagger**
> interface{} GetMetaSwagger(ctx, optional)
Get meta swagger spec

The meta ESI swagger specification.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SwaggerApiGetMetaSwaggerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SwaggerApiGetMetaSwaggerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAgent** | **optional.String**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **optional.String**| Client identifier, takes precedence over User-Agent | 
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1Swagger**
> interface{} GetV1Swagger(ctx, optional)
Get v1 swagger spec

The v1 ESI swagger specification.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SwaggerApiGetV1SwaggerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SwaggerApiGetV1SwaggerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAgent** | **optional.String**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **optional.String**| Client identifier, takes precedence over User-Agent | 
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV2Swagger**
> interface{} GetV2Swagger(ctx, optional)
Get v2 swagger spec

The v2 ESI swagger specification.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SwaggerApiGetV2SwaggerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SwaggerApiGetV2SwaggerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAgent** | **optional.String**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **optional.String**| Client identifier, takes precedence over User-Agent | 
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV3Swagger**
> interface{} GetV3Swagger(ctx, optional)
Get v3 swagger spec

The v3 ESI swagger specification.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SwaggerApiGetV3SwaggerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SwaggerApiGetV3SwaggerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAgent** | **optional.String**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **optional.String**| Client identifier, takes precedence over User-Agent | 
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV4Swagger**
> interface{} GetV4Swagger(ctx, optional)
Get v4 swagger spec

The v4 ESI swagger specification.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SwaggerApiGetV4SwaggerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SwaggerApiGetV4SwaggerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAgent** | **optional.String**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **optional.String**| Client identifier, takes precedence over User-Agent | 
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5Swagger**
> interface{} GetV5Swagger(ctx, optional)
Get v5 swagger spec

The v5 ESI swagger specification.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SwaggerApiGetV5SwaggerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SwaggerApiGetV5SwaggerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAgent** | **optional.String**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **optional.String**| Client identifier, takes precedence over User-Agent | 
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV6Swagger**
> interface{} GetV6Swagger(ctx, optional)
Get v6 swagger spec

The v6 ESI swagger specification.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SwaggerApiGetV6SwaggerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SwaggerApiGetV6SwaggerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAgent** | **optional.String**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **optional.String**| Client identifier, takes precedence over User-Agent | 
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

