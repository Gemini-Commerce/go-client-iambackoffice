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

// checks if the IambackofficeUnassignRolesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IambackofficeUnassignRolesResponse{}

// IambackofficeUnassignRolesResponse struct for IambackofficeUnassignRolesResponse
type IambackofficeUnassignRolesResponse struct {
	RoleCodes []string `json:"roleCodes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IambackofficeUnassignRolesResponse IambackofficeUnassignRolesResponse

// NewIambackofficeUnassignRolesResponse instantiates a new IambackofficeUnassignRolesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIambackofficeUnassignRolesResponse() *IambackofficeUnassignRolesResponse {
	this := IambackofficeUnassignRolesResponse{}
	return &this
}

// NewIambackofficeUnassignRolesResponseWithDefaults instantiates a new IambackofficeUnassignRolesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIambackofficeUnassignRolesResponseWithDefaults() *IambackofficeUnassignRolesResponse {
	this := IambackofficeUnassignRolesResponse{}
	return &this
}

// GetRoleCodes returns the RoleCodes field value if set, zero value otherwise.
func (o *IambackofficeUnassignRolesResponse) GetRoleCodes() []string {
	if o == nil || IsNil(o.RoleCodes) {
		var ret []string
		return ret
	}
	return o.RoleCodes
}

// GetRoleCodesOk returns a tuple with the RoleCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeUnassignRolesResponse) GetRoleCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.RoleCodes) {
		return nil, false
	}
	return o.RoleCodes, true
}

// &#39;Has&#39;RoleCodes returns a boolean if a field has been set.
func (o *IambackofficeUnassignRolesResponse) &#39;Has&#39;RoleCodes() bool {
	if o != nil && !IsNil(o.RoleCodes) {
		return true
	}

	return false
}

// SetRoleCodes gets a reference to the given []string and assigns it to the RoleCodes field.
func (o *IambackofficeUnassignRolesResponse) SetRoleCodes(v []string) {
	o.RoleCodes = v
}

func (o IambackofficeUnassignRolesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IambackofficeUnassignRolesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoleCodes) {
		toSerialize["roleCodes"] = o.RoleCodes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IambackofficeUnassignRolesResponse) UnmarshalJSON(data []byte) (err error) {
	varIambackofficeUnassignRolesResponse := _IambackofficeUnassignRolesResponse{}

	err = json.Unmarshal(data, &varIambackofficeUnassignRolesResponse)

	if err != nil {
		return err
	}

	*o = IambackofficeUnassignRolesResponse(varIambackofficeUnassignRolesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "roleCodes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *IambackofficeUnassignRolesResponse) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *IambackofficeUnassignRolesResponse) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableIambackofficeUnassignRolesResponse struct {
	value *IambackofficeUnassignRolesResponse
	isSet bool
}

func (v NullableIambackofficeUnassignRolesResponse) Get() *IambackofficeUnassignRolesResponse {
	return v.value
}

func (v *NullableIambackofficeUnassignRolesResponse) Set(val *IambackofficeUnassignRolesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIambackofficeUnassignRolesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIambackofficeUnassignRolesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIambackofficeUnassignRolesResponse(val *IambackofficeUnassignRolesResponse) *NullableIambackofficeUnassignRolesResponse {
	return &NullableIambackofficeUnassignRolesResponse{value: val, isSet: true}
}

func (v NullableIambackofficeUnassignRolesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIambackofficeUnassignRolesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


