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

// checks if the IambackofficeGetUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IambackofficeGetUserRequest{}

// IambackofficeGetUserRequest struct for IambackofficeGetUserRequest
type IambackofficeGetUserRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	UserId *string `json:"userId,omitempty"`
	Email *string `json:"email,omitempty"`
	Username *string `json:"username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IambackofficeGetUserRequest IambackofficeGetUserRequest

// NewIambackofficeGetUserRequest instantiates a new IambackofficeGetUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIambackofficeGetUserRequest() *IambackofficeGetUserRequest {
	this := IambackofficeGetUserRequest{}
	return &this
}

// NewIambackofficeGetUserRequestWithDefaults instantiates a new IambackofficeGetUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIambackofficeGetUserRequestWithDefaults() *IambackofficeGetUserRequest {
	this := IambackofficeGetUserRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *IambackofficeGetUserRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeGetUserRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *IambackofficeGetUserRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *IambackofficeGetUserRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *IambackofficeGetUserRequest) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeGetUserRequest) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *IambackofficeGetUserRequest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *IambackofficeGetUserRequest) SetUserId(v string) {
	o.UserId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *IambackofficeGetUserRequest) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeGetUserRequest) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *IambackofficeGetUserRequest) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *IambackofficeGetUserRequest) SetEmail(v string) {
	o.Email = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *IambackofficeGetUserRequest) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeGetUserRequest) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *IambackofficeGetUserRequest) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *IambackofficeGetUserRequest) SetUsername(v string) {
	o.Username = &v
}

func (o IambackofficeGetUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IambackofficeGetUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IambackofficeGetUserRequest) UnmarshalJSON(data []byte) (err error) {
	varIambackofficeGetUserRequest := _IambackofficeGetUserRequest{}

	err = json.Unmarshal(data, &varIambackofficeGetUserRequest)

	if err != nil {
		return err
	}

	*o = IambackofficeGetUserRequest(varIambackofficeGetUserRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "userId")
		delete(additionalProperties, "email")
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *IambackofficeGetUserRequest) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *IambackofficeGetUserRequest) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableIambackofficeGetUserRequest struct {
	value *IambackofficeGetUserRequest
	isSet bool
}

func (v NullableIambackofficeGetUserRequest) Get() *IambackofficeGetUserRequest {
	return v.value
}

func (v *NullableIambackofficeGetUserRequest) Set(val *IambackofficeGetUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIambackofficeGetUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIambackofficeGetUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIambackofficeGetUserRequest(val *IambackofficeGetUserRequest) *NullableIambackofficeGetUserRequest {
	return &NullableIambackofficeGetUserRequest{value: val, isSet: true}
}

func (v NullableIambackofficeGetUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIambackofficeGetUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


