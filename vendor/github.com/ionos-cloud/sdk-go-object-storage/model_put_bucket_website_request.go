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
	"encoding/json"
)

import "encoding/xml"

// PutBucketWebsiteRequest Specifies website configuration parameters for an IONOS Object Storage bucket.
type PutBucketWebsiteRequest struct {
	XMLName               xml.Name               `xml:"WebsiteConfiguration"`
	ErrorDocument         *ErrorDocument         `json:"ErrorDocument,omitempty" xml:"ErrorDocument"`
	IndexDocument         *IndexDocument         `json:"IndexDocument,omitempty" xml:"IndexDocument"`
	RedirectAllRequestsTo *RedirectAllRequestsTo `json:"RedirectAllRequestsTo,omitempty" xml:"RedirectAllRequestsTo"`
	RoutingRules          *[]RoutingRule         `json:"RoutingRules,omitempty" xml:"RoutingRules"`
}

// NewPutBucketWebsiteRequest instantiates a new PutBucketWebsiteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutBucketWebsiteRequest() *PutBucketWebsiteRequest {
	this := PutBucketWebsiteRequest{}

	return &this
}

// NewPutBucketWebsiteRequestWithDefaults instantiates a new PutBucketWebsiteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutBucketWebsiteRequestWithDefaults() *PutBucketWebsiteRequest {
	this := PutBucketWebsiteRequest{}
	return &this
}

// GetErrorDocument returns the ErrorDocument field value
// If the value is explicit nil, the zero value for ErrorDocument will be returned
func (o *PutBucketWebsiteRequest) GetErrorDocument() *ErrorDocument {
	if o == nil {
		return nil
	}

	return o.ErrorDocument

}

// GetErrorDocumentOk returns a tuple with the ErrorDocument field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PutBucketWebsiteRequest) GetErrorDocumentOk() (*ErrorDocument, bool) {
	if o == nil {
		return nil, false
	}

	return o.ErrorDocument, true
}

// SetErrorDocument sets field value
func (o *PutBucketWebsiteRequest) SetErrorDocument(v ErrorDocument) {

	o.ErrorDocument = &v

}

// HasErrorDocument returns a boolean if a field has been set.
func (o *PutBucketWebsiteRequest) HasErrorDocument() bool {
	if o != nil && o.ErrorDocument != nil {
		return true
	}

	return false
}

// GetIndexDocument returns the IndexDocument field value
// If the value is explicit nil, the zero value for IndexDocument will be returned
func (o *PutBucketWebsiteRequest) GetIndexDocument() *IndexDocument {
	if o == nil {
		return nil
	}

	return o.IndexDocument

}

// GetIndexDocumentOk returns a tuple with the IndexDocument field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PutBucketWebsiteRequest) GetIndexDocumentOk() (*IndexDocument, bool) {
	if o == nil {
		return nil, false
	}

	return o.IndexDocument, true
}

// SetIndexDocument sets field value
func (o *PutBucketWebsiteRequest) SetIndexDocument(v IndexDocument) {

	o.IndexDocument = &v

}

// HasIndexDocument returns a boolean if a field has been set.
func (o *PutBucketWebsiteRequest) HasIndexDocument() bool {
	if o != nil && o.IndexDocument != nil {
		return true
	}

	return false
}

// GetRedirectAllRequestsTo returns the RedirectAllRequestsTo field value
// If the value is explicit nil, the zero value for RedirectAllRequestsTo will be returned
func (o *PutBucketWebsiteRequest) GetRedirectAllRequestsTo() *RedirectAllRequestsTo {
	if o == nil {
		return nil
	}

	return o.RedirectAllRequestsTo

}

// GetRedirectAllRequestsToOk returns a tuple with the RedirectAllRequestsTo field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PutBucketWebsiteRequest) GetRedirectAllRequestsToOk() (*RedirectAllRequestsTo, bool) {
	if o == nil {
		return nil, false
	}

	return o.RedirectAllRequestsTo, true
}

// SetRedirectAllRequestsTo sets field value
func (o *PutBucketWebsiteRequest) SetRedirectAllRequestsTo(v RedirectAllRequestsTo) {

	o.RedirectAllRequestsTo = &v

}

// HasRedirectAllRequestsTo returns a boolean if a field has been set.
func (o *PutBucketWebsiteRequest) HasRedirectAllRequestsTo() bool {
	if o != nil && o.RedirectAllRequestsTo != nil {
		return true
	}

	return false
}

// GetRoutingRules returns the RoutingRules field value
// If the value is explicit nil, the zero value for []RoutingRule will be returned
func (o *PutBucketWebsiteRequest) GetRoutingRules() *[]RoutingRule {
	if o == nil {
		return nil
	}

	return o.RoutingRules

}

// GetRoutingRulesOk returns a tuple with the RoutingRules field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PutBucketWebsiteRequest) GetRoutingRulesOk() (*[]RoutingRule, bool) {
	if o == nil {
		return nil, false
	}

	return o.RoutingRules, true
}

// SetRoutingRules sets field value
func (o *PutBucketWebsiteRequest) SetRoutingRules(v []RoutingRule) {

	o.RoutingRules = &v

}

// HasRoutingRules returns a boolean if a field has been set.
func (o *PutBucketWebsiteRequest) HasRoutingRules() bool {
	if o != nil && o.RoutingRules != nil {
		return true
	}

	return false
}

func (o PutBucketWebsiteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorDocument != nil {
		toSerialize["ErrorDocument"] = o.ErrorDocument
	}

	if o.IndexDocument != nil {
		toSerialize["IndexDocument"] = o.IndexDocument
	}

	if o.RedirectAllRequestsTo != nil {
		toSerialize["RedirectAllRequestsTo"] = o.RedirectAllRequestsTo
	}

	if o.RoutingRules != nil {
		toSerialize["RoutingRules"] = o.RoutingRules
	}

	return json.Marshal(toSerialize)
}

type NullablePutBucketWebsiteRequest struct {
	value *PutBucketWebsiteRequest
	isSet bool
}

func (v NullablePutBucketWebsiteRequest) Get() *PutBucketWebsiteRequest {
	return v.value
}

func (v *NullablePutBucketWebsiteRequest) Set(val *PutBucketWebsiteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutBucketWebsiteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutBucketWebsiteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutBucketWebsiteRequest(val *PutBucketWebsiteRequest) *NullablePutBucketWebsiteRequest {
	return &NullablePutBucketWebsiteRequest{value: val, isSet: true}
}

func (v NullablePutBucketWebsiteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutBucketWebsiteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}