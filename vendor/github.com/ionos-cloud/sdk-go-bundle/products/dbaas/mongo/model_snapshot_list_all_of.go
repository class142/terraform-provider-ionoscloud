/*
 * IONOS DBaaS MongoDB REST API
 *
 * With IONOS Cloud Database as a Service, you have the ability to quickly set up and manage a MongoDB database. You can also delete clusters, manage backups and users via the API.   MongoDB is an open source, cross-platform, document-oriented database program. Classified as a NoSQL database program, it uses JSON-like documents with optional schemas.  The MongoDB API allows you to create additional database clusters or modify existing ones. Both tools, the Data Center Designer (DCD) and the API use the same concepts consistently and are well suited for smooth and intuitive use.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongo

import (
	"encoding/json"
)

// SnapshotListAllOf struct for SnapshotListAllOf
type SnapshotListAllOf struct {
	Type *ResourceType `json:"type,omitempty"`
	// The unique ID of the resource.
	Id    *string             `json:"id,omitempty"`
	Items *[]SnapshotResponse `json:"items,omitempty"`
}

// NewSnapshotListAllOf instantiates a new SnapshotListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotListAllOf() *SnapshotListAllOf {
	this := SnapshotListAllOf{}

	return &this
}

// NewSnapshotListAllOfWithDefaults instantiates a new SnapshotListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotListAllOfWithDefaults() *SnapshotListAllOf {
	this := SnapshotListAllOf{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for ResourceType will be returned
func (o *SnapshotListAllOf) GetType() *ResourceType {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotListAllOf) GetTypeOk() (*ResourceType, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *SnapshotListAllOf) SetType(v ResourceType) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *SnapshotListAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SnapshotListAllOf) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotListAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *SnapshotListAllOf) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *SnapshotListAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []SnapshotResponse will be returned
func (o *SnapshotListAllOf) GetItems() *[]SnapshotResponse {
	if o == nil {
		return nil
	}

	return o.Items

}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotListAllOf) GetItemsOk() (*[]SnapshotResponse, bool) {
	if o == nil {
		return nil, false
	}

	return o.Items, true
}

// SetItems sets field value
func (o *SnapshotListAllOf) SetItems(v []SnapshotResponse) {

	o.Items = &v

}

// HasItems returns a boolean if a field has been set.
func (o *SnapshotListAllOf) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

func (o SnapshotListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	if o.Items != nil {
		toSerialize["items"] = o.Items
	}

	return json.Marshal(toSerialize)
}

type NullableSnapshotListAllOf struct {
	value *SnapshotListAllOf
	isSet bool
}

func (v NullableSnapshotListAllOf) Get() *SnapshotListAllOf {
	return v.value
}

func (v *NullableSnapshotListAllOf) Set(val *SnapshotListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotListAllOf(val *SnapshotListAllOf) *NullableSnapshotListAllOf {
	return &NullableSnapshotListAllOf{value: val, isSet: true}
}

func (v NullableSnapshotListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
