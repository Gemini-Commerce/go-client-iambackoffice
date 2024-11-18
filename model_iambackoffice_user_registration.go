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

// checks if the IambackofficeUserRegistration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IambackofficeUserRegistration{}

// IambackofficeUserRegistration struct for IambackofficeUserRegistration
type IambackofficeUserRegistration struct {
    TenantId *string `json:"tenantId,omitempty"`
    Data *map[string]string `json:"data,omitempty"`
    PreferredLanguages []string `json:"preferredLanguages,omitempty"`
    Roles []string `json:"roles,omitempty"`
    Timezone *string `json:"timezone,omitempty"`
    AdditionalProperties map[string]interface{}
}

    type _IambackofficeUserRegistration IambackofficeUserRegistration

// NewIambackofficeUserRegistration instantiates a new IambackofficeUserRegistration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIambackofficeUserRegistration() *IambackofficeUserRegistration {
this := IambackofficeUserRegistration{}
return &this
}

// NewIambackofficeUserRegistrationWithDefaults instantiates a new IambackofficeUserRegistration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIambackofficeUserRegistrationWithDefaults() *IambackofficeUserRegistration {
this := IambackofficeUserRegistration{}
return &this
}

        // GetTenantId returns the TenantId field value if set, zero value otherwise.
        func (o *IambackofficeUserRegistration) GetTenantId() string {
        if o == nil || IsNil(o.TenantId) {
        var ret string
        return ret
        }
            return *o.TenantId
        }

        // GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
        // and a boolean to check if the value has been set.
        func (o *IambackofficeUserRegistration) GetTenantIdOk() (*string, bool) {
        if o == nil || IsNil(o.TenantId) {
            return nil, false
        }
            return o.TenantId, true
        }

        // HasTenantId returns a boolean if a field has been set.
        func (o *IambackofficeUserRegistration) HasTenantId() bool {
        if o != nil && !IsNil(o.TenantId) {
        return true
        }

        return false
        }

        // SetTenantId gets a reference to the given string and assigns it to the TenantId field.
        func (o *IambackofficeUserRegistration) SetTenantId(v string) {
            o.TenantId = &v
        }

        // GetData returns the Data field value if set, zero value otherwise.
        func (o *IambackofficeUserRegistration) GetData() map[string]string {
        if o == nil || IsNil(o.Data) {
        var ret map[string]string
        return ret
        }
            return *o.Data
        }

        // GetDataOk returns a tuple with the Data field value if set, nil otherwise
        // and a boolean to check if the value has been set.
        func (o *IambackofficeUserRegistration) GetDataOk() (*map[string]string, bool) {
        if o == nil || IsNil(o.Data) {
            return nil, false
        }
            return o.Data, true
        }

        // HasData returns a boolean if a field has been set.
        func (o *IambackofficeUserRegistration) HasData() bool {
        if o != nil && !IsNil(o.Data) {
        return true
        }

        return false
        }

        // SetData gets a reference to the given map[string]string and assigns it to the Data field.
        func (o *IambackofficeUserRegistration) SetData(v map[string]string) {
            o.Data = &v
        }

        // GetPreferredLanguages returns the PreferredLanguages field value if set, zero value otherwise.
        func (o *IambackofficeUserRegistration) GetPreferredLanguages() []string {
        if o == nil || IsNil(o.PreferredLanguages) {
        var ret []string
        return ret
        }
            return o.PreferredLanguages
        }

        // GetPreferredLanguagesOk returns a tuple with the PreferredLanguages field value if set, nil otherwise
        // and a boolean to check if the value has been set.
        func (o *IambackofficeUserRegistration) GetPreferredLanguagesOk() ([]string, bool) {
        if o == nil || IsNil(o.PreferredLanguages) {
            return nil, false
        }
            return o.PreferredLanguages, true
        }

        // HasPreferredLanguages returns a boolean if a field has been set.
        func (o *IambackofficeUserRegistration) HasPreferredLanguages() bool {
        if o != nil && !IsNil(o.PreferredLanguages) {
        return true
        }

        return false
        }

        // SetPreferredLanguages gets a reference to the given []string and assigns it to the PreferredLanguages field.
        func (o *IambackofficeUserRegistration) SetPreferredLanguages(v []string) {
            o.PreferredLanguages = v
        }

        // GetRoles returns the Roles field value if set, zero value otherwise.
        func (o *IambackofficeUserRegistration) GetRoles() []string {
        if o == nil || IsNil(o.Roles) {
        var ret []string
        return ret
        }
            return o.Roles
        }

        // GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
        // and a boolean to check if the value has been set.
        func (o *IambackofficeUserRegistration) GetRolesOk() ([]string, bool) {
        if o == nil || IsNil(o.Roles) {
            return nil, false
        }
            return o.Roles, true
        }

        // HasRoles returns a boolean if a field has been set.
        func (o *IambackofficeUserRegistration) HasRoles() bool {
        if o != nil && !IsNil(o.Roles) {
        return true
        }

        return false
        }

        // SetRoles gets a reference to the given []string and assigns it to the Roles field.
        func (o *IambackofficeUserRegistration) SetRoles(v []string) {
            o.Roles = v
        }

        // GetTimezone returns the Timezone field value if set, zero value otherwise.
        func (o *IambackofficeUserRegistration) GetTimezone() string {
        if o == nil || IsNil(o.Timezone) {
        var ret string
        return ret
        }
            return *o.Timezone
        }

        // GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
        // and a boolean to check if the value has been set.
        func (o *IambackofficeUserRegistration) GetTimezoneOk() (*string, bool) {
        if o == nil || IsNil(o.Timezone) {
            return nil, false
        }
            return o.Timezone, true
        }

        // HasTimezone returns a boolean if a field has been set.
        func (o *IambackofficeUserRegistration) HasTimezone() bool {
        if o != nil && !IsNil(o.Timezone) {
        return true
        }

        return false
        }

        // SetTimezone gets a reference to the given string and assigns it to the Timezone field.
        func (o *IambackofficeUserRegistration) SetTimezone(v string) {
            o.Timezone = &v
        }

    func (o IambackofficeUserRegistration) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
    return []byte{}, err
    }
    return json.Marshal(toSerialize)
    }

func (o IambackofficeUserRegistration) ToMap() (map[string]interface{}, error) {
toSerialize := map[string]interface{}{}
            if !IsNil(o.TenantId) {
            toSerialize["tenantId"] = o.TenantId
            }
            if !IsNil(o.Data) {
            toSerialize["data"] = o.Data
            }
            if !IsNil(o.PreferredLanguages) {
            toSerialize["preferredLanguages"] = o.PreferredLanguages
            }
            if !IsNil(o.Roles) {
            toSerialize["roles"] = o.Roles
            }
            if !IsNil(o.Timezone) {
            toSerialize["timezone"] = o.Timezone
            }

    for key, value := range o.AdditionalProperties {
    toSerialize[key] = value
    }

return toSerialize, nil
}

        func (o *IambackofficeUserRegistration) UnmarshalJSON(data []byte) (err error) {
            varIambackofficeUserRegistration := _IambackofficeUserRegistration{}

            err = json.Unmarshal(data, &varIambackofficeUserRegistration)

            if err != nil {
            return err
            }

            *o = IambackofficeUserRegistration(varIambackofficeUserRegistration)

            additionalProperties := make(map[string]interface{})

            if err = json.Unmarshal(data, &additionalProperties); err == nil {
                delete(additionalProperties, "tenantId")
                delete(additionalProperties, "data")
                delete(additionalProperties, "preferredLanguages")
                delete(additionalProperties, "roles")
                delete(additionalProperties, "timezone")
            o.AdditionalProperties = additionalProperties
            }

            return err
        }

    // GetValue returns the value of well-known types
    func (o *IambackofficeUserRegistration) GetValue() interface{} {
    if o == nil || IsNil(o.AdditionalProperties) {
    return nil
    }
    return o.AdditionalProperties["value"]
    }
    // SetValue populate the value of well-known types
    func (o *IambackofficeUserRegistration) SetValue(value interface{}) {
    if o == nil || IsNil(value) {
    return
    }
    if IsNil(o.AdditionalProperties) {
    o.AdditionalProperties = map[string]interface{}{}
    }
    o.AdditionalProperties["value"] = value
    return
    }
type NullableIambackofficeUserRegistration struct {
	value *IambackofficeUserRegistration
	isSet bool
}

func (v NullableIambackofficeUserRegistration) Get() *IambackofficeUserRegistration {
	return v.value
}

func (v *NullableIambackofficeUserRegistration) Set(val *IambackofficeUserRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableIambackofficeUserRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableIambackofficeUserRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIambackofficeUserRegistration(val *IambackofficeUserRegistration) *NullableIambackofficeUserRegistration {
	return &NullableIambackofficeUserRegistration{value: val, isSet: true}
}

func (v NullableIambackofficeUserRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIambackofficeUserRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

