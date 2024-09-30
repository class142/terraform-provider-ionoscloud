/*
 * IONOS Cloud - S3 Management API
 *
 * S3 Management API is a RESTful API that manages the S3 service configuration for IONOS Cloud.
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// MetadataWithSupportedRegionsAllOf struct for MetadataWithSupportedRegionsAllOf
type MetadataWithSupportedRegionsAllOf struct {
	// The list of supported regions.
	SupportedRegions *[]string `json:"supportedRegions"`
}

// NewMetadataWithSupportedRegionsAllOf instantiates a new MetadataWithSupportedRegionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataWithSupportedRegionsAllOf(supportedRegions []string) *MetadataWithSupportedRegionsAllOf {
	this := MetadataWithSupportedRegionsAllOf{}

	this.SupportedRegions = &supportedRegions

	return &this
}

// NewMetadataWithSupportedRegionsAllOfWithDefaults instantiates a new MetadataWithSupportedRegionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataWithSupportedRegionsAllOfWithDefaults() *MetadataWithSupportedRegionsAllOf {
	this := MetadataWithSupportedRegionsAllOf{}
	return &this
}

// GetSupportedRegions returns the SupportedRegions field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *MetadataWithSupportedRegionsAllOf) GetSupportedRegions() *[]string {
	if o == nil {
		return nil
	}

	return o.SupportedRegions

}

// GetSupportedRegionsOk returns a tuple with the SupportedRegions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataWithSupportedRegionsAllOf) GetSupportedRegionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}

	return o.SupportedRegions, true
}

// SetSupportedRegions sets field value
func (o *MetadataWithSupportedRegionsAllOf) SetSupportedRegions(v []string) {

	o.SupportedRegions = &v

}

// HasSupportedRegions returns a boolean if a field has been set.
func (o *MetadataWithSupportedRegionsAllOf) HasSupportedRegions() bool {
	if o != nil && o.SupportedRegions != nil {
		return true
	}

	return false
}

func (o MetadataWithSupportedRegionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SupportedRegions != nil {
		toSerialize["supportedRegions"] = o.SupportedRegions
	}

	return json.Marshal(toSerialize)
}

type NullableMetadataWithSupportedRegionsAllOf struct {
	value *MetadataWithSupportedRegionsAllOf
	isSet bool
}

func (v NullableMetadataWithSupportedRegionsAllOf) Get() *MetadataWithSupportedRegionsAllOf {
	return v.value
}

func (v *NullableMetadataWithSupportedRegionsAllOf) Set(val *MetadataWithSupportedRegionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataWithSupportedRegionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataWithSupportedRegionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataWithSupportedRegionsAllOf(val *MetadataWithSupportedRegionsAllOf) *NullableMetadataWithSupportedRegionsAllOf {
	return &NullableMetadataWithSupportedRegionsAllOf{value: val, isSet: true}
}

func (v NullableMetadataWithSupportedRegionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataWithSupportedRegionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
