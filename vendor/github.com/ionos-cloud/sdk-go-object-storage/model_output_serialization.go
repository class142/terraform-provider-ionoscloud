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

// OutputSerialization Describes how results of the Select job are serialized.
type OutputSerialization struct {
	XMLName xml.Name    `xml:"OutputSerialization"`
	CSV     *CSVOutput  `json:"CSV,omitempty" xml:"CSV"`
	JSON    *JSONOutput `json:"JSON,omitempty" xml:"JSON"`
}

// NewOutputSerialization instantiates a new OutputSerialization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputSerialization() *OutputSerialization {
	this := OutputSerialization{}

	return &this
}

// NewOutputSerializationWithDefaults instantiates a new OutputSerialization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputSerializationWithDefaults() *OutputSerialization {
	this := OutputSerialization{}
	return &this
}

// GetCSV returns the CSV field value
// If the value is explicit nil, the zero value for CSVOutput will be returned
func (o *OutputSerialization) GetCSV() *CSVOutput {
	if o == nil {
		return nil
	}

	return o.CSV

}

// GetCSVOk returns a tuple with the CSV field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OutputSerialization) GetCSVOk() (*CSVOutput, bool) {
	if o == nil {
		return nil, false
	}

	return o.CSV, true
}

// SetCSV sets field value
func (o *OutputSerialization) SetCSV(v CSVOutput) {

	o.CSV = &v

}

// HasCSV returns a boolean if a field has been set.
func (o *OutputSerialization) HasCSV() bool {
	if o != nil && o.CSV != nil {
		return true
	}

	return false
}

// GetJSON returns the JSON field value
// If the value is explicit nil, the zero value for JSONOutput will be returned
func (o *OutputSerialization) GetJSON() *JSONOutput {
	if o == nil {
		return nil
	}

	return o.JSON

}

// GetJSONOk returns a tuple with the JSON field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OutputSerialization) GetJSONOk() (*JSONOutput, bool) {
	if o == nil {
		return nil, false
	}

	return o.JSON, true
}

// SetJSON sets field value
func (o *OutputSerialization) SetJSON(v JSONOutput) {

	o.JSON = &v

}

// HasJSON returns a boolean if a field has been set.
func (o *OutputSerialization) HasJSON() bool {
	if o != nil && o.JSON != nil {
		return true
	}

	return false
}

func (o OutputSerialization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CSV != nil {
		toSerialize["CSV"] = o.CSV
	}

	if o.JSON != nil {
		toSerialize["JSON"] = o.JSON
	}

	return json.Marshal(toSerialize)
}

type NullableOutputSerialization struct {
	value *OutputSerialization
	isSet bool
}

func (v NullableOutputSerialization) Get() *OutputSerialization {
	return v.value
}

func (v *NullableOutputSerialization) Set(val *OutputSerialization) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputSerialization) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputSerialization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputSerialization(val *OutputSerialization) *NullableOutputSerialization {
	return &NullableOutputSerialization{value: val, isSet: true}
}

func (v NullableOutputSerialization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputSerialization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
