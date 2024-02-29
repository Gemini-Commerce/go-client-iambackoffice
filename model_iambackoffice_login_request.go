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

// checks if the IambackofficeLoginRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IambackofficeLoginRequest{}

// IambackofficeLoginRequest struct for IambackofficeLoginRequest
type IambackofficeLoginRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	IpAddress *string `json:"ipAddress,omitempty"`
	LoginId *string `json:"loginId,omitempty"`
	MetaData *IambackofficeMetaData `json:"metaData,omitempty"`
	NoTokens *bool `json:"noTokens,omitempty"`
	Password *string `json:"password,omitempty"`
	TwoFactorTrustId *string `json:"twoFactorTrustId,omitempty"`
}

// NewIambackofficeLoginRequest instantiates a new IambackofficeLoginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIambackofficeLoginRequest() *IambackofficeLoginRequest {
	this := IambackofficeLoginRequest{}
	return &this
}

// NewIambackofficeLoginRequestWithDefaults instantiates a new IambackofficeLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIambackofficeLoginRequestWithDefaults() *IambackofficeLoginRequest {
	this := IambackofficeLoginRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *IambackofficeLoginRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeLoginRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *IambackofficeLoginRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *IambackofficeLoginRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *IambackofficeLoginRequest) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeLoginRequest) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *IambackofficeLoginRequest) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *IambackofficeLoginRequest) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetLoginId returns the LoginId field value if set, zero value otherwise.
func (o *IambackofficeLoginRequest) GetLoginId() string {
	if o == nil || IsNil(o.LoginId) {
		var ret string
		return ret
	}
	return *o.LoginId
}

// GetLoginIdOk returns a tuple with the LoginId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeLoginRequest) GetLoginIdOk() (*string, bool) {
	if o == nil || IsNil(o.LoginId) {
		return nil, false
	}
	return o.LoginId, true
}

// HasLoginId returns a boolean if a field has been set.
func (o *IambackofficeLoginRequest) HasLoginId() bool {
	if o != nil && !IsNil(o.LoginId) {
		return true
	}

	return false
}

// SetLoginId gets a reference to the given string and assigns it to the LoginId field.
func (o *IambackofficeLoginRequest) SetLoginId(v string) {
	o.LoginId = &v
}

// GetMetaData returns the MetaData field value if set, zero value otherwise.
func (o *IambackofficeLoginRequest) GetMetaData() IambackofficeMetaData {
	if o == nil || IsNil(o.MetaData) {
		var ret IambackofficeMetaData
		return ret
	}
	return *o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeLoginRequest) GetMetaDataOk() (*IambackofficeMetaData, bool) {
	if o == nil || IsNil(o.MetaData) {
		return nil, false
	}
	return o.MetaData, true
}

// HasMetaData returns a boolean if a field has been set.
func (o *IambackofficeLoginRequest) HasMetaData() bool {
	if o != nil && !IsNil(o.MetaData) {
		return true
	}

	return false
}

// SetMetaData gets a reference to the given IambackofficeMetaData and assigns it to the MetaData field.
func (o *IambackofficeLoginRequest) SetMetaData(v IambackofficeMetaData) {
	o.MetaData = &v
}

// GetNoTokens returns the NoTokens field value if set, zero value otherwise.
func (o *IambackofficeLoginRequest) GetNoTokens() bool {
	if o == nil || IsNil(o.NoTokens) {
		var ret bool
		return ret
	}
	return *o.NoTokens
}

// GetNoTokensOk returns a tuple with the NoTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeLoginRequest) GetNoTokensOk() (*bool, bool) {
	if o == nil || IsNil(o.NoTokens) {
		return nil, false
	}
	return o.NoTokens, true
}

// HasNoTokens returns a boolean if a field has been set.
func (o *IambackofficeLoginRequest) HasNoTokens() bool {
	if o != nil && !IsNil(o.NoTokens) {
		return true
	}

	return false
}

// SetNoTokens gets a reference to the given bool and assigns it to the NoTokens field.
func (o *IambackofficeLoginRequest) SetNoTokens(v bool) {
	o.NoTokens = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *IambackofficeLoginRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeLoginRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *IambackofficeLoginRequest) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *IambackofficeLoginRequest) SetPassword(v string) {
	o.Password = &v
}

// GetTwoFactorTrustId returns the TwoFactorTrustId field value if set, zero value otherwise.
func (o *IambackofficeLoginRequest) GetTwoFactorTrustId() string {
	if o == nil || IsNil(o.TwoFactorTrustId) {
		var ret string
		return ret
	}
	return *o.TwoFactorTrustId
}

// GetTwoFactorTrustIdOk returns a tuple with the TwoFactorTrustId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeLoginRequest) GetTwoFactorTrustIdOk() (*string, bool) {
	if o == nil || IsNil(o.TwoFactorTrustId) {
		return nil, false
	}
	return o.TwoFactorTrustId, true
}

// HasTwoFactorTrustId returns a boolean if a field has been set.
func (o *IambackofficeLoginRequest) HasTwoFactorTrustId() bool {
	if o != nil && !IsNil(o.TwoFactorTrustId) {
		return true
	}

	return false
}

// SetTwoFactorTrustId gets a reference to the given string and assigns it to the TwoFactorTrustId field.
func (o *IambackofficeLoginRequest) SetTwoFactorTrustId(v string) {
	o.TwoFactorTrustId = &v
}

func (o IambackofficeLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IambackofficeLoginRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if !IsNil(o.LoginId) {
		toSerialize["loginId"] = o.LoginId
	}
	if !IsNil(o.MetaData) {
		toSerialize["metaData"] = o.MetaData
	}
	if !IsNil(o.NoTokens) {
		toSerialize["noTokens"] = o.NoTokens
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.TwoFactorTrustId) {
		toSerialize["twoFactorTrustId"] = o.TwoFactorTrustId
	}
	return toSerialize, nil
}

type NullableIambackofficeLoginRequest struct {
	value *IambackofficeLoginRequest
	isSet bool
}

func (v NullableIambackofficeLoginRequest) Get() *IambackofficeLoginRequest {
	return v.value
}

func (v *NullableIambackofficeLoginRequest) Set(val *IambackofficeLoginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIambackofficeLoginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIambackofficeLoginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIambackofficeLoginRequest(val *IambackofficeLoginRequest) *NullableIambackofficeLoginRequest {
	return &NullableIambackofficeLoginRequest{value: val, isSet: true}
}

func (v NullableIambackofficeLoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIambackofficeLoginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

