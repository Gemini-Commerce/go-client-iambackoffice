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

// checks if the IambackofficeSearchGroupsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IambackofficeSearchGroupsRequest{}

// IambackofficeSearchGroupsRequest struct for IambackofficeSearchGroupsRequest
type IambackofficeSearchGroupsRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Filters *SearchGroupsRequestFilters `json:"filters,omitempty"`
	FiltersMask *string `json:"filtersMask,omitempty"`
}

// NewIambackofficeSearchGroupsRequest instantiates a new IambackofficeSearchGroupsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIambackofficeSearchGroupsRequest() *IambackofficeSearchGroupsRequest {
	this := IambackofficeSearchGroupsRequest{}
	return &this
}

// NewIambackofficeSearchGroupsRequestWithDefaults instantiates a new IambackofficeSearchGroupsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIambackofficeSearchGroupsRequestWithDefaults() *IambackofficeSearchGroupsRequest {
	this := IambackofficeSearchGroupsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *IambackofficeSearchGroupsRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeSearchGroupsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *IambackofficeSearchGroupsRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *IambackofficeSearchGroupsRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *IambackofficeSearchGroupsRequest) GetFilters() SearchGroupsRequestFilters {
	if o == nil || IsNil(o.Filters) {
		var ret SearchGroupsRequestFilters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeSearchGroupsRequest) GetFiltersOk() (*SearchGroupsRequestFilters, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *IambackofficeSearchGroupsRequest) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given SearchGroupsRequestFilters and assigns it to the Filters field.
func (o *IambackofficeSearchGroupsRequest) SetFilters(v SearchGroupsRequestFilters) {
	o.Filters = &v
}

// GetFiltersMask returns the FiltersMask field value if set, zero value otherwise.
func (o *IambackofficeSearchGroupsRequest) GetFiltersMask() string {
	if o == nil || IsNil(o.FiltersMask) {
		var ret string
		return ret
	}
	return *o.FiltersMask
}

// GetFiltersMaskOk returns a tuple with the FiltersMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeSearchGroupsRequest) GetFiltersMaskOk() (*string, bool) {
	if o == nil || IsNil(o.FiltersMask) {
		return nil, false
	}
	return o.FiltersMask, true
}

// HasFiltersMask returns a boolean if a field has been set.
func (o *IambackofficeSearchGroupsRequest) HasFiltersMask() bool {
	if o != nil && !IsNil(o.FiltersMask) {
		return true
	}

	return false
}

// SetFiltersMask gets a reference to the given string and assigns it to the FiltersMask field.
func (o *IambackofficeSearchGroupsRequest) SetFiltersMask(v string) {
	o.FiltersMask = &v
}

func (o IambackofficeSearchGroupsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IambackofficeSearchGroupsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.FiltersMask) {
		toSerialize["filtersMask"] = o.FiltersMask
	}
	return toSerialize, nil
}

type NullableIambackofficeSearchGroupsRequest struct {
	value *IambackofficeSearchGroupsRequest
	isSet bool
}

func (v NullableIambackofficeSearchGroupsRequest) Get() *IambackofficeSearchGroupsRequest {
	return v.value
}

func (v *NullableIambackofficeSearchGroupsRequest) Set(val *IambackofficeSearchGroupsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIambackofficeSearchGroupsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIambackofficeSearchGroupsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIambackofficeSearchGroupsRequest(val *IambackofficeSearchGroupsRequest) *NullableIambackofficeSearchGroupsRequest {
	return &NullableIambackofficeSearchGroupsRequest{value: val, isSet: true}
}

func (v NullableIambackofficeSearchGroupsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIambackofficeSearchGroupsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


