/*
 * VM Auto Scaling API
 *
 * The VM Auto Scaling Service enables IONOS clients to horizontally scale the number of VM replicas based on configured rules. You can use VM Auto Scaling to ensure that you have a sufficient number of replicas to handle your application loads at all times.  For this purpose, create a VM Auto Scaling Group that contains the server replicas. The VM Auto Scaling Service ensures that the number of replicas in the group is always within the defined limits.   When scaling policies are set, VM Auto Scaling creates or deletes replicas according to the requirements of your applications. For each policy, specified 'scale-in' and 'scale-out' actions are performed when the corresponding thresholds are reached.
 *
 * API version: 1-SDK.1
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// Error401Message struct for Error401Message
type Error401Message struct {
	ErrorCode *string `json:"errorCode,omitempty"`
	Message   *string `json:"message,omitempty"`
}

// NewError401Message instantiates a new Error401Message object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewError401Message() *Error401Message {
	this := Error401Message{}

	return &this
}

// NewError401MessageWithDefaults instantiates a new Error401Message object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewError401MessageWithDefaults() *Error401Message {
	this := Error401Message{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Error401Message) GetErrorCode() *string {
	if o == nil {
		return nil
	}

	return o.ErrorCode

}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Error401Message) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *Error401Message) SetErrorCode(v string) {

	o.ErrorCode = &v

}

// HasErrorCode returns a boolean if a field has been set.
func (o *Error401Message) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// GetMessage returns the Message field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Error401Message) GetMessage() *string {
	if o == nil {
		return nil
	}

	return o.Message

}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Error401Message) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Message, true
}

// SetMessage sets field value
func (o *Error401Message) SetMessage(v string) {

	o.Message = &v

}

// HasMessage returns a boolean if a field has been set.
func (o *Error401Message) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

func (o Error401Message) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}

	if o.Message != nil {
		toSerialize["message"] = o.Message
	}

	return json.Marshal(toSerialize)
}

type NullableError401Message struct {
	value *Error401Message
	isSet bool
}

func (v NullableError401Message) Get() *Error401Message {
	return v.value
}

func (v *NullableError401Message) Set(val *Error401Message) {
	v.value = val
	v.isSet = true
}

func (v NullableError401Message) IsSet() bool {
	return v.isSet
}

func (v *NullableError401Message) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableError401Message(val *Error401Message) *NullableError401Message {
	return &NullableError401Message{value: val, isSet: true}
}

func (v NullableError401Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableError401Message) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}