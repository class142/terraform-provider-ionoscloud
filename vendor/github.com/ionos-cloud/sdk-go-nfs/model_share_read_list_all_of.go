/*
 * IONOS Cloud - Network File Storage API
 *
 * The RESTful API for managing Network File Storage.
 *
 * API version: 0.1.1
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// ShareReadListAllOf struct for ShareReadListAllOf
type ShareReadListAllOf struct {
	// The Share identifier (UUID)
	Id *string `json:"id"`
	// The resource type
	Type *string `json:"type"`
	// The URL of the Share.
	Href *string `json:"href"`
	// The list of share resources.
	Items *[]ShareRead `json:"items,omitempty"`
}

// NewShareReadListAllOf instantiates a new ShareReadListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShareReadListAllOf(id string, type_ string, href string) *ShareReadListAllOf {
	this := ShareReadListAllOf{}

	this.Id = &id
	this.Type = &type_
	this.Href = &href

	return &this
}

// NewShareReadListAllOfWithDefaults instantiates a new ShareReadListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShareReadListAllOfWithDefaults() *ShareReadListAllOf {
	this := ShareReadListAllOf{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ShareReadListAllOf) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShareReadListAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *ShareReadListAllOf) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *ShareReadListAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ShareReadListAllOf) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShareReadListAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *ShareReadListAllOf) SetType(v string) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *ShareReadListAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ShareReadListAllOf) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShareReadListAllOf) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *ShareReadListAllOf) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *ShareReadListAllOf) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []ShareRead will be returned
func (o *ShareReadListAllOf) GetItems() *[]ShareRead {
	if o == nil {
		return nil
	}

	return o.Items

}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShareReadListAllOf) GetItemsOk() (*[]ShareRead, bool) {
	if o == nil {
		return nil, false
	}

	return o.Items, true
}

// SetItems sets field value
func (o *ShareReadListAllOf) SetItems(v []ShareRead) {

	o.Items = &v

}

// HasItems returns a boolean if a field has been set.
func (o *ShareReadListAllOf) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

func (o ShareReadListAllOf) MarshalJSON() ([]byte, error) {
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

type NullableShareReadListAllOf struct {
	value *ShareReadListAllOf
	isSet bool
}

func (v NullableShareReadListAllOf) Get() *ShareReadListAllOf {
	return v.value
}

func (v *NullableShareReadListAllOf) Set(val *ShareReadListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableShareReadListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableShareReadListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShareReadListAllOf(val *ShareReadListAllOf) *NullableShareReadListAllOf {
	return &NullableShareReadListAllOf{value: val, isSet: true}
}

func (v NullableShareReadListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShareReadListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
