/*
 * Container Registry service
 *
 * ## Overview Container Registry service enables IONOS clients to manage docker and OCI compliant registries for use by their managed Kubernetes clusters. Use a Container Registry to ensure you have a privately accessed registry to efficiently support image pulls. ## Changelog ### 1.1.0  - Added new endpoints for Repositories  - Added new endpoints for Artifacts  - Added new endpoints for Vulnerabilities  - Added registry vulnerabilityScanning feature
 *
 * API version: 1.1.0
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// RepositoryRead struct for RepositoryRead
type RepositoryRead struct {
	Id         *string             `json:"id"`
	Type       *string             `json:"type"`
	Href       *string             `json:"href"`
	Metadata   *RepositoryMetadata `json:"metadata"`
	Properties *Repository         `json:"properties"`
}

// NewRepositoryRead instantiates a new RepositoryRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryRead(id string, type_ string, href string, metadata RepositoryMetadata, properties Repository) *RepositoryRead {
	this := RepositoryRead{}

	this.Id = &id
	this.Type = &type_
	this.Href = &href
	this.Metadata = &metadata
	this.Properties = &properties

	return &this
}

// NewRepositoryReadWithDefaults instantiates a new RepositoryRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryReadWithDefaults() *RepositoryRead {
	this := RepositoryRead{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RepositoryRead) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryRead) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *RepositoryRead) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *RepositoryRead) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RepositoryRead) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryRead) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *RepositoryRead) SetType(v string) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *RepositoryRead) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RepositoryRead) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryRead) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *RepositoryRead) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *RepositoryRead) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for RepositoryMetadata will be returned
func (o *RepositoryRead) GetMetadata() *RepositoryMetadata {
	if o == nil {
		return nil
	}

	return o.Metadata

}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryRead) GetMetadataOk() (*RepositoryMetadata, bool) {
	if o == nil {
		return nil, false
	}

	return o.Metadata, true
}

// SetMetadata sets field value
func (o *RepositoryRead) SetMetadata(v RepositoryMetadata) {

	o.Metadata = &v

}

// HasMetadata returns a boolean if a field has been set.
func (o *RepositoryRead) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for Repository will be returned
func (o *RepositoryRead) GetProperties() *Repository {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryRead) GetPropertiesOk() (*Repository, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *RepositoryRead) SetProperties(v Repository) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *RepositoryRead) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

func (o RepositoryRead) MarshalJSON() ([]byte, error) {
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

	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}

	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}

	return json.Marshal(toSerialize)
}

type NullableRepositoryRead struct {
	value *RepositoryRead
	isSet bool
}

func (v NullableRepositoryRead) Get() *RepositoryRead {
	return v.value
}

func (v *NullableRepositoryRead) Set(val *RepositoryRead) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryRead) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryRead(val *RepositoryRead) *NullableRepositoryRead {
	return &NullableRepositoryRead{value: val, isSet: true}
}

func (v NullableRepositoryRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}