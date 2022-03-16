/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// ServerProperties struct for ServerProperties
type ServerProperties struct {
	// The ID of the template for creating a CUBE server; the available templates for CUBE servers can be found on the templates resource.
	TemplateUuid *string `json:"templateUuid,omitempty"`
	// The name of the  resource.
	Name *string `json:"name,omitempty"`
	// The total number of cores for the server.
	Cores *int32 `json:"cores"`
	// The memory size for the server in MB, such as 2048. Size must be specified in multiples of 256 MB with a minimum of 256 MB; however, if you set ramHotPlug to TRUE then you must use a minimum of 1024 MB. If you set the RAM size more than 240GB, then ramHotPlug will be set to FALSE and can not be set to TRUE unless RAM size not set to less than 240GB.
	Ram *int32 `json:"ram"`
	// The availability zone in which the server should be provisioned.
	AvailabilityZone *string `json:"availabilityZone,omitempty"`
	// Status of the virtual machine.
	VmState    *string            `json:"vmState,omitempty"`
	BootCdrom  *ResourceReference `json:"bootCdrom,omitempty"`
	BootVolume *ResourceReference `json:"bootVolume,omitempty"`
	// CPU architecture on which server gets provisioned; not all CPU architectures are available in all datacenter regions; available CPU architectures can be retrieved from the datacenter resource.
	CpuFamily *string `json:"cpuFamily,omitempty"`
	// server usages: ENTERPRISE or CUBE
	Type *string `json:"type,omitempty"`
}

// NewServerProperties instantiates a new ServerProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerProperties(cores int32, ram int32) *ServerProperties {
	this := ServerProperties{}

	this.Cores = &cores
	this.Ram = &ram

	return &this
}

// NewServerPropertiesWithDefaults instantiates a new ServerProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerPropertiesWithDefaults() *ServerProperties {
	this := ServerProperties{}
	return &this
}

// GetTemplateUuid returns the TemplateUuid field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServerProperties) GetTemplateUuid() *string {
	if o == nil {
		return nil
	}

	return o.TemplateUuid

}

// GetTemplateUuidOk returns a tuple with the TemplateUuid field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerProperties) GetTemplateUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.TemplateUuid, true
}

// SetTemplateUuid sets field value
func (o *ServerProperties) SetTemplateUuid(v string) {

	o.TemplateUuid = &v

}

// HasTemplateUuid returns a boolean if a field has been set.
func (o *ServerProperties) HasTemplateUuid() bool {
	if o != nil && o.TemplateUuid != nil {
		return true
	}

	return false
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServerProperties) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Name, true
}

// SetName sets field value
func (o *ServerProperties) SetName(v string) {

	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *ServerProperties) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// GetCores returns the Cores field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ServerProperties) GetCores() *int32 {
	if o == nil {
		return nil
	}

	return o.Cores

}

// GetCoresOk returns a tuple with the Cores field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerProperties) GetCoresOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Cores, true
}

// SetCores sets field value
func (o *ServerProperties) SetCores(v int32) {

	o.Cores = &v

}

// HasCores returns a boolean if a field has been set.
func (o *ServerProperties) HasCores() bool {
	if o != nil && o.Cores != nil {
		return true
	}

	return false
}

// GetRam returns the Ram field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ServerProperties) GetRam() *int32 {
	if o == nil {
		return nil
	}

	return o.Ram

}

// GetRamOk returns a tuple with the Ram field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerProperties) GetRamOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Ram, true
}

// SetRam sets field value
func (o *ServerProperties) SetRam(v int32) {

	o.Ram = &v

}

// HasRam returns a boolean if a field has been set.
func (o *ServerProperties) HasRam() bool {
	if o != nil && o.Ram != nil {
		return true
	}

	return false
}

// GetAvailabilityZone returns the AvailabilityZone field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServerProperties) GetAvailabilityZone() *string {
	if o == nil {
		return nil
	}

	return o.AvailabilityZone

}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerProperties) GetAvailabilityZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.AvailabilityZone, true
}

// SetAvailabilityZone sets field value
func (o *ServerProperties) SetAvailabilityZone(v string) {

	o.AvailabilityZone = &v

}

// HasAvailabilityZone returns a boolean if a field has been set.
func (o *ServerProperties) HasAvailabilityZone() bool {
	if o != nil && o.AvailabilityZone != nil {
		return true
	}

	return false
}

// GetVmState returns the VmState field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServerProperties) GetVmState() *string {
	if o == nil {
		return nil
	}

	return o.VmState

}

// GetVmStateOk returns a tuple with the VmState field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerProperties) GetVmStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.VmState, true
}

// SetVmState sets field value
func (o *ServerProperties) SetVmState(v string) {

	o.VmState = &v

}

// HasVmState returns a boolean if a field has been set.
func (o *ServerProperties) HasVmState() bool {
	if o != nil && o.VmState != nil {
		return true
	}

	return false
}

// GetBootCdrom returns the BootCdrom field value
// If the value is explicit nil, the zero value for ResourceReference will be returned
func (o *ServerProperties) GetBootCdrom() *ResourceReference {
	if o == nil {
		return nil
	}

	return o.BootCdrom

}

// GetBootCdromOk returns a tuple with the BootCdrom field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerProperties) GetBootCdromOk() (*ResourceReference, bool) {
	if o == nil {
		return nil, false
	}

	return o.BootCdrom, true
}

// SetBootCdrom sets field value
func (o *ServerProperties) SetBootCdrom(v ResourceReference) {

	o.BootCdrom = &v

}

// HasBootCdrom returns a boolean if a field has been set.
func (o *ServerProperties) HasBootCdrom() bool {
	if o != nil && o.BootCdrom != nil {
		return true
	}

	return false
}

// GetBootVolume returns the BootVolume field value
// If the value is explicit nil, the zero value for ResourceReference will be returned
func (o *ServerProperties) GetBootVolume() *ResourceReference {
	if o == nil {
		return nil
	}

	return o.BootVolume

}

// GetBootVolumeOk returns a tuple with the BootVolume field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerProperties) GetBootVolumeOk() (*ResourceReference, bool) {
	if o == nil {
		return nil, false
	}

	return o.BootVolume, true
}

// SetBootVolume sets field value
func (o *ServerProperties) SetBootVolume(v ResourceReference) {

	o.BootVolume = &v

}

// HasBootVolume returns a boolean if a field has been set.
func (o *ServerProperties) HasBootVolume() bool {
	if o != nil && o.BootVolume != nil {
		return true
	}

	return false
}

// GetCpuFamily returns the CpuFamily field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServerProperties) GetCpuFamily() *string {
	if o == nil {
		return nil
	}

	return o.CpuFamily

}

// GetCpuFamilyOk returns a tuple with the CpuFamily field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerProperties) GetCpuFamilyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.CpuFamily, true
}

// SetCpuFamily sets field value
func (o *ServerProperties) SetCpuFamily(v string) {

	o.CpuFamily = &v

}

// HasCpuFamily returns a boolean if a field has been set.
func (o *ServerProperties) HasCpuFamily() bool {
	if o != nil && o.CpuFamily != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServerProperties) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerProperties) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *ServerProperties) SetType(v string) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *ServerProperties) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

func (o ServerProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TemplateUuid != nil {
		toSerialize["templateUuid"] = o.TemplateUuid
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Cores != nil {
		toSerialize["cores"] = o.Cores
	}
	if o.Ram != nil {
		toSerialize["ram"] = o.Ram
	}
	if o.AvailabilityZone != nil {
		toSerialize["availabilityZone"] = o.AvailabilityZone
	}
	if o.VmState != nil {
		toSerialize["vmState"] = o.VmState
	}
	if o.BootCdrom != nil {
		toSerialize["bootCdrom"] = o.BootCdrom
	}
	if o.BootVolume != nil {
		toSerialize["bootVolume"] = o.BootVolume
	}
	if o.CpuFamily != nil {
		toSerialize["cpuFamily"] = o.CpuFamily
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableServerProperties struct {
	value *ServerProperties
	isSet bool
}

func (v NullableServerProperties) Get() *ServerProperties {
	return v.value
}

func (v *NullableServerProperties) Set(val *ServerProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableServerProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableServerProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerProperties(val *ServerProperties) *NullableServerProperties {
	return &NullableServerProperties{value: val, isSet: true}
}

func (v NullableServerProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
