/*
 * Metabase
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package metabase

type DatasetQueryResultsData struct {
	Columns         []string                      `json:"columns,omitempty"`
	Rows            [][]interface{}               `json:"rows,omitempty"`
	NativeForm      DatasetQueryResultsNativeForm `json:"native_form,omitempty"`
	Cols            []DatasetQueryResultsCol      `json:"cols,omitempty"`
	ResultsMetadata DatasetQueryResultsMetadata   `json:"results_metadata,omitempty"`
	RowsTruncated   int64                         `json:"rows_truncated,omitempty"`
}
