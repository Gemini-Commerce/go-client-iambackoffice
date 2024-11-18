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

// checks if the IambackofficeTwoFactorMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IambackofficeTwoFactorMethod{}

// IambackofficeTwoFactorMethod struct for IambackofficeTwoFactorMethod
type IambackofficeTwoFactorMethod struct {
	Authenticator        *IambackofficeAuthenticatorConfiguration `json:"authenticator,omitempty"`
	Email                *string                                  `json:"email,omitempty"`
	Id                   *string                                  `json:"id,omitempty"`
	Method               *string                                  `json:"method,omitempty"`
	MobilePhone          *string                                  `json:"mobilePhone,omitempty"`
	Secret               *string                                  `json:"secret,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IambackofficeTwoFactorMethod IambackofficeTwoFactorMethod

// NewIambackofficeTwoFactorMethod instantiates a new IambackofficeTwoFactorMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIambackofficeTwoFactorMethod() *IambackofficeTwoFactorMethod {
	this := IambackofficeTwoFactorMethod{}
	return &this
}

// NewIambackofficeTwoFactorMethodWithDefaults instantiates a new IambackofficeTwoFactorMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIambackofficeTwoFactorMethodWithDefaults() *IambackofficeTwoFactorMethod {
	this := IambackofficeTwoFactorMethod{}
	return &this
}

// GetAuthenticator returns the Authenticator field value if set, zero value otherwise.
func (o *IambackofficeTwoFactorMethod) GetAuthenticator() IambackofficeAuthenticatorConfiguration {
	if o == nil || IsNil(o.Authenticator) {
		var ret IambackofficeAuthenticatorConfiguration
		return ret
	}
	return *o.Authenticator
}

// GetAuthenticatorOk returns a tuple with the Authenticator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeTwoFactorMethod) GetAuthenticatorOk() (*IambackofficeAuthenticatorConfiguration, bool) {
	if o == nil || IsNil(o.Authenticator) {
		return nil, false
	}
	return o.Authenticator, true
}

// HasAuthenticator returns a boolean if a field has been set.
func (o *IambackofficeTwoFactorMethod) HasAuthenticator() bool {
	if o != nil && !IsNil(o.Authenticator) {
		return true
	}

	return false
}

// SetAuthenticator gets a reference to the given IambackofficeAuthenticatorConfiguration and assigns it to the Authenticator field.
func (o *IambackofficeTwoFactorMethod) SetAuthenticator(v IambackofficeAuthenticatorConfiguration) {
	o.Authenticator = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *IambackofficeTwoFactorMethod) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeTwoFactorMethod) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *IambackofficeTwoFactorMethod) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *IambackofficeTwoFactorMethod) SetEmail(v string) {
	o.Email = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IambackofficeTwoFactorMethod) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeTwoFactorMethod) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IambackofficeTwoFactorMethod) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IambackofficeTwoFactorMethod) SetId(v string) {
	o.Id = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *IambackofficeTwoFactorMethod) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeTwoFactorMethod) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *IambackofficeTwoFactorMethod) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *IambackofficeTwoFactorMethod) SetMethod(v string) {
	o.Method = &v
}

// GetMobilePhone returns the MobilePhone field value if set, zero value otherwise.
func (o *IambackofficeTwoFactorMethod) GetMobilePhone() string {
	if o == nil || IsNil(o.MobilePhone) {
		var ret string
		return ret
	}
	return *o.MobilePhone
}

// GetMobilePhoneOk returns a tuple with the MobilePhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeTwoFactorMethod) GetMobilePhoneOk() (*string, bool) {
	if o == nil || IsNil(o.MobilePhone) {
		return nil, false
	}
	return o.MobilePhone, true
}

// HasMobilePhone returns a boolean if a field has been set.
func (o *IambackofficeTwoFactorMethod) HasMobilePhone() bool {
	if o != nil && !IsNil(o.MobilePhone) {
		return true
	}

	return false
}

// SetMobilePhone gets a reference to the given string and assigns it to the MobilePhone field.
func (o *IambackofficeTwoFactorMethod) SetMobilePhone(v string) {
	o.MobilePhone = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *IambackofficeTwoFactorMethod) GetSecret() string {
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeTwoFactorMethod) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *IambackofficeTwoFactorMethod) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *IambackofficeTwoFactorMethod) SetSecret(v string) {
	o.Secret = &v
}

func (o IambackofficeTwoFactorMethod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IambackofficeTwoFactorMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Authenticator) {
		toSerialize["authenticator"] = o.Authenticator
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.MobilePhone) {
		toSerialize["mobilePhone"] = o.MobilePhone
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IambackofficeTwoFactorMethod) UnmarshalJSON(data []byte) (err error) {
	varIambackofficeTwoFactorMethod := _IambackofficeTwoFactorMethod{}

	err = json.Unmarshal(data, &varIambackofficeTwoFactorMethod)

	if err != nil {
		return err
	}

	*o = IambackofficeTwoFactorMethod(varIambackofficeTwoFactorMethod)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authenticator")
		delete(additionalProperties, "email")
		delete(additionalProperties, "id")
		delete(additionalProperties, "method")
		delete(additionalProperties, "mobilePhone")
		delete(additionalProperties, "secret")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *IambackofficeTwoFactorMethod) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populate the value of well-known types
func (o *IambackofficeTwoFactorMethod) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableIambackofficeTwoFactorMethod struct {
	value *IambackofficeTwoFactorMethod
	isSet bool
}

func (v NullableIambackofficeTwoFactorMethod) Get() *IambackofficeTwoFactorMethod {
	return v.value
}

func (v *NullableIambackofficeTwoFactorMethod) Set(val *IambackofficeTwoFactorMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableIambackofficeTwoFactorMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableIambackofficeTwoFactorMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIambackofficeTwoFactorMethod(val *IambackofficeTwoFactorMethod) *NullableIambackofficeTwoFactorMethod {
	return &NullableIambackofficeTwoFactorMethod{value: val, isSet: true}
}

func (v NullableIambackofficeTwoFactorMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIambackofficeTwoFactorMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
