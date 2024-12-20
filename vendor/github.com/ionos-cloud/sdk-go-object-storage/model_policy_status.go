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

// PolicyStatus The container element for a bucket's policy status.
type PolicyStatus struct {
	XMLName xml.Name `xml:"PolicyStatus"`
	// The policy status for this bucket: - `true` indicates that this bucket is public. - `false` indicates that this bucket is private.
	IsPublic *bool `json:"IsPublic,omitempty" xml:"IsPublic"`
}

// NewPolicyStatus instantiates a new PolicyStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyStatus() *PolicyStatus {
	this := PolicyStatus{}

	return &this
}

// NewPolicyStatusWithDefaults instantiates a new PolicyStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyStatusWithDefaults() *PolicyStatus {
	this := PolicyStatus{}
	return &this
}

// GetIsPublic returns the IsPublic field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *PolicyStatus) GetIsPublic() *bool {
	if o == nil {
		return nil
	}

	return o.IsPublic

}

// GetIsPublicOk returns a tuple with the IsPublic field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyStatus) GetIsPublicOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.IsPublic, true
}

// SetIsPublic sets field value
func (o *PolicyStatus) SetIsPublic(v bool) {

	o.IsPublic = &v

}

// HasIsPublic returns a boolean if a field has been set.
func (o *PolicyStatus) HasIsPublic() bool {
	if o != nil && o.IsPublic != nil {
		return true
	}

	return false
}

func (o PolicyStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsPublic != nil {
		toSerialize["IsPublic"] = o.IsPublic
	}

	return json.Marshal(toSerialize)
}

type NullablePolicyStatus struct {
	value *PolicyStatus
	isSet bool
}

func (v NullablePolicyStatus) Get() *PolicyStatus {
	return v.value
}

func (v *NullablePolicyStatus) Set(val *PolicyStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyStatus(val *PolicyStatus) *NullablePolicyStatus {
	return &NullablePolicyStatus{value: val, isSet: true}
}

func (v NullablePolicyStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
