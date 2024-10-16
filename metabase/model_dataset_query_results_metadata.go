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

// checks if the DatasetQueryResultsMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatasetQueryResultsMetadata{}

// DatasetQueryResultsMetadata struct for DatasetQueryResultsMetadata
type DatasetQueryResultsMetadata struct {
	Checksum *string                             `json:"checksum,omitempty"`
	Columns  []DatasetQueryResultsMetadataColumn `json:"columns,omitempty"`
}

// NewDatasetQueryResultsMetadata instantiates a new DatasetQueryResultsMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatasetQueryResultsMetadata() *DatasetQueryResultsMetadata {
	this := DatasetQueryResultsMetadata{}
	return &this
}

// NewDatasetQueryResultsMetadataWithDefaults instantiates a new DatasetQueryResultsMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatasetQueryResultsMetadataWithDefaults() *DatasetQueryResultsMetadata {
	this := DatasetQueryResultsMetadata{}
	return &this
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *DatasetQueryResultsMetadata) GetChecksum() string {
	if o == nil || IsNil(o.Checksum) {
		var ret string
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetQueryResultsMetadata) GetChecksumOk() (*string, bool) {
	if o == nil || IsNil(o.Checksum) {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *DatasetQueryResultsMetadata) HasChecksum() bool {
	if o != nil && !IsNil(o.Checksum) {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given string and assigns it to the Checksum field.
func (o *DatasetQueryResultsMetadata) SetChecksum(v string) {
	o.Checksum = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *DatasetQueryResultsMetadata) GetColumns() []DatasetQueryResultsMetadataColumn {
	if o == nil || IsNil(o.Columns) {
		var ret []DatasetQueryResultsMetadataColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetQueryResultsMetadata) GetColumnsOk() ([]DatasetQueryResultsMetadataColumn, bool) {
	if o == nil || IsNil(o.Columns) {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *DatasetQueryResultsMetadata) HasColumns() bool {
	if o != nil && !IsNil(o.Columns) {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []DatasetQueryResultsMetadataColumn and assigns it to the Columns field.
func (o *DatasetQueryResultsMetadata) SetColumns(v []DatasetQueryResultsMetadataColumn) {
	o.Columns = v
}

func (o DatasetQueryResultsMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatasetQueryResultsMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Checksum) {
		toSerialize["checksum"] = o.Checksum
	}
	if !IsNil(o.Columns) {
		toSerialize["columns"] = o.Columns
	}
	return toSerialize, nil
}

type NullableDatasetQueryResultsMetadata struct {
	value *DatasetQueryResultsMetadata
	isSet bool
}

func (v NullableDatasetQueryResultsMetadata) Get() *DatasetQueryResultsMetadata {
	return v.value
}

func (v *NullableDatasetQueryResultsMetadata) Set(val *DatasetQueryResultsMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableDatasetQueryResultsMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableDatasetQueryResultsMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatasetQueryResultsMetadata(val *DatasetQueryResultsMetadata) *NullableDatasetQueryResultsMetadata {
	return &NullableDatasetQueryResultsMetadata{value: val, isSet: true}
}

func (v NullableDatasetQueryResultsMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatasetQueryResultsMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
