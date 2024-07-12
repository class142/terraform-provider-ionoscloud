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

// checks if the Encryption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Encryption{}

// Encryption struct for Encryption
type Encryption struct {
	XMLName xml.Name `xml:"Encryption"`
	// The server-side encryption algorithm used when storing job results in IONOS S3 Object Storage (AES256).
	EncryptionType string `json:"EncryptionType" xml:"EncryptionType"`
}

// NewEncryption instantiates a new Encryption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncryption(encryptionType string) *Encryption {
	this := Encryption{}

	this.EncryptionType = encryptionType

	return &this
}

// NewEncryptionWithDefaults instantiates a new Encryption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptionWithDefaults() *Encryption {
	this := Encryption{}
	return &this
}

// GetEncryptionType returns the EncryptionType field value
func (o *Encryption) GetEncryptionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EncryptionType
}

// GetEncryptionTypeOk returns a tuple with the EncryptionType field value
// and a boolean to check if the value has been set.
func (o *Encryption) GetEncryptionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EncryptionType, true
}

// SetEncryptionType sets field value
func (o *Encryption) SetEncryptionType(v string) {
	o.EncryptionType = v
}

func (o Encryption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Encryption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsZero(o.EncryptionType) {
		toSerialize["EncryptionType"] = o.EncryptionType
	}
	return toSerialize, nil
}

type NullableEncryption struct {
	value *Encryption
	isSet bool
}

func (v NullableEncryption) Get() *Encryption {
	return v.value
}

func (v *NullableEncryption) Set(val *Encryption) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryption) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryption(val *Encryption) *NullableEncryption {
	return &NullableEncryption{value: val, isSet: true}
}

func (v NullableEncryption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}