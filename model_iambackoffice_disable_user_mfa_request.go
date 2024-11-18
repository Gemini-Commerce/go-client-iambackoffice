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

// checks if the IambackofficeDisableUserMfaRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IambackofficeDisableUserMfaRequest{}

// IambackofficeDisableUserMfaRequest struct for IambackofficeDisableUserMfaRequest
type IambackofficeDisableUserMfaRequest struct {
	UserId               *string `json:"userId,omitempty"`
	Code                 *string `json:"code,omitempty"`
	MethodId             *string `json:"methodId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IambackofficeDisableUserMfaRequest IambackofficeDisableUserMfaRequest

// NewIambackofficeDisableUserMfaRequest instantiates a new IambackofficeDisableUserMfaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIambackofficeDisableUserMfaRequest() *IambackofficeDisableUserMfaRequest {
	this := IambackofficeDisableUserMfaRequest{}
	return &this
}

// NewIambackofficeDisableUserMfaRequestWithDefaults instantiates a new IambackofficeDisableUserMfaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIambackofficeDisableUserMfaRequestWithDefaults() *IambackofficeDisableUserMfaRequest {
	this := IambackofficeDisableUserMfaRequest{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *IambackofficeDisableUserMfaRequest) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeDisableUserMfaRequest) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *IambackofficeDisableUserMfaRequest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *IambackofficeDisableUserMfaRequest) SetUserId(v string) {
	o.UserId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *IambackofficeDisableUserMfaRequest) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeDisableUserMfaRequest) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *IambackofficeDisableUserMfaRequest) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *IambackofficeDisableUserMfaRequest) SetCode(v string) {
	o.Code = &v
}

// GetMethodId returns the MethodId field value if set, zero value otherwise.
func (o *IambackofficeDisableUserMfaRequest) GetMethodId() string {
	if o == nil || IsNil(o.MethodId) {
		var ret string
		return ret
	}
	return *o.MethodId
}

// GetMethodIdOk returns a tuple with the MethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeDisableUserMfaRequest) GetMethodIdOk() (*string, bool) {
	if o == nil || IsNil(o.MethodId) {
		return nil, false
	}
	return o.MethodId, true
}

// HasMethodId returns a boolean if a field has been set.
func (o *IambackofficeDisableUserMfaRequest) HasMethodId() bool {
	if o != nil && !IsNil(o.MethodId) {
		return true
	}

	return false
}

// SetMethodId gets a reference to the given string and assigns it to the MethodId field.
func (o *IambackofficeDisableUserMfaRequest) SetMethodId(v string) {
	o.MethodId = &v
}

func (o IambackofficeDisableUserMfaRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IambackofficeDisableUserMfaRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.MethodId) {
		toSerialize["methodId"] = o.MethodId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IambackofficeDisableUserMfaRequest) UnmarshalJSON(data []byte) (err error) {
	varIambackofficeDisableUserMfaRequest := _IambackofficeDisableUserMfaRequest{}

	err = json.Unmarshal(data, &varIambackofficeDisableUserMfaRequest)

	if err != nil {
		return err
	}

	*o = IambackofficeDisableUserMfaRequest(varIambackofficeDisableUserMfaRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "userId")
		delete(additionalProperties, "code")
		delete(additionalProperties, "methodId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *IambackofficeDisableUserMfaRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populate the value of well-known types
func (o *IambackofficeDisableUserMfaRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableIambackofficeDisableUserMfaRequest struct {
	value *IambackofficeDisableUserMfaRequest
	isSet bool
}

func (v NullableIambackofficeDisableUserMfaRequest) Get() *IambackofficeDisableUserMfaRequest {
	return v.value
}

func (v *NullableIambackofficeDisableUserMfaRequest) Set(val *IambackofficeDisableUserMfaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIambackofficeDisableUserMfaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIambackofficeDisableUserMfaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIambackofficeDisableUserMfaRequest(val *IambackofficeDisableUserMfaRequest) *NullableIambackofficeDisableUserMfaRequest {
	return &NullableIambackofficeDisableUserMfaRequest{value: val, isSet: true}
}

func (v NullableIambackofficeDisableUserMfaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIambackofficeDisableUserMfaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
