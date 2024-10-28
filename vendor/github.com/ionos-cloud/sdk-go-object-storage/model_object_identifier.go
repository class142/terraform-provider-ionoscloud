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

// ObjectIdentifier Object Identifier is unique value to identify objects.
type ObjectIdentifier struct {
	XMLName xml.Name `xml:"Object"`
	// The object key.
	Key *string `json:"Key" xml:"Key"`
	// VersionId for the specific version of the object to delete.
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId"`
}

// NewObjectIdentifier instantiates a new ObjectIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectIdentifier(key string) *ObjectIdentifier {
	this := ObjectIdentifier{}

	this.Key = &key

	return &this
}

// NewObjectIdentifierWithDefaults instantiates a new ObjectIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectIdentifierWithDefaults() *ObjectIdentifier {
	this := ObjectIdentifier{}
	return &this
}

// GetKey returns the Key field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ObjectIdentifier) GetKey() *string {
	if o == nil {
		return nil
	}

	return o.Key

}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectIdentifier) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Key, true
}

// SetKey sets field value
func (o *ObjectIdentifier) SetKey(v string) {

	o.Key = &v

}

// HasKey returns a boolean if a field has been set.
func (o *ObjectIdentifier) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// GetVersionId returns the VersionId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ObjectIdentifier) GetVersionId() *string {
	if o == nil {
		return nil
	}

	return o.VersionId

}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectIdentifier) GetVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.VersionId, true
}

// SetVersionId sets field value
func (o *ObjectIdentifier) SetVersionId(v string) {

	o.VersionId = &v

}

// HasVersionId returns a boolean if a field has been set.
func (o *ObjectIdentifier) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

func (o ObjectIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}

	if o.VersionId != nil {
		toSerialize["VersionId"] = o.VersionId
	}

	return json.Marshal(toSerialize)
}

type NullableObjectIdentifier struct {
	value *ObjectIdentifier
	isSet bool
}

func (v NullableObjectIdentifier) Get() *ObjectIdentifier {
	return v.value
}

func (v *NullableObjectIdentifier) Set(val *ObjectIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectIdentifier(val *ObjectIdentifier) *NullableObjectIdentifier {
	return &NullableObjectIdentifier{value: val, isSet: true}
}

func (v NullableObjectIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
