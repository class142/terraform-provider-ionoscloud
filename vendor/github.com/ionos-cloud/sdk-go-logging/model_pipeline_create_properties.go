/*
 * IONOS Logging REST API
 *
 * Logging as a Service (LaaS) is a service that provides a centralized logging system where users are able to push and aggregate their system or application logs. This service also provides a visualization platform where users are able to observe, search and filter the logs and also create dashboards and alerts for their data points. This service can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an API. The API allows you to create logging pipelines or modify existing ones. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// PipelineCreateProperties Create pipeline properties
type PipelineCreateProperties struct {
	// The friendly name of your pipeline.
	Name *string `json:"name"`
	// The information of the log pipelines
	Logs *[]PipelineCreatePropertiesLogs `json:"logs"`
}

// NewPipelineCreateProperties instantiates a new PipelineCreateProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineCreateProperties(name string, logs []PipelineCreatePropertiesLogs) *PipelineCreateProperties {
	this := PipelineCreateProperties{}

	this.Name = &name
	this.Logs = &logs

	return &this
}

// NewPipelineCreatePropertiesWithDefaults instantiates a new PipelineCreateProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineCreatePropertiesWithDefaults() *PipelineCreateProperties {
	this := PipelineCreateProperties{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PipelineCreateProperties) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PipelineCreateProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Name, true
}

// SetName sets field value
func (o *PipelineCreateProperties) SetName(v string) {

	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *PipelineCreateProperties) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// GetLogs returns the Logs field value
// If the value is explicit nil, the zero value for []PipelineCreatePropertiesLogs will be returned
func (o *PipelineCreateProperties) GetLogs() *[]PipelineCreatePropertiesLogs {
	if o == nil {
		return nil
	}

	return o.Logs

}

// GetLogsOk returns a tuple with the Logs field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PipelineCreateProperties) GetLogsOk() (*[]PipelineCreatePropertiesLogs, bool) {
	if o == nil {
		return nil, false
	}

	return o.Logs, true
}

// SetLogs sets field value
func (o *PipelineCreateProperties) SetLogs(v []PipelineCreatePropertiesLogs) {

	o.Logs = &v

}

// HasLogs returns a boolean if a field has been set.
func (o *PipelineCreateProperties) HasLogs() bool {
	if o != nil && o.Logs != nil {
		return true
	}

	return false
}

func (o PipelineCreateProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	if o.Logs != nil {
		toSerialize["logs"] = o.Logs
	}

	return json.Marshal(toSerialize)
}

type NullablePipelineCreateProperties struct {
	value *PipelineCreateProperties
	isSet bool
}

func (v NullablePipelineCreateProperties) Get() *PipelineCreateProperties {
	return v.value
}

func (v *NullablePipelineCreateProperties) Set(val *PipelineCreateProperties) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineCreateProperties) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineCreateProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineCreateProperties(val *PipelineCreateProperties) *NullablePipelineCreateProperties {
	return &NullablePipelineCreateProperties{value: val, isSet: true}
}

func (v NullablePipelineCreateProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineCreateProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}