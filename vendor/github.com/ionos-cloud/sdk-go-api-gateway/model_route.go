/*
 * IONOS Cloud - API Gateway
 *
 * API Gateway is an application that acts as a \"front door\" for backend services and APIs, handling client requests and routing them to the appropriate backend.
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// Route A route is a rule that maps an incoming request to a specific backend service.
type Route struct {
	// The name of the route.
	Name *string `json:"name"`
	// This field specifies the protocol used by the ingress to route traffic to the backend service.
	Type *string `json:"type"`
	// The paths that the route should match.
	Paths *[]string `json:"paths"`
	// The HTTP methods that the route should match.
	Methods *[]string `json:"methods"`
	// To enable websocket support.
	Websocket *bool             `json:"websocket,omitempty"`
	Upstreams *[]RouteUpstreams `json:"upstreams"`
}

// NewRoute instantiates a new Route object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoute(name string, type_ string, paths []string, methods []string, upstreams []RouteUpstreams) *Route {
	this := Route{}

	this.Name = &name
	this.Type = &type_
	this.Paths = &paths
	this.Methods = &methods
	var websocket bool = false
	this.Websocket = &websocket
	this.Upstreams = &upstreams

	return &this
}

// NewRouteWithDefaults instantiates a new Route object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteWithDefaults() *Route {
	this := Route{}
	var type_ string = "http"
	this.Type = &type_
	var websocket bool = false
	this.Websocket = &websocket
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Route) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Route) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Name, true
}

// SetName sets field value
func (o *Route) SetName(v string) {

	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *Route) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Route) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Route) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *Route) SetType(v string) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *Route) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetPaths returns the Paths field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *Route) GetPaths() *[]string {
	if o == nil {
		return nil
	}

	return o.Paths

}

// GetPathsOk returns a tuple with the Paths field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Route) GetPathsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Paths, true
}

// SetPaths sets field value
func (o *Route) SetPaths(v []string) {

	o.Paths = &v

}

// HasPaths returns a boolean if a field has been set.
func (o *Route) HasPaths() bool {
	if o != nil && o.Paths != nil {
		return true
	}

	return false
}

// GetMethods returns the Methods field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *Route) GetMethods() *[]string {
	if o == nil {
		return nil
	}

	return o.Methods

}

// GetMethodsOk returns a tuple with the Methods field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Route) GetMethodsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Methods, true
}

// SetMethods sets field value
func (o *Route) SetMethods(v []string) {

	o.Methods = &v

}

// HasMethods returns a boolean if a field has been set.
func (o *Route) HasMethods() bool {
	if o != nil && o.Methods != nil {
		return true
	}

	return false
}

// GetWebsocket returns the Websocket field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *Route) GetWebsocket() *bool {
	if o == nil {
		return nil
	}

	return o.Websocket

}

// GetWebsocketOk returns a tuple with the Websocket field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Route) GetWebsocketOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.Websocket, true
}

// SetWebsocket sets field value
func (o *Route) SetWebsocket(v bool) {

	o.Websocket = &v

}

// HasWebsocket returns a boolean if a field has been set.
func (o *Route) HasWebsocket() bool {
	if o != nil && o.Websocket != nil {
		return true
	}

	return false
}

// GetUpstreams returns the Upstreams field value
// If the value is explicit nil, the zero value for []RouteUpstreams will be returned
func (o *Route) GetUpstreams() *[]RouteUpstreams {
	if o == nil {
		return nil
	}

	return o.Upstreams

}

// GetUpstreamsOk returns a tuple with the Upstreams field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Route) GetUpstreamsOk() (*[]RouteUpstreams, bool) {
	if o == nil {
		return nil, false
	}

	return o.Upstreams, true
}

// SetUpstreams sets field value
func (o *Route) SetUpstreams(v []RouteUpstreams) {

	o.Upstreams = &v

}

// HasUpstreams returns a boolean if a field has been set.
func (o *Route) HasUpstreams() bool {
	if o != nil && o.Upstreams != nil {
		return true
	}

	return false
}

func (o Route) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	if o.Paths != nil {
		toSerialize["paths"] = o.Paths
	}

	if o.Methods != nil {
		toSerialize["methods"] = o.Methods
	}

	if o.Websocket != nil {
		toSerialize["websocket"] = o.Websocket
	}

	if o.Upstreams != nil {
		toSerialize["upstreams"] = o.Upstreams
	}

	return json.Marshal(toSerialize)
}

type NullableRoute struct {
	value *Route
	isSet bool
}

func (v NullableRoute) Get() *Route {
	return v.value
}

func (v *NullableRoute) Set(val *Route) {
	v.value = val
	v.isSet = true
}

func (v NullableRoute) IsSet() bool {
	return v.isSet
}

func (v *NullableRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoute(val *Route) *NullableRoute {
	return &NullableRoute{value: val, isSet: true}
}

func (v NullableRoute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}