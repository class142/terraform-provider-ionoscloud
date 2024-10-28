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

// DeleteMarkerEntry Information about the delete marker.
type DeleteMarkerEntry struct {
	XMLName xml.Name `xml:"DeleteMarker"`
	Owner   *Owner   `json:"Owner,omitempty" xml:"Owner"`
	// The object key.
	Key *string `json:"Key,omitempty" xml:"Key"`
	// Version ID of the Deletion Marker
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId"`
	// Specifies whether the object is (true) or is not (false) the latest version of an object.
	IsLatest *bool `json:"IsLatest,omitempty" xml:"IsLatest"`
	// Creation date of the object.
	LastModified *IonosTime `json:"LastModified,omitempty" xml:"LastModified"`
}

// NewDeleteMarkerEntry instantiates a new DeleteMarkerEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteMarkerEntry() *DeleteMarkerEntry {
	this := DeleteMarkerEntry{}

	return &this
}

// NewDeleteMarkerEntryWithDefaults instantiates a new DeleteMarkerEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteMarkerEntryWithDefaults() *DeleteMarkerEntry {
	this := DeleteMarkerEntry{}
	return &this
}

// GetOwner returns the Owner field value
// If the value is explicit nil, the zero value for Owner will be returned
func (o *DeleteMarkerEntry) GetOwner() *Owner {
	if o == nil {
		return nil
	}

	return o.Owner

}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteMarkerEntry) GetOwnerOk() (*Owner, bool) {
	if o == nil {
		return nil, false
	}

	return o.Owner, true
}

// SetOwner sets field value
func (o *DeleteMarkerEntry) SetOwner(v Owner) {

	o.Owner = &v

}

// HasOwner returns a boolean if a field has been set.
func (o *DeleteMarkerEntry) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// GetKey returns the Key field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DeleteMarkerEntry) GetKey() *string {
	if o == nil {
		return nil
	}

	return o.Key

}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteMarkerEntry) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Key, true
}

// SetKey sets field value
func (o *DeleteMarkerEntry) SetKey(v string) {

	o.Key = &v

}

// HasKey returns a boolean if a field has been set.
func (o *DeleteMarkerEntry) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// GetVersionId returns the VersionId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DeleteMarkerEntry) GetVersionId() *string {
	if o == nil {
		return nil
	}

	return o.VersionId

}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteMarkerEntry) GetVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.VersionId, true
}

// SetVersionId sets field value
func (o *DeleteMarkerEntry) SetVersionId(v string) {

	o.VersionId = &v

}

// HasVersionId returns a boolean if a field has been set.
func (o *DeleteMarkerEntry) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// GetIsLatest returns the IsLatest field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *DeleteMarkerEntry) GetIsLatest() *bool {
	if o == nil {
		return nil
	}

	return o.IsLatest

}

// GetIsLatestOk returns a tuple with the IsLatest field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteMarkerEntry) GetIsLatestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.IsLatest, true
}

// SetIsLatest sets field value
func (o *DeleteMarkerEntry) SetIsLatest(v bool) {

	o.IsLatest = &v

}

// HasIsLatest returns a boolean if a field has been set.
func (o *DeleteMarkerEntry) HasIsLatest() bool {
	if o != nil && o.IsLatest != nil {
		return true
	}

	return false
}

// GetLastModified returns the LastModified field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *DeleteMarkerEntry) GetLastModified() *time.Time {
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
func (o *DeleteMarkerEntry) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}

	if o.LastModified == nil {
		return nil, false
	}
	return &o.LastModified.Time, true

}

// SetLastModified sets field value
func (o *DeleteMarkerEntry) SetLastModified(v time.Time) {

	o.LastModified = &IonosTime{v}

}

// HasLastModified returns a boolean if a field has been set.
func (o *DeleteMarkerEntry) HasLastModified() bool {
	if o != nil && o.LastModified != nil {
		return true
	}

	return false
}

func (o DeleteMarkerEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Owner != nil {
		toSerialize["Owner"] = o.Owner
	}

	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}

	if o.VersionId != nil {
		toSerialize["VersionId"] = o.VersionId
	}

	if o.IsLatest != nil {
		toSerialize["IsLatest"] = o.IsLatest
	}

	if o.LastModified != nil {
		toSerialize["LastModified"] = o.LastModified
	}

	return json.Marshal(toSerialize)
}

type NullableDeleteMarkerEntry struct {
	value *DeleteMarkerEntry
	isSet bool
}

func (v NullableDeleteMarkerEntry) Get() *DeleteMarkerEntry {
	return v.value
}

func (v *NullableDeleteMarkerEntry) Set(val *DeleteMarkerEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteMarkerEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteMarkerEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteMarkerEntry(val *DeleteMarkerEntry) *NullableDeleteMarkerEntry {
	return &NullableDeleteMarkerEntry{value: val, isSet: true}
}

func (v NullableDeleteMarkerEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteMarkerEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
