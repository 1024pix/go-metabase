/*
Metabase

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metabase

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// DatabaseApiService DatabaseApi service
type DatabaseApiService service

type ApiDatabaseMetadataRequest struct {
	ctx                      context.Context
	ApiService               *DatabaseApiService
	databaseId               int32
	includeHidden            *bool
	includeEditableDataModel *bool
}

// show hidden tables and fields.
func (r ApiDatabaseMetadataRequest) IncludeHidden(includeHidden bool) ApiDatabaseMetadataRequest {
	r.includeHidden = &includeHidden
	return r
}

// return tables for which the current user has data model editing permissions.
func (r ApiDatabaseMetadataRequest) IncludeEditableDataModel(includeEditableDataModel bool) ApiDatabaseMetadataRequest {
	r.includeEditableDataModel = &includeEditableDataModel
	return r
}

func (r ApiDatabaseMetadataRequest) Execute() (*Database, *http.Response, error) {
	return r.ApiService.DatabaseMetadataExecute(r)
}

/*
DatabaseMetadata Get metadata about a Database.

Get metadata about a Database, including all of its Tables and Fields. Returns DB, fields, and field values. By default only non-hidden tables and fields are returned.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param databaseId The database ID
	@return ApiDatabaseMetadataRequest
*/
func (a *DatabaseApiService) DatabaseMetadata(ctx context.Context, databaseId int32) ApiDatabaseMetadataRequest {
	return ApiDatabaseMetadataRequest{
		ApiService: a,
		ctx:        ctx,
		databaseId: databaseId,
	}
}

// Execute executes the request
//
//	@return Database
func (a *DatabaseApiService) DatabaseMetadataExecute(r ApiDatabaseMetadataRequest) (*Database, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Database
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseApiService.DatabaseMetadata")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/database/{databaseId}/metadata"
	localVarPath = strings.Replace(localVarPath, "{"+"databaseId"+"}", url.PathEscape(parameterValueToString(r.databaseId, "databaseId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.databaseId < 1 {
		return localVarReturnValue, nil, reportError("databaseId must be greater than 1")
	}

	if r.includeHidden != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_hidden", r.includeHidden, "")
	}
	if r.includeEditableDataModel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_editable_data_model", r.includeEditableDataModel, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListDatabasesRequest struct {
	ctx           context.Context
	ApiService    *DatabaseApiService
	includeTables *bool
	includeCards  *bool
}

// value may be nil, or if non-nil, value must be a valid boolean string (&#39;true&#39; or &#39;false&#39;).
func (r ApiListDatabasesRequest) IncludeTables(includeTables bool) ApiListDatabasesRequest {
	r.includeTables = &includeTables
	return r
}

// value may be nil, or if non-nil, value must be a valid boolean string (&#39;true&#39; or &#39;false&#39;).
func (r ApiListDatabasesRequest) IncludeCards(includeCards bool) ApiListDatabasesRequest {
	r.includeCards = &includeCards
	return r
}

func (r ApiListDatabasesRequest) Execute() ([]Database, *http.Response, error) {
	return r.ApiService.ListDatabasesExecute(r)
}

/*
ListDatabases List Databases

Fetch all Databases. include_tables means we should hydrate the Tables belonging to each DB. include_cards here means we should also include virtual Table entries for saved Questions, e.g. so we can easily use them as source Tables in queries. Default for both is false.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListDatabasesRequest
*/
func (a *DatabaseApiService) ListDatabases(ctx context.Context) ApiListDatabasesRequest {
	return ApiListDatabasesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []Database
func (a *DatabaseApiService) ListDatabasesExecute(r ApiListDatabasesRequest) ([]Database, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []Database
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseApiService.ListDatabases")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/database"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeTables != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_tables", r.includeTables, "")
	}
	if r.includeCards != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_cards", r.includeCards, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
