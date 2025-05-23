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

// LinkProfileEligibilityCheckResponse LinkProfileEligibilityCheckResponse defines the response schema for `/link/profile/eligibility/check`
type LinkProfileEligibilityCheckResponse struct {
	// Indicates whether Plaid has a profile matching the customer's eligibility requirements for this user
	ProfileMatches bool `json:"profile_matches"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _LinkProfileEligibilityCheckResponse LinkProfileEligibilityCheckResponse

// NewLinkProfileEligibilityCheckResponse instantiates a new LinkProfileEligibilityCheckResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkProfileEligibilityCheckResponse(profileMatches bool, requestId string) *LinkProfileEligibilityCheckResponse {
	this := LinkProfileEligibilityCheckResponse{}
	this.ProfileMatches = profileMatches
	this.RequestId = requestId
	return &this
}

// NewLinkProfileEligibilityCheckResponseWithDefaults instantiates a new LinkProfileEligibilityCheckResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkProfileEligibilityCheckResponseWithDefaults() *LinkProfileEligibilityCheckResponse {
	this := LinkProfileEligibilityCheckResponse{}
	return &this
}

// GetProfileMatches returns the ProfileMatches field value
func (o *LinkProfileEligibilityCheckResponse) GetProfileMatches() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ProfileMatches
}

// GetProfileMatchesOk returns a tuple with the ProfileMatches field value
// and a boolean to check if the value has been set.
func (o *LinkProfileEligibilityCheckResponse) GetProfileMatchesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProfileMatches, true
}

// SetProfileMatches sets field value
func (o *LinkProfileEligibilityCheckResponse) SetProfileMatches(v bool) {
	o.ProfileMatches = v
}

// GetRequestId returns the RequestId field value
func (o *LinkProfileEligibilityCheckResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *LinkProfileEligibilityCheckResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *LinkProfileEligibilityCheckResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o LinkProfileEligibilityCheckResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["profile_matches"] = o.ProfileMatches
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinkProfileEligibilityCheckResponse) UnmarshalJSON(bytes []byte) (err error) {
	varLinkProfileEligibilityCheckResponse := _LinkProfileEligibilityCheckResponse{}

	if err = json.Unmarshal(bytes, &varLinkProfileEligibilityCheckResponse); err == nil {
		*o = LinkProfileEligibilityCheckResponse(varLinkProfileEligibilityCheckResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "profile_matches")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkProfileEligibilityCheckResponse struct {
	value *LinkProfileEligibilityCheckResponse
	isSet bool
}

func (v NullableLinkProfileEligibilityCheckResponse) Get() *LinkProfileEligibilityCheckResponse {
	return v.value
}

func (v *NullableLinkProfileEligibilityCheckResponse) Set(val *LinkProfileEligibilityCheckResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkProfileEligibilityCheckResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkProfileEligibilityCheckResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkProfileEligibilityCheckResponse(val *LinkProfileEligibilityCheckResponse) *NullableLinkProfileEligibilityCheckResponse {
	return &NullableLinkProfileEligibilityCheckResponse{value: val, isSet: true}
}

func (v NullableLinkProfileEligibilityCheckResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkProfileEligibilityCheckResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


