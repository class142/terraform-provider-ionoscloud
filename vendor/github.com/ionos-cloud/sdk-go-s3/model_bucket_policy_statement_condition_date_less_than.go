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

import (
	"gopkg.in/validator.v2"
)

// BucketPolicyStatementConditionDateLessThan - struct for BucketPolicyStatementConditionDateLessThan
type BucketPolicyStatementConditionDateLessThan struct {
	BucketPolicyStatementConditionDateGreaterThanOneOf *BucketPolicyStatementConditionDateGreaterThanOneOf
	BucketPolicyStatementConditionDateLessThanOneOf    *BucketPolicyStatementConditionDateLessThanOneOf
}

// BucketPolicyStatementConditionDateGreaterThanOneOfAsBucketPolicyStatementConditionDateLessThan is a convenience function that returns BucketPolicyStatementConditionDateGreaterThanOneOf wrapped in BucketPolicyStatementConditionDateLessThan
func BucketPolicyStatementConditionDateGreaterThanOneOfAsBucketPolicyStatementConditionDateLessThan(v *BucketPolicyStatementConditionDateGreaterThanOneOf) BucketPolicyStatementConditionDateLessThan {
	return BucketPolicyStatementConditionDateLessThan{
		BucketPolicyStatementConditionDateGreaterThanOneOf: v,
	}
}

// BucketPolicyStatementConditionDateLessThanOneOfAsBucketPolicyStatementConditionDateLessThan is a convenience function that returns BucketPolicyStatementConditionDateLessThanOneOf wrapped in BucketPolicyStatementConditionDateLessThan
func BucketPolicyStatementConditionDateLessThanOneOfAsBucketPolicyStatementConditionDateLessThan(v *BucketPolicyStatementConditionDateLessThanOneOf) BucketPolicyStatementConditionDateLessThan {
	return BucketPolicyStatementConditionDateLessThan{
		BucketPolicyStatementConditionDateLessThanOneOf: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BucketPolicyStatementConditionDateLessThan) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BucketPolicyStatementConditionDateGreaterThanOneOf
	err = newStrictDecoder(data).Decode(&dst.BucketPolicyStatementConditionDateGreaterThanOneOf)
	if err == nil {
		jsonBucketPolicyStatementConditionDateGreaterThanOneOf, _ := json.Marshal(dst.BucketPolicyStatementConditionDateGreaterThanOneOf)
		if string(jsonBucketPolicyStatementConditionDateGreaterThanOneOf) == "{}" { // empty struct
			dst.BucketPolicyStatementConditionDateGreaterThanOneOf = nil
		} else {
			if err = validator.Validate(dst.BucketPolicyStatementConditionDateGreaterThanOneOf); err != nil {
				dst.BucketPolicyStatementConditionDateGreaterThanOneOf = nil
			} else {
				match++
			}
		}
	} else {
		dst.BucketPolicyStatementConditionDateGreaterThanOneOf = nil
	}

	// try to unmarshal data into BucketPolicyStatementConditionDateLessThanOneOf
	err = newStrictDecoder(data).Decode(&dst.BucketPolicyStatementConditionDateLessThanOneOf)
	if err == nil {
		jsonBucketPolicyStatementConditionDateLessThanOneOf, _ := json.Marshal(dst.BucketPolicyStatementConditionDateLessThanOneOf)
		if string(jsonBucketPolicyStatementConditionDateLessThanOneOf) == "{}" { // empty struct
			dst.BucketPolicyStatementConditionDateLessThanOneOf = nil
		} else {
			if err = validator.Validate(dst.BucketPolicyStatementConditionDateLessThanOneOf); err != nil {
				dst.BucketPolicyStatementConditionDateLessThanOneOf = nil
			} else {
				match++
			}
		}
	} else {
		dst.BucketPolicyStatementConditionDateLessThanOneOf = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BucketPolicyStatementConditionDateGreaterThanOneOf = nil
		dst.BucketPolicyStatementConditionDateLessThanOneOf = nil

		return fmt.Errorf("data matches more than one schema in oneOf(BucketPolicyStatementConditionDateLessThan)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(BucketPolicyStatementConditionDateLessThan)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BucketPolicyStatementConditionDateLessThan) MarshalJSON() ([]byte, error) {
	if src.BucketPolicyStatementConditionDateGreaterThanOneOf != nil {
		return json.Marshal(&src.BucketPolicyStatementConditionDateGreaterThanOneOf)
	}

	if src.BucketPolicyStatementConditionDateLessThanOneOf != nil {
		return json.Marshal(&src.BucketPolicyStatementConditionDateLessThanOneOf)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *BucketPolicyStatementConditionDateLessThan) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BucketPolicyStatementConditionDateGreaterThanOneOf != nil {
		return obj.BucketPolicyStatementConditionDateGreaterThanOneOf
	}

	if obj.BucketPolicyStatementConditionDateLessThanOneOf != nil {
		return obj.BucketPolicyStatementConditionDateLessThanOneOf
	}

	// all schemas are nil
	return nil
}

type NullableBucketPolicyStatementConditionDateLessThan struct {
	value *BucketPolicyStatementConditionDateLessThan
	isSet bool
}

func (v NullableBucketPolicyStatementConditionDateLessThan) Get() *BucketPolicyStatementConditionDateLessThan {
	return v.value
}

func (v *NullableBucketPolicyStatementConditionDateLessThan) Set(val *BucketPolicyStatementConditionDateLessThan) {
	v.value = val
	v.isSet = true
}

func (v NullableBucketPolicyStatementConditionDateLessThan) IsSet() bool {
	return v.isSet
}

func (v *NullableBucketPolicyStatementConditionDateLessThan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBucketPolicyStatementConditionDateLessThan(val *BucketPolicyStatementConditionDateLessThan) *NullableBucketPolicyStatementConditionDateLessThan {
	return &NullableBucketPolicyStatementConditionDateLessThan{value: val, isSet: true}
}

func (v NullableBucketPolicyStatementConditionDateLessThan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBucketPolicyStatementConditionDateLessThan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}