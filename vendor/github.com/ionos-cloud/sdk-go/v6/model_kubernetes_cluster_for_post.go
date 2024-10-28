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

// KubernetesClusterForPost struct for KubernetesClusterForPost
type KubernetesClusterForPost struct {
	// The resource unique identifier.
	Id *string `json:"id,omitempty"`
	// The object type.
	Type *string `json:"type,omitempty"`
	// The URL to the object representation (absolute path).
	Href       *string                             `json:"href,omitempty"`
	Metadata   *DatacenterElementMetadata          `json:"metadata,omitempty"`
	Properties *KubernetesClusterPropertiesForPost `json:"properties"`
	Entities   *KubernetesClusterEntities          `json:"entities,omitempty"`
}

// NewKubernetesClusterForPost instantiates a new KubernetesClusterForPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesClusterForPost(properties KubernetesClusterPropertiesForPost) *KubernetesClusterForPost {
	this := KubernetesClusterForPost{}

	this.Properties = &properties

	return &this
}

// NewKubernetesClusterForPostWithDefaults instantiates a new KubernetesClusterForPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesClusterForPostWithDefaults() *KubernetesClusterForPost {
	this := KubernetesClusterForPost{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, nil is returned
func (o *KubernetesClusterForPost) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterForPost) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *KubernetesClusterForPost) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *KubernetesClusterForPost) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, nil is returned
func (o *KubernetesClusterForPost) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterForPost) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *KubernetesClusterForPost) SetType(v string) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *KubernetesClusterForPost) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetHref returns the Href field value
// If the value is explicit nil, nil is returned
func (o *KubernetesClusterForPost) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterForPost) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *KubernetesClusterForPost) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *KubernetesClusterForPost) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, nil is returned
func (o *KubernetesClusterForPost) GetMetadata() *DatacenterElementMetadata {
	if o == nil {
		return nil
	}

	return o.Metadata

}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterForPost) GetMetadataOk() (*DatacenterElementMetadata, bool) {
	if o == nil {
		return nil, false
	}

	return o.Metadata, true
}

// SetMetadata sets field value
func (o *KubernetesClusterForPost) SetMetadata(v DatacenterElementMetadata) {

	o.Metadata = &v

}

// HasMetadata returns a boolean if a field has been set.
func (o *KubernetesClusterForPost) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// GetProperties returns the Properties field value
// If the value is explicit nil, nil is returned
func (o *KubernetesClusterForPost) GetProperties() *KubernetesClusterPropertiesForPost {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterForPost) GetPropertiesOk() (*KubernetesClusterPropertiesForPost, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *KubernetesClusterForPost) SetProperties(v KubernetesClusterPropertiesForPost) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *KubernetesClusterForPost) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// GetEntities returns the Entities field value
// If the value is explicit nil, nil is returned
func (o *KubernetesClusterForPost) GetEntities() *KubernetesClusterEntities {
	if o == nil {
		return nil
	}

	return o.Entities

}

// GetEntitiesOk returns a tuple with the Entities field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterForPost) GetEntitiesOk() (*KubernetesClusterEntities, bool) {
	if o == nil {
		return nil, false
	}

	return o.Entities, true
}

// SetEntities sets field value
func (o *KubernetesClusterForPost) SetEntities(v KubernetesClusterEntities) {

	o.Entities = &v

}

// HasEntities returns a boolean if a field has been set.
func (o *KubernetesClusterForPost) HasEntities() bool {
	if o != nil && o.Entities != nil {
		return true
	}

	return false
}

func (o KubernetesClusterForPost) MarshalJSON() ([]byte, error) {
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

	if o.Entities != nil {
		toSerialize["entities"] = o.Entities
	}

	return json.Marshal(toSerialize)
}

type NullableKubernetesClusterForPost struct {
	value *KubernetesClusterForPost
	isSet bool
}

func (v NullableKubernetesClusterForPost) Get() *KubernetesClusterForPost {
	return v.value
}

func (v *NullableKubernetesClusterForPost) Set(val *KubernetesClusterForPost) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesClusterForPost) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesClusterForPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesClusterForPost(val *KubernetesClusterForPost) *NullableKubernetesClusterForPost {
	return &NullableKubernetesClusterForPost{value: val, isSet: true}
}

func (v NullableKubernetesClusterForPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesClusterForPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
