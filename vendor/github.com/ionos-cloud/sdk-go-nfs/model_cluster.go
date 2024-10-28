/*
 * IONOS Cloud - Network File Storage API
 *
 * The RESTful API for managing Network File Storage.
 *
 * API version: 0.1.1
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// Cluster Network File Storage cluster
type Cluster struct {
	Name        *string               `json:"name"`
	Connections *[]ClusterConnections `json:"connections"`
	Nfs         *ClusterNfs           `json:"nfs,omitempty"`
	// The size of the Network File Storage cluster in TiB. Note that the cluster size cannot be reduced after provisioning. This value determines the billing fees.
	Size *int32 `json:"size,omitempty"`
}

// NewCluster instantiates a new Cluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCluster(name string, connections []ClusterConnections) *Cluster {
	this := Cluster{}

	this.Name = &name
	this.Connections = &connections
	var size int32 = 2
	this.Size = &size

	return &this
}

// NewClusterWithDefaults instantiates a new Cluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterWithDefaults() *Cluster {
	this := Cluster{}
	var size int32 = 2
	this.Size = &size
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Cluster) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Name, true
}

// SetName sets field value
func (o *Cluster) SetName(v string) {

	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *Cluster) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// GetConnections returns the Connections field value
// If the value is explicit nil, the zero value for []ClusterConnections will be returned
func (o *Cluster) GetConnections() *[]ClusterConnections {
	if o == nil {
		return nil
	}

	return o.Connections

}

// GetConnectionsOk returns a tuple with the Connections field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetConnectionsOk() (*[]ClusterConnections, bool) {
	if o == nil {
		return nil, false
	}

	return o.Connections, true
}

// SetConnections sets field value
func (o *Cluster) SetConnections(v []ClusterConnections) {

	o.Connections = &v

}

// HasConnections returns a boolean if a field has been set.
func (o *Cluster) HasConnections() bool {
	if o != nil && o.Connections != nil {
		return true
	}

	return false
}

// GetNfs returns the Nfs field value
// If the value is explicit nil, the zero value for ClusterNfs will be returned
func (o *Cluster) GetNfs() *ClusterNfs {
	if o == nil {
		return nil
	}

	return o.Nfs

}

// GetNfsOk returns a tuple with the Nfs field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetNfsOk() (*ClusterNfs, bool) {
	if o == nil {
		return nil, false
	}

	return o.Nfs, true
}

// SetNfs sets field value
func (o *Cluster) SetNfs(v ClusterNfs) {

	o.Nfs = &v

}

// HasNfs returns a boolean if a field has been set.
func (o *Cluster) HasNfs() bool {
	if o != nil && o.Nfs != nil {
		return true
	}

	return false
}

// GetSize returns the Size field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Cluster) GetSize() *int32 {
	if o == nil {
		return nil
	}

	return o.Size

}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Size, true
}

// SetSize sets field value
func (o *Cluster) SetSize(v int32) {

	o.Size = &v

}

// HasSize returns a boolean if a field has been set.
func (o *Cluster) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

func (o Cluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	if o.Connections != nil {
		toSerialize["connections"] = o.Connections
	}

	if o.Nfs != nil {
		toSerialize["nfs"] = o.Nfs
	}

	if o.Size != nil {
		toSerialize["size"] = o.Size
	}

	return json.Marshal(toSerialize)
}

type NullableCluster struct {
	value *Cluster
	isSet bool
}

func (v NullableCluster) Get() *Cluster {
	return v.value
}

func (v *NullableCluster) Set(val *Cluster) {
	v.value = val
	v.isSet = true
}

func (v NullableCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCluster(val *Cluster) *NullableCluster {
	return &NullableCluster{value: val, isSet: true}
}

func (v NullableCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
