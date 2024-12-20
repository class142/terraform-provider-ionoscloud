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
	"fmt"
)

// StorageClass Currently, only one storage class is available – `STANDARD`.
type StorageClass string

// List of StorageClass
const (
	STORAGECLASS_STANDARD StorageClass = "STANDARD"
)

func (v *StorageClass) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StorageClass(value)
	for _, existing := range []StorageClass{"STANDARD"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StorageClass", value)
}

// Ptr returns reference to StorageClass value
func (v StorageClass) Ptr() *StorageClass {
	return &v
}

type NullableStorageClass struct {
	value *StorageClass
	isSet bool
}

func (v NullableStorageClass) Get() *StorageClass {
	return v.value
}

func (v *NullableStorageClass) Set(val *StorageClass) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageClass) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageClass(val *StorageClass) *NullableStorageClass {
	return &NullableStorageClass{value: val, isSet: true}
}

func (v NullableStorageClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
