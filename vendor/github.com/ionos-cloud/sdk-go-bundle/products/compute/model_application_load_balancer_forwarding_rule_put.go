/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package compute

import (
	"encoding/json"
)

// ApplicationLoadBalancerForwardingRulePut struct for ApplicationLoadBalancerForwardingRulePut
type ApplicationLoadBalancerForwardingRulePut struct {
	// The resource's unique identifier.
	Id *string `json:"id,omitempty"`
	// The type of object that has been created.
	Type *Type `json:"type,omitempty"`
	// URL to the object representation (absolute path).
	Href       *string                                          `json:"href,omitempty"`
	Properties *ApplicationLoadBalancerForwardingRuleProperties `json:"properties"`
}

// NewApplicationLoadBalancerForwardingRulePut instantiates a new ApplicationLoadBalancerForwardingRulePut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationLoadBalancerForwardingRulePut(properties ApplicationLoadBalancerForwardingRuleProperties) *ApplicationLoadBalancerForwardingRulePut {
	this := ApplicationLoadBalancerForwardingRulePut{}

	this.Properties = &properties

	return &this
}

// NewApplicationLoadBalancerForwardingRulePutWithDefaults instantiates a new ApplicationLoadBalancerForwardingRulePut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationLoadBalancerForwardingRulePutWithDefaults() *ApplicationLoadBalancerForwardingRulePut {
	this := ApplicationLoadBalancerForwardingRulePut{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ApplicationLoadBalancerForwardingRulePut) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationLoadBalancerForwardingRulePut) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *ApplicationLoadBalancerForwardingRulePut) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *ApplicationLoadBalancerForwardingRulePut) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for Type will be returned
func (o *ApplicationLoadBalancerForwardingRulePut) GetType() *Type {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationLoadBalancerForwardingRulePut) GetTypeOk() (*Type, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *ApplicationLoadBalancerForwardingRulePut) SetType(v Type) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *ApplicationLoadBalancerForwardingRulePut) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ApplicationLoadBalancerForwardingRulePut) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationLoadBalancerForwardingRulePut) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *ApplicationLoadBalancerForwardingRulePut) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *ApplicationLoadBalancerForwardingRulePut) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for ApplicationLoadBalancerForwardingRuleProperties will be returned
func (o *ApplicationLoadBalancerForwardingRulePut) GetProperties() *ApplicationLoadBalancerForwardingRuleProperties {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationLoadBalancerForwardingRulePut) GetPropertiesOk() (*ApplicationLoadBalancerForwardingRuleProperties, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *ApplicationLoadBalancerForwardingRulePut) SetProperties(v ApplicationLoadBalancerForwardingRuleProperties) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *ApplicationLoadBalancerForwardingRulePut) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

func (o ApplicationLoadBalancerForwardingRulePut) MarshalJSON() ([]byte, error) {
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

	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}

	return json.Marshal(toSerialize)
}

type NullableApplicationLoadBalancerForwardingRulePut struct {
	value *ApplicationLoadBalancerForwardingRulePut
	isSet bool
}

func (v NullableApplicationLoadBalancerForwardingRulePut) Get() *ApplicationLoadBalancerForwardingRulePut {
	return v.value
}

func (v *NullableApplicationLoadBalancerForwardingRulePut) Set(val *ApplicationLoadBalancerForwardingRulePut) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationLoadBalancerForwardingRulePut) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationLoadBalancerForwardingRulePut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationLoadBalancerForwardingRulePut(val *ApplicationLoadBalancerForwardingRulePut) *NullableApplicationLoadBalancerForwardingRulePut {
	return &NullableApplicationLoadBalancerForwardingRulePut{value: val, isSet: true}
}

func (v NullableApplicationLoadBalancerForwardingRulePut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationLoadBalancerForwardingRulePut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
