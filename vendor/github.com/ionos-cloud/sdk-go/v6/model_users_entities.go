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

// UsersEntities struct for UsersEntities
type UsersEntities struct {
	Owns   *ResourcesUsers `json:"owns,omitempty"`
	Groups *GroupUsers     `json:"groups,omitempty"`
}

// NewUsersEntities instantiates a new UsersEntities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersEntities() *UsersEntities {
	this := UsersEntities{}

	return &this
}

// NewUsersEntitiesWithDefaults instantiates a new UsersEntities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersEntitiesWithDefaults() *UsersEntities {
	this := UsersEntities{}
	return &this
}

// GetOwns returns the Owns field value
// If the value is explicit nil, nil is returned
func (o *UsersEntities) GetOwns() *ResourcesUsers {
	if o == nil {
		return nil
	}

	return o.Owns

}

// GetOwnsOk returns a tuple with the Owns field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsersEntities) GetOwnsOk() (*ResourcesUsers, bool) {
	if o == nil {
		return nil, false
	}

	return o.Owns, true
}

// SetOwns sets field value
func (o *UsersEntities) SetOwns(v ResourcesUsers) {

	o.Owns = &v

}

// HasOwns returns a boolean if a field has been set.
func (o *UsersEntities) HasOwns() bool {
	if o != nil && o.Owns != nil {
		return true
	}

	return false
}

// GetGroups returns the Groups field value
// If the value is explicit nil, nil is returned
func (o *UsersEntities) GetGroups() *GroupUsers {
	if o == nil {
		return nil
	}

	return o.Groups

}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsersEntities) GetGroupsOk() (*GroupUsers, bool) {
	if o == nil {
		return nil, false
	}

	return o.Groups, true
}

// SetGroups sets field value
func (o *UsersEntities) SetGroups(v GroupUsers) {

	o.Groups = &v

}

// HasGroups returns a boolean if a field has been set.
func (o *UsersEntities) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

func (o UsersEntities) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Owns != nil {
		toSerialize["owns"] = o.Owns
	}

	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}

	return json.Marshal(toSerialize)
}

type NullableUsersEntities struct {
	value *UsersEntities
	isSet bool
}

func (v NullableUsersEntities) Get() *UsersEntities {
	return v.value
}

func (v *NullableUsersEntities) Set(val *UsersEntities) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersEntities) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersEntities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersEntities(val *UsersEntities) *NullableUsersEntities {
	return &NullableUsersEntities{value: val, isSet: true}
}

func (v NullableUsersEntities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersEntities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
