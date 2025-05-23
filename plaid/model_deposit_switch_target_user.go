/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.586.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// DepositSwitchTargetUser (Deprecated) The deposit switch target user
type DepositSwitchTargetUser struct {
	// The given name (first name) of the user.
	GivenName string `json:"given_name"`
	// The family name (last name) of the user.
	FamilyName string `json:"family_name"`
	// The phone number of the user. The endpoint can accept a variety of phone number formats, including E.164.
	Phone string `json:"phone"`
	// The email address of the user.
	Email string `json:"email"`
	Address *DepositSwitchAddressData `json:"address,omitempty"`
	// The taxpayer ID of the user, generally their SSN, EIN, or TIN.
	TaxPayerId *string `json:"tax_payer_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DepositSwitchTargetUser DepositSwitchTargetUser

// NewDepositSwitchTargetUser instantiates a new DepositSwitchTargetUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositSwitchTargetUser(givenName string, familyName string, phone string, email string) *DepositSwitchTargetUser {
	this := DepositSwitchTargetUser{}
	this.GivenName = givenName
	this.FamilyName = familyName
	this.Phone = phone
	this.Email = email
	return &this
}

// NewDepositSwitchTargetUserWithDefaults instantiates a new DepositSwitchTargetUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositSwitchTargetUserWithDefaults() *DepositSwitchTargetUser {
	this := DepositSwitchTargetUser{}
	return &this
}

// GetGivenName returns the GivenName field value
func (o *DepositSwitchTargetUser) GetGivenName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchTargetUser) GetGivenNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GivenName, true
}

// SetGivenName sets field value
func (o *DepositSwitchTargetUser) SetGivenName(v string) {
	o.GivenName = v
}

// GetFamilyName returns the FamilyName field value
func (o *DepositSwitchTargetUser) GetFamilyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FamilyName
}

// GetFamilyNameOk returns a tuple with the FamilyName field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchTargetUser) GetFamilyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FamilyName, true
}

// SetFamilyName sets field value
func (o *DepositSwitchTargetUser) SetFamilyName(v string) {
	o.FamilyName = v
}

// GetPhone returns the Phone field value
func (o *DepositSwitchTargetUser) GetPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchTargetUser) GetPhoneOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *DepositSwitchTargetUser) SetPhone(v string) {
	o.Phone = v
}

// GetEmail returns the Email field value
func (o *DepositSwitchTargetUser) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchTargetUser) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *DepositSwitchTargetUser) SetEmail(v string) {
	o.Email = v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *DepositSwitchTargetUser) GetAddress() DepositSwitchAddressData {
	if o == nil || o.Address == nil {
		var ret DepositSwitchAddressData
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositSwitchTargetUser) GetAddressOk() (*DepositSwitchAddressData, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *DepositSwitchTargetUser) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given DepositSwitchAddressData and assigns it to the Address field.
func (o *DepositSwitchTargetUser) SetAddress(v DepositSwitchAddressData) {
	o.Address = &v
}

// GetTaxPayerId returns the TaxPayerId field value if set, zero value otherwise.
func (o *DepositSwitchTargetUser) GetTaxPayerId() string {
	if o == nil || o.TaxPayerId == nil {
		var ret string
		return ret
	}
	return *o.TaxPayerId
}

// GetTaxPayerIdOk returns a tuple with the TaxPayerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositSwitchTargetUser) GetTaxPayerIdOk() (*string, bool) {
	if o == nil || o.TaxPayerId == nil {
		return nil, false
	}
	return o.TaxPayerId, true
}

// HasTaxPayerId returns a boolean if a field has been set.
func (o *DepositSwitchTargetUser) HasTaxPayerId() bool {
	if o != nil && o.TaxPayerId != nil {
		return true
	}

	return false
}

// SetTaxPayerId gets a reference to the given string and assigns it to the TaxPayerId field.
func (o *DepositSwitchTargetUser) SetTaxPayerId(v string) {
	o.TaxPayerId = &v
}

func (o DepositSwitchTargetUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["given_name"] = o.GivenName
	}
	if true {
		toSerialize["family_name"] = o.FamilyName
	}
	if true {
		toSerialize["phone"] = o.Phone
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.TaxPayerId != nil {
		toSerialize["tax_payer_id"] = o.TaxPayerId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DepositSwitchTargetUser) UnmarshalJSON(bytes []byte) (err error) {
	varDepositSwitchTargetUser := _DepositSwitchTargetUser{}

	if err = json.Unmarshal(bytes, &varDepositSwitchTargetUser); err == nil {
		*o = DepositSwitchTargetUser(varDepositSwitchTargetUser)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "given_name")
		delete(additionalProperties, "family_name")
		delete(additionalProperties, "phone")
		delete(additionalProperties, "email")
		delete(additionalProperties, "address")
		delete(additionalProperties, "tax_payer_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDepositSwitchTargetUser struct {
	value *DepositSwitchTargetUser
	isSet bool
}

func (v NullableDepositSwitchTargetUser) Get() *DepositSwitchTargetUser {
	return v.value
}

func (v *NullableDepositSwitchTargetUser) Set(val *DepositSwitchTargetUser) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositSwitchTargetUser) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositSwitchTargetUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositSwitchTargetUser(val *DepositSwitchTargetUser) *NullableDepositSwitchTargetUser {
	return &NullableDepositSwitchTargetUser{value: val, isSet: true}
}

func (v NullableDepositSwitchTargetUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositSwitchTargetUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


