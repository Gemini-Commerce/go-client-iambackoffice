/*
IamBackoffice Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iambackoffice

import (
	"encoding/json"
)

// checks if the IambackofficeCreateGroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IambackofficeCreateGroupResponse{}

// IambackofficeCreateGroupResponse struct for IambackofficeCreateGroupResponse
type IambackofficeCreateGroupResponse struct {
	Group                *IambackofficeGroup `json:"group,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IambackofficeCreateGroupResponse IambackofficeCreateGroupResponse

// NewIambackofficeCreateGroupResponse instantiates a new IambackofficeCreateGroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIambackofficeCreateGroupResponse() *IambackofficeCreateGroupResponse {
	this := IambackofficeCreateGroupResponse{}
	return &this
}

// NewIambackofficeCreateGroupResponseWithDefaults instantiates a new IambackofficeCreateGroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIambackofficeCreateGroupResponseWithDefaults() *IambackofficeCreateGroupResponse {
	this := IambackofficeCreateGroupResponse{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *IambackofficeCreateGroupResponse) GetGroup() IambackofficeGroup {
	if o == nil || IsNil(o.Group) {
		var ret IambackofficeGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeCreateGroupResponse) GetGroupOk() (*IambackofficeGroup, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *IambackofficeCreateGroupResponse) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given IambackofficeGroup and assigns it to the Group field.
func (o *IambackofficeCreateGroupResponse) SetGroup(v IambackofficeGroup) {
	o.Group = &v
}

func (o IambackofficeCreateGroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IambackofficeCreateGroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IambackofficeCreateGroupResponse) UnmarshalJSON(data []byte) (err error) {
	varIambackofficeCreateGroupResponse := _IambackofficeCreateGroupResponse{}

	err = json.Unmarshal(data, &varIambackofficeCreateGroupResponse)

	if err != nil {
		return err
	}

	*o = IambackofficeCreateGroupResponse(varIambackofficeCreateGroupResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "group")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *IambackofficeCreateGroupResponse) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *IambackofficeCreateGroupResponse) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableIambackofficeCreateGroupResponse struct {
	value *IambackofficeCreateGroupResponse
	isSet bool
}

func (v NullableIambackofficeCreateGroupResponse) Get() *IambackofficeCreateGroupResponse {
	return v.value
}

func (v *NullableIambackofficeCreateGroupResponse) Set(val *IambackofficeCreateGroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIambackofficeCreateGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIambackofficeCreateGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIambackofficeCreateGroupResponse(val *IambackofficeCreateGroupResponse) *NullableIambackofficeCreateGroupResponse {
	return &NullableIambackofficeCreateGroupResponse{value: val, isSet: true}
}

func (v NullableIambackofficeCreateGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIambackofficeCreateGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
