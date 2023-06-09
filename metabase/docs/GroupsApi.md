# \GroupsApi

All URIs are relative to *http://example.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPermissionsGroup**](GroupsApi.md#GetPermissionsGroup) | **Get** /api/permissions/group | Fetch all permissions group.



## GetPermissionsGroup

> []PermissionGroup GetPermissionsGroup(ctx).Execute()

Fetch all permissions group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/grokify/go-metabase"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetPermissionsGroup(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetPermissionsGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPermissionsGroup`: []PermissionGroup
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetPermissionsGroup`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPermissionsGroupRequest struct via the builder pattern


### Return type

[**[]PermissionGroup**](PermissionGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

