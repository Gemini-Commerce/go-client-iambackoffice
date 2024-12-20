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

// checks if the IambackofficeAssignUserToGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IambackofficeAssignUserToGroupRequest{}

// IambackofficeAssignUserToGroupRequest struct for IambackofficeAssignUserToGroupRequest
type IambackofficeAssignUserToGroupRequest struct {
	TenantId             *string `json:"tenantId,omitempty"`
	UserId               *string `json:"userId,omitempty"`
	GroupId              *string `json:"groupId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IambackofficeAssignUserToGroupRequest IambackofficeAssignUserToGroupRequest

// NewIambackofficeAssignUserToGroupRequest instantiates a new IambackofficeAssignUserToGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIambackofficeAssignUserToGroupRequest() *IambackofficeAssignUserToGroupRequest {
	this := IambackofficeAssignUserToGroupRequest{}
	return &this
}

// NewIambackofficeAssignUserToGroupRequestWithDefaults instantiates a new IambackofficeAssignUserToGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIambackofficeAssignUserToGroupRequestWithDefaults() *IambackofficeAssignUserToGroupRequest {
	this := IambackofficeAssignUserToGroupRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *IambackofficeAssignUserToGroupRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeAssignUserToGroupRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *IambackofficeAssignUserToGroupRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *IambackofficeAssignUserToGroupRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *IambackofficeAssignUserToGroupRequest) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeAssignUserToGroupRequest) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *IambackofficeAssignUserToGroupRequest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *IambackofficeAssignUserToGroupRequest) SetUserId(v string) {
	o.UserId = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *IambackofficeAssignUserToGroupRequest) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeAssignUserToGroupRequest) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *IambackofficeAssignUserToGroupRequest) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *IambackofficeAssignUserToGroupRequest) SetGroupId(v string) {
	o.GroupId = &v
}

func (o IambackofficeAssignUserToGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IambackofficeAssignUserToGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IambackofficeAssignUserToGroupRequest) UnmarshalJSON(data []byte) (err error) {
	varIambackofficeAssignUserToGroupRequest := _IambackofficeAssignUserToGroupRequest{}

	err = json.Unmarshal(data, &varIambackofficeAssignUserToGroupRequest)

	if err != nil {
		return err
	}

	*o = IambackofficeAssignUserToGroupRequest(varIambackofficeAssignUserToGroupRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "userId")
		delete(additionalProperties, "groupId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *IambackofficeAssignUserToGroupRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *IambackofficeAssignUserToGroupRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableIambackofficeAssignUserToGroupRequest struct {
	value *IambackofficeAssignUserToGroupRequest
	isSet bool
}

func (v NullableIambackofficeAssignUserToGroupRequest) Get() *IambackofficeAssignUserToGroupRequest {
	return v.value
}

func (v *NullableIambackofficeAssignUserToGroupRequest) Set(val *IambackofficeAssignUserToGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIambackofficeAssignUserToGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIambackofficeAssignUserToGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIambackofficeAssignUserToGroupRequest(val *IambackofficeAssignUserToGroupRequest) *NullableIambackofficeAssignUserToGroupRequest {
	return &NullableIambackofficeAssignUserToGroupRequest{value: val, isSet: true}
}

func (v NullableIambackofficeAssignUserToGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIambackofficeAssignUserToGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
