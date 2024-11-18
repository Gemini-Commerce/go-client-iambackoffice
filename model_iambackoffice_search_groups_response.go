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

// checks if the IambackofficeSearchGroupsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IambackofficeSearchGroupsResponse{}

// IambackofficeSearchGroupsResponse struct for IambackofficeSearchGroupsResponse
type IambackofficeSearchGroupsResponse struct {
    Groups []IambackofficeGroup `json:"groups,omitempty"`
    AdditionalProperties map[string]interface{}
}

    type _IambackofficeSearchGroupsResponse IambackofficeSearchGroupsResponse

// NewIambackofficeSearchGroupsResponse instantiates a new IambackofficeSearchGroupsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIambackofficeSearchGroupsResponse() *IambackofficeSearchGroupsResponse {
this := IambackofficeSearchGroupsResponse{}
return &this
}

// NewIambackofficeSearchGroupsResponseWithDefaults instantiates a new IambackofficeSearchGroupsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIambackofficeSearchGroupsResponseWithDefaults() *IambackofficeSearchGroupsResponse {
this := IambackofficeSearchGroupsResponse{}
return &this
}

        // GetGroups returns the Groups field value if set, zero value otherwise.
        func (o *IambackofficeSearchGroupsResponse) GetGroups() []IambackofficeGroup {
        if o == nil || IsNil(o.Groups) {
        var ret []IambackofficeGroup
        return ret
        }
            return o.Groups
        }

        // GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
        // and a boolean to check if the value has been set.
        func (o *IambackofficeSearchGroupsResponse) GetGroupsOk() ([]IambackofficeGroup, bool) {
        if o == nil || IsNil(o.Groups) {
            return nil, false
        }
            return o.Groups, true
        }

        // HasGroups returns a boolean if a field has been set.
        func (o *IambackofficeSearchGroupsResponse) HasGroups() bool {
        if o != nil && !IsNil(o.Groups) {
        return true
        }

        return false
        }

        // SetGroups gets a reference to the given []IambackofficeGroup and assigns it to the Groups field.
        func (o *IambackofficeSearchGroupsResponse) SetGroups(v []IambackofficeGroup) {
            o.Groups = v
        }

    func (o IambackofficeSearchGroupsResponse) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
    return []byte{}, err
    }
    return json.Marshal(toSerialize)
    }

func (o IambackofficeSearchGroupsResponse) ToMap() (map[string]interface{}, error) {
toSerialize := map[string]interface{}{}
            if !IsNil(o.Groups) {
            toSerialize["groups"] = o.Groups
            }

    for key, value := range o.AdditionalProperties {
    toSerialize[key] = value
    }

return toSerialize, nil
}

        func (o *IambackofficeSearchGroupsResponse) UnmarshalJSON(data []byte) (err error) {
            varIambackofficeSearchGroupsResponse := _IambackofficeSearchGroupsResponse{}

            err = json.Unmarshal(data, &varIambackofficeSearchGroupsResponse)

            if err != nil {
            return err
            }

            *o = IambackofficeSearchGroupsResponse(varIambackofficeSearchGroupsResponse)

            additionalProperties := make(map[string]interface{})

            if err = json.Unmarshal(data, &additionalProperties); err == nil {
                delete(additionalProperties, "groups")
            o.AdditionalProperties = additionalProperties
            }

            return err
        }

    // GetValue returns the value of well-known types
    func (o *IambackofficeSearchGroupsResponse) GetValue() interface{} {
    if o == nil || IsNil(o.AdditionalProperties) {
    return nil
    }
    return o.AdditionalProperties["value"]
    }
    // SetValue populate the value of well-known types
    func (o *IambackofficeSearchGroupsResponse) SetValue(value interface{}) {
    if o == nil || IsNil(value) {
    return
    }
    if IsNil(o.AdditionalProperties) {
    o.AdditionalProperties = map[string]interface{}{}
    }
    o.AdditionalProperties["value"] = value
    return
    }
type NullableIambackofficeSearchGroupsResponse struct {
	value *IambackofficeSearchGroupsResponse
	isSet bool
}

func (v NullableIambackofficeSearchGroupsResponse) Get() *IambackofficeSearchGroupsResponse {
	return v.value
}

func (v *NullableIambackofficeSearchGroupsResponse) Set(val *IambackofficeSearchGroupsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIambackofficeSearchGroupsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIambackofficeSearchGroupsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIambackofficeSearchGroupsResponse(val *IambackofficeSearchGroupsResponse) *NullableIambackofficeSearchGroupsResponse {
	return &NullableIambackofficeSearchGroupsResponse{value: val, isSet: true}
}

func (v NullableIambackofficeSearchGroupsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIambackofficeSearchGroupsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

