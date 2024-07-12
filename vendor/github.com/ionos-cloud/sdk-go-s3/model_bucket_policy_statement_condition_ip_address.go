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

// checks if the BucketPolicyStatementConditionIpAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BucketPolicyStatementConditionIpAddress{}

// BucketPolicyStatementConditionIpAddress struct for BucketPolicyStatementConditionIpAddress
type BucketPolicyStatementConditionIpAddress struct {
	XMLName     xml.Name `xml:"BucketPolicyStatementConditionIpAddress"`
	AwsSourceIp []string `json:"aws:SourceIp,omitempty" xml:"aws:SourceIp"`
}

// NewBucketPolicyStatementConditionIpAddress instantiates a new BucketPolicyStatementConditionIpAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBucketPolicyStatementConditionIpAddress() *BucketPolicyStatementConditionIpAddress {
	this := BucketPolicyStatementConditionIpAddress{}

	return &this
}

// NewBucketPolicyStatementConditionIpAddressWithDefaults instantiates a new BucketPolicyStatementConditionIpAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBucketPolicyStatementConditionIpAddressWithDefaults() *BucketPolicyStatementConditionIpAddress {
	this := BucketPolicyStatementConditionIpAddress{}
	return &this
}

// GetAwsSourceIp returns the AwsSourceIp field value if set, zero value otherwise.
func (o *BucketPolicyStatementConditionIpAddress) GetAwsSourceIp() []string {
	if o == nil || IsNil(o.AwsSourceIp) {
		var ret []string
		return ret
	}
	return o.AwsSourceIp
}

// GetAwsSourceIpOk returns a tuple with the AwsSourceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketPolicyStatementConditionIpAddress) GetAwsSourceIpOk() ([]string, bool) {
	if o == nil || IsNil(o.AwsSourceIp) {
		return nil, false
	}
	return o.AwsSourceIp, true
}

// HasAwsSourceIp returns a boolean if a field has been set.
func (o *BucketPolicyStatementConditionIpAddress) HasAwsSourceIp() bool {
	if o != nil && !IsNil(o.AwsSourceIp) {
		return true
	}

	return false
}

// SetAwsSourceIp gets a reference to the given []string and assigns it to the AwsSourceIp field.
func (o *BucketPolicyStatementConditionIpAddress) SetAwsSourceIp(v []string) {
	o.AwsSourceIp = v
}

func (o BucketPolicyStatementConditionIpAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BucketPolicyStatementConditionIpAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AwsSourceIp) {
		toSerialize["aws:SourceIp"] = o.AwsSourceIp
	}
	return toSerialize, nil
}

type NullableBucketPolicyStatementConditionIpAddress struct {
	value *BucketPolicyStatementConditionIpAddress
	isSet bool
}

func (v NullableBucketPolicyStatementConditionIpAddress) Get() *BucketPolicyStatementConditionIpAddress {
	return v.value
}

func (v *NullableBucketPolicyStatementConditionIpAddress) Set(val *BucketPolicyStatementConditionIpAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableBucketPolicyStatementConditionIpAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableBucketPolicyStatementConditionIpAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBucketPolicyStatementConditionIpAddress(val *BucketPolicyStatementConditionIpAddress) *NullableBucketPolicyStatementConditionIpAddress {
	return &NullableBucketPolicyStatementConditionIpAddress{value: val, isSet: true}
}

func (v NullableBucketPolicyStatementConditionIpAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBucketPolicyStatementConditionIpAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}