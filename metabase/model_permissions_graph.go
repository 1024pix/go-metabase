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

// checks if the PermissionsGraph type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionsGraph{}

// PermissionsGraph struct for PermissionsGraph
type PermissionsGraph struct {
	Revision *int32                                     `json:"revision,omitempty"`
	Groups   *map[string]map[string]PermissionGraphData `json:"groups,omitempty"`
}

// NewPermissionsGraph instantiates a new PermissionsGraph object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionsGraph() *PermissionsGraph {
	this := PermissionsGraph{}
	return &this
}

// NewPermissionsGraphWithDefaults instantiates a new PermissionsGraph object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionsGraphWithDefaults() *PermissionsGraph {
	this := PermissionsGraph{}
	return &this
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *PermissionsGraph) GetRevision() int32 {
	if o == nil || IsNil(o.Revision) {
		var ret int32
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsGraph) GetRevisionOk() (*int32, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *PermissionsGraph) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given int32 and assigns it to the Revision field.
func (o *PermissionsGraph) SetRevision(v int32) {
	o.Revision = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *PermissionsGraph) GetGroups() map[string]map[string]PermissionGraphData {
	if o == nil || IsNil(o.Groups) {
		var ret map[string]map[string]PermissionGraphData
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsGraph) GetGroupsOk() (*map[string]map[string]PermissionGraphData, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *PermissionsGraph) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given map[string]map[string]PermissionGraphData and assigns it to the Groups field.
func (o *PermissionsGraph) SetGroups(v map[string]map[string]PermissionGraphData) {
	o.Groups = &v
}

func (o PermissionsGraph) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionsGraph) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	return toSerialize, nil
}

type NullablePermissionsGraph struct {
	value *PermissionsGraph
	isSet bool
}

func (v NullablePermissionsGraph) Get() *PermissionsGraph {
	return v.value
}

func (v *NullablePermissionsGraph) Set(val *PermissionsGraph) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionsGraph) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionsGraph) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionsGraph(val *PermissionsGraph) *NullablePermissionsGraph {
	return &NullablePermissionsGraph{value: val, isSet: true}
}

func (v NullablePermissionsGraph) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionsGraph) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
