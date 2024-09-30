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

// StorageClassListAllOf struct for StorageClassListAllOf
type StorageClassListAllOf struct {
	// ID of the list of StorageClass resources.
	Id *string `json:"id"`
	// The type of the resource.
	Type *string `json:"type"`
	// The URL of the list of StorageClass resources.
	Href *string `json:"href"`
	// The list of StorageClass resources.
	Items *[]StorageClass `json:"items,omitempty"`
}

// NewStorageClassListAllOf instantiates a new StorageClassListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageClassListAllOf(id string, type_ string, href string) *StorageClassListAllOf {
	this := StorageClassListAllOf{}

	this.Id = &id
	this.Type = &type_
	this.Href = &href

	return &this
}

// NewStorageClassListAllOfWithDefaults instantiates a new StorageClassListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageClassListAllOfWithDefaults() *StorageClassListAllOf {
	this := StorageClassListAllOf{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StorageClassListAllOf) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageClassListAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *StorageClassListAllOf) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *StorageClassListAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StorageClassListAllOf) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageClassListAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *StorageClassListAllOf) SetType(v string) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *StorageClassListAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StorageClassListAllOf) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageClassListAllOf) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *StorageClassListAllOf) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *StorageClassListAllOf) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []StorageClass will be returned
func (o *StorageClassListAllOf) GetItems() *[]StorageClass {
	if o == nil {
		return nil
	}

	return o.Items

}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageClassListAllOf) GetItemsOk() (*[]StorageClass, bool) {
	if o == nil {
		return nil, false
	}

	return o.Items, true
}

// SetItems sets field value
func (o *StorageClassListAllOf) SetItems(v []StorageClass) {

	o.Items = &v

}

// HasItems returns a boolean if a field has been set.
func (o *StorageClassListAllOf) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

func (o StorageClassListAllOf) MarshalJSON() ([]byte, error) {
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

type NullableStorageClassListAllOf struct {
	value *StorageClassListAllOf
	isSet bool
}

func (v NullableStorageClassListAllOf) Get() *StorageClassListAllOf {
	return v.value
}

func (v *NullableStorageClassListAllOf) Set(val *StorageClassListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageClassListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageClassListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageClassListAllOf(val *StorageClassListAllOf) *NullableStorageClassListAllOf {
	return &NullableStorageClassListAllOf{value: val, isSet: true}
}

func (v NullableStorageClassListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageClassListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}