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

// checks if the IambackofficeAuthenticatorConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IambackofficeAuthenticatorConfiguration{}

// IambackofficeAuthenticatorConfiguration struct for IambackofficeAuthenticatorConfiguration
type IambackofficeAuthenticatorConfiguration struct {
	Algorithm            *string `json:"algorithm,omitempty"`
	CodeLength           *int32  `json:"codeLength,omitempty"`
	TimeStep             *int32  `json:"timeStep,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IambackofficeAuthenticatorConfiguration IambackofficeAuthenticatorConfiguration

// NewIambackofficeAuthenticatorConfiguration instantiates a new IambackofficeAuthenticatorConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIambackofficeAuthenticatorConfiguration() *IambackofficeAuthenticatorConfiguration {
	this := IambackofficeAuthenticatorConfiguration{}
	return &this
}

// NewIambackofficeAuthenticatorConfigurationWithDefaults instantiates a new IambackofficeAuthenticatorConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIambackofficeAuthenticatorConfigurationWithDefaults() *IambackofficeAuthenticatorConfiguration {
	this := IambackofficeAuthenticatorConfiguration{}
	return &this
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *IambackofficeAuthenticatorConfiguration) GetAlgorithm() string {
	if o == nil || IsNil(o.Algorithm) {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeAuthenticatorConfiguration) GetAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.Algorithm) {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *IambackofficeAuthenticatorConfiguration) HasAlgorithm() bool {
	if o != nil && !IsNil(o.Algorithm) {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *IambackofficeAuthenticatorConfiguration) SetAlgorithm(v string) {
	o.Algorithm = &v
}

// GetCodeLength returns the CodeLength field value if set, zero value otherwise.
func (o *IambackofficeAuthenticatorConfiguration) GetCodeLength() int32 {
	if o == nil || IsNil(o.CodeLength) {
		var ret int32
		return ret
	}
	return *o.CodeLength
}

// GetCodeLengthOk returns a tuple with the CodeLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeAuthenticatorConfiguration) GetCodeLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.CodeLength) {
		return nil, false
	}
	return o.CodeLength, true
}

// HasCodeLength returns a boolean if a field has been set.
func (o *IambackofficeAuthenticatorConfiguration) HasCodeLength() bool {
	if o != nil && !IsNil(o.CodeLength) {
		return true
	}

	return false
}

// SetCodeLength gets a reference to the given int32 and assigns it to the CodeLength field.
func (o *IambackofficeAuthenticatorConfiguration) SetCodeLength(v int32) {
	o.CodeLength = &v
}

// GetTimeStep returns the TimeStep field value if set, zero value otherwise.
func (o *IambackofficeAuthenticatorConfiguration) GetTimeStep() int32 {
	if o == nil || IsNil(o.TimeStep) {
		var ret int32
		return ret
	}
	return *o.TimeStep
}

// GetTimeStepOk returns a tuple with the TimeStep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeAuthenticatorConfiguration) GetTimeStepOk() (*int32, bool) {
	if o == nil || IsNil(o.TimeStep) {
		return nil, false
	}
	return o.TimeStep, true
}

// HasTimeStep returns a boolean if a field has been set.
func (o *IambackofficeAuthenticatorConfiguration) HasTimeStep() bool {
	if o != nil && !IsNil(o.TimeStep) {
		return true
	}

	return false
}

// SetTimeStep gets a reference to the given int32 and assigns it to the TimeStep field.
func (o *IambackofficeAuthenticatorConfiguration) SetTimeStep(v int32) {
	o.TimeStep = &v
}

func (o IambackofficeAuthenticatorConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IambackofficeAuthenticatorConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Algorithm) {
		toSerialize["algorithm"] = o.Algorithm
	}
	if !IsNil(o.CodeLength) {
		toSerialize["codeLength"] = o.CodeLength
	}
	if !IsNil(o.TimeStep) {
		toSerialize["timeStep"] = o.TimeStep
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IambackofficeAuthenticatorConfiguration) UnmarshalJSON(data []byte) (err error) {
	varIambackofficeAuthenticatorConfiguration := _IambackofficeAuthenticatorConfiguration{}

	err = json.Unmarshal(data, &varIambackofficeAuthenticatorConfiguration)

	if err != nil {
		return err
	}

	*o = IambackofficeAuthenticatorConfiguration(varIambackofficeAuthenticatorConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "algorithm")
		delete(additionalProperties, "codeLength")
		delete(additionalProperties, "timeStep")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIambackofficeAuthenticatorConfiguration struct {
	value *IambackofficeAuthenticatorConfiguration
	isSet bool
}

func (v NullableIambackofficeAuthenticatorConfiguration) Get() *IambackofficeAuthenticatorConfiguration {
	return v.value
}

func (v *NullableIambackofficeAuthenticatorConfiguration) Set(val *IambackofficeAuthenticatorConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableIambackofficeAuthenticatorConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableIambackofficeAuthenticatorConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIambackofficeAuthenticatorConfiguration(val *IambackofficeAuthenticatorConfiguration) *NullableIambackofficeAuthenticatorConfiguration {
	return &NullableIambackofficeAuthenticatorConfiguration{value: val, isSet: true}
}

func (v NullableIambackofficeAuthenticatorConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIambackofficeAuthenticatorConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
