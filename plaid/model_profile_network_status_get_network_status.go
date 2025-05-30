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
	"fmt"
)

// ProfileNetworkStatusGetNetworkStatus Enum representing the overall network status of the user
type ProfileNetworkStatusGetNetworkStatus string

var _ = fmt.Printf

// List of ProfileNetworkStatusGetNetworkStatus
const (
	PROFILENETWORKSTATUSGETNETWORKSTATUS_UNKNOWN ProfileNetworkStatusGetNetworkStatus = "UNKNOWN"
	PROFILENETWORKSTATUSGETNETWORKSTATUS_RETURNING_USER ProfileNetworkStatusGetNetworkStatus = "RETURNING_USER"
)

var allowedProfileNetworkStatusGetNetworkStatusEnumValues = []ProfileNetworkStatusGetNetworkStatus{
	"UNKNOWN",
	"RETURNING_USER",
}

func (v *ProfileNetworkStatusGetNetworkStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := ProfileNetworkStatusGetNetworkStatus(value)


	*v = enumTypeValue
	return nil
}

// NewProfileNetworkStatusGetNetworkStatusFromValue returns a pointer to a valid ProfileNetworkStatusGetNetworkStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProfileNetworkStatusGetNetworkStatusFromValue(v string) (*ProfileNetworkStatusGetNetworkStatus, error) {
	ev := ProfileNetworkStatusGetNetworkStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProfileNetworkStatusGetNetworkStatus) IsValid() bool {
	for _, existing := range allowedProfileNetworkStatusGetNetworkStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProfileNetworkStatusGetNetworkStatus value
func (v ProfileNetworkStatusGetNetworkStatus) Ptr() *ProfileNetworkStatusGetNetworkStatus {
	return &v
}

type NullableProfileNetworkStatusGetNetworkStatus struct {
	value *ProfileNetworkStatusGetNetworkStatus
	isSet bool
}

func (v NullableProfileNetworkStatusGetNetworkStatus) Get() *ProfileNetworkStatusGetNetworkStatus {
	return v.value
}

func (v *NullableProfileNetworkStatusGetNetworkStatus) Set(val *ProfileNetworkStatusGetNetworkStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileNetworkStatusGetNetworkStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileNetworkStatusGetNetworkStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileNetworkStatusGetNetworkStatus(val *ProfileNetworkStatusGetNetworkStatus) *NullableProfileNetworkStatusGetNetworkStatus {
	return &NullableProfileNetworkStatusGetNetworkStatus{value: val, isSet: true}
}

func (v NullableProfileNetworkStatusGetNetworkStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileNetworkStatusGetNetworkStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

