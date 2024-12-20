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

// ServerSideEncryptionByDefault Describes the default server-side encryption to apply to new objects in the bucket. If a PUT Object request doesn't specify any server-side encryption, this default encryption will be applied.
type ServerSideEncryptionByDefault struct {
	XMLName      xml.Name              `xml:"ApplyServerSideEncryptionByDefault"`
	SSEAlgorithm *ServerSideEncryption `json:"SSEAlgorithm" xml:"SSEAlgorithm"`
}

// NewServerSideEncryptionByDefault instantiates a new ServerSideEncryptionByDefault object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerSideEncryptionByDefault(sSEAlgorithm ServerSideEncryption) *ServerSideEncryptionByDefault {
	this := ServerSideEncryptionByDefault{}

	this.SSEAlgorithm = &sSEAlgorithm

	return &this
}

// NewServerSideEncryptionByDefaultWithDefaults instantiates a new ServerSideEncryptionByDefault object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerSideEncryptionByDefaultWithDefaults() *ServerSideEncryptionByDefault {
	this := ServerSideEncryptionByDefault{}
	return &this
}

// GetSSEAlgorithm returns the SSEAlgorithm field value
// If the value is explicit nil, the zero value for ServerSideEncryption will be returned
func (o *ServerSideEncryptionByDefault) GetSSEAlgorithm() *ServerSideEncryption {
	if o == nil {
		return nil
	}

	return o.SSEAlgorithm

}

// GetSSEAlgorithmOk returns a tuple with the SSEAlgorithm field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerSideEncryptionByDefault) GetSSEAlgorithmOk() (*ServerSideEncryption, bool) {
	if o == nil {
		return nil, false
	}

	return o.SSEAlgorithm, true
}

// SetSSEAlgorithm sets field value
func (o *ServerSideEncryptionByDefault) SetSSEAlgorithm(v ServerSideEncryption) {

	o.SSEAlgorithm = &v

}

// HasSSEAlgorithm returns a boolean if a field has been set.
func (o *ServerSideEncryptionByDefault) HasSSEAlgorithm() bool {
	if o != nil && o.SSEAlgorithm != nil {
		return true
	}

	return false
}

func (o ServerSideEncryptionByDefault) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SSEAlgorithm != nil {
		toSerialize["SSEAlgorithm"] = o.SSEAlgorithm
	}

	return json.Marshal(toSerialize)
}

type NullableServerSideEncryptionByDefault struct {
	value *ServerSideEncryptionByDefault
	isSet bool
}

func (v NullableServerSideEncryptionByDefault) Get() *ServerSideEncryptionByDefault {
	return v.value
}

func (v *NullableServerSideEncryptionByDefault) Set(val *ServerSideEncryptionByDefault) {
	v.value = val
	v.isSet = true
}

func (v NullableServerSideEncryptionByDefault) IsSet() bool {
	return v.isSet
}

func (v *NullableServerSideEncryptionByDefault) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerSideEncryptionByDefault(val *ServerSideEncryptionByDefault) *NullableServerSideEncryptionByDefault {
	return &NullableServerSideEncryptionByDefault{value: val, isSet: true}
}

func (v NullableServerSideEncryptionByDefault) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerSideEncryptionByDefault) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
