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

// AccessKeyListAllOf struct for AccessKeyListAllOf
type AccessKeyListAllOf struct {
	// ID of the list of AccessKey resources.
	Id *string `json:"id"`
	// The type of the resource.
	Type *string `json:"type"`
	// The URL of the list of AccessKey resources.
	Href *string `json:"href"`
	// The list of AccessKey resources.
	Items *[]AccessKey `json:"items,omitempty"`
}

// NewAccessKeyListAllOf instantiates a new AccessKeyListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessKeyListAllOf(id string, type_ string, href string) *AccessKeyListAllOf {
	this := AccessKeyListAllOf{}

	this.Id = &id
	this.Type = &type_
	this.Href = &href

	return &this
}

// NewAccessKeyListAllOfWithDefaults instantiates a new AccessKeyListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessKeyListAllOfWithDefaults() *AccessKeyListAllOf {
	this := AccessKeyListAllOf{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AccessKeyListAllOf) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessKeyListAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *AccessKeyListAllOf) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *AccessKeyListAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AccessKeyListAllOf) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessKeyListAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *AccessKeyListAllOf) SetType(v string) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *AccessKeyListAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AccessKeyListAllOf) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessKeyListAllOf) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *AccessKeyListAllOf) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *AccessKeyListAllOf) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []AccessKey will be returned
func (o *AccessKeyListAllOf) GetItems() *[]AccessKey {
	if o == nil {
		return nil
	}

	return o.Items

}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessKeyListAllOf) GetItemsOk() (*[]AccessKey, bool) {
	if o == nil {
		return nil, false
	}

	return o.Items, true
}

// SetItems sets field value
func (o *AccessKeyListAllOf) SetItems(v []AccessKey) {

	o.Items = &v

}

// HasItems returns a boolean if a field has been set.
func (o *AccessKeyListAllOf) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

func (o AccessKeyListAllOf) MarshalJSON() ([]byte, error) {
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

type NullableAccessKeyListAllOf struct {
	value *AccessKeyListAllOf
	isSet bool
}

func (v NullableAccessKeyListAllOf) Get() *AccessKeyListAllOf {
	return v.value
}

func (v *NullableAccessKeyListAllOf) Set(val *AccessKeyListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessKeyListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessKeyListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessKeyListAllOf(val *AccessKeyListAllOf) *NullableAccessKeyListAllOf {
	return &NullableAccessKeyListAllOf{value: val, isSet: true}
}

func (v NullableAccessKeyListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessKeyListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}