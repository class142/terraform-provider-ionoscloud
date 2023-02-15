/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package compute

import (
	"encoding/json"

	"fmt"
)

// NatGatewayRuleType the model 'NatGatewayRuleType'
type NatGatewayRuleType string

// List of NatGatewayRuleType
const (
	SNAT NatGatewayRuleType = "SNAT"
)

func (v *NatGatewayRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NatGatewayRuleType(value)
	for _, existing := range []NatGatewayRuleType{"SNAT"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NatGatewayRuleType", value)
}

// Ptr returns reference to NatGatewayRuleType value
func (v NatGatewayRuleType) Ptr() *NatGatewayRuleType {
	return &v
}

type NullableNatGatewayRuleType struct {
	value *NatGatewayRuleType
	isSet bool
}

func (v NullableNatGatewayRuleType) Get() *NatGatewayRuleType {
	return v.value
}

func (v *NullableNatGatewayRuleType) Set(val *NatGatewayRuleType) {
	v.value = val
	v.isSet = true
}

func (v NullableNatGatewayRuleType) IsSet() bool {
	return v.isSet
}

func (v *NullableNatGatewayRuleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNatGatewayRuleType(val *NatGatewayRuleType) *NullableNatGatewayRuleType {
	return &NullableNatGatewayRuleType{value: val, isSet: true}
}

func (v NullableNatGatewayRuleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNatGatewayRuleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
