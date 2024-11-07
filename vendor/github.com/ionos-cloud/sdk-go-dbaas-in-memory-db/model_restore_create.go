/*
 * In-Memory DB API
 *
 * API description for the IONOS In-Memory DB
 *
 * API version: 1.0.0
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// RestoreCreate struct for RestoreCreate
type RestoreCreate struct {
	// Metadata
	Metadata   *map[string]interface{} `json:"metadata,omitempty"`
	Properties *Restore                `json:"properties"`
}

// NewRestoreCreate instantiates a new RestoreCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoreCreate(properties Restore) *RestoreCreate {
	this := RestoreCreate{}

	this.Properties = &properties

	return &this
}

// NewRestoreCreateWithDefaults instantiates a new RestoreCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreCreateWithDefaults() *RestoreCreate {
	this := RestoreCreate{}
	return &this
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *RestoreCreate) GetMetadata() *map[string]interface{} {
	if o == nil {
		return nil
	}

	return o.Metadata

}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreCreate) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}

	return o.Metadata, true
}

// SetMetadata sets field value
func (o *RestoreCreate) SetMetadata(v map[string]interface{}) {

	o.Metadata = &v

}

// HasMetadata returns a boolean if a field has been set.
func (o *RestoreCreate) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for Restore will be returned
func (o *RestoreCreate) GetProperties() *Restore {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreCreate) GetPropertiesOk() (*Restore, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *RestoreCreate) SetProperties(v Restore) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *RestoreCreate) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

func (o RestoreCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}

	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}

	return json.Marshal(toSerialize)
}

type NullableRestoreCreate struct {
	value *RestoreCreate
	isSet bool
}

func (v NullableRestoreCreate) Get() *RestoreCreate {
	return v.value
}

func (v *NullableRestoreCreate) Set(val *RestoreCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreCreate(val *RestoreCreate) *NullableRestoreCreate {
	return &NullableRestoreCreate{value: val, isSet: true}
}

func (v NullableRestoreCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}