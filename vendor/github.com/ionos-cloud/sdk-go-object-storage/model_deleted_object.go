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

// DeletedObject Information about the deleted object.
type DeletedObject struct {
	XMLName xml.Name `xml:"Deleted"`
	// The object key.
	Key *string `json:"Key,omitempty" xml:"Key"`
	// Version ID of the deleted object
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId"`
	// Specifies whether the versioned object that was permanently deleted was (true) or was not (false) a delete marker. In a simple DELETE, this header indicates whether (true) or not (false) a delete marker was created.
	DeleteMarker *bool `json:"DeleteMarker,omitempty" xml:"DeleteMarker"`
	// The version ID of the delete marker created as a result of the DELETE operation. If you delete a specific object version, the value returned by this header is the version ID of the object version deleted.
	DeleteMarkerVersionId *string `json:"DeleteMarkerVersionId,omitempty" xml:"DeleteMarkerVersionId"`
}

// NewDeletedObject instantiates a new DeletedObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeletedObject() *DeletedObject {
	this := DeletedObject{}

	return &this
}

// NewDeletedObjectWithDefaults instantiates a new DeletedObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeletedObjectWithDefaults() *DeletedObject {
	this := DeletedObject{}
	return &this
}

// GetKey returns the Key field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DeletedObject) GetKey() *string {
	if o == nil {
		return nil
	}

	return o.Key

}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeletedObject) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Key, true
}

// SetKey sets field value
func (o *DeletedObject) SetKey(v string) {

	o.Key = &v

}

// HasKey returns a boolean if a field has been set.
func (o *DeletedObject) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// GetVersionId returns the VersionId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DeletedObject) GetVersionId() *string {
	if o == nil {
		return nil
	}

	return o.VersionId

}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeletedObject) GetVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.VersionId, true
}

// SetVersionId sets field value
func (o *DeletedObject) SetVersionId(v string) {

	o.VersionId = &v

}

// HasVersionId returns a boolean if a field has been set.
func (o *DeletedObject) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// GetDeleteMarker returns the DeleteMarker field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *DeletedObject) GetDeleteMarker() *bool {
	if o == nil {
		return nil
	}

	return o.DeleteMarker

}

// GetDeleteMarkerOk returns a tuple with the DeleteMarker field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeletedObject) GetDeleteMarkerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.DeleteMarker, true
}

// SetDeleteMarker sets field value
func (o *DeletedObject) SetDeleteMarker(v bool) {

	o.DeleteMarker = &v

}

// HasDeleteMarker returns a boolean if a field has been set.
func (o *DeletedObject) HasDeleteMarker() bool {
	if o != nil && o.DeleteMarker != nil {
		return true
	}

	return false
}

// GetDeleteMarkerVersionId returns the DeleteMarkerVersionId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DeletedObject) GetDeleteMarkerVersionId() *string {
	if o == nil {
		return nil
	}

	return o.DeleteMarkerVersionId

}

// GetDeleteMarkerVersionIdOk returns a tuple with the DeleteMarkerVersionId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeletedObject) GetDeleteMarkerVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.DeleteMarkerVersionId, true
}

// SetDeleteMarkerVersionId sets field value
func (o *DeletedObject) SetDeleteMarkerVersionId(v string) {

	o.DeleteMarkerVersionId = &v

}

// HasDeleteMarkerVersionId returns a boolean if a field has been set.
func (o *DeletedObject) HasDeleteMarkerVersionId() bool {
	if o != nil && o.DeleteMarkerVersionId != nil {
		return true
	}

	return false
}

func (o DeletedObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}

	if o.VersionId != nil {
		toSerialize["VersionId"] = o.VersionId
	}

	if o.DeleteMarker != nil {
		toSerialize["DeleteMarker"] = o.DeleteMarker
	}

	if o.DeleteMarkerVersionId != nil {
		toSerialize["DeleteMarkerVersionId"] = o.DeleteMarkerVersionId
	}

	return json.Marshal(toSerialize)
}

type NullableDeletedObject struct {
	value *DeletedObject
	isSet bool
}

func (v NullableDeletedObject) Get() *DeletedObject {
	return v.value
}

func (v *NullableDeletedObject) Set(val *DeletedObject) {
	v.value = val
	v.isSet = true
}

func (v NullableDeletedObject) IsSet() bool {
	return v.isSet
}

func (v *NullableDeletedObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeletedObject(val *DeletedObject) *NullableDeletedObject {
	return &NullableDeletedObject{value: val, isSet: true}
}

func (v NullableDeletedObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeletedObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
