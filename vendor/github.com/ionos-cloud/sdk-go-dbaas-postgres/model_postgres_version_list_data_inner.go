/*
 * IONOS DBaaS PostgreSQL REST API
 *
 * An enterprise-grade Database is provided as a Service (DBaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.  The API allows you to create additional PostgreSQL database clusters or modify existing ones. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// PostgresVersionListDataInner struct for PostgresVersionListDataInner
type PostgresVersionListDataInner struct {
	Name *string `json:"name,omitempty"`
}

// NewPostgresVersionListDataInner instantiates a new PostgresVersionListDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostgresVersionListDataInner() *PostgresVersionListDataInner {
	this := PostgresVersionListDataInner{}

	return &this
}

// NewPostgresVersionListDataInnerWithDefaults instantiates a new PostgresVersionListDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostgresVersionListDataInnerWithDefaults() *PostgresVersionListDataInner {
	this := PostgresVersionListDataInner{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PostgresVersionListDataInner) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostgresVersionListDataInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Name, true
}

// SetName sets field value
func (o *PostgresVersionListDataInner) SetName(v string) {

	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *PostgresVersionListDataInner) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

func (o PostgresVersionListDataInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	return json.Marshal(toSerialize)
}

type NullablePostgresVersionListDataInner struct {
	value *PostgresVersionListDataInner
	isSet bool
}

func (v NullablePostgresVersionListDataInner) Get() *PostgresVersionListDataInner {
	return v.value
}

func (v *NullablePostgresVersionListDataInner) Set(val *PostgresVersionListDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePostgresVersionListDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePostgresVersionListDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostgresVersionListDataInner(val *PostgresVersionListDataInner) *NullablePostgresVersionListDataInner {
	return &NullablePostgresVersionListDataInner{value: val, isSet: true}
}

func (v NullablePostgresVersionListDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostgresVersionListDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}