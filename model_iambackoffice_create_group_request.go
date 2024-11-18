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

// checks if the IambackofficeCreateGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IambackofficeCreateGroupRequest{}

// IambackofficeCreateGroupRequest struct for IambackofficeCreateGroupRequest
type IambackofficeCreateGroupRequest struct {
	TenantId             *string            `json:"tenantId,omitempty"`
	Name                 *string            `json:"name,omitempty"`
	Data                 *map[string]string `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IambackofficeCreateGroupRequest IambackofficeCreateGroupRequest

// NewIambackofficeCreateGroupRequest instantiates a new IambackofficeCreateGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIambackofficeCreateGroupRequest() *IambackofficeCreateGroupRequest {
	this := IambackofficeCreateGroupRequest{}
	return &this
}

// NewIambackofficeCreateGroupRequestWithDefaults instantiates a new IambackofficeCreateGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIambackofficeCreateGroupRequestWithDefaults() *IambackofficeCreateGroupRequest {
	this := IambackofficeCreateGroupRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *IambackofficeCreateGroupRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeCreateGroupRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *IambackofficeCreateGroupRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *IambackofficeCreateGroupRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IambackofficeCreateGroupRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeCreateGroupRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IambackofficeCreateGroupRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IambackofficeCreateGroupRequest) SetName(v string) {
	o.Name = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *IambackofficeCreateGroupRequest) GetData() map[string]string {
	if o == nil || IsNil(o.Data) {
		var ret map[string]string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeCreateGroupRequest) GetDataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *IambackofficeCreateGroupRequest) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]string and assigns it to the Data field.
func (o *IambackofficeCreateGroupRequest) SetData(v map[string]string) {
	o.Data = &v
}

func (o IambackofficeCreateGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IambackofficeCreateGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IambackofficeCreateGroupRequest) UnmarshalJSON(data []byte) (err error) {
	varIambackofficeCreateGroupRequest := _IambackofficeCreateGroupRequest{}

	err = json.Unmarshal(data, &varIambackofficeCreateGroupRequest)

	if err != nil {
		return err
	}

	*o = IambackofficeCreateGroupRequest(varIambackofficeCreateGroupRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "name")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIambackofficeCreateGroupRequest struct {
	value *IambackofficeCreateGroupRequest
	isSet bool
}

func (v NullableIambackofficeCreateGroupRequest) Get() *IambackofficeCreateGroupRequest {
	return v.value
}

func (v *NullableIambackofficeCreateGroupRequest) Set(val *IambackofficeCreateGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIambackofficeCreateGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIambackofficeCreateGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIambackofficeCreateGroupRequest(val *IambackofficeCreateGroupRequest) *NullableIambackofficeCreateGroupRequest {
	return &NullableIambackofficeCreateGroupRequest{value: val, isSet: true}
}

func (v NullableIambackofficeCreateGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIambackofficeCreateGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
