/*
 * CLOUD API
 *
 *  IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	_context "context"
	"fmt"
	"io"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// TargetGroupsApiService TargetGroupsApi service
type TargetGroupsApiService service

type ApiTargetGroupsDeleteRequest struct {
	ctx             _context.Context
	ApiService      *TargetGroupsApiService
	targetGroupId   string
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiTargetGroupsDeleteRequest) Pretty(pretty bool) ApiTargetGroupsDeleteRequest {
	r.pretty = &pretty
	return r
}
func (r ApiTargetGroupsDeleteRequest) Depth(depth int32) ApiTargetGroupsDeleteRequest {
	r.depth = &depth
	return r
}
func (r ApiTargetGroupsDeleteRequest) XContractNumber(xContractNumber int32) ApiTargetGroupsDeleteRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiTargetGroupsDeleteRequest) Execute() (*APIResponse, error) {
	return r.ApiService.TargetGroupsDeleteExecute(r)
}

/*
 * TargetGroupsDelete Delete a Target Group by ID
 * Deletes the target group specified by its ID.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param targetGroupId The unique ID of the target group.
 * @return ApiTargetGroupsDeleteRequest
 */
func (a *TargetGroupsApiService) TargetGroupsDelete(ctx _context.Context, targetGroupId string) ApiTargetGroupsDeleteRequest {
	return ApiTargetGroupsDeleteRequest{
		ApiService:    a,
		ctx:           ctx,
		targetGroupId: targetGroupId,
	}
}

/*
 * Execute executes the request
 */
func (a *TargetGroupsApiService) TargetGroupsDeleteExecute(r ApiTargetGroupsDeleteRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetGroupsApiService.TargetGroupsDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/targetgroups/{targetGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"targetGroupId"+"}", _neturl.PathEscape(parameterToString(r.targetGroupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("pretty")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("pretty", parameterToString(true, ""))
		}
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("depth")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("depth", parameterToString(0, ""))
		}
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xContractNumber != nil {
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["TokenAuthentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestURL:  localVarPath,
		RequestTime: httpRequestTime,
		Operation:   "TargetGroupsDelete",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarAPIResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, string(localVarBody)),
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, err.Error())
			return localVarAPIResponse, newErr
		}
		newErr.model = v
		return localVarAPIResponse, newErr
	}

	return localVarAPIResponse, nil
}

type ApiTargetgroupsFindByTargetGroupIdRequest struct {
	ctx             _context.Context
	ApiService      *TargetGroupsApiService
	targetGroupId   string
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiTargetgroupsFindByTargetGroupIdRequest) Pretty(pretty bool) ApiTargetgroupsFindByTargetGroupIdRequest {
	r.pretty = &pretty
	return r
}
func (r ApiTargetgroupsFindByTargetGroupIdRequest) Depth(depth int32) ApiTargetgroupsFindByTargetGroupIdRequest {
	r.depth = &depth
	return r
}
func (r ApiTargetgroupsFindByTargetGroupIdRequest) XContractNumber(xContractNumber int32) ApiTargetgroupsFindByTargetGroupIdRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiTargetgroupsFindByTargetGroupIdRequest) Execute() (TargetGroup, *APIResponse, error) {
	return r.ApiService.TargetgroupsFindByTargetGroupIdExecute(r)
}

/*
 * TargetgroupsFindByTargetGroupId Get a Target Group by ID
 * Retrieves the properties of the target group specified by its ID.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param targetGroupId The unique ID of the target group.
 * @return ApiTargetgroupsFindByTargetGroupIdRequest
 */
func (a *TargetGroupsApiService) TargetgroupsFindByTargetGroupId(ctx _context.Context, targetGroupId string) ApiTargetgroupsFindByTargetGroupIdRequest {
	return ApiTargetgroupsFindByTargetGroupIdRequest{
		ApiService:    a,
		ctx:           ctx,
		targetGroupId: targetGroupId,
	}
}

/*
 * Execute executes the request
 * @return TargetGroup
 */
func (a *TargetGroupsApiService) TargetgroupsFindByTargetGroupIdExecute(r ApiTargetgroupsFindByTargetGroupIdRequest) (TargetGroup, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TargetGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetGroupsApiService.TargetgroupsFindByTargetGroupId")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/targetgroups/{targetGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"targetGroupId"+"}", _neturl.PathEscape(parameterToString(r.targetGroupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("pretty")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("pretty", parameterToString(true, ""))
		}
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("depth")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("depth", parameterToString(0, ""))
		}
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xContractNumber != nil {
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["TokenAuthentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestURL:  localVarPath,
		RequestTime: httpRequestTime,
		Operation:   "TargetgroupsFindByTargetGroupId",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, string(localVarBody)),
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, err.Error())
			return localVarReturnValue, localVarAPIResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      err.Error(),
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}

type ApiTargetgroupsGetRequest struct {
	ctx             _context.Context
	ApiService      *TargetGroupsApiService
	filters         _neturl.Values
	orderBy         *string
	maxResults      *int32
	pretty          *bool
	depth           *int32
	xContractNumber *int32
	offset          *int32
	limit           *int32
}

func (r ApiTargetgroupsGetRequest) Pretty(pretty bool) ApiTargetgroupsGetRequest {
	r.pretty = &pretty
	return r
}
func (r ApiTargetgroupsGetRequest) Depth(depth int32) ApiTargetgroupsGetRequest {
	r.depth = &depth
	return r
}
func (r ApiTargetgroupsGetRequest) XContractNumber(xContractNumber int32) ApiTargetgroupsGetRequest {
	r.xContractNumber = &xContractNumber
	return r
}
func (r ApiTargetgroupsGetRequest) Offset(offset int32) ApiTargetgroupsGetRequest {
	r.offset = &offset
	return r
}
func (r ApiTargetgroupsGetRequest) Limit(limit int32) ApiTargetgroupsGetRequest {
	r.limit = &limit
	return r
}

// Filters query parameters limit results to those containing a matching value for a specific property.
func (r ApiTargetgroupsGetRequest) Filter(key string, value string) ApiTargetgroupsGetRequest {
	filterKey := fmt.Sprintf(FilterQueryParam, key)
	r.filters[filterKey] = append(r.filters[filterKey], value)
	return r
}

// OrderBy query param sorts the results alphanumerically in ascending order based on the specified property.
func (r ApiTargetgroupsGetRequest) OrderBy(orderBy string) ApiTargetgroupsGetRequest {
	r.orderBy = &orderBy
	return r
}

// MaxResults query param limits the number of results returned.
func (r ApiTargetgroupsGetRequest) MaxResults(maxResults int32) ApiTargetgroupsGetRequest {
	r.maxResults = &maxResults
	return r
}

func (r ApiTargetgroupsGetRequest) Execute() (TargetGroups, *APIResponse, error) {
	return r.ApiService.TargetgroupsGetExecute(r)
}

/*
  - TargetgroupsGet Get Target Groups
  - Lists target groups.

A target group is a set of one or more registered targets. You must specify an IP address, a port number, and a weight for each target. Any object with an IP address in your VDC can be a target, for example, a VM, another load balancer, etc. You can register a target with multiple target groups.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @return ApiTargetgroupsGetRequest
*/
func (a *TargetGroupsApiService) TargetgroupsGet(ctx _context.Context) ApiTargetgroupsGetRequest {
	return ApiTargetgroupsGetRequest{
		ApiService: a,
		ctx:        ctx,
		filters:    _neturl.Values{},
	}
}

/*
 * Execute executes the request
 * @return TargetGroups
 */
func (a *TargetGroupsApiService) TargetgroupsGetExecute(r ApiTargetgroupsGetRequest) (TargetGroups, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TargetGroups
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetGroupsApiService.TargetgroupsGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/targetgroups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("pretty")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("pretty", parameterToString(true, ""))
		}
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("depth")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("depth", parameterToString(0, ""))
		}
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("offset")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("offset", parameterToString(0, ""))
		}
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("limit")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("limit", parameterToString(100, ""))
		}
	}
	if r.orderBy != nil {
		localVarQueryParams.Add("orderBy", parameterToString(*r.orderBy, ""))
	}
	if r.maxResults != nil {
		localVarQueryParams.Add("maxResults", parameterToString(*r.maxResults, ""))
	}
	if len(r.filters) > 0 {
		for k, v := range r.filters {
			for _, iv := range v {
				localVarQueryParams.Add(k, iv)
			}
		}
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xContractNumber != nil {
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["TokenAuthentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestURL:  localVarPath,
		RequestTime: httpRequestTime,
		Operation:   "TargetgroupsGet",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, string(localVarBody)),
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, err.Error())
			return localVarReturnValue, localVarAPIResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      err.Error(),
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}

type ApiTargetgroupsPatchRequest struct {
	ctx                   _context.Context
	ApiService            *TargetGroupsApiService
	targetGroupId         string
	targetGroupProperties *TargetGroupProperties
	pretty                *bool
	depth                 *int32
	xContractNumber       *int32
}

func (r ApiTargetgroupsPatchRequest) TargetGroupProperties(targetGroupProperties TargetGroupProperties) ApiTargetgroupsPatchRequest {
	r.targetGroupProperties = &targetGroupProperties
	return r
}
func (r ApiTargetgroupsPatchRequest) Pretty(pretty bool) ApiTargetgroupsPatchRequest {
	r.pretty = &pretty
	return r
}
func (r ApiTargetgroupsPatchRequest) Depth(depth int32) ApiTargetgroupsPatchRequest {
	r.depth = &depth
	return r
}
func (r ApiTargetgroupsPatchRequest) XContractNumber(xContractNumber int32) ApiTargetgroupsPatchRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiTargetgroupsPatchRequest) Execute() (TargetGroup, *APIResponse, error) {
	return r.ApiService.TargetgroupsPatchExecute(r)
}

/*
 * TargetgroupsPatch Partially Modify a Target Group by ID
 * Updates the properties of the target group specified by its ID.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param targetGroupId The unique ID of the target group.
 * @return ApiTargetgroupsPatchRequest
 */
func (a *TargetGroupsApiService) TargetgroupsPatch(ctx _context.Context, targetGroupId string) ApiTargetgroupsPatchRequest {
	return ApiTargetgroupsPatchRequest{
		ApiService:    a,
		ctx:           ctx,
		targetGroupId: targetGroupId,
	}
}

/*
 * Execute executes the request
 * @return TargetGroup
 */
func (a *TargetGroupsApiService) TargetgroupsPatchExecute(r ApiTargetgroupsPatchRequest) (TargetGroup, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TargetGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetGroupsApiService.TargetgroupsPatch")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/targetgroups/{targetGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"targetGroupId"+"}", _neturl.PathEscape(parameterToString(r.targetGroupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.targetGroupProperties == nil {
		return localVarReturnValue, nil, reportError("targetGroupProperties is required and must be specified")
	}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("pretty")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("pretty", parameterToString(true, ""))
		}
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("depth")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("depth", parameterToString(0, ""))
		}
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xContractNumber != nil {
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	// body params
	localVarPostBody = r.targetGroupProperties
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["TokenAuthentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestURL:  localVarPath,
		RequestTime: httpRequestTime,
		Operation:   "TargetgroupsPatch",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, string(localVarBody)),
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, err.Error())
			return localVarReturnValue, localVarAPIResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      err.Error(),
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}

type ApiTargetgroupsPostRequest struct {
	ctx             _context.Context
	ApiService      *TargetGroupsApiService
	targetGroup     *TargetGroup
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiTargetgroupsPostRequest) TargetGroup(targetGroup TargetGroup) ApiTargetgroupsPostRequest {
	r.targetGroup = &targetGroup
	return r
}
func (r ApiTargetgroupsPostRequest) Pretty(pretty bool) ApiTargetgroupsPostRequest {
	r.pretty = &pretty
	return r
}
func (r ApiTargetgroupsPostRequest) Depth(depth int32) ApiTargetgroupsPostRequest {
	r.depth = &depth
	return r
}
func (r ApiTargetgroupsPostRequest) XContractNumber(xContractNumber int32) ApiTargetgroupsPostRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiTargetgroupsPostRequest) Execute() (TargetGroup, *APIResponse, error) {
	return r.ApiService.TargetgroupsPostExecute(r)
}

/*
 * TargetgroupsPost Create a Target Group
 * Creates a target group.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiTargetgroupsPostRequest
 */
func (a *TargetGroupsApiService) TargetgroupsPost(ctx _context.Context) ApiTargetgroupsPostRequest {
	return ApiTargetgroupsPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return TargetGroup
 */
func (a *TargetGroupsApiService) TargetgroupsPostExecute(r ApiTargetgroupsPostRequest) (TargetGroup, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TargetGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetGroupsApiService.TargetgroupsPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/targetgroups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.targetGroup == nil {
		return localVarReturnValue, nil, reportError("targetGroup is required and must be specified")
	}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("pretty")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("pretty", parameterToString(true, ""))
		}
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("depth")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("depth", parameterToString(0, ""))
		}
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xContractNumber != nil {
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	// body params
	localVarPostBody = r.targetGroup
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["TokenAuthentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestURL:  localVarPath,
		RequestTime: httpRequestTime,
		Operation:   "TargetgroupsPost",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, string(localVarBody)),
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, err.Error())
			return localVarReturnValue, localVarAPIResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      err.Error(),
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}

type ApiTargetgroupsPutRequest struct {
	ctx             _context.Context
	ApiService      *TargetGroupsApiService
	targetGroupId   string
	targetGroup     *TargetGroupPut
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiTargetgroupsPutRequest) TargetGroup(targetGroup TargetGroupPut) ApiTargetgroupsPutRequest {
	r.targetGroup = &targetGroup
	return r
}
func (r ApiTargetgroupsPutRequest) Pretty(pretty bool) ApiTargetgroupsPutRequest {
	r.pretty = &pretty
	return r
}
func (r ApiTargetgroupsPutRequest) Depth(depth int32) ApiTargetgroupsPutRequest {
	r.depth = &depth
	return r
}
func (r ApiTargetgroupsPutRequest) XContractNumber(xContractNumber int32) ApiTargetgroupsPutRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiTargetgroupsPutRequest) Execute() (TargetGroup, *APIResponse, error) {
	return r.ApiService.TargetgroupsPutExecute(r)
}

/*
 * TargetgroupsPut Modify a Target Group by ID
 * Modifies the properties of the target group specified by its ID.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param targetGroupId The unique ID of the target group.
 * @return ApiTargetgroupsPutRequest
 */
func (a *TargetGroupsApiService) TargetgroupsPut(ctx _context.Context, targetGroupId string) ApiTargetgroupsPutRequest {
	return ApiTargetgroupsPutRequest{
		ApiService:    a,
		ctx:           ctx,
		targetGroupId: targetGroupId,
	}
}

/*
 * Execute executes the request
 * @return TargetGroup
 */
func (a *TargetGroupsApiService) TargetgroupsPutExecute(r ApiTargetgroupsPutRequest) (TargetGroup, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TargetGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetGroupsApiService.TargetgroupsPut")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/targetgroups/{targetGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"targetGroupId"+"}", _neturl.PathEscape(parameterToString(r.targetGroupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.targetGroup == nil {
		return localVarReturnValue, nil, reportError("targetGroup is required and must be specified")
	}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("pretty")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("pretty", parameterToString(true, ""))
		}
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("depth")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("depth", parameterToString(0, ""))
		}
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xContractNumber != nil {
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	// body params
	localVarPostBody = r.targetGroup
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["TokenAuthentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestURL:  localVarPath,
		RequestTime: httpRequestTime,
		Operation:   "TargetgroupsPut",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, string(localVarBody)),
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, err.Error())
			return localVarReturnValue, localVarAPIResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      err.Error(),
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}
