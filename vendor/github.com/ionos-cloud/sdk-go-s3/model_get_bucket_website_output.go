/*
 * IONOS S3 Object Storage API for contract-owned buckets
 *
 * ## Overview The IONOS S3 Object Storage API for contract-owned buckets is a REST-based API that allows developers and applications to interact directly with IONOS' scalable storage solution, leveraging the S3 protocol for object storage operations. Its design ensures seamless compatibility with existing tools and libraries tailored for S3 systems.  ### API References - [S3 Management API Reference](https://api.ionos.com/docs/s3-management/v1/) for managing Access Keys - S3 API Reference for contract-owned buckets - current document - [S3 API Reference for user-owned buckets](https://api.ionos.com/docs/s3-user-owned-buckets/v2/)  ### User documentation [IONOS S3 Object Storage User Guide](https://docs.ionos.com/cloud/managed-services/s3-object-storage) * [Documentation on user-owned and contract-owned buckets](https://docs.ionos.com/cloud/managed-services/s3-object-storage/concepts/buckets) * [Documentation on S3 API Compatibility](https://docs.ionos.com/cloud/managed-services/s3-object-storage/concepts/s3-api-compatibility) * [S3 Tools](https://docs.ionos.com/cloud/managed-services/s3-object-storage/s3-tools)  ## Endpoints for contract-owned buckets | Location | Region Name | Bucket Type | Endpoint | | --- | --- | --- | --- | | **Berlin, Germany** | **eu-central-3** | Contract-owned | `https://s3.eu-central-3.ionoscloud.com` |  ## Changelog - 30.05.2024 Initial version
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

// GetBucketWebsiteOutput struct for GetBucketWebsiteOutput
type GetBucketWebsiteOutput struct {
	XMLName               xml.Name               `xml:"GetBucketWebsiteOutput"`
	RedirectAllRequestsTo *RedirectAllRequestsTo `json:"RedirectAllRequestsTo,omitempty" xml:"RedirectAllRequestsTo"`
	IndexDocument         *IndexDocument         `json:"IndexDocument,omitempty" xml:"IndexDocument"`
	ErrorDocument         *ErrorDocument         `json:"ErrorDocument,omitempty" xml:"ErrorDocument"`
	RoutingRules          *[]RoutingRule         `json:"RoutingRules,omitempty" xml:"RoutingRules"`
}

// NewGetBucketWebsiteOutput instantiates a new GetBucketWebsiteOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBucketWebsiteOutput() *GetBucketWebsiteOutput {
	this := GetBucketWebsiteOutput{}

	return &this
}

// NewGetBucketWebsiteOutputWithDefaults instantiates a new GetBucketWebsiteOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBucketWebsiteOutputWithDefaults() *GetBucketWebsiteOutput {
	this := GetBucketWebsiteOutput{}
	return &this
}

// GetRedirectAllRequestsTo returns the RedirectAllRequestsTo field value
// If the value is explicit nil, the zero value for RedirectAllRequestsTo will be returned
func (o *GetBucketWebsiteOutput) GetRedirectAllRequestsTo() *RedirectAllRequestsTo {
	if o == nil {
		return nil
	}

	return o.RedirectAllRequestsTo

}

// GetRedirectAllRequestsToOk returns a tuple with the RedirectAllRequestsTo field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetBucketWebsiteOutput) GetRedirectAllRequestsToOk() (*RedirectAllRequestsTo, bool) {
	if o == nil {
		return nil, false
	}

	return o.RedirectAllRequestsTo, true
}

// SetRedirectAllRequestsTo sets field value
func (o *GetBucketWebsiteOutput) SetRedirectAllRequestsTo(v RedirectAllRequestsTo) {

	o.RedirectAllRequestsTo = &v

}

// HasRedirectAllRequestsTo returns a boolean if a field has been set.
func (o *GetBucketWebsiteOutput) HasRedirectAllRequestsTo() bool {
	if o != nil && o.RedirectAllRequestsTo != nil {
		return true
	}

	return false
}

// GetIndexDocument returns the IndexDocument field value
// If the value is explicit nil, the zero value for IndexDocument will be returned
func (o *GetBucketWebsiteOutput) GetIndexDocument() *IndexDocument {
	if o == nil {
		return nil
	}

	return o.IndexDocument

}

// GetIndexDocumentOk returns a tuple with the IndexDocument field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetBucketWebsiteOutput) GetIndexDocumentOk() (*IndexDocument, bool) {
	if o == nil {
		return nil, false
	}

	return o.IndexDocument, true
}

// SetIndexDocument sets field value
func (o *GetBucketWebsiteOutput) SetIndexDocument(v IndexDocument) {

	o.IndexDocument = &v

}

// HasIndexDocument returns a boolean if a field has been set.
func (o *GetBucketWebsiteOutput) HasIndexDocument() bool {
	if o != nil && o.IndexDocument != nil {
		return true
	}

	return false
}

// GetErrorDocument returns the ErrorDocument field value
// If the value is explicit nil, the zero value for ErrorDocument will be returned
func (o *GetBucketWebsiteOutput) GetErrorDocument() *ErrorDocument {
	if o == nil {
		return nil
	}

	return o.ErrorDocument

}

// GetErrorDocumentOk returns a tuple with the ErrorDocument field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetBucketWebsiteOutput) GetErrorDocumentOk() (*ErrorDocument, bool) {
	if o == nil {
		return nil, false
	}

	return o.ErrorDocument, true
}

// SetErrorDocument sets field value
func (o *GetBucketWebsiteOutput) SetErrorDocument(v ErrorDocument) {

	o.ErrorDocument = &v

}

// HasErrorDocument returns a boolean if a field has been set.
func (o *GetBucketWebsiteOutput) HasErrorDocument() bool {
	if o != nil && o.ErrorDocument != nil {
		return true
	}

	return false
}

// GetRoutingRules returns the RoutingRules field value
// If the value is explicit nil, the zero value for []RoutingRule will be returned
func (o *GetBucketWebsiteOutput) GetRoutingRules() *[]RoutingRule {
	if o == nil {
		return nil
	}

	return o.RoutingRules

}

// GetRoutingRulesOk returns a tuple with the RoutingRules field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetBucketWebsiteOutput) GetRoutingRulesOk() (*[]RoutingRule, bool) {
	if o == nil {
		return nil, false
	}

	return o.RoutingRules, true
}

// SetRoutingRules sets field value
func (o *GetBucketWebsiteOutput) SetRoutingRules(v []RoutingRule) {

	o.RoutingRules = &v

}

// HasRoutingRules returns a boolean if a field has been set.
func (o *GetBucketWebsiteOutput) HasRoutingRules() bool {
	if o != nil && o.RoutingRules != nil {
		return true
	}

	return false
}

func (o GetBucketWebsiteOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RedirectAllRequestsTo != nil {
		toSerialize["RedirectAllRequestsTo"] = o.RedirectAllRequestsTo
	}

	if o.IndexDocument != nil {
		toSerialize["IndexDocument"] = o.IndexDocument
	}

	if o.ErrorDocument != nil {
		toSerialize["ErrorDocument"] = o.ErrorDocument
	}

	if o.RoutingRules != nil {
		toSerialize["RoutingRules"] = o.RoutingRules
	}

	return json.Marshal(toSerialize)
}

type NullableGetBucketWebsiteOutput struct {
	value *GetBucketWebsiteOutput
	isSet bool
}

func (v NullableGetBucketWebsiteOutput) Get() *GetBucketWebsiteOutput {
	return v.value
}

func (v *NullableGetBucketWebsiteOutput) Set(val *GetBucketWebsiteOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBucketWebsiteOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBucketWebsiteOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBucketWebsiteOutput(val *GetBucketWebsiteOutput) *NullableGetBucketWebsiteOutput {
	return &NullableGetBucketWebsiteOutput{value: val, isSet: true}
}

func (v NullableGetBucketWebsiteOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBucketWebsiteOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}