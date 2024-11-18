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

// checks if the IambackofficeLogoutResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IambackofficeLogoutResponse{}

// IambackofficeLogoutResponse struct for IambackofficeLogoutResponse
type IambackofficeLogoutResponse struct {
	Success              *bool `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IambackofficeLogoutResponse IambackofficeLogoutResponse

// NewIambackofficeLogoutResponse instantiates a new IambackofficeLogoutResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIambackofficeLogoutResponse() *IambackofficeLogoutResponse {
	this := IambackofficeLogoutResponse{}
	return &this
}

// NewIambackofficeLogoutResponseWithDefaults instantiates a new IambackofficeLogoutResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIambackofficeLogoutResponseWithDefaults() *IambackofficeLogoutResponse {
	this := IambackofficeLogoutResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *IambackofficeLogoutResponse) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeLogoutResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *IambackofficeLogoutResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *IambackofficeLogoutResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o IambackofficeLogoutResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IambackofficeLogoutResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IambackofficeLogoutResponse) UnmarshalJSON(data []byte) (err error) {
	varIambackofficeLogoutResponse := _IambackofficeLogoutResponse{}

	err = json.Unmarshal(data, &varIambackofficeLogoutResponse)

	if err != nil {
		return err
	}

	*o = IambackofficeLogoutResponse(varIambackofficeLogoutResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *IambackofficeLogoutResponse) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *IambackofficeLogoutResponse) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableIambackofficeLogoutResponse struct {
	value *IambackofficeLogoutResponse
	isSet bool
}

func (v NullableIambackofficeLogoutResponse) Get() *IambackofficeLogoutResponse {
	return v.value
}

func (v *NullableIambackofficeLogoutResponse) Set(val *IambackofficeLogoutResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIambackofficeLogoutResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIambackofficeLogoutResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIambackofficeLogoutResponse(val *IambackofficeLogoutResponse) *NullableIambackofficeLogoutResponse {
	return &NullableIambackofficeLogoutResponse{value: val, isSet: true}
}

func (v NullableIambackofficeLogoutResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIambackofficeLogoutResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
