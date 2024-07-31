/*
 * IONOS S3 Object Storage API for contract-owned buckets
 *
 * ## Overview The IONOS S3 Object Storage API for contract-owned buckets is a REST-based API that allows developers and applications to interact directly with IONOS' scalable storage solution, leveraging the S3 protocol for object storage operations. Its design ensures seamless compatibility with existing tools and libraries tailored for S3 systems.  ### API References - [S3 Management API Reference](https://api.ionos.com/docs/s3-management/v1/) for managing Access Keys - S3 API Reference for contract-owned buckets - current document - [S3 API Reference for user-owned buckets](https://api.ionos.com/docs/s3-user-owned-buckets/v2/)  ### User documentation [IONOS S3 Object Storage User Guide](https://docs.ionos.com/cloud/managed-services/s3-object-storage) * [Documentation on user-owned and contract-owned buckets](https://docs.ionos.com/cloud/managed-services/s3-object-storage/concepts/buckets) * [Documentation on S3 API Compatibility](https://docs.ionos.com/cloud/managed-services/s3-object-storage/concepts/s3-api-compatibility) * [S3 Tools](https://docs.ionos.com/cloud/managed-services/s3-object-storage/s3-tools)  ## Endpoints for contract-owned buckets | Location | Region Name | Bucket Type | Endpoint | | --- | --- | --- | --- | | **Berlin, Germany** | **eu-central-3** | Contract-owned | `https://s3.eu-central-3.ionoscloud.com` |  ## Changelog - 30.05.2024 Initial version
 *
 * API version: 2.0.2
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
	"fmt"
)

// EncodingType Encoding type used by IONOS S3 Object Storage to encode object key names in the XML response. If you specify encoding-type request parameter, IONOS S3 Object Storage includes this element in the response, and returns encoded key name values in the following response elements: `KeyMarker`, `NextKeyMarker`, `Prefix`, `Key`, and `Delimiter`.
type EncodingType string

// List of EncodingType
const (
	ENCODINGTYPE_URL EncodingType = "url"
)

func (v *EncodingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EncodingType(value)
	for _, existing := range []EncodingType{"url"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EncodingType", value)
}

// Ptr returns reference to EncodingType value
func (v EncodingType) Ptr() *EncodingType {
	return &v
}

type NullableEncodingType struct {
	value *EncodingType
	isSet bool
}

func (v NullableEncodingType) Get() *EncodingType {
	return v.value
}

func (v *NullableEncodingType) Set(val *EncodingType) {
	v.value = val
	v.isSet = true
}

func (v NullableEncodingType) IsSet() bool {
	return v.isSet
}

func (v *NullableEncodingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncodingType(val *EncodingType) *NullableEncodingType {
	return &NullableEncodingType{value: val, isSet: true}
}

func (v NullableEncodingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncodingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
