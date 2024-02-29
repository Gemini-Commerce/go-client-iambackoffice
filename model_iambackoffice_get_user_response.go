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

// checks if the IambackofficeGetUserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IambackofficeGetUserResponse{}

// IambackofficeGetUserResponse struct for IambackofficeGetUserResponse
type IambackofficeGetUserResponse struct {
	User *IambackofficeUser `json:"user,omitempty"`
}

// NewIambackofficeGetUserResponse instantiates a new IambackofficeGetUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIambackofficeGetUserResponse() *IambackofficeGetUserResponse {
	this := IambackofficeGetUserResponse{}
	return &this
}

// NewIambackofficeGetUserResponseWithDefaults instantiates a new IambackofficeGetUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIambackofficeGetUserResponseWithDefaults() *IambackofficeGetUserResponse {
	this := IambackofficeGetUserResponse{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *IambackofficeGetUserResponse) GetUser() IambackofficeUser {
	if o == nil || IsNil(o.User) {
		var ret IambackofficeUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeGetUserResponse) GetUserOk() (*IambackofficeUser, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *IambackofficeGetUserResponse) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given IambackofficeUser and assigns it to the User field.
func (o *IambackofficeGetUserResponse) SetUser(v IambackofficeUser) {
	o.User = &v
}

func (o IambackofficeGetUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IambackofficeGetUserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableIambackofficeGetUserResponse struct {
	value *IambackofficeGetUserResponse
	isSet bool
}

func (v NullableIambackofficeGetUserResponse) Get() *IambackofficeGetUserResponse {
	return v.value
}

func (v *NullableIambackofficeGetUserResponse) Set(val *IambackofficeGetUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIambackofficeGetUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIambackofficeGetUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIambackofficeGetUserResponse(val *IambackofficeGetUserResponse) *NullableIambackofficeGetUserResponse {
	return &NullableIambackofficeGetUserResponse{value: val, isSet: true}
}

func (v NullableIambackofficeGetUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIambackofficeGetUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

