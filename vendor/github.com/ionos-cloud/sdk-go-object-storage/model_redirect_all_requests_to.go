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

// RedirectAllRequestsTo Specifies the redirect behavior of all requests to a website endpoint of an IONOS Object Storage bucket.
type RedirectAllRequestsTo struct {
	XMLName xml.Name `xml:"RedirectAllRequestsTo"`
	// Name of the host where requests are redirected.
	HostName *string `json:"HostName" xml:"HostName"`
	// Protocol to use when redirecting requests. The default is the protocol that is used in the original request.
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol"`
}

// NewRedirectAllRequestsTo instantiates a new RedirectAllRequestsTo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedirectAllRequestsTo(hostName string) *RedirectAllRequestsTo {
	this := RedirectAllRequestsTo{}

	this.HostName = &hostName

	return &this
}

// NewRedirectAllRequestsToWithDefaults instantiates a new RedirectAllRequestsTo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedirectAllRequestsToWithDefaults() *RedirectAllRequestsTo {
	this := RedirectAllRequestsTo{}
	return &this
}

// GetHostName returns the HostName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RedirectAllRequestsTo) GetHostName() *string {
	if o == nil {
		return nil
	}

	return o.HostName

}

// GetHostNameOk returns a tuple with the HostName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RedirectAllRequestsTo) GetHostNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.HostName, true
}

// SetHostName sets field value
func (o *RedirectAllRequestsTo) SetHostName(v string) {

	o.HostName = &v

}

// HasHostName returns a boolean if a field has been set.
func (o *RedirectAllRequestsTo) HasHostName() bool {
	if o != nil && o.HostName != nil {
		return true
	}

	return false
}

// GetProtocol returns the Protocol field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RedirectAllRequestsTo) GetProtocol() *string {
	if o == nil {
		return nil
	}

	return o.Protocol

}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RedirectAllRequestsTo) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Protocol, true
}

// SetProtocol sets field value
func (o *RedirectAllRequestsTo) SetProtocol(v string) {

	o.Protocol = &v

}

// HasProtocol returns a boolean if a field has been set.
func (o *RedirectAllRequestsTo) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

func (o RedirectAllRequestsTo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HostName != nil {
		toSerialize["HostName"] = o.HostName
	}

	if o.Protocol != nil {
		toSerialize["Protocol"] = o.Protocol
	}

	return json.Marshal(toSerialize)
}

type NullableRedirectAllRequestsTo struct {
	value *RedirectAllRequestsTo
	isSet bool
}

func (v NullableRedirectAllRequestsTo) Get() *RedirectAllRequestsTo {
	return v.value
}

func (v *NullableRedirectAllRequestsTo) Set(val *RedirectAllRequestsTo) {
	v.value = val
	v.isSet = true
}

func (v NullableRedirectAllRequestsTo) IsSet() bool {
	return v.isSet
}

func (v *NullableRedirectAllRequestsTo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedirectAllRequestsTo(val *RedirectAllRequestsTo) *NullableRedirectAllRequestsTo {
	return &NullableRedirectAllRequestsTo{value: val, isSet: true}
}

func (v NullableRedirectAllRequestsTo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedirectAllRequestsTo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
