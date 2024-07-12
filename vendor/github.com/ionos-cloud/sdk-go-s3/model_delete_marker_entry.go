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

	"time"
)

import "encoding/xml"

// checks if the DeleteMarkerEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteMarkerEntry{}

// DeleteMarkerEntry Information about the delete marker.
type DeleteMarkerEntry struct {
	XMLName xml.Name `xml:"DeleteMarkerEntry"`
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

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *DeleteMarkerEntry) GetOwner() Owner {
	if o == nil || IsNil(o.Owner) {
		var ret Owner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteMarkerEntry) GetOwnerOk() (*Owner, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *DeleteMarkerEntry) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given Owner and assigns it to the Owner field.
func (o *DeleteMarkerEntry) SetOwner(v Owner) {
	o.Owner = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *DeleteMarkerEntry) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteMarkerEntry) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *DeleteMarkerEntry) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *DeleteMarkerEntry) SetKey(v string) {
	o.Key = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *DeleteMarkerEntry) GetVersionId() string {
	if o == nil || IsNil(o.VersionId) {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteMarkerEntry) GetVersionIdOk() (*string, bool) {
	if o == nil || IsNil(o.VersionId) {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *DeleteMarkerEntry) HasVersionId() bool {
	if o != nil && !IsNil(o.VersionId) {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *DeleteMarkerEntry) SetVersionId(v string) {
	o.VersionId = &v
}

// GetIsLatest returns the IsLatest field value if set, zero value otherwise.
func (o *DeleteMarkerEntry) GetIsLatest() bool {
	if o == nil || IsNil(o.IsLatest) {
		var ret bool
		return ret
	}
	return *o.IsLatest
}

// GetIsLatestOk returns a tuple with the IsLatest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteMarkerEntry) GetIsLatestOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLatest) {
		return nil, false
	}
	return o.IsLatest, true
}

// HasIsLatest returns a boolean if a field has been set.
func (o *DeleteMarkerEntry) HasIsLatest() bool {
	if o != nil && !IsNil(o.IsLatest) {
		return true
	}

	return false
}

// SetIsLatest gets a reference to the given bool and assigns it to the IsLatest field.
func (o *DeleteMarkerEntry) SetIsLatest(v bool) {
	o.IsLatest = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *DeleteMarkerEntry) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return o.LastModified.Time
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteMarkerEntry) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return &o.LastModified.Time, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *DeleteMarkerEntry) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *DeleteMarkerEntry) SetLastModified(v time.Time) {
	o.LastModified = &IonosTime{v}
}

func (o DeleteMarkerEntry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteMarkerEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Owner) {
		toSerialize["Owner"] = o.Owner
	}
	if !IsNil(o.Key) {
		toSerialize["Key"] = o.Key
	}
	if !IsNil(o.VersionId) {
		toSerialize["VersionId"] = o.VersionId
	}
	if !IsNil(o.IsLatest) {
		toSerialize["IsLatest"] = o.IsLatest
	}
	if !IsNil(o.LastModified) {
		toSerialize["LastModified"] = o.LastModified
	}
	return toSerialize, nil
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