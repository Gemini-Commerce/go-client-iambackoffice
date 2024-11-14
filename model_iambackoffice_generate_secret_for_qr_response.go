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

// checks if the IambackofficeGenerateSecretForQRResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IambackofficeGenerateSecretForQRResponse{}

// IambackofficeGenerateSecretForQRResponse struct for IambackofficeGenerateSecretForQRResponse
type IambackofficeGenerateSecretForQRResponse struct {
	Secret *string `json:"secret,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IambackofficeGenerateSecretForQRResponse IambackofficeGenerateSecretForQRResponse

// NewIambackofficeGenerateSecretForQRResponse instantiates a new IambackofficeGenerateSecretForQRResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIambackofficeGenerateSecretForQRResponse() *IambackofficeGenerateSecretForQRResponse {
	this := IambackofficeGenerateSecretForQRResponse{}
	return &this
}

// NewIambackofficeGenerateSecretForQRResponseWithDefaults instantiates a new IambackofficeGenerateSecretForQRResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIambackofficeGenerateSecretForQRResponseWithDefaults() *IambackofficeGenerateSecretForQRResponse {
	this := IambackofficeGenerateSecretForQRResponse{}
	return &this
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *IambackofficeGenerateSecretForQRResponse) GetSecret() string {
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeGenerateSecretForQRResponse) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// &#39;Has&#39;Secret returns a boolean if a field has been set.
func (o *IambackofficeGenerateSecretForQRResponse) &#39;Has&#39;Secret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *IambackofficeGenerateSecretForQRResponse) SetSecret(v string) {
	o.Secret = &v
}

func (o IambackofficeGenerateSecretForQRResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IambackofficeGenerateSecretForQRResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IambackofficeGenerateSecretForQRResponse) UnmarshalJSON(data []byte) (err error) {
	varIambackofficeGenerateSecretForQRResponse := _IambackofficeGenerateSecretForQRResponse{}

	err = json.Unmarshal(data, &varIambackofficeGenerateSecretForQRResponse)

	if err != nil {
		return err
	}

	*o = IambackofficeGenerateSecretForQRResponse(varIambackofficeGenerateSecretForQRResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "secret")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *IambackofficeGenerateSecretForQRResponse) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *IambackofficeGenerateSecretForQRResponse) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableIambackofficeGenerateSecretForQRResponse struct {
	value *IambackofficeGenerateSecretForQRResponse
	isSet bool
}

func (v NullableIambackofficeGenerateSecretForQRResponse) Get() *IambackofficeGenerateSecretForQRResponse {
	return v.value
}

func (v *NullableIambackofficeGenerateSecretForQRResponse) Set(val *IambackofficeGenerateSecretForQRResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIambackofficeGenerateSecretForQRResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIambackofficeGenerateSecretForQRResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIambackofficeGenerateSecretForQRResponse(val *IambackofficeGenerateSecretForQRResponse) *NullableIambackofficeGenerateSecretForQRResponse {
	return &NullableIambackofficeGenerateSecretForQRResponse{value: val, isSet: true}
}

func (v NullableIambackofficeGenerateSecretForQRResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIambackofficeGenerateSecretForQRResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


