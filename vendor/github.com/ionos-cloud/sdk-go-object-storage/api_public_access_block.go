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

// PublicAccessBlockApiService PublicAccessBlockApi service
type PublicAccessBlockApiService service

type ApiDeletePublicAccessBlockRequest struct {
	ctx        context.Context
	ApiService *PublicAccessBlockApiService
	bucket     string
}

func (r ApiDeletePublicAccessBlockRequest) Execute() (*APIResponse, error) {
	return r.ApiService.DeletePublicAccessBlockExecute(r)
}

/*
DeletePublicAccessBlock DeletePublicAccessBlock

Deletes the public access configuration for an Object Storage bucket.

#### Permissions
You must be the contract owner or an administrator to perform this operation. If not, they can grant you permission
to perform the `s3:DeleteBucketPublicAccessBlock` operation using [Bucket Policy](#tag/Policy/operation/PutBucketPolicy).

#### S3 API Compatibility
- The `x-amz-expected-bucket-owner` header isn't supported.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bucket
	@return ApiDeletePublicAccessBlockRequest
*/
func (a *PublicAccessBlockApiService) DeletePublicAccessBlock(ctx context.Context, bucket string) ApiDeletePublicAccessBlockRequest {
	return ApiDeletePublicAccessBlockRequest{
		ApiService: a,
		ctx:        ctx,
		bucket:     bucket,
	}
}

// Execute executes the request
func (a *PublicAccessBlockApiService) DeletePublicAccessBlockExecute(r ApiDeletePublicAccessBlockRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicAccessBlockApiService.DeletePublicAccessBlock")
	if err != nil {
		gerr := GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return nil, gerr
	}

	localVarPath := localBasePath + "/{Bucket}?publicAccessBlock"
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
		Operation:   "DeletePublicAccessBlock",
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.SetError(err.Error())
				return localVarAPIResponse, newErr
			}
			newErr.SetModel(v)
			return localVarAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiGetPublicAccessBlockRequest struct {
	ctx        context.Context
	ApiService *PublicAccessBlockApiService
	bucket     string
}

func (r ApiGetPublicAccessBlockRequest) Execute() (*BlockPublicAccessOutput, *APIResponse, error) {
	return r.ApiService.GetPublicAccessBlockExecute(r)
}

/*
GetPublicAccessBlock GetPublicAccessBlock

Retrieves the public access configuration for a bucket.

#### Permissions
You must be the contract owner or an administrator to perform this operation. If not, they can grant you permission
to perform the `s3:GetBucketPublicAccessBlock` operation using [Bucket Policy](#tag/Policy/operation/PutBucketPolicy).

#### S3 API Compatibility
- The `x-amz-expected-bucket-owner` header isn't supported.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bucket
	@return ApiGetPublicAccessBlockRequest
*/
func (a *PublicAccessBlockApiService) GetPublicAccessBlock(ctx context.Context, bucket string) ApiGetPublicAccessBlockRequest {
	return ApiGetPublicAccessBlockRequest{
		ApiService: a,
		ctx:        ctx,
		bucket:     bucket,
	}
}

// Execute executes the request
//
//	@return BlockPublicAccessOutput
func (a *PublicAccessBlockApiService) GetPublicAccessBlockExecute(r ApiGetPublicAccessBlockRequest) (*BlockPublicAccessOutput, *APIResponse, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BlockPublicAccessOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicAccessBlockApiService.GetPublicAccessBlock")
	if err != nil {
		gerr := GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/{Bucket}?publicAccessBlock"
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
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml"}

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
		Operation:   "GetPublicAccessBlock",
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.SetError(err.Error())
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.SetModel(v)
			return localVarReturnValue, localVarAPIResponse, newErr
		}
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

type ApiPutPublicAccessBlockRequest struct {
	ctx                      context.Context
	ApiService               *PublicAccessBlockApiService
	bucket                   string
	blockPublicAccessPayload *BlockPublicAccessPayload
	contentMD5               *string
}

func (r ApiPutPublicAccessBlockRequest) BlockPublicAccessPayload(blockPublicAccessPayload BlockPublicAccessPayload) ApiPutPublicAccessBlockRequest {
	r.blockPublicAccessPayload = &blockPublicAccessPayload
	return r
}

func (r ApiPutPublicAccessBlockRequest) ContentMD5(contentMD5 string) ApiPutPublicAccessBlockRequest {
	r.contentMD5 = &contentMD5
	return r
}

func (r ApiPutPublicAccessBlockRequest) Execute() (*APIResponse, error) {
	return r.ApiService.PutPublicAccessBlockExecute(r)
}

/*
PutPublicAccessBlock PutPublicAccessBlock

Blocks public access to an Object Storage bucket based on the specified parameters.

This operation modifies the bucket's settings to either prevent public access entirely
or impose restrictions based on specific conditions.

#### Permissions
You must be the contract owner or an administrator to perform this operation. If not, they can grant you permission
to perform the `s3:PutBucketPublicAccessBlock` operation using [Bucket Policy](#tag/Policy/operation/PutBucketPolicy).
**Note:** The bucket owner can always perform this operation, even if the policy explicitly denies access to it.

#### S3 API Compatibility
- The `x-amz-expected-bucket-owner` header isn't supported.
- The `x-amz-confirm-remove-self-bucket-access` header isn't supported.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bucket
	@return ApiPutPublicAccessBlockRequest
*/
func (a *PublicAccessBlockApiService) PutPublicAccessBlock(ctx context.Context, bucket string) ApiPutPublicAccessBlockRequest {
	return ApiPutPublicAccessBlockRequest{
		ApiService: a,
		ctx:        ctx,
		bucket:     bucket,
	}
}

// Execute executes the request
func (a *PublicAccessBlockApiService) PutPublicAccessBlockExecute(r ApiPutPublicAccessBlockRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicAccessBlockApiService.PutPublicAccessBlock")
	if err != nil {
		gerr := GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return nil, gerr
	}

	localVarPath := localBasePath + "/{Bucket}?publicAccessBlock"
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
	if r.blockPublicAccessPayload == nil {
		return nil, reportError("blockPublicAccessPayload is required and must be specified")
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
	if r.contentMD5 != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-MD5", r.contentMD5, "")
	}
	// body params
	localVarPostBody = r.blockPublicAccessPayload
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
		Operation:   "PutPublicAccessBlock",
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
			return localVarAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.SetError(err.Error())
				return localVarAPIResponse, newErr
			}
			newErr.SetModel(v)
			return localVarAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
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
