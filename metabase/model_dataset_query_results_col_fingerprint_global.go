/*
Metabase

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metabase

import (
	"encoding/json"
)

// checks if the DatasetQueryResultsColFingerprintGlobal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatasetQueryResultsColFingerprintGlobal{}

// DatasetQueryResultsColFingerprintGlobal struct for DatasetQueryResultsColFingerprintGlobal
type DatasetQueryResultsColFingerprintGlobal struct {
	DistinctCount *int64 `json:"distinct-count,omitempty"`
}

// NewDatasetQueryResultsColFingerprintGlobal instantiates a new DatasetQueryResultsColFingerprintGlobal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatasetQueryResultsColFingerprintGlobal() *DatasetQueryResultsColFingerprintGlobal {
	this := DatasetQueryResultsColFingerprintGlobal{}
	return &this
}

// NewDatasetQueryResultsColFingerprintGlobalWithDefaults instantiates a new DatasetQueryResultsColFingerprintGlobal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatasetQueryResultsColFingerprintGlobalWithDefaults() *DatasetQueryResultsColFingerprintGlobal {
	this := DatasetQueryResultsColFingerprintGlobal{}
	return &this
}

// GetDistinctCount returns the DistinctCount field value if set, zero value otherwise.
func (o *DatasetQueryResultsColFingerprintGlobal) GetDistinctCount() int64 {
	if o == nil || IsNil(o.DistinctCount) {
		var ret int64
		return ret
	}
	return *o.DistinctCount
}

// GetDistinctCountOk returns a tuple with the DistinctCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetQueryResultsColFingerprintGlobal) GetDistinctCountOk() (*int64, bool) {
	if o == nil || IsNil(o.DistinctCount) {
		return nil, false
	}
	return o.DistinctCount, true
}

// HasDistinctCount returns a boolean if a field has been set.
func (o *DatasetQueryResultsColFingerprintGlobal) HasDistinctCount() bool {
	if o != nil && !IsNil(o.DistinctCount) {
		return true
	}

	return false
}

// SetDistinctCount gets a reference to the given int64 and assigns it to the DistinctCount field.
func (o *DatasetQueryResultsColFingerprintGlobal) SetDistinctCount(v int64) {
	o.DistinctCount = &v
}

func (o DatasetQueryResultsColFingerprintGlobal) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatasetQueryResultsColFingerprintGlobal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DistinctCount) {
		toSerialize["distinct-count"] = o.DistinctCount
	}
	return toSerialize, nil
}

type NullableDatasetQueryResultsColFingerprintGlobal struct {
	value *DatasetQueryResultsColFingerprintGlobal
	isSet bool
}

func (v NullableDatasetQueryResultsColFingerprintGlobal) Get() *DatasetQueryResultsColFingerprintGlobal {
	return v.value
}

func (v *NullableDatasetQueryResultsColFingerprintGlobal) Set(val *DatasetQueryResultsColFingerprintGlobal) {
	v.value = val
	v.isSet = true
}

func (v NullableDatasetQueryResultsColFingerprintGlobal) IsSet() bool {
	return v.isSet
}

func (v *NullableDatasetQueryResultsColFingerprintGlobal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatasetQueryResultsColFingerprintGlobal(val *DatasetQueryResultsColFingerprintGlobal) *NullableDatasetQueryResultsColFingerprintGlobal {
	return &NullableDatasetQueryResultsColFingerprintGlobal{value: val, isSet: true}
}

func (v NullableDatasetQueryResultsColFingerprintGlobal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatasetQueryResultsColFingerprintGlobal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
