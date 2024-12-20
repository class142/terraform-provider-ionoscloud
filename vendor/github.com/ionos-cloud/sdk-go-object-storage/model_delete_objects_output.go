/*
 * IONOS Object Storage API for contract-owned buckets
 *
 * ## Overview The IONOS Object Storage API for contract-owned buckets is a REST-based API that allows developers and applications to interact directly with IONOS' scalable storage solution, leveraging the S3 protocol for object storage operations. Its design ensures seamless compatibility with existing tools and libraries tailored for S3 systems.  ### API References - [S3 API Reference for contract-owned buckets](https://api.ionos.com/docs/s3-contract-owned-buckets/v2/) ### User documentation [IONOS Object Storage User Guide](https://docs.ionos.com/cloud/managed-services/s3-object-storage) * [Documentation on user-owned and contract-owned buckets](https://docs.ionos.com/cloud/managed-services/s3-object-storage/concepts/buckets) * [Documentation on S3 API Compatibility](https://docs.ionos.com/cloud/managed-services/s3-object-storage/concepts/s3-api-compatibility) * [S3 Tools](https://docs.ionos.com/cloud/managed-services/s3-object-storage/s3-tools)  ## Endpoints for contract-owned buckets | Location | Region Name | Bucket Type | Endpoint | | --- | --- | --- | --- | | **Berlin, Germany** | **eu-central-3** | Contract-owned | `https://s3.eu-central-3.ionoscloud.com` |  ## Changelog - 30.05.2024 Initial version
 *
 * API version: 2.0.2
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

import "encoding/xml"

// DeleteObjectsOutput struct for DeleteObjectsOutput
type DeleteObjectsOutput struct {
	XMLName xml.Name `xml:"DeleteResult"`
	// Container element for a successful delete. It identifies the object that was successfully deleted.
	Deleted *[]DeletedObject `json:"Deleted,omitempty" xml:"Deleted"`
	Errors  *[]DeletionError `json:"Errors,omitempty" xml:"Error"`
}

// NewDeleteObjectsOutput instantiates a new DeleteObjectsOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteObjectsOutput() *DeleteObjectsOutput {
	this := DeleteObjectsOutput{}

	return &this
}

// NewDeleteObjectsOutputWithDefaults instantiates a new DeleteObjectsOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteObjectsOutputWithDefaults() *DeleteObjectsOutput {
	this := DeleteObjectsOutput{}
	return &this
}

// GetDeleted returns the Deleted field value
// If the value is explicit nil, the zero value for []DeletedObject will be returned
func (o *DeleteObjectsOutput) GetDeleted() *[]DeletedObject {
	if o == nil {
		return nil
	}

	return o.Deleted

}

// GetDeletedOk returns a tuple with the Deleted field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteObjectsOutput) GetDeletedOk() (*[]DeletedObject, bool) {
	if o == nil {
		return nil, false
	}

	return o.Deleted, true
}

// SetDeleted sets field value
func (o *DeleteObjectsOutput) SetDeleted(v []DeletedObject) {

	o.Deleted = &v

}

// HasDeleted returns a boolean if a field has been set.
func (o *DeleteObjectsOutput) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// GetErrors returns the Errors field value
// If the value is explicit nil, the zero value for []DeletionError will be returned
func (o *DeleteObjectsOutput) GetErrors() *[]DeletionError {
	if o == nil {
		return nil
	}

	return o.Errors

}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteObjectsOutput) GetErrorsOk() (*[]DeletionError, bool) {
	if o == nil {
		return nil, false
	}

	return o.Errors, true
}

// SetErrors sets field value
func (o *DeleteObjectsOutput) SetErrors(v []DeletionError) {

	o.Errors = &v

}

// HasErrors returns a boolean if a field has been set.
func (o *DeleteObjectsOutput) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

func (o DeleteObjectsOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Deleted != nil {
		toSerialize["Deleted"] = o.Deleted
	}

	if o.Errors != nil {
		toSerialize["Errors"] = o.Errors
	}

	return json.Marshal(toSerialize)
}

type NullableDeleteObjectsOutput struct {
	value *DeleteObjectsOutput
	isSet bool
}

func (v NullableDeleteObjectsOutput) Get() *DeleteObjectsOutput {
	return v.value
}

func (v *NullableDeleteObjectsOutput) Set(val *DeleteObjectsOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteObjectsOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteObjectsOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteObjectsOutput(val *DeleteObjectsOutput) *NullableDeleteObjectsOutput {
	return &NullableDeleteObjectsOutput{value: val, isSet: true}
}

func (v NullableDeleteObjectsOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteObjectsOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
