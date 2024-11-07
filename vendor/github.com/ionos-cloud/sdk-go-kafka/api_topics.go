/*
 * Kafka as a Service API
 *
 * An managed Apache Kafka cluster is designed to be highly fault-tolerant and scalable, allowing large volumes of data to be ingested, stored, and processed in real-time. By distributing data across multiple brokers, Kafka achieves high throughput and low latency, making it suitable for applications requiring real-time data processing and analytics.
 *
 * API version: 1.7.1
 * Contact: support@cloud.ionos.com
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

// TopicsApiService TopicsApi service
type TopicsApiService service

type ApiClustersTopicsDeleteRequest struct {
	ctx        _context.Context
	ApiService *TopicsApiService
	clusterId  string
	topicId    string
}

func (r ApiClustersTopicsDeleteRequest) Execute() (*APIResponse, error) {
	return r.ApiService.ClustersTopicsDeleteExecute(r)
}

/*
 * ClustersTopicsDelete Delete Topic
 * Deletes the specified Topic.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The ID (UUID) of the Cluster.
 * @param topicId The ID (UUID) of the Topic.
 * @return ApiClustersTopicsDeleteRequest
 */
func (a *TopicsApiService) ClustersTopicsDelete(ctx _context.Context, clusterId string, topicId string) ApiClustersTopicsDeleteRequest {
	return ApiClustersTopicsDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		clusterId:  clusterId,
		topicId:    topicId,
	}
}

/*
 * Execute executes the request
 */
func (a *TopicsApiService) ClustersTopicsDeleteExecute(r ApiClustersTopicsDeleteRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopicsApiService.ClustersTopicsDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/clusters/{clusterId}/topics/{topicId}"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterId"+"}", _neturl.PathEscape(parameterToString(r.clusterId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topicId"+"}", _neturl.PathEscape(parameterToString(r.topicId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestTime: httpRequestTime,
		RequestURL:  localVarPath,
		Operation:   "ClustersTopicsDelete",
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
			error:      fmt.Sprintf("%s: %s", localVarHTTPResponse.Status, string(localVarBody)),
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarAPIResponse, newErr
		}
		newErr.model = v
		return localVarAPIResponse, newErr
	}

	return localVarAPIResponse, nil
}

type ApiClustersTopicsFindByIdRequest struct {
	ctx        _context.Context
	ApiService *TopicsApiService
	clusterId  string
	topicId    string
}

func (r ApiClustersTopicsFindByIdRequest) Execute() (TopicRead, *APIResponse, error) {
	return r.ApiService.ClustersTopicsFindByIdExecute(r)
}

/*
 * ClustersTopicsFindById Retrieve Topic
 * Returns the Topic by ID.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The ID (UUID) of the Cluster.
 * @param topicId The ID (UUID) of the Topic.
 * @return ApiClustersTopicsFindByIdRequest
 */
func (a *TopicsApiService) ClustersTopicsFindById(ctx _context.Context, clusterId string, topicId string) ApiClustersTopicsFindByIdRequest {
	return ApiClustersTopicsFindByIdRequest{
		ApiService: a,
		ctx:        ctx,
		clusterId:  clusterId,
		topicId:    topicId,
	}
}

/*
 * Execute executes the request
 * @return TopicRead
 */
func (a *TopicsApiService) ClustersTopicsFindByIdExecute(r ApiClustersTopicsFindByIdRequest) (TopicRead, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TopicRead
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopicsApiService.ClustersTopicsFindById")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/clusters/{clusterId}/topics/{topicId}"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterId"+"}", _neturl.PathEscape(parameterToString(r.clusterId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topicId"+"}", _neturl.PathEscape(parameterToString(r.topicId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestTime: httpRequestTime,
		RequestURL:  localVarPath,
		Operation:   "ClustersTopicsFindById",
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
			error:      fmt.Sprintf("%s: %s", localVarHTTPResponse.Status, string(localVarBody)),
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
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

type ApiClustersTopicsGetRequest struct {
	ctx        _context.Context
	ApiService *TopicsApiService
	clusterId  string
}

func (r ApiClustersTopicsGetRequest) Execute() (TopicReadList, *APIResponse, error) {
	return r.ApiService.ClustersTopicsGetExecute(r)
}

/*
  - ClustersTopicsGet Retrieve all Topics
  - This endpoint enables retrieving all Topics using

pagination and optional filters.

  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param clusterId The ID (UUID) of the Cluster.
  - @return ApiClustersTopicsGetRequest
*/
func (a *TopicsApiService) ClustersTopicsGet(ctx _context.Context, clusterId string) ApiClustersTopicsGetRequest {
	return ApiClustersTopicsGetRequest{
		ApiService: a,
		ctx:        ctx,
		clusterId:  clusterId,
	}
}

/*
 * Execute executes the request
 * @return TopicReadList
 */
func (a *TopicsApiService) ClustersTopicsGetExecute(r ApiClustersTopicsGetRequest) (TopicReadList, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TopicReadList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopicsApiService.ClustersTopicsGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/clusters/{clusterId}/topics"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterId"+"}", _neturl.PathEscape(parameterToString(r.clusterId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestTime: httpRequestTime,
		RequestURL:  localVarPath,
		Operation:   "ClustersTopicsGet",
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
			error:      fmt.Sprintf("%s: %s", localVarHTTPResponse.Status, string(localVarBody)),
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
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

type ApiClustersTopicsPostRequest struct {
	ctx         _context.Context
	ApiService  *TopicsApiService
	clusterId   string
	topicCreate *TopicCreate
}

func (r ApiClustersTopicsPostRequest) TopicCreate(topicCreate TopicCreate) ApiClustersTopicsPostRequest {
	r.topicCreate = &topicCreate
	return r
}

func (r ApiClustersTopicsPostRequest) Execute() (TopicRead, *APIResponse, error) {
	return r.ApiService.ClustersTopicsPostExecute(r)
}

/*
  - ClustersTopicsPost Create Topic
  - Creates a new Topic.

The full Topic needs to be provided to create the object.
Optional data will be filled with defaults or left empty.

  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param clusterId The ID (UUID) of the Cluster.
  - @return ApiClustersTopicsPostRequest
*/
func (a *TopicsApiService) ClustersTopicsPost(ctx _context.Context, clusterId string) ApiClustersTopicsPostRequest {
	return ApiClustersTopicsPostRequest{
		ApiService: a,
		ctx:        ctx,
		clusterId:  clusterId,
	}
}

/*
 * Execute executes the request
 * @return TopicRead
 */
func (a *TopicsApiService) ClustersTopicsPostExecute(r ApiClustersTopicsPostRequest) (TopicRead, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TopicRead
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopicsApiService.ClustersTopicsPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/clusters/{clusterId}/topics"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterId"+"}", _neturl.PathEscape(parameterToString(r.clusterId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.topicCreate == nil {
		return localVarReturnValue, nil, reportError("topicCreate is required and must be specified")
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
	// body params
	localVarPostBody = r.topicCreate
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestTime: httpRequestTime,
		RequestURL:  localVarPath,
		Operation:   "ClustersTopicsPost",
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
			error:      fmt.Sprintf("%s: %s", localVarHTTPResponse.Status, string(localVarBody)),
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 415 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
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