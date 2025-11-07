# User

## Properties

Name 			| Type 					| Description | Notes
------------ 	| ------------- 		| ----------- | -------------
**Id** 			| **int32** 			|  | [optional] 
**Email** 		| Pointer to **string** |  | [optional] 
**CommonName** 	| Pointer to **string** |  | [optional] 
**FirstName** 	| Pointer to **string** |  | [optional] 
**LastName** 	| Pointer to **string** |  | [optional] 
**Locale** 		| Pointer to **string** |  | [optional] 
**LastLogin** 	| Pointer to **time.Time**	|  | [optional] 
**IsActive** 	| Pointer to **bool** 	|  | [optional] 
**UpdatedAt** 	| Pointer to **time.Time** |  | [optional] 
**DateJoined** 	| Pointer to **time.Time** |  | [optional] 
**GroupIDs** 	| Slice of **int32** 	|  | [optional] 
**IsSuperuser** | Pointer to **bool**  	|  | [optional] 
**PersonalCollectionID** | Pointer to **time.Time** |  | [optional] 
**IsQbnewb** 	| Pointer to **bool** 	|  | [optional] 
**JwtAttributes** | Pointer to **string** |  | [optional] 
**LoginAttributes** | ? 				| type struct unknown | [optional] 
**Tenant_id** 	| Pointer to **string** | no confirmation of type yet | [optional] 
**SSOSource** 	| Pointer to **string** |  | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object  
This constructor will assign default values to properties that have it defined,  
and ensures all required properties by API are set, if any.

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object  
This constructor will only assign default values to properties that have it defined,  
but does not guarantee that properties required by the API are set.

### MarshalJSON

`func (o User) MarshalJSON() ([]byte, error)`

MarshalJSON encodes the User struct into JSON format using the internal `ToMap()` logic.

### ToMap

`func (o User) ToMap() (map[string]interface{}, error)`

ToMap converts the User struct into a map for easier JSON serialization and manipulation.

---

## Nullable Type

### NullableUser

`type NullableUser struct`

Wrapper type to handle optional presence of a `User` in API responses.

### Methods

#### Get

`func (v NullableUser) Get() *User`

Get returns the value.

#### Set

`func (v *NullableUser) Set(val *User)`

Set assigns the provided value and marks the field as set.

#### IsSet

`func (v NullableUser) IsSet() bool`

IsSet returns whether the field has been explicitly set.

#### Unset

`func (v *NullableUser) Unset()`

Unset clears the stored value and marks it as unset.

#### NewNullableUser

`func NewNullableUser(val *User) *NullableUser`

NewNullableUser returns a new instance of NullableUser wrapping the given User.

#### MarshalJSON

`func (v NullableUser) MarshalJSON() ([]byte, error)`

MarshalJSON serializes the underlying User value if set.

#### UnmarshalJSON

`func (v *NullableUser) UnmarshalJSON(src []byte) error`

UnmarshalJSON deserializes the JSON input into the underlying User value and marks it as set.

---

## Related Types

### Users

`type Users struct`

Wrapper type for API responses that return multiple users.

#### Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **[]User** | List of user objects returned by the API | [optional] 
**Total** | Pointer to **int32** | Total number of users | [optional] 
**Limit** | Pointer to **int32** | Pagination limit | [optional] 
**Offset** | Pointer to **int32** | Pagination offset | [optional] 

#### Methods

##### NewUsers

`func NewUsers(data []User) *Users`

NewUsers instantiates a new Users object with provided user data.

##### NewUsersWithDefaults

`func NewUsersWithDefaults() *Users`

NewUsersWithDefaults instantiates a new empty Users object.

##### MarshalJSON

`func (u Users) MarshalJSON() ([]byte, error)`

MarshalJSON serializes the Users struct into JSON.

##### UnmarshalJSON

`func (u *Users) UnmarshalJSON(b []byte) error`

UnmarshalJSON decodes JSON into a Users object, mapping `data`, `total`, `limit`, and `offset` fields.

##### ToMap

`func (u Users) ToMap() (map[string]interface{}, error)`

ToMap converts the Users struct into a generic map.

---

### NullableUsers

`type NullableUsers struct`

Wrapper for optional presence of a `Users` object in API responses.

#### Methods

##### Get

`func (v NullableUsers) Get() *Users`

Returns the Users value.

##### Set

`func (v *NullableUsers) Set(val *Users)`

Assigns a value and marks it as set.

##### IsSet

`func (v NullableUsers) IsSet() bool`

Checks whether the value is set.

##### Unset

`func (v *NullableUsers) Unset()`

Clears the value.

##### NewNullableUsers

`func NewNullableUsers(val *Users) *NullableUsers`

Returns a new NullableUsers containing the given value.

##### MarshalJSON

`func (v NullableUsers) MarshalJSON() ([]byte, error)`

Serializes the wrapped value if set.

##### UnmarshalJSON

`func (v *NullableUsers) UnmarshalJSON(src []byte) error`

Deserializes JSON input and marks the object as set.

---

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


