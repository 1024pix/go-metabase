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

// checks if the DatabaseList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseList{}

// DatabaseList struct for DatabaseList
type DatabaseList struct {
	Data  []Database `json:"data,omitempty"`
	Total *int32     `json:"total,omitempty"`
}

// NewDatabaseList instantiates a new DatabaseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseList() *DatabaseList {
	this := DatabaseList{}
	return &this
}

// NewDatabaseListWithDefaults instantiates a new DatabaseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseListWithDefaults() *DatabaseList {
	this := DatabaseList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DatabaseList) GetData() []Database {
	if o == nil || IsNil(o.Data) {
		var ret []Database
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseList) GetDataOk() ([]Database, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DatabaseList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Database and assigns it to the Data field.
func (o *DatabaseList) SetData(v []Database) {
	o.Data = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *DatabaseList) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseList) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *DatabaseList) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *DatabaseList) SetTotal(v int32) {
	o.Total = &v
}

func (o DatabaseList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableDatabaseList struct {
	value *DatabaseList
	isSet bool
}

func (v NullableDatabaseList) Get() *DatabaseList {
	return v.value
}

func (v *NullableDatabaseList) Set(val *DatabaseList) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseList) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseList(val *DatabaseList) *NullableDatabaseList {
	return &NullableDatabaseList{value: val, isSet: true}
}

func (v NullableDatabaseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
