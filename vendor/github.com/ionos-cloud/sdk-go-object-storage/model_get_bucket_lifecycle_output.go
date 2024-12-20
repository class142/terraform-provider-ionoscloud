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

// GetBucketLifecycleOutput struct for GetBucketLifecycleOutput
type GetBucketLifecycleOutput struct {
	XMLName xml.Name `xml:"LifecycleConfiguration"`
	// Container for a lifecycle rules.
	Rules *[]Rule `json:"Rules,omitempty" xml:"Rule"`
}

// NewGetBucketLifecycleOutput instantiates a new GetBucketLifecycleOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBucketLifecycleOutput() *GetBucketLifecycleOutput {
	this := GetBucketLifecycleOutput{}

	return &this
}

// NewGetBucketLifecycleOutputWithDefaults instantiates a new GetBucketLifecycleOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBucketLifecycleOutputWithDefaults() *GetBucketLifecycleOutput {
	this := GetBucketLifecycleOutput{}
	return &this
}

// GetRules returns the Rules field value
// If the value is explicit nil, the zero value for []Rule will be returned
func (o *GetBucketLifecycleOutput) GetRules() *[]Rule {
	if o == nil {
		return nil
	}

	return o.Rules

}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetBucketLifecycleOutput) GetRulesOk() (*[]Rule, bool) {
	if o == nil {
		return nil, false
	}

	return o.Rules, true
}

// SetRules sets field value
func (o *GetBucketLifecycleOutput) SetRules(v []Rule) {

	o.Rules = &v

}

// HasRules returns a boolean if a field has been set.
func (o *GetBucketLifecycleOutput) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

func (o GetBucketLifecycleOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rules != nil {
		toSerialize["Rules"] = o.Rules
	}

	return json.Marshal(toSerialize)
}

type NullableGetBucketLifecycleOutput struct {
	value *GetBucketLifecycleOutput
	isSet bool
}

func (v NullableGetBucketLifecycleOutput) Get() *GetBucketLifecycleOutput {
	return v.value
}

func (v *NullableGetBucketLifecycleOutput) Set(val *GetBucketLifecycleOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBucketLifecycleOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBucketLifecycleOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBucketLifecycleOutput(val *GetBucketLifecycleOutput) *NullableGetBucketLifecycleOutput {
	return &NullableGetBucketLifecycleOutput{value: val, isSet: true}
}

func (v NullableGetBucketLifecycleOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBucketLifecycleOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
