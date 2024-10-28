/*
 * Certificate Manager Service API
 *
 * Using the Certificate Manager Service, you can conveniently provision and manage SSL certificates  with IONOS services and your internal connected resources.   For the [Application Load Balancer](https://api.ionos.com/docs/cloud/v6/#Application-Load-Balancers-get-datacenters-datacenterId-applicationloadbalancers), you usually need a certificate to encrypt your HTTPS traffic. The service provides the basic functions of uploading and deleting your certificates for this purpose.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// ProviderCreate struct for ProviderCreate
type ProviderCreate struct {
	// Metadata
	Metadata   *map[string]interface{} `json:"metadata,omitempty"`
	Properties *Provider               `json:"properties"`
}

// NewProviderCreate instantiates a new ProviderCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderCreate(properties Provider) *ProviderCreate {
	this := ProviderCreate{}

	this.Properties = &properties

	return &this
}

// NewProviderCreateWithDefaults instantiates a new ProviderCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderCreateWithDefaults() *ProviderCreate {
	this := ProviderCreate{}
	return &this
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *ProviderCreate) GetMetadata() *map[string]interface{} {
	if o == nil {
		return nil
	}

	return o.Metadata

}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderCreate) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}

	return o.Metadata, true
}

// SetMetadata sets field value
func (o *ProviderCreate) SetMetadata(v map[string]interface{}) {

	o.Metadata = &v

}

// HasMetadata returns a boolean if a field has been set.
func (o *ProviderCreate) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for Provider will be returned
func (o *ProviderCreate) GetProperties() *Provider {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderCreate) GetPropertiesOk() (*Provider, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *ProviderCreate) SetProperties(v Provider) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *ProviderCreate) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

func (o ProviderCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}

	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}

	return json.Marshal(toSerialize)
}

type NullableProviderCreate struct {
	value *ProviderCreate
	isSet bool
}

func (v NullableProviderCreate) Get() *ProviderCreate {
	return v.value
}

func (v *NullableProviderCreate) Set(val *ProviderCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderCreate(val *ProviderCreate) *NullableProviderCreate {
	return &NullableProviderCreate{value: val, isSet: true}
}

func (v NullableProviderCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
