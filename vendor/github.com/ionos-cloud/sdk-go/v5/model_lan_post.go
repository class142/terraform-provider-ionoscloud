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

// LanPost struct for LanPost
type LanPost struct {
	// The resource's unique identifier
	Id *string `json:"id,omitempty"`
	// The type of object that has been created
	Type *Type `json:"type,omitempty"`
	// URL to the object representation (absolute path)
	Href *string `json:"href,omitempty"`
	Metadata *DatacenterElementMetadata `json:"metadata,omitempty"`
	Entities *LanEntities `json:"entities,omitempty"`
	Properties *LanPropertiesPost `json:"properties"`
}



// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LanPost) GetId() *string {
	if o == nil {
		return nil
	}


	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LanPost) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Id, true
}

// SetId sets field value
func (o *LanPost) SetId(v string) {


	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *LanPost) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}


// GetType returns the Type field value
// If the value is explicit nil, the zero value for Type will be returned
func (o *LanPost) GetType() *Type {
	if o == nil {
		return nil
	}


	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LanPost) GetTypeOk() (*Type, bool) {
	if o == nil {
		return nil, false
	}


	return o.Type, true
}

// SetType sets field value
func (o *LanPost) SetType(v Type) {


	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *LanPost) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}


// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LanPost) GetHref() *string {
	if o == nil {
		return nil
	}


	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LanPost) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Href, true
}

// SetHref sets field value
func (o *LanPost) SetHref(v string) {


	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *LanPost) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}


// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for DatacenterElementMetadata will be returned
func (o *LanPost) GetMetadata() *DatacenterElementMetadata {
	if o == nil {
		return nil
	}


	return o.Metadata

}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LanPost) GetMetadataOk() (*DatacenterElementMetadata, bool) {
	if o == nil {
		return nil, false
	}


	return o.Metadata, true
}

// SetMetadata sets field value
func (o *LanPost) SetMetadata(v DatacenterElementMetadata) {


	o.Metadata = &v

}

// HasMetadata returns a boolean if a field has been set.
func (o *LanPost) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}


// GetEntities returns the Entities field value
// If the value is explicit nil, the zero value for LanEntities will be returned
func (o *LanPost) GetEntities() *LanEntities {
	if o == nil {
		return nil
	}


	return o.Entities

}

// GetEntitiesOk returns a tuple with the Entities field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LanPost) GetEntitiesOk() (*LanEntities, bool) {
	if o == nil {
		return nil, false
	}


	return o.Entities, true
}

// SetEntities sets field value
func (o *LanPost) SetEntities(v LanEntities) {


	o.Entities = &v

}

// HasEntities returns a boolean if a field has been set.
func (o *LanPost) HasEntities() bool {
	if o != nil && o.Entities != nil {
		return true
	}

	return false
}


// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for LanPropertiesPost will be returned
func (o *LanPost) GetProperties() *LanPropertiesPost {
	if o == nil {
		return nil
	}


	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LanPost) GetPropertiesOk() (*LanPropertiesPost, bool) {
	if o == nil {
		return nil, false
	}


	return o.Properties, true
}

// SetProperties sets field value
func (o *LanPost) SetProperties(v LanPropertiesPost) {


	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *LanPost) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

func (o LanPost) MarshalJSON() ([]byte, error) {
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

	if o.Entities != nil {
		toSerialize["entities"] = o.Entities
	}

	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableLanPost struct {
	value *LanPost
	isSet bool
}

func (v NullableLanPost) Get() *LanPost {
	return v.value
}

func (v *NullableLanPost) Set(val *LanPost) {
	v.value = val
	v.isSet = true
}

func (v NullableLanPost) IsSet() bool {
	return v.isSet
}

func (v *NullableLanPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLanPost(val *LanPost) *NullableLanPost {
	return &NullableLanPost{value: val, isSet: true}
}

func (v NullableLanPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLanPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


