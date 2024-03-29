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

// checks if the IambackofficeMetaData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IambackofficeMetaData{}

// IambackofficeMetaData struct for IambackofficeMetaData
type IambackofficeMetaData struct {
	Device *IambackofficeDevice `json:"device,omitempty"`
}

// NewIambackofficeMetaData instantiates a new IambackofficeMetaData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIambackofficeMetaData() *IambackofficeMetaData {
	this := IambackofficeMetaData{}
	return &this
}

// NewIambackofficeMetaDataWithDefaults instantiates a new IambackofficeMetaData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIambackofficeMetaDataWithDefaults() *IambackofficeMetaData {
	this := IambackofficeMetaData{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *IambackofficeMetaData) GetDevice() IambackofficeDevice {
	if o == nil || IsNil(o.Device) {
		var ret IambackofficeDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeMetaData) GetDeviceOk() (*IambackofficeDevice, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *IambackofficeMetaData) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given IambackofficeDevice and assigns it to the Device field.
func (o *IambackofficeMetaData) SetDevice(v IambackofficeDevice) {
	o.Device = &v
}

func (o IambackofficeMetaData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IambackofficeMetaData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	return toSerialize, nil
}

type NullableIambackofficeMetaData struct {
	value *IambackofficeMetaData
	isSet bool
}

func (v NullableIambackofficeMetaData) Get() *IambackofficeMetaData {
	return v.value
}

func (v *NullableIambackofficeMetaData) Set(val *IambackofficeMetaData) {
	v.value = val
	v.isSet = true
}

func (v NullableIambackofficeMetaData) IsSet() bool {
	return v.isSet
}

func (v *NullableIambackofficeMetaData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIambackofficeMetaData(val *IambackofficeMetaData) *NullableIambackofficeMetaData {
	return &NullableIambackofficeMetaData{value: val, isSet: true}
}

func (v NullableIambackofficeMetaData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIambackofficeMetaData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


