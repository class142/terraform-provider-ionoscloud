/*
 * IONOS Logging REST API
 *
 * The logging service offers a centralized platform to collect and store logs from various systems and applications. It includes tools to search, filter, visualize, and create alerts based on your log data.  This API provides programmatic control over logging pipelines, enabling you to create new pipelines or modify existing ones. It mirrors the functionality of the DCD visual tool, ensuring a consistent experience regardless of your chosen interface.
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logging

import (
	"encoding/json"
)

// checks if the Destination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Destination{}

// Destination The information of the logging aggregator storage
type Destination struct {
	// The internal output stream to send logs to
	Type *string `json:"type,omitempty"`
	// defines the number of days a log record should be kept in loki. Works with loki destination type only.
	RetentionInDays *int32 `json:"retentionInDays,omitempty"`
}

// NewDestination instantiates a new Destination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestination() *Destination {
	this := Destination{}

	var type_ string = "loki"
	this.Type = &type_
	var retentionInDays int32 = 30
	this.RetentionInDays = &retentionInDays

	return &this
}

// NewDestinationWithDefaults instantiates a new Destination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationWithDefaults() *Destination {
	this := Destination{}
	var type_ string = "loki"
	this.Type = &type_
	var retentionInDays int32 = 30
	this.RetentionInDays = &retentionInDays
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Destination) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Destination) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Destination) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Destination) SetType(v string) {
	o.Type = &v
}

// GetRetentionInDays returns the RetentionInDays field value if set, zero value otherwise.
func (o *Destination) GetRetentionInDays() int32 {
	if o == nil || IsNil(o.RetentionInDays) {
		var ret int32
		return ret
	}
	return *o.RetentionInDays
}

// GetRetentionInDaysOk returns a tuple with the RetentionInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Destination) GetRetentionInDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.RetentionInDays) {
		return nil, false
	}
	return o.RetentionInDays, true
}

// HasRetentionInDays returns a boolean if a field has been set.
func (o *Destination) HasRetentionInDays() bool {
	if o != nil && !IsNil(o.RetentionInDays) {
		return true
	}

	return false
}

// SetRetentionInDays gets a reference to the given int32 and assigns it to the RetentionInDays field.
func (o *Destination) SetRetentionInDays(v int32) {
	o.RetentionInDays = &v
}

func (o Destination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.RetentionInDays) {
		toSerialize["retentionInDays"] = o.RetentionInDays
	}
	return toSerialize, nil
}

type NullableDestination struct {
	value *Destination
	isSet bool
}

func (v NullableDestination) Get() *Destination {
	return v.value
}

func (v *NullableDestination) Set(val *Destination) {
	v.value = val
	v.isSet = true
}

func (v NullableDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestination(val *Destination) *NullableDestination {
	return &NullableDestination{value: val, isSet: true}
}

func (v NullableDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}