/*
 * Container Registry service
 *
 * Container Registry service enables IONOS clients to manage docker and OCI compliant registries for use by their managed Kubernetes clusters. Use a Container Registry to ensure you have a privately accessed registry to efficiently support image pulls.
 *
 * API version: 1.0
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package containerregistry

import (
	"encoding/json"
)

// TokenResponse struct for TokenResponse
type TokenResponse struct {
	Href       *string              `json:"href,omitempty"`
	Id         *string              `json:"id,omitempty"`
	Metadata   *ApiResourceMetadata `json:"metadata"`
	Properties *TokenProperties     `json:"properties"`
	Type       *string              `json:"type,omitempty"`
}

// NewTokenResponse instantiates a new TokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenResponse(metadata ApiResourceMetadata, properties TokenProperties) *TokenResponse {
	this := TokenResponse{}

	this.Metadata = &metadata
	this.Properties = &properties

	return &this
}

// NewTokenResponseWithDefaults instantiates a new TokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenResponseWithDefaults() *TokenResponse {
	this := TokenResponse{}
	return &this
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TokenResponse) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenResponse) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *TokenResponse) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *TokenResponse) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TokenResponse) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *TokenResponse) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *TokenResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for ApiResourceMetadata will be returned
func (o *TokenResponse) GetMetadata() *ApiResourceMetadata {
	if o == nil {
		return nil
	}

	return o.Metadata

}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenResponse) GetMetadataOk() (*ApiResourceMetadata, bool) {
	if o == nil {
		return nil, false
	}

	return o.Metadata, true
}

// SetMetadata sets field value
func (o *TokenResponse) SetMetadata(v ApiResourceMetadata) {

	o.Metadata = &v

}

// HasMetadata returns a boolean if a field has been set.
func (o *TokenResponse) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for TokenProperties will be returned
func (o *TokenResponse) GetProperties() *TokenProperties {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenResponse) GetPropertiesOk() (*TokenProperties, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *TokenResponse) SetProperties(v TokenProperties) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *TokenResponse) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TokenResponse) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *TokenResponse) SetType(v string) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *TokenResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

func (o TokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}

	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	toSerialize["metadata"] = o.Metadata

	toSerialize["properties"] = o.Properties

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	return json.Marshal(toSerialize)
}

type NullableTokenResponse struct {
	value *TokenResponse
	isSet bool
}

func (v NullableTokenResponse) Get() *TokenResponse {
	return v.value
}

func (v *NullableTokenResponse) Set(val *TokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenResponse(val *TokenResponse) *NullableTokenResponse {
	return &NullableTokenResponse{value: val, isSet: true}
}

func (v NullableTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
