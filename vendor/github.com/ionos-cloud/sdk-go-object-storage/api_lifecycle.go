/*
 * IONOS Object Storage API for contract-owned buckets
 *
 * ## Overview The IONOS Object Storage API for contract-owned buckets is a REST-based API that allows developers and applications to interact directly with IONOS' scalable storage solution, leveraging the S3 protocol for object storage operations. Its design ensures seamless compatibility with existing tools and libraries tailored for S3 systems.  ### API References - [S3 API Reference for contract-owned buckets](https://api.ionos.com/docs/s3-contract-owned-buckets/v2/) ### User documentation [IONOS Object Storage User Guide](https://docs.ionos.com/cloud/managed-services/s3-object-storage) * [Documentation on user-owned and contract-owned buckets](https://docs.ionos.com/cloud/managed-services/s3-object-storage/concepts/buckets) * [Documentation on S3 API Compatibility](https://docs.ionos.com/cloud/managed-services/s3-object-storage/concepts/s3-api-compatibility) * [S3 Tools](https://docs.ionos.com/cloud/managed-services/s3-object-storage/s3-tools)  ## Endpoints for contract-owned buckets | Location | Region Name | Bucket Type | Endpoint | | --- | --- | --- | --- | | **Berlin, Germany** | **eu-central-3** | Contract-owned | `https://s3.eu-central-3.ionoscloud.com` |  ## Changelog - 30.05.2024 Initial version
 *
 * API version: 2.0.2
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// LifecycleApiService LifecycleApi service
type LifecycleApiService service

type ApiDeleteBucketLifecycleRequest struct {
	ctx        context.Context
	ApiService *LifecycleApiService
	bucket     string
}

func (r ApiDeleteBucketLifecycleRequest) Execute() (*APIResponse, error) {
	return r.ApiService.DeleteBucketLifecycleExecute(r)
}

/*
DeleteBucketLifecycle DeleteBucketLifecycle

Deletes the lifecycle configuration from the specified bucket.
As a result, objects within the bucket will neither expire nor be automatically deleted based
on any rules from the deleted configuration.

#### Permissions
You must be the contract owner or an administrator to perform this operation. If not, they can grant you permission
to perform the `s3:PutLifecycleConfiguration` operation using [Bucket Policy](#tag/Policy/operation/PutBucketPolicy).

**Note:** A brief delay may occur before the lifecycle configuration deletion is fully
propagated across all IONOS Object Storage systems. During this time, lifecycle rules may remain temporarily active.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bucket
	@return ApiDeleteBucketLifecycleRequest
*/
func (a *LifecycleApiService) DeleteBucketLifecycle(ctx context.Context, bucket string) ApiDeleteBucketLifecycleRequest {
	return ApiDeleteBucketLifecycleRequest{
		ApiService: a,
		ctx:        ctx,
		bucket:     bucket,
	}
}

// Execute executes the request
func (a *LifecycleApiService) DeleteBucketLifecycleExecute(r ApiDeleteBucketLifecycleRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LifecycleApiService.DeleteBucketLifecycle")
	if err != nil {
		gerr := GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return nil, gerr
	}

	localVarPath := localBasePath + "/{Bucket}?lifecycle"
	localVarPath = strings.Replace(localVarPath, "{"+"Bucket"+"}", parameterValueToString(r.bucket, "bucket"), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if Strlen(r.bucket) < 3 {
		return nil, reportError("bucket must have at least 3 elements")
	}
	if Strlen(r.bucket) > 63 {
		return nil, reportError("bucket must have less than 63 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["hmac"]; ok {
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)
	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestTime: httpRequestTime,
		RequestURL:  localVarPath,
		Operation:   "DeleteBucketLifecycle",
	}
	if err != nil || localVarHTTPResponse == nil {
		return localVarAPIResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{}
		newErr.SetStatusCode(localVarHTTPResponse.StatusCode)
		newErr.SetBody(localVarBody)
		newErr.SetError(fmt.Sprintf("%s: %s", localVarHTTPResponse.Status, string(localVarBody)))
		return localVarAPIResponse, newErr
	}

	return localVarAPIResponse, nil
}

type ApiGetBucketLifecycleRequest struct {
	ctx        context.Context
	ApiService *LifecycleApiService
	bucket     string
}

func (r ApiGetBucketLifecycleRequest) Execute() (*GetBucketLifecycleOutput, *APIResponse, error) {
	return r.ApiService.GetBucketLifecycleExecute(r)
}

/*
GetBucketLifecycle GetBucketLifecycle

Returns the lifecycle configuration for the specified Object Storage bucket.

#### Permissions
You must be the contract owner or an administrator to perform this operation. If not, they can grant you permission
to perform the `s3:GetLifecycleConfiguration` operation using [Bucket Policy](#tag/Policy/operation/PutBucketPolicy).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bucket
	@return ApiGetBucketLifecycleRequest
*/
func (a *LifecycleApiService) GetBucketLifecycle(ctx context.Context, bucket string) ApiGetBucketLifecycleRequest {
	return ApiGetBucketLifecycleRequest{
		ApiService: a,
		ctx:        ctx,
		bucket:     bucket,
	}
}

// Execute executes the request
//
//	@return GetBucketLifecycleOutput
func (a *LifecycleApiService) GetBucketLifecycleExecute(r ApiGetBucketLifecycleRequest) (*GetBucketLifecycleOutput, *APIResponse, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetBucketLifecycleOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LifecycleApiService.GetBucketLifecycle")
	if err != nil {
		gerr := GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/{Bucket}?lifecycle"
	localVarPath = strings.Replace(localVarPath, "{"+"Bucket"+"}", parameterValueToString(r.bucket, "bucket"), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if Strlen(r.bucket) < 3 {
		return localVarReturnValue, nil, reportError("bucket must have at least 3 elements")
	}
	if Strlen(r.bucket) > 63 {
		return localVarReturnValue, nil, reportError("bucket must have less than 63 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["hmac"]; ok {
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)
	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestTime: httpRequestTime,
		RequestURL:  localVarPath,
		Operation:   "GetBucketLifecycle",
	}
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{}
		newErr.SetStatusCode(localVarHTTPResponse.StatusCode)
		newErr.SetBody(localVarBody)
		newErr.SetError(fmt.Sprintf("%s: %s", localVarHTTPResponse.Status, string(localVarBody)))
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.SetError(err.Error())
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.SetModel(v)
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{}
		newErr.SetStatusCode(localVarHTTPResponse.StatusCode)
		newErr.SetBody(localVarBody)
		newErr.SetError(err.Error())
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}

type ApiPutBucketLifecycleRequest struct {
	ctx                       context.Context
	ApiService                *LifecycleApiService
	bucket                    string
	contentMD5                *string
	putBucketLifecycleRequest *PutBucketLifecycleRequest
}

func (r ApiPutBucketLifecycleRequest) ContentMD5(contentMD5 string) ApiPutBucketLifecycleRequest {
	r.contentMD5 = &contentMD5
	return r
}

func (r ApiPutBucketLifecycleRequest) PutBucketLifecycleRequest(putBucketLifecycleRequest PutBucketLifecycleRequest) ApiPutBucketLifecycleRequest {
	r.putBucketLifecycleRequest = &putBucketLifecycleRequest
	return r
}

func (r ApiPutBucketLifecycleRequest) Execute() (*APIResponse, error) {
	return r.ApiService.PutBucketLifecycleExecute(r)
}

/*
PutBucketLifecycle PutBucketLifecycle

Creates a new lifecycle configuration for a specified bucket, or replaces an existing configuration.</p>
This lifecycle configuration allows automatic management of the objects within the bucket.
Typical actions can include the deletion of objects after a certain period or deletion of non-current
versions of objects.

#### Permissions
You must be the contract owner or an administrator to perform this operation. If not, they can grant you permission
to perform the `s3:PutLifecycleConfiguration` operation using [Bucket Policy](#tag/Policy/operation/PutBucketPolicy).

#### S3 API Compatibility
- The `NewerNoncurrentVersions` setting is not supported for the `NoncurrentVersionExpiration` option.
- The `Transition` and the `NoncurrentVersionTransition` options are omitted as only the `STANDARD` storage class is currenly supported.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bucket
	@return ApiPutBucketLifecycleRequest
*/
func (a *LifecycleApiService) PutBucketLifecycle(ctx context.Context, bucket string) ApiPutBucketLifecycleRequest {
	return ApiPutBucketLifecycleRequest{
		ApiService: a,
		ctx:        ctx,
		bucket:     bucket,
	}
}

// Execute executes the request
func (a *LifecycleApiService) PutBucketLifecycleExecute(r ApiPutBucketLifecycleRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LifecycleApiService.PutBucketLifecycle")
	if err != nil {
		gerr := GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return nil, gerr
	}

	localVarPath := localBasePath + "/{Bucket}?lifecycle"
	localVarPath = strings.Replace(localVarPath, "{"+"Bucket"+"}", parameterValueToString(r.bucket, "bucket"), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if Strlen(r.bucket) < 3 {
		return nil, reportError("bucket must have at least 3 elements")
	}
	if Strlen(r.bucket) > 63 {
		return nil, reportError("bucket must have less than 63 elements")
	}
	if r.contentMD5 == nil {
		return nil, reportError("contentMD5 is required and must be specified")
	}
	if r.putBucketLifecycleRequest == nil {
		return nil, reportError("putBucketLifecycleRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/xml"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-MD5", r.contentMD5, "")
	// body params
	localVarPostBody = r.putBucketLifecycleRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["hmac"]; ok {
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)
	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestTime: httpRequestTime,
		RequestURL:  localVarPath,
		Operation:   "PutBucketLifecycle",
	}
	if err != nil || localVarHTTPResponse == nil {
		return localVarAPIResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{}
		newErr.SetStatusCode(localVarHTTPResponse.StatusCode)
		newErr.SetBody(localVarBody)
		newErr.SetError(fmt.Sprintf("%s: %s", localVarHTTPResponse.Status, string(localVarBody)))
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.SetError(err.Error())
				return localVarAPIResponse, newErr
			}
			newErr.SetModel(v)
		}
		return localVarAPIResponse, newErr
	}

	return localVarAPIResponse, nil
}
