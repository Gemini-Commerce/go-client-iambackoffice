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

// checks if the SearchGroupsRequestFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchGroupsRequestFilters{}

// SearchGroupsRequestFilters struct for SearchGroupsRequestFilters
type SearchGroupsRequestFilters struct {
	Data *map[string]string `json:"data,omitempty"`
}

// NewSearchGroupsRequestFilters instantiates a new SearchGroupsRequestFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchGroupsRequestFilters() *SearchGroupsRequestFilters {
	this := SearchGroupsRequestFilters{}
	return &this
}

// NewSearchGroupsRequestFiltersWithDefaults instantiates a new SearchGroupsRequestFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchGroupsRequestFiltersWithDefaults() *SearchGroupsRequestFilters {
	this := SearchGroupsRequestFilters{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SearchGroupsRequestFilters) GetData() map[string]string {
	if o == nil || IsNil(o.Data) {
		var ret map[string]string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchGroupsRequestFilters) GetDataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SearchGroupsRequestFilters) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]string and assigns it to the Data field.
func (o *SearchGroupsRequestFilters) SetData(v map[string]string) {
	o.Data = &v
}

func (o SearchGroupsRequestFilters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchGroupsRequestFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableSearchGroupsRequestFilters struct {
	value *SearchGroupsRequestFilters
	isSet bool
}

func (v NullableSearchGroupsRequestFilters) Get() *SearchGroupsRequestFilters {
	return v.value
}

func (v *NullableSearchGroupsRequestFilters) Set(val *SearchGroupsRequestFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchGroupsRequestFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchGroupsRequestFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchGroupsRequestFilters(val *SearchGroupsRequestFilters) *NullableSearchGroupsRequestFilters {
	return &NullableSearchGroupsRequestFilters{value: val, isSet: true}
}

func (v NullableSearchGroupsRequestFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchGroupsRequestFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

