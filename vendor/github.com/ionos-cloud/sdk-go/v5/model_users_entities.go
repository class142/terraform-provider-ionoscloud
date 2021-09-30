/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// UsersEntities struct for UsersEntities
type UsersEntities struct {
	Owns *ResourcesUsers `json:"owns,omitempty"`
	Groups *GroupUsers `json:"groups,omitempty"`
}



// GetOwns returns the Owns field value
// If the value is explicit nil, the zero value for ResourcesUsers will be returned
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
// If the value is explicit nil, the zero value for GroupUsers will be returned
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


