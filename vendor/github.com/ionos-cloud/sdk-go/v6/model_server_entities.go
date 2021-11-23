/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0-SDK.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// ServerEntities struct for ServerEntities
type ServerEntities struct {
	Cdroms  *Cdroms          `json:"cdroms,omitempty"`
	Volumes *AttachedVolumes `json:"volumes,omitempty"`
	Nics    *Nics            `json:"nics,omitempty"`
}

// GetCdroms returns the Cdroms field value
// If the value is explicit nil, the zero value for Cdroms will be returned
func (o *ServerEntities) GetCdroms() *Cdroms {
	if o == nil {
		return nil
	}

	return o.Cdroms

}

// GetCdromsOk returns a tuple with the Cdroms field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerEntities) GetCdromsOk() (*Cdroms, bool) {
	if o == nil {
		return nil, false
	}

	return o.Cdroms, true
}

// SetCdroms sets field value
func (o *ServerEntities) SetCdroms(v Cdroms) {

	o.Cdroms = &v

}

// HasCdroms returns a boolean if a field has been set.
func (o *ServerEntities) HasCdroms() bool {
	if o != nil && o.Cdroms != nil {
		return true
	}

	return false
}

// GetVolumes returns the Volumes field value
// If the value is explicit nil, the zero value for AttachedVolumes will be returned
func (o *ServerEntities) GetVolumes() *AttachedVolumes {
	if o == nil {
		return nil
	}

	return o.Volumes

}

// GetVolumesOk returns a tuple with the Volumes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerEntities) GetVolumesOk() (*AttachedVolumes, bool) {
	if o == nil {
		return nil, false
	}

	return o.Volumes, true
}

// SetVolumes sets field value
func (o *ServerEntities) SetVolumes(v AttachedVolumes) {

	o.Volumes = &v

}

// HasVolumes returns a boolean if a field has been set.
func (o *ServerEntities) HasVolumes() bool {
	if o != nil && o.Volumes != nil {
		return true
	}

	return false
}

// GetNics returns the Nics field value
// If the value is explicit nil, the zero value for Nics will be returned
func (o *ServerEntities) GetNics() *Nics {
	if o == nil {
		return nil
	}

	return o.Nics

}

// GetNicsOk returns a tuple with the Nics field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerEntities) GetNicsOk() (*Nics, bool) {
	if o == nil {
		return nil, false
	}

	return o.Nics, true
}

// SetNics sets field value
func (o *ServerEntities) SetNics(v Nics) {

	o.Nics = &v

}

// HasNics returns a boolean if a field has been set.
func (o *ServerEntities) HasNics() bool {
	if o != nil && o.Nics != nil {
		return true
	}

	return false
}

func (o ServerEntities) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Cdroms != nil {
		toSerialize["cdroms"] = o.Cdroms
	}

	if o.Volumes != nil {
		toSerialize["volumes"] = o.Volumes
	}

	if o.Nics != nil {
		toSerialize["nics"] = o.Nics
	}
	return json.Marshal(toSerialize)
}

type NullableServerEntities struct {
	value *ServerEntities
	isSet bool
}

func (v NullableServerEntities) Get() *ServerEntities {
	return v.value
}

func (v *NullableServerEntities) Set(val *ServerEntities) {
	v.value = val
	v.isSet = true
}

func (v NullableServerEntities) IsSet() bool {
	return v.isSet
}

func (v *NullableServerEntities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerEntities(val *ServerEntities) *NullableServerEntities {
	return &NullableServerEntities{value: val, isSet: true}
}

func (v NullableServerEntities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerEntities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}