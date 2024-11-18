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

// checks if the IambackofficeLoginResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IambackofficeLoginResponse{}

// IambackofficeLoginResponse struct for IambackofficeLoginResponse
type IambackofficeLoginResponse struct {
	Tokens               *ProtobufAny                   `json:"tokens,omitempty"`
	User                 *IambackofficeUser             `json:"user,omitempty"`
	Methods              []IambackofficeTwoFactorMethod `json:"methods,omitempty"`
	TwoFactorId          *string                        `json:"twoFactorId,omitempty"`
	TenantId             *string                        `json:"tenantId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IambackofficeLoginResponse IambackofficeLoginResponse

// NewIambackofficeLoginResponse instantiates a new IambackofficeLoginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIambackofficeLoginResponse() *IambackofficeLoginResponse {
	this := IambackofficeLoginResponse{}
	return &this
}

// NewIambackofficeLoginResponseWithDefaults instantiates a new IambackofficeLoginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIambackofficeLoginResponseWithDefaults() *IambackofficeLoginResponse {
	this := IambackofficeLoginResponse{}
	return &this
}

// GetTokens returns the Tokens field value if set, zero value otherwise.
func (o *IambackofficeLoginResponse) GetTokens() ProtobufAny {
	if o == nil || IsNil(o.Tokens) {
		var ret ProtobufAny
		return ret
	}
	return *o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeLoginResponse) GetTokensOk() (*ProtobufAny, bool) {
	if o == nil || IsNil(o.Tokens) {
		return nil, false
	}
	return o.Tokens, true
}

// HasTokens returns a boolean if a field has been set.
func (o *IambackofficeLoginResponse) HasTokens() bool {
	if o != nil && !IsNil(o.Tokens) {
		return true
	}

	return false
}

// SetTokens gets a reference to the given ProtobufAny and assigns it to the Tokens field.
func (o *IambackofficeLoginResponse) SetTokens(v ProtobufAny) {
	o.Tokens = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *IambackofficeLoginResponse) GetUser() IambackofficeUser {
	if o == nil || IsNil(o.User) {
		var ret IambackofficeUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeLoginResponse) GetUserOk() (*IambackofficeUser, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *IambackofficeLoginResponse) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given IambackofficeUser and assigns it to the User field.
func (o *IambackofficeLoginResponse) SetUser(v IambackofficeUser) {
	o.User = &v
}

// GetMethods returns the Methods field value if set, zero value otherwise.
func (o *IambackofficeLoginResponse) GetMethods() []IambackofficeTwoFactorMethod {
	if o == nil || IsNil(o.Methods) {
		var ret []IambackofficeTwoFactorMethod
		return ret
	}
	return o.Methods
}

// GetMethodsOk returns a tuple with the Methods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeLoginResponse) GetMethodsOk() ([]IambackofficeTwoFactorMethod, bool) {
	if o == nil || IsNil(o.Methods) {
		return nil, false
	}
	return o.Methods, true
}

// HasMethods returns a boolean if a field has been set.
func (o *IambackofficeLoginResponse) HasMethods() bool {
	if o != nil && !IsNil(o.Methods) {
		return true
	}

	return false
}

// SetMethods gets a reference to the given []IambackofficeTwoFactorMethod and assigns it to the Methods field.
func (o *IambackofficeLoginResponse) SetMethods(v []IambackofficeTwoFactorMethod) {
	o.Methods = v
}

// GetTwoFactorId returns the TwoFactorId field value if set, zero value otherwise.
func (o *IambackofficeLoginResponse) GetTwoFactorId() string {
	if o == nil || IsNil(o.TwoFactorId) {
		var ret string
		return ret
	}
	return *o.TwoFactorId
}

// GetTwoFactorIdOk returns a tuple with the TwoFactorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeLoginResponse) GetTwoFactorIdOk() (*string, bool) {
	if o == nil || IsNil(o.TwoFactorId) {
		return nil, false
	}
	return o.TwoFactorId, true
}

// HasTwoFactorId returns a boolean if a field has been set.
func (o *IambackofficeLoginResponse) HasTwoFactorId() bool {
	if o != nil && !IsNil(o.TwoFactorId) {
		return true
	}

	return false
}

// SetTwoFactorId gets a reference to the given string and assigns it to the TwoFactorId field.
func (o *IambackofficeLoginResponse) SetTwoFactorId(v string) {
	o.TwoFactorId = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *IambackofficeLoginResponse) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeLoginResponse) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *IambackofficeLoginResponse) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *IambackofficeLoginResponse) SetTenantId(v string) {
	o.TenantId = &v
}

func (o IambackofficeLoginResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IambackofficeLoginResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tokens) {
		toSerialize["tokens"] = o.Tokens
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Methods) {
		toSerialize["methods"] = o.Methods
	}
	if !IsNil(o.TwoFactorId) {
		toSerialize["twoFactorId"] = o.TwoFactorId
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IambackofficeLoginResponse) UnmarshalJSON(data []byte) (err error) {
	varIambackofficeLoginResponse := _IambackofficeLoginResponse{}

	err = json.Unmarshal(data, &varIambackofficeLoginResponse)

	if err != nil {
		return err
	}

	*o = IambackofficeLoginResponse(varIambackofficeLoginResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tokens")
		delete(additionalProperties, "user")
		delete(additionalProperties, "methods")
		delete(additionalProperties, "twoFactorId")
		delete(additionalProperties, "tenantId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *IambackofficeLoginResponse) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *IambackofficeLoginResponse) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableIambackofficeLoginResponse struct {
	value *IambackofficeLoginResponse
	isSet bool
}

func (v NullableIambackofficeLoginResponse) Get() *IambackofficeLoginResponse {
	return v.value
}

func (v *NullableIambackofficeLoginResponse) Set(val *IambackofficeLoginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIambackofficeLoginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIambackofficeLoginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIambackofficeLoginResponse(val *IambackofficeLoginResponse) *NullableIambackofficeLoginResponse {
	return &NullableIambackofficeLoginResponse{value: val, isSet: true}
}

func (v NullableIambackofficeLoginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIambackofficeLoginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
