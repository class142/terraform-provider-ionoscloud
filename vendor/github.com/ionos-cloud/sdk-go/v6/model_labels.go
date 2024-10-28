/*
 * CLOUD API
 *
 *  IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// Labels struct for Labels
type Labels struct {
	// A unique representation of the label as a resource collection.
	Id *string `json:"id,omitempty"`
	// The type of resource within a collection.
	Type *string `json:"type,omitempty"`
	// URL to the collection representation (absolute path).
	Href *string `json:"href,omitempty"`
	// Array of items in the collection.
	Items *[]Label `json:"items,omitempty"`
}

// NewLabels instantiates a new Labels object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabels() *Labels {
	this := Labels{}

	return &this
}

// NewLabelsWithDefaults instantiates a new Labels object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelsWithDefaults() *Labels {
	this := Labels{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, nil is returned
func (o *Labels) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Labels) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *Labels) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *Labels) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, nil is returned
func (o *Labels) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Labels) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *Labels) SetType(v string) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *Labels) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetHref returns the Href field value
// If the value is explicit nil, nil is returned
func (o *Labels) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Labels) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *Labels) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *Labels) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetItems returns the Items field value
// If the value is explicit nil, nil is returned
func (o *Labels) GetItems() *[]Label {
	if o == nil {
		return nil
	}

	return o.Items

}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Labels) GetItemsOk() (*[]Label, bool) {
	if o == nil {
		return nil, false
	}

	return o.Items, true
}

// SetItems sets field value
func (o *Labels) SetItems(v []Label) {

	o.Items = &v

}

// HasItems returns a boolean if a field has been set.
func (o *Labels) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

func (o Labels) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	if o.Href != nil {
		toSerialize["href"] = o.Href
	}

	if o.Items != nil {
		toSerialize["items"] = o.Items
	}

	return json.Marshal(toSerialize)
}

type NullableLabels struct {
	value *Labels
	isSet bool
}

func (v NullableLabels) Get() *Labels {
	return v.value
}

func (v *NullableLabels) Set(val *Labels) {
	v.value = val
	v.isSet = true
}

func (v NullableLabels) IsSet() bool {
	return v.isSet
}

func (v *NullableLabels) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabels(val *Labels) *NullableLabels {
	return &NullableLabels{value: val, isSet: true}
}

func (v NullableLabels) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabels) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
