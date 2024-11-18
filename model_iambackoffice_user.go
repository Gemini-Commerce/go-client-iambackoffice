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

// checks if the IambackofficeUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IambackofficeUser{}

// IambackofficeUser struct for IambackofficeUser
type IambackofficeUser struct {
	Active                    *bool                                    `json:"active,omitempty"`
	BirthDate                 *string                                  `json:"birthDate,omitempty"`
	Data                      *map[string]string                       `json:"data,omitempty"`
	Email                     *string                                  `json:"email,omitempty"`
	FirstName                 *string                                  `json:"firstName,omitempty"`
	FullName                  *string                                  `json:"fullName,omitempty"`
	Id                        *string                                  `json:"id,omitempty"`
	ImageUrl                  *string                                  `json:"imageUrl,omitempty"`
	InsertInstant             *string                                  `json:"insertInstant,omitempty"`
	LastLoginInstant          *string                                  `json:"lastLoginInstant,omitempty"`
	LastName                  *string                                  `json:"lastName,omitempty"`
	LastUpdateInstant         *string                                  `json:"lastUpdateInstant,omitempty"`
	MiddleName                *string                                  `json:"middleName,omitempty"`
	MobilePhone               *string                                  `json:"mobilePhone,omitempty"`
	Password                  *string                                  `json:"password,omitempty"`
	PasswordChangeRequired    *bool                                    `json:"passwordChangeRequired,omitempty"`
	PasswordLastUpdateInstant *string                                  `json:"passwordLastUpdateInstant,omitempty"`
	PreferredLanguages        []string                                 `json:"preferredLanguages,omitempty"`
	Registrations             []IambackofficeUserRegistration          `json:"registrations,omitempty"`
	Timezone                  *string                                  `json:"timezone,omitempty"`
	TwoFactor                 *IambackofficeUserTwoFactorConfiguration `json:"twoFactor,omitempty"`
	Username                  *string                                  `json:"username,omitempty"`
	Verified                  *bool                                    `json:"verified,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _IambackofficeUser IambackofficeUser

// NewIambackofficeUser instantiates a new IambackofficeUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIambackofficeUser() *IambackofficeUser {
	this := IambackofficeUser{}
	return &this
}

// NewIambackofficeUserWithDefaults instantiates a new IambackofficeUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIambackofficeUserWithDefaults() *IambackofficeUser {
	this := IambackofficeUser{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *IambackofficeUser) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeUser) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *IambackofficeUser) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *IambackofficeUser) SetActive(v bool) {
	o.Active = &v
}

// GetBirthDate returns the BirthDate field value if set, zero value otherwise.
func (o *IambackofficeUser) GetBirthDate() string {
	if o == nil || IsNil(o.BirthDate) {
		var ret string
		return ret
	}
	return *o.BirthDate
}

// GetBirthDateOk returns a tuple with the BirthDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeUser) GetBirthDateOk() (*string, bool) {
	if o == nil || IsNil(o.BirthDate) {
		return nil, false
	}
	return o.BirthDate, true
}

// HasBirthDate returns a boolean if a field has been set.
func (o *IambackofficeUser) HasBirthDate() bool {
	if o != nil && !IsNil(o.BirthDate) {
		return true
	}

	return false
}

// SetBirthDate gets a reference to the given string and assigns it to the BirthDate field.
func (o *IambackofficeUser) SetBirthDate(v string) {
	o.BirthDate = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *IambackofficeUser) GetData() map[string]string {
	if o == nil || IsNil(o.Data) {
		var ret map[string]string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeUser) GetDataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *IambackofficeUser) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]string and assigns it to the Data field.
func (o *IambackofficeUser) SetData(v map[string]string) {
	o.Data = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *IambackofficeUser) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeUser) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *IambackofficeUser) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *IambackofficeUser) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *IambackofficeUser) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeUser) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *IambackofficeUser) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *IambackofficeUser) SetFirstName(v string) {
	o.FirstName = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *IambackofficeUser) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeUser) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *IambackofficeUser) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *IambackofficeUser) SetFullName(v string) {
	o.FullName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IambackofficeUser) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeUser) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IambackofficeUser) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IambackofficeUser) SetId(v string) {
	o.Id = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *IambackofficeUser) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeUser) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *IambackofficeUser) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *IambackofficeUser) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetInsertInstant returns the InsertInstant field value if set, zero value otherwise.
func (o *IambackofficeUser) GetInsertInstant() string {
	if o == nil || IsNil(o.InsertInstant) {
		var ret string
		return ret
	}
	return *o.InsertInstant
}

// GetInsertInstantOk returns a tuple with the InsertInstant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeUser) GetInsertInstantOk() (*string, bool) {
	if o == nil || IsNil(o.InsertInstant) {
		return nil, false
	}
	return o.InsertInstant, true
}

// HasInsertInstant returns a boolean if a field has been set.
func (o *IambackofficeUser) HasInsertInstant() bool {
	if o != nil && !IsNil(o.InsertInstant) {
		return true
	}

	return false
}

// SetInsertInstant gets a reference to the given string and assigns it to the InsertInstant field.
func (o *IambackofficeUser) SetInsertInstant(v string) {
	o.InsertInstant = &v
}

// GetLastLoginInstant returns the LastLoginInstant field value if set, zero value otherwise.
func (o *IambackofficeUser) GetLastLoginInstant() string {
	if o == nil || IsNil(o.LastLoginInstant) {
		var ret string
		return ret
	}
	return *o.LastLoginInstant
}

// GetLastLoginInstantOk returns a tuple with the LastLoginInstant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeUser) GetLastLoginInstantOk() (*string, bool) {
	if o == nil || IsNil(o.LastLoginInstant) {
		return nil, false
	}
	return o.LastLoginInstant, true
}

// HasLastLoginInstant returns a boolean if a field has been set.
func (o *IambackofficeUser) HasLastLoginInstant() bool {
	if o != nil && !IsNil(o.LastLoginInstant) {
		return true
	}

	return false
}

// SetLastLoginInstant gets a reference to the given string and assigns it to the LastLoginInstant field.
func (o *IambackofficeUser) SetLastLoginInstant(v string) {
	o.LastLoginInstant = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *IambackofficeUser) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeUser) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *IambackofficeUser) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *IambackofficeUser) SetLastName(v string) {
	o.LastName = &v
}

// GetLastUpdateInstant returns the LastUpdateInstant field value if set, zero value otherwise.
func (o *IambackofficeUser) GetLastUpdateInstant() string {
	if o == nil || IsNil(o.LastUpdateInstant) {
		var ret string
		return ret
	}
	return *o.LastUpdateInstant
}

// GetLastUpdateInstantOk returns a tuple with the LastUpdateInstant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeUser) GetLastUpdateInstantOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdateInstant) {
		return nil, false
	}
	return o.LastUpdateInstant, true
}

// HasLastUpdateInstant returns a boolean if a field has been set.
func (o *IambackofficeUser) HasLastUpdateInstant() bool {
	if o != nil && !IsNil(o.LastUpdateInstant) {
		return true
	}

	return false
}

// SetLastUpdateInstant gets a reference to the given string and assigns it to the LastUpdateInstant field.
func (o *IambackofficeUser) SetLastUpdateInstant(v string) {
	o.LastUpdateInstant = &v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *IambackofficeUser) GetMiddleName() string {
	if o == nil || IsNil(o.MiddleName) {
		var ret string
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeUser) GetMiddleNameOk() (*string, bool) {
	if o == nil || IsNil(o.MiddleName) {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *IambackofficeUser) HasMiddleName() bool {
	if o != nil && !IsNil(o.MiddleName) {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.
func (o *IambackofficeUser) SetMiddleName(v string) {
	o.MiddleName = &v
}

// GetMobilePhone returns the MobilePhone field value if set, zero value otherwise.
func (o *IambackofficeUser) GetMobilePhone() string {
	if o == nil || IsNil(o.MobilePhone) {
		var ret string
		return ret
	}
	return *o.MobilePhone
}

// GetMobilePhoneOk returns a tuple with the MobilePhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeUser) GetMobilePhoneOk() (*string, bool) {
	if o == nil || IsNil(o.MobilePhone) {
		return nil, false
	}
	return o.MobilePhone, true
}

// HasMobilePhone returns a boolean if a field has been set.
func (o *IambackofficeUser) HasMobilePhone() bool {
	if o != nil && !IsNil(o.MobilePhone) {
		return true
	}

	return false
}

// SetMobilePhone gets a reference to the given string and assigns it to the MobilePhone field.
func (o *IambackofficeUser) SetMobilePhone(v string) {
	o.MobilePhone = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *IambackofficeUser) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeUser) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *IambackofficeUser) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *IambackofficeUser) SetPassword(v string) {
	o.Password = &v
}

// GetPasswordChangeRequired returns the PasswordChangeRequired field value if set, zero value otherwise.
func (o *IambackofficeUser) GetPasswordChangeRequired() bool {
	if o == nil || IsNil(o.PasswordChangeRequired) {
		var ret bool
		return ret
	}
	return *o.PasswordChangeRequired
}

// GetPasswordChangeRequiredOk returns a tuple with the PasswordChangeRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeUser) GetPasswordChangeRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.PasswordChangeRequired) {
		return nil, false
	}
	return o.PasswordChangeRequired, true
}

// HasPasswordChangeRequired returns a boolean if a field has been set.
func (o *IambackofficeUser) HasPasswordChangeRequired() bool {
	if o != nil && !IsNil(o.PasswordChangeRequired) {
		return true
	}

	return false
}

// SetPasswordChangeRequired gets a reference to the given bool and assigns it to the PasswordChangeRequired field.
func (o *IambackofficeUser) SetPasswordChangeRequired(v bool) {
	o.PasswordChangeRequired = &v
}

// GetPasswordLastUpdateInstant returns the PasswordLastUpdateInstant field value if set, zero value otherwise.
func (o *IambackofficeUser) GetPasswordLastUpdateInstant() string {
	if o == nil || IsNil(o.PasswordLastUpdateInstant) {
		var ret string
		return ret
	}
	return *o.PasswordLastUpdateInstant
}

// GetPasswordLastUpdateInstantOk returns a tuple with the PasswordLastUpdateInstant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeUser) GetPasswordLastUpdateInstantOk() (*string, bool) {
	if o == nil || IsNil(o.PasswordLastUpdateInstant) {
		return nil, false
	}
	return o.PasswordLastUpdateInstant, true
}

// HasPasswordLastUpdateInstant returns a boolean if a field has been set.
func (o *IambackofficeUser) HasPasswordLastUpdateInstant() bool {
	if o != nil && !IsNil(o.PasswordLastUpdateInstant) {
		return true
	}

	return false
}

// SetPasswordLastUpdateInstant gets a reference to the given string and assigns it to the PasswordLastUpdateInstant field.
func (o *IambackofficeUser) SetPasswordLastUpdateInstant(v string) {
	o.PasswordLastUpdateInstant = &v
}

// GetPreferredLanguages returns the PreferredLanguages field value if set, zero value otherwise.
func (o *IambackofficeUser) GetPreferredLanguages() []string {
	if o == nil || IsNil(o.PreferredLanguages) {
		var ret []string
		return ret
	}
	return o.PreferredLanguages
}

// GetPreferredLanguagesOk returns a tuple with the PreferredLanguages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeUser) GetPreferredLanguagesOk() ([]string, bool) {
	if o == nil || IsNil(o.PreferredLanguages) {
		return nil, false
	}
	return o.PreferredLanguages, true
}

// HasPreferredLanguages returns a boolean if a field has been set.
func (o *IambackofficeUser) HasPreferredLanguages() bool {
	if o != nil && !IsNil(o.PreferredLanguages) {
		return true
	}

	return false
}

// SetPreferredLanguages gets a reference to the given []string and assigns it to the PreferredLanguages field.
func (o *IambackofficeUser) SetPreferredLanguages(v []string) {
	o.PreferredLanguages = v
}

// GetRegistrations returns the Registrations field value if set, zero value otherwise.
func (o *IambackofficeUser) GetRegistrations() []IambackofficeUserRegistration {
	if o == nil || IsNil(o.Registrations) {
		var ret []IambackofficeUserRegistration
		return ret
	}
	return o.Registrations
}

// GetRegistrationsOk returns a tuple with the Registrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeUser) GetRegistrationsOk() ([]IambackofficeUserRegistration, bool) {
	if o == nil || IsNil(o.Registrations) {
		return nil, false
	}
	return o.Registrations, true
}

// HasRegistrations returns a boolean if a field has been set.
func (o *IambackofficeUser) HasRegistrations() bool {
	if o != nil && !IsNil(o.Registrations) {
		return true
	}

	return false
}

// SetRegistrations gets a reference to the given []IambackofficeUserRegistration and assigns it to the Registrations field.
func (o *IambackofficeUser) SetRegistrations(v []IambackofficeUserRegistration) {
	o.Registrations = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *IambackofficeUser) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeUser) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *IambackofficeUser) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *IambackofficeUser) SetTimezone(v string) {
	o.Timezone = &v
}

// GetTwoFactor returns the TwoFactor field value if set, zero value otherwise.
func (o *IambackofficeUser) GetTwoFactor() IambackofficeUserTwoFactorConfiguration {
	if o == nil || IsNil(o.TwoFactor) {
		var ret IambackofficeUserTwoFactorConfiguration
		return ret
	}
	return *o.TwoFactor
}

// GetTwoFactorOk returns a tuple with the TwoFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeUser) GetTwoFactorOk() (*IambackofficeUserTwoFactorConfiguration, bool) {
	if o == nil || IsNil(o.TwoFactor) {
		return nil, false
	}
	return o.TwoFactor, true
}

// HasTwoFactor returns a boolean if a field has been set.
func (o *IambackofficeUser) HasTwoFactor() bool {
	if o != nil && !IsNil(o.TwoFactor) {
		return true
	}

	return false
}

// SetTwoFactor gets a reference to the given IambackofficeUserTwoFactorConfiguration and assigns it to the TwoFactor field.
func (o *IambackofficeUser) SetTwoFactor(v IambackofficeUserTwoFactorConfiguration) {
	o.TwoFactor = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *IambackofficeUser) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeUser) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *IambackofficeUser) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *IambackofficeUser) SetUsername(v string) {
	o.Username = &v
}

// GetVerified returns the Verified field value if set, zero value otherwise.
func (o *IambackofficeUser) GetVerified() bool {
	if o == nil || IsNil(o.Verified) {
		var ret bool
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IambackofficeUser) GetVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.Verified) {
		return nil, false
	}
	return o.Verified, true
}

// HasVerified returns a boolean if a field has been set.
func (o *IambackofficeUser) HasVerified() bool {
	if o != nil && !IsNil(o.Verified) {
		return true
	}

	return false
}

// SetVerified gets a reference to the given bool and assigns it to the Verified field.
func (o *IambackofficeUser) SetVerified(v bool) {
	o.Verified = &v
}

func (o IambackofficeUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IambackofficeUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.BirthDate) {
		toSerialize["birthDate"] = o.BirthDate
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.FullName) {
		toSerialize["fullName"] = o.FullName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	if !IsNil(o.InsertInstant) {
		toSerialize["insertInstant"] = o.InsertInstant
	}
	if !IsNil(o.LastLoginInstant) {
		toSerialize["lastLoginInstant"] = o.LastLoginInstant
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.LastUpdateInstant) {
		toSerialize["lastUpdateInstant"] = o.LastUpdateInstant
	}
	if !IsNil(o.MiddleName) {
		toSerialize["middleName"] = o.MiddleName
	}
	if !IsNil(o.MobilePhone) {
		toSerialize["mobilePhone"] = o.MobilePhone
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.PasswordChangeRequired) {
		toSerialize["passwordChangeRequired"] = o.PasswordChangeRequired
	}
	if !IsNil(o.PasswordLastUpdateInstant) {
		toSerialize["passwordLastUpdateInstant"] = o.PasswordLastUpdateInstant
	}
	if !IsNil(o.PreferredLanguages) {
		toSerialize["preferredLanguages"] = o.PreferredLanguages
	}
	if !IsNil(o.Registrations) {
		toSerialize["registrations"] = o.Registrations
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !IsNil(o.TwoFactor) {
		toSerialize["twoFactor"] = o.TwoFactor
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Verified) {
		toSerialize["verified"] = o.Verified
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IambackofficeUser) UnmarshalJSON(data []byte) (err error) {
	varIambackofficeUser := _IambackofficeUser{}

	err = json.Unmarshal(data, &varIambackofficeUser)

	if err != nil {
		return err
	}

	*o = IambackofficeUser(varIambackofficeUser)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "active")
		delete(additionalProperties, "birthDate")
		delete(additionalProperties, "data")
		delete(additionalProperties, "email")
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "fullName")
		delete(additionalProperties, "id")
		delete(additionalProperties, "imageUrl")
		delete(additionalProperties, "insertInstant")
		delete(additionalProperties, "lastLoginInstant")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "lastUpdateInstant")
		delete(additionalProperties, "middleName")
		delete(additionalProperties, "mobilePhone")
		delete(additionalProperties, "password")
		delete(additionalProperties, "passwordChangeRequired")
		delete(additionalProperties, "passwordLastUpdateInstant")
		delete(additionalProperties, "preferredLanguages")
		delete(additionalProperties, "registrations")
		delete(additionalProperties, "timezone")
		delete(additionalProperties, "twoFactor")
		delete(additionalProperties, "username")
		delete(additionalProperties, "verified")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIambackofficeUser struct {
	value *IambackofficeUser
	isSet bool
}

func (v NullableIambackofficeUser) Get() *IambackofficeUser {
	return v.value
}

func (v *NullableIambackofficeUser) Set(val *IambackofficeUser) {
	v.value = val
	v.isSet = true
}

func (v NullableIambackofficeUser) IsSet() bool {
	return v.isSet
}

func (v *NullableIambackofficeUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIambackofficeUser(val *IambackofficeUser) *NullableIambackofficeUser {
	return &NullableIambackofficeUser{value: val, isSet: true}
}

func (v NullableIambackofficeUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIambackofficeUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
