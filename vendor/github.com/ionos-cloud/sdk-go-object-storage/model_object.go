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
	"time"
)

import "encoding/xml"

// Object An object consists of data and its descriptive metadata.
type Object struct {
	XMLName xml.Name `xml:"Contents"`
	// The object key.
	Key *string `json:"Key,omitempty" xml:"Key"`
	// Creation date of the object.
	LastModified *IonosTime          `json:"LastModified,omitempty" xml:"LastModified"`
	StorageClass *ObjectStorageClass `json:"StorageClass,omitempty" xml:"StorageClass"`
	// Size in bytes of the object
	Size *int32 `json:"Size,omitempty" xml:"Size"`
	// Entity tag that identifies the object's data. Objects with different object data will have different entity tags. The entity tag is an opaque string. The entity tag may or may not be an MD5 digest of the object data. If the entity tag is not an MD5 digest of the object data, it will contain one or more nonhexadecimal characters and/or will consist of less than 32 or more than 32 hexadecimal digits.
	ETag  *string `json:"ETag,omitempty" xml:"ETag"`
	Owner *Owner  `json:"Owner,omitempty" xml:"Owner"`
}

// NewObject instantiates a new Object object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObject() *Object {
	this := Object{}

	return &this
}

// NewObjectWithDefaults instantiates a new Object object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectWithDefaults() *Object {
	this := Object{}
	return &this
}

// GetKey returns the Key field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Object) GetKey() *string {
	if o == nil {
		return nil
	}

	return o.Key

}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Object) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Key, true
}

// SetKey sets field value
func (o *Object) SetKey(v string) {

	o.Key = &v

}

// HasKey returns a boolean if a field has been set.
func (o *Object) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// GetLastModified returns the LastModified field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Object) GetLastModified() *time.Time {
	if o == nil {
		return nil
	}

	if o.LastModified == nil {
		return nil
	}
	return &o.LastModified.Time

}

// GetLastModifiedOk returns a tuple with the LastModified field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Object) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}

	if o.LastModified == nil {
		return nil, false
	}
	return &o.LastModified.Time, true

}

// SetLastModified sets field value
func (o *Object) SetLastModified(v time.Time) {

	o.LastModified = &IonosTime{v}

}

// HasLastModified returns a boolean if a field has been set.
func (o *Object) HasLastModified() bool {
	if o != nil && o.LastModified != nil {
		return true
	}

	return false
}

// GetStorageClass returns the StorageClass field value
// If the value is explicit nil, the zero value for ObjectStorageClass will be returned
func (o *Object) GetStorageClass() *ObjectStorageClass {
	if o == nil {
		return nil
	}

	return o.StorageClass

}

// GetStorageClassOk returns a tuple with the StorageClass field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Object) GetStorageClassOk() (*ObjectStorageClass, bool) {
	if o == nil {
		return nil, false
	}

	return o.StorageClass, true
}

// SetStorageClass sets field value
func (o *Object) SetStorageClass(v ObjectStorageClass) {

	o.StorageClass = &v

}

// HasStorageClass returns a boolean if a field has been set.
func (o *Object) HasStorageClass() bool {
	if o != nil && o.StorageClass != nil {
		return true
	}

	return false
}

// GetSize returns the Size field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Object) GetSize() *int32 {
	if o == nil {
		return nil
	}

	return o.Size

}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Object) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Size, true
}

// SetSize sets field value
func (o *Object) SetSize(v int32) {

	o.Size = &v

}

// HasSize returns a boolean if a field has been set.
func (o *Object) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// GetETag returns the ETag field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Object) GetETag() *string {
	if o == nil {
		return nil
	}

	return o.ETag

}

// GetETagOk returns a tuple with the ETag field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Object) GetETagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.ETag, true
}

// SetETag sets field value
func (o *Object) SetETag(v string) {

	o.ETag = &v

}

// HasETag returns a boolean if a field has been set.
func (o *Object) HasETag() bool {
	if o != nil && o.ETag != nil {
		return true
	}

	return false
}

// GetOwner returns the Owner field value
// If the value is explicit nil, the zero value for Owner will be returned
func (o *Object) GetOwner() *Owner {
	if o == nil {
		return nil
	}

	return o.Owner

}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Object) GetOwnerOk() (*Owner, bool) {
	if o == nil {
		return nil, false
	}

	return o.Owner, true
}

// SetOwner sets field value
func (o *Object) SetOwner(v Owner) {

	o.Owner = &v

}

// HasOwner returns a boolean if a field has been set.
func (o *Object) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

func (o Object) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}

	if o.LastModified != nil {
		toSerialize["LastModified"] = o.LastModified
	}

	if o.StorageClass != nil {
		toSerialize["StorageClass"] = o.StorageClass
	}

	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}

	if o.ETag != nil {
		toSerialize["ETag"] = o.ETag
	}

	if o.Owner != nil {
		toSerialize["Owner"] = o.Owner
	}

	return json.Marshal(toSerialize)
}

type NullableObject struct {
	value *Object
	isSet bool
}

func (v NullableObject) Get() *Object {
	return v.value
}

func (v *NullableObject) Set(val *Object) {
	v.value = val
	v.isSet = true
}

func (v NullableObject) IsSet() bool {
	return v.isSet
}

func (v *NullableObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObject(val *Object) *NullableObject {
	return &NullableObject{value: val, isSet: true}
}

func (v NullableObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
