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

// DepositSwitchTokenCreateRequest (Deprecated) DepositSwitchTokenCreateRequest defines the request schema for `/deposit_switch/token/create`
type DepositSwitchTokenCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The ID of the deposit switch
	DepositSwitchId string `json:"deposit_switch_id"`
}

// NewDepositSwitchTokenCreateRequest instantiates a new DepositSwitchTokenCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositSwitchTokenCreateRequest(depositSwitchId string) *DepositSwitchTokenCreateRequest {
	this := DepositSwitchTokenCreateRequest{}
	this.DepositSwitchId = depositSwitchId
	return &this
}

// NewDepositSwitchTokenCreateRequestWithDefaults instantiates a new DepositSwitchTokenCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositSwitchTokenCreateRequestWithDefaults() *DepositSwitchTokenCreateRequest {
	this := DepositSwitchTokenCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *DepositSwitchTokenCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositSwitchTokenCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *DepositSwitchTokenCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *DepositSwitchTokenCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *DepositSwitchTokenCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositSwitchTokenCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *DepositSwitchTokenCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *DepositSwitchTokenCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetDepositSwitchId returns the DepositSwitchId field value
func (o *DepositSwitchTokenCreateRequest) GetDepositSwitchId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DepositSwitchId
}

// GetDepositSwitchIdOk returns a tuple with the DepositSwitchId field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchTokenCreateRequest) GetDepositSwitchIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DepositSwitchId, true
}

// SetDepositSwitchId sets field value
func (o *DepositSwitchTokenCreateRequest) SetDepositSwitchId(v string) {
	o.DepositSwitchId = v
}

func (o DepositSwitchTokenCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["deposit_switch_id"] = o.DepositSwitchId
	}
	return json.Marshal(toSerialize)
}

type NullableDepositSwitchTokenCreateRequest struct {
	value *DepositSwitchTokenCreateRequest
	isSet bool
}

func (v NullableDepositSwitchTokenCreateRequest) Get() *DepositSwitchTokenCreateRequest {
	return v.value
}

func (v *NullableDepositSwitchTokenCreateRequest) Set(val *DepositSwitchTokenCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositSwitchTokenCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositSwitchTokenCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositSwitchTokenCreateRequest(val *DepositSwitchTokenCreateRequest) *NullableDepositSwitchTokenCreateRequest {
	return &NullableDepositSwitchTokenCreateRequest{value: val, isSet: true}
}

func (v NullableDepositSwitchTokenCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositSwitchTokenCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


