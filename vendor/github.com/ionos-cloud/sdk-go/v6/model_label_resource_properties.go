/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// LabelResourceProperties struct for LabelResourceProperties
type LabelResourceProperties struct {
	// A label key
	Key *string `json:"key,omitempty"`
	// A label value
	Value *string `json:"value,omitempty"`
}

// GetKey returns the Key field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LabelResourceProperties) GetKey() *string {
	if o == nil {
		return nil
	}

	return o.Key

}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LabelResourceProperties) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Key, true
}

// SetKey sets field value
func (o *LabelResourceProperties) SetKey(v string) {

	o.Key = &v

}

// HasKey returns a boolean if a field has been set.
func (o *LabelResourceProperties) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LabelResourceProperties) GetValue() *string {
	if o == nil {
		return nil
	}

	return o.Value

}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LabelResourceProperties) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Value, true
}

// SetValue sets field value
func (o *LabelResourceProperties) SetValue(v string) {

	o.Value = &v

}

// HasValue returns a boolean if a field has been set.
func (o *LabelResourceProperties) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

func (o LabelResourceProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableLabelResourceProperties struct {
	value *LabelResourceProperties
	isSet bool
}

func (v NullableLabelResourceProperties) Get() *LabelResourceProperties {
	return v.value
}

func (v *NullableLabelResourceProperties) Set(val *LabelResourceProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelResourceProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelResourceProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelResourceProperties(val *LabelResourceProperties) *NullableLabelResourceProperties {
	return &NullableLabelResourceProperties{value: val, isSet: true}
}

func (v NullableLabelResourceProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelResourceProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
