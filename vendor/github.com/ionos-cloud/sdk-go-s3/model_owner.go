/*
 * IONOS S3 Object Storage API for contract-owned buckets
 *
 * ## Overview The IONOS S3 Object Storage API for contract-owned buckets is a REST-based API that allows developers and applications to interact directly with IONOS' scalable storage solution, leveraging the S3 protocol for object storage operations. Its design ensures seamless compatibility with existing tools and libraries tailored for S3 systems.  ### API References - [S3 API Reference for contract-owned buckets](https://api.ionos.com/docs/s3-contract-owned-buckets/v2/) ### User documentation [IONOS S3 Object Storage User Guide](https://docs.ionos.com/cloud/managed-services/s3-object-storage) * [Documentation on user-owned and contract-owned buckets](https://docs.ionos.com/cloud/managed-services/s3-object-storage/concepts/buckets) * [Documentation on S3 API Compatibility](https://docs.ionos.com/cloud/managed-services/s3-object-storage/concepts/s3-api-compatibility) * [S3 Tools](https://docs.ionos.com/cloud/managed-services/s3-object-storage/s3-tools)  ## Endpoints for contract-owned buckets | Location | Region Name | Bucket Type | Endpoint | | --- | --- | --- | --- | | **Berlin, Germany** | **eu-central-3** | Contract-owned | `https://s3.eu-central-3.ionoscloud.com` |  ## Changelog - 30.05.2024 Initial version
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

// Owner Container for the owner's ID and display name.
type Owner struct {
	XMLName xml.Name `xml:"Owner"`
	// Container for the Contract Number of the owner.
	ID *int32 `json:"ID,omitempty" xml:"ID"`
	// Container for the display name of the owner.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName"`
}

// NewOwner instantiates a new Owner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOwner() *Owner {
	this := Owner{}

	return &this
}

// NewOwnerWithDefaults instantiates a new Owner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOwnerWithDefaults() *Owner {
	this := Owner{}
	return &this
}

// GetID returns the ID field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Owner) GetID() *int32 {
	if o == nil {
		return nil
	}

	return o.ID

}

// GetIDOk returns a tuple with the ID field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Owner) GetIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.ID, true
}

// SetID sets field value
func (o *Owner) SetID(v int32) {

	o.ID = &v

}

// HasID returns a boolean if a field has been set.
func (o *Owner) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// GetDisplayName returns the DisplayName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Owner) GetDisplayName() *string {
	if o == nil {
		return nil
	}

	return o.DisplayName

}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Owner) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.DisplayName, true
}

// SetDisplayName sets field value
func (o *Owner) SetDisplayName(v string) {

	o.DisplayName = &v

}

// HasDisplayName returns a boolean if a field has been set.
func (o *Owner) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

func (o Owner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ID != nil {
		toSerialize["ID"] = o.ID
	}

	if o.DisplayName != nil {
		toSerialize["DisplayName"] = o.DisplayName
	}

	return json.Marshal(toSerialize)
}

type NullableOwner struct {
	value *Owner
	isSet bool
}

func (v NullableOwner) Get() *Owner {
	return v.value
}

func (v *NullableOwner) Set(val *Owner) {
	v.value = val
	v.isSet = true
}

func (v NullableOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOwner(val *Owner) *NullableOwner {
	return &NullableOwner{value: val, isSet: true}
}

func (v NullableOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
