/*
 * Nudsf_DataRepository
 *
 * Nudsf Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * Source file: 3GPP TS 29.598 UDSF Services, V17.6.0.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.598/
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nudsf_DataRepository

import (
    "github.com/free5gc/openapi"
    "github.com/free5gc/openapi/models"

	"context"
	"io/ioutil"
	"net/url"
	"strings"
	"fmt"
)

// Linger please
var (
	_ context.Context
)

type BlockCRUDApiService service

/*
BlockCRUDApiService Create or Update a specific Block in a Record.
Create or update a specific Block, related to a Record
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param RealmId - Identifier of the Realm
 * @param StorageId - Identifier of the Storage
 * @param RecordId - Identifier of the Record
 * @param BlockId - Id of the Block
 * @param Body - information on the Block to create
 * @param GetPrevious - Retrieve the Block before update
 * @param IfNoneMatch - Validator for conditional requests, as described in RFC 7232, 3.2
 * @param IfMatch - Record validator for conditional requests, as described in RFC 7232, 3.2
 * @param SupportedFeatures - Features required to be supported by the target NF

@return CreateOrModifyBlockResponse
*/

// CreateOrModifyBlockRequest
type CreateOrModifyBlockRequest struct {
    RealmId *string
    StorageId *string
    RecordId *string
    BlockId *string
    Body map[string]interface{}
    GetPrevious *bool
    IfNoneMatch *string
    IfMatch *string
    SupportedFeatures *string
}

func (r *CreateOrModifyBlockRequest) SetRealmId(RealmId string) {
    r.RealmId = &RealmId
}
func (r *CreateOrModifyBlockRequest) SetStorageId(StorageId string) {
    r.StorageId = &StorageId
}
func (r *CreateOrModifyBlockRequest) SetRecordId(RecordId string) {
    r.RecordId = &RecordId
}
func (r *CreateOrModifyBlockRequest) SetBlockId(BlockId string) {
    r.BlockId = &BlockId
}
func (r *CreateOrModifyBlockRequest) SetBody(Body map[string]interface{}) {
    r.Body = Body
}
func (r *CreateOrModifyBlockRequest) SetGetPrevious(GetPrevious bool) {
    r.GetPrevious = &GetPrevious
}
func (r *CreateOrModifyBlockRequest) SetIfNoneMatch(IfNoneMatch string) {
    r.IfNoneMatch = &IfNoneMatch
}
func (r *CreateOrModifyBlockRequest) SetIfMatch(IfMatch string) {
    r.IfMatch = &IfMatch
}
func (r *CreateOrModifyBlockRequest) SetSupportedFeatures(SupportedFeatures string) {
    r.SupportedFeatures = &SupportedFeatures
}

func (a *BlockCRUDApiService) CreateOrModifyBlock(ctx context.Context, request *CreateOrModifyBlockRequest) (*CreateOrModifyBlockResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Put")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CreateOrModifyBlockResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/{realmId}/{storageId}/records/{recordId}/blocks/{blockId}"
	localVarPath = strings.Replace(localVarPath, "{"+"realmId"+"}", fmt.Sprintf("%v", request.RealmId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"storageId"+"}", fmt.Sprintf("%v", request.StorageId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"recordId"+"}", fmt.Sprintf("%v", request.RecordId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"blockId"+"}", fmt.Sprintf("%v", request.BlockId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	
    if request.GetPrevious != nil {
        
        localVarQueryParams.Add("get-previous", openapi.ParameterToString(request.GetPrevious, ""))
    } 
    if request.SupportedFeatures != nil {
        
        localVarQueryParams.Add("supported-features", openapi.ParameterToString(request.SupportedFeatures, ""))
    } 

    localVarHTTPContentTypes := []string{"application/json"}

    localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{ "*/*", "application/problem+json" }

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}


	if request.IfNoneMatch != nil {
        
		localVarHeaderParams["If-None-Match"] = openapi.ParameterToString(request.IfNoneMatch, "")
	} 

	if request.IfMatch != nil {
        
		localVarHeaderParams["If-Match"] = openapi.ParameterToString(request.IfMatch, "")
	} 


	// body params
	localVarPostBody = request.Body


	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, err
	}

    apiError := openapi.GenericOpenAPIError{
        RawBody:     localVarBody,
        ErrorStatus: localVarHTTPResponse.Status,
    }

    switch localVarHTTPResponse.StatusCode {
        case 200:
            err = openapi.Deserialize(&localVarReturnValue.map[string]interface{}, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            localVarReturnValue.CacheControl = localVarHTTPResponse.Header.Get("Cache-Control")
            localVarReturnValue.ETag = localVarHTTPResponse.Header.Get("ETag")
            localVarReturnValue.LastModified = localVarHTTPResponse.Header.Get("Last-Modified")
            return &localVarReturnValue, nil
        case 201:
        return &localVarReturnValue, apiError
        case 204:
        return &localVarReturnValue, apiError
        case 400:
            var v CreateOrModifyBlockError
            err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            apiError.ErrorModel = v
            return nil, apiError
        case 401:
            var v CreateOrModifyBlockError
            err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            apiError.ErrorModel = v
            return nil, apiError
        case 403:
            var v CreateOrModifyBlockError
            err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            apiError.ErrorModel = v
            return nil, apiError
        case 404:
            var v CreateOrModifyBlockError
            err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            apiError.ErrorModel = v
            return nil, apiError
        case 408:
            var v CreateOrModifyBlockError
            err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            apiError.ErrorModel = v
            return nil, apiError
        case 412:
            var v CreateOrModifyBlockError
            err = openapi.Deserialize(&v.map[string]interface{}, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            v.CacheControl = localVarHTTPResponse.Header.Get("Cache-Control")
            v.ETag = localVarHTTPResponse.Header.Get("ETag")
            v.LastModified = localVarHTTPResponse.Header.Get("Last-Modified")
            apiError.ErrorModel = v
            return nil, apiError
        case 413:
            var v CreateOrModifyBlockError
            err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            apiError.ErrorModel = v
            return nil, apiError
        case 500:
            var v CreateOrModifyBlockError
            err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            apiError.ErrorModel = v
            return nil, apiError
        case 503:
            var v CreateOrModifyBlockError
            err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            apiError.ErrorModel = v
            return nil, apiError
        default:
        return &localVarReturnValue, apiError
    }
}

type CreateOrModifyBlockResponse struct {
    CacheControl string
ETag string
LastModified string
Location string
    map[string]interface{} map[string]interface{}
}

type CreateOrModifyBlockError struct {
    CacheControl string
ETag string
LastModified string
    map[string]interface{} models.ProblemDetails
map[string]interface{} map[string]interface{}
}

/*
BlockCRUDApiService Delete a specific Block. Then update the Record
delete a specific Block, related to a Record
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param RealmId - Identifier of the Realm
 * @param StorageId - Identifier of the Storage
 * @param RecordId - Identifier of the Record
 * @param BlockId - Id of the Block
 * @param GetPrevious - Retrieve the Block before delete
 * @param IfMatch - Record validator for conditional requests, as described in RFC 7232, 3.2
 * @param SupportedFeatures - Features required to be supported by the target NF

@return DeleteBlockResponse
*/

// DeleteBlockRequest
type DeleteBlockRequest struct {
    RealmId *string
    StorageId *string
    RecordId *string
    BlockId *string
    GetPrevious *bool
    IfMatch *string
    SupportedFeatures *string
}

func (r *DeleteBlockRequest) SetRealmId(RealmId string) {
    r.RealmId = &RealmId
}
func (r *DeleteBlockRequest) SetStorageId(StorageId string) {
    r.StorageId = &StorageId
}
func (r *DeleteBlockRequest) SetRecordId(RecordId string) {
    r.RecordId = &RecordId
}
func (r *DeleteBlockRequest) SetBlockId(BlockId string) {
    r.BlockId = &BlockId
}
func (r *DeleteBlockRequest) SetGetPrevious(GetPrevious bool) {
    r.GetPrevious = &GetPrevious
}
func (r *DeleteBlockRequest) SetIfMatch(IfMatch string) {
    r.IfMatch = &IfMatch
}
func (r *DeleteBlockRequest) SetSupportedFeatures(SupportedFeatures string) {
    r.SupportedFeatures = &SupportedFeatures
}

func (a *BlockCRUDApiService) DeleteBlock(ctx context.Context, request *DeleteBlockRequest) (*DeleteBlockResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Delete")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DeleteBlockResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/{realmId}/{storageId}/records/{recordId}/blocks/{blockId}"
	localVarPath = strings.Replace(localVarPath, "{"+"realmId"+"}", fmt.Sprintf("%v", request.RealmId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"storageId"+"}", fmt.Sprintf("%v", request.StorageId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"recordId"+"}", fmt.Sprintf("%v", request.RecordId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"blockId"+"}", fmt.Sprintf("%v", request.BlockId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	
    if request.GetPrevious != nil {
        
        localVarQueryParams.Add("get-previous", openapi.ParameterToString(request.GetPrevious, ""))
    } 
    if request.SupportedFeatures != nil {
        
        localVarQueryParams.Add("supported-features", openapi.ParameterToString(request.SupportedFeatures, ""))
    } 

    localVarHTTPContentTypes := []string{"application/json"}

    localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{ "*/*", "application/problem+json" }

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}


	if request.IfMatch != nil {
        
		localVarHeaderParams["If-Match"] = openapi.ParameterToString(request.IfMatch, "")
	} 



	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, err
	}

    apiError := openapi.GenericOpenAPIError{
        RawBody:     localVarBody,
        ErrorStatus: localVarHTTPResponse.Status,
    }

    switch localVarHTTPResponse.StatusCode {
        case 200:
            err = openapi.Deserialize(&localVarReturnValue.map[string]interface{}, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            localVarReturnValue.ETag = localVarHTTPResponse.Header.Get("ETag")
            localVarReturnValue.LastModified = localVarHTTPResponse.Header.Get("Last-Modified")
            return &localVarReturnValue, nil
        case 204:
        return &localVarReturnValue, apiError
        case 400:
            var v DeleteBlockError
            err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            apiError.ErrorModel = v
            return nil, apiError
        case 401:
            var v DeleteBlockError
            err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            apiError.ErrorModel = v
            return nil, apiError
        case 403:
            var v DeleteBlockError
            err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            apiError.ErrorModel = v
            return nil, apiError
        case 404:
            var v DeleteBlockError
            err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            apiError.ErrorModel = v
            return nil, apiError
        case 408:
            var v DeleteBlockError
            err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            apiError.ErrorModel = v
            return nil, apiError
        case 412:
            var v DeleteBlockError
            err = openapi.Deserialize(&v.map[string]interface{}, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            v.CacheControl = localVarHTTPResponse.Header.Get("Cache-Control")
            v.ETag = localVarHTTPResponse.Header.Get("ETag")
            v.LastModified = localVarHTTPResponse.Header.Get("Last-Modified")
            apiError.ErrorModel = v
            return nil, apiError
        case 500:
            var v DeleteBlockError
            err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            apiError.ErrorModel = v
            return nil, apiError
        case 503:
            var v DeleteBlockError
            err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            apiError.ErrorModel = v
            return nil, apiError
        default:
        return &localVarReturnValue, apiError
    }
}

type DeleteBlockResponse struct {
    ETag string
LastModified string
    map[string]interface{} map[string]interface{}
}

type DeleteBlockError struct {
    CacheControl string
ETag string
LastModified string
    map[string]interface{} models.ProblemDetails
map[string]interface{} map[string]interface{}
}

/*
BlockCRUDApiService Retrieve a specific Block
retrieve a specific Block
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param RealmId - Identifier of the Realm
 * @param StorageId - Identifier of the Storage
 * @param RecordId - Identifier of the Record
 * @param BlockId - Id of the Block
 * @param IfNoneMatch - Validator for conditional requests, as described in RFC 7232, 3.2
 * @param IfModifiedSince - Validator for conditional requests, as described in RFC 7232, 3.3
 * @param SupportedFeatures - Features required to be supported by the target NF

@return GetBlockResponse
*/

// GetBlockRequest
type GetBlockRequest struct {
    RealmId *string
    StorageId *string
    RecordId *string
    BlockId *string
    IfNoneMatch *string
    IfModifiedSince *string
    SupportedFeatures *string
}

func (r *GetBlockRequest) SetRealmId(RealmId string) {
    r.RealmId = &RealmId
}
func (r *GetBlockRequest) SetStorageId(StorageId string) {
    r.StorageId = &StorageId
}
func (r *GetBlockRequest) SetRecordId(RecordId string) {
    r.RecordId = &RecordId
}
func (r *GetBlockRequest) SetBlockId(BlockId string) {
    r.BlockId = &BlockId
}
func (r *GetBlockRequest) SetIfNoneMatch(IfNoneMatch string) {
    r.IfNoneMatch = &IfNoneMatch
}
func (r *GetBlockRequest) SetIfModifiedSince(IfModifiedSince string) {
    r.IfModifiedSince = &IfModifiedSince
}
func (r *GetBlockRequest) SetSupportedFeatures(SupportedFeatures string) {
    r.SupportedFeatures = &SupportedFeatures
}

func (a *BlockCRUDApiService) GetBlock(ctx context.Context, request *GetBlockRequest) (*GetBlockResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetBlockResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/{realmId}/{storageId}/records/{recordId}/blocks/{blockId}"
	localVarPath = strings.Replace(localVarPath, "{"+"realmId"+"}", fmt.Sprintf("%v", request.RealmId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"storageId"+"}", fmt.Sprintf("%v", request.StorageId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"recordId"+"}", fmt.Sprintf("%v", request.RecordId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"blockId"+"}", fmt.Sprintf("%v", request.BlockId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	
    if request.SupportedFeatures != nil {
        
        localVarQueryParams.Add("supported-features", openapi.ParameterToString(request.SupportedFeatures, ""))
    } 

    localVarHTTPContentTypes := []string{"application/json"}

    localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{ "*/*", "application/problem+json" }

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}


	if request.IfNoneMatch != nil {
        
		localVarHeaderParams["If-None-Match"] = openapi.ParameterToString(request.IfNoneMatch, "")
	} 

	if request.IfModifiedSince != nil {
        
		localVarHeaderParams["If-Modified-Since"] = openapi.ParameterToString(request.IfModifiedSince, "")
	} 



	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, err
	}

    apiError := openapi.GenericOpenAPIError{
        RawBody:     localVarBody,
        ErrorStatus: localVarHTTPResponse.Status,
    }

    switch localVarHTTPResponse.StatusCode {
        case 200:
            err = openapi.Deserialize(&localVarReturnValue.map[string]interface{}, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            localVarReturnValue.CacheControl = localVarHTTPResponse.Header.Get("Cache-Control")
            localVarReturnValue.ETag = localVarHTTPResponse.Header.Get("ETag")
            localVarReturnValue.LastModified = localVarHTTPResponse.Header.Get("Last-Modified")
            return &localVarReturnValue, nil
        case 304:
            var v GetBlockError
            err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            v.CacheControl = localVarHTTPResponse.Header.Get("Cache-Control")
            v.ETag = localVarHTTPResponse.Header.Get("ETag")
            v.RetryAfter = localVarHTTPResponse.Header.Get("Retry-After")
            apiError.ErrorModel = v
            return nil, apiError
        case 400:
            var v GetBlockError
            err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            apiError.ErrorModel = v
            return nil, apiError
        case 401:
            var v GetBlockError
            err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            apiError.ErrorModel = v
            return nil, apiError
        case 403:
            var v GetBlockError
            err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            apiError.ErrorModel = v
            return nil, apiError
        case 404:
            var v GetBlockError
            err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            apiError.ErrorModel = v
            return nil, apiError
        case 500:
            var v GetBlockError
            err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            apiError.ErrorModel = v
            return nil, apiError
        case 503:
            var v GetBlockError
            err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            apiError.ErrorModel = v
            return nil, apiError
        default:
        return &localVarReturnValue, apiError
    }
}

type GetBlockResponse struct {
    CacheControl string
ETag string
LastModified string
    map[string]interface{} map[string]interface{}
}

type GetBlockError struct {
    CacheControl string
ETag string
RetryAfter interface{}
    map[string]interface{} models.ProblemDetails
}

/*
BlockCRUDApiService Record's Blocks access
retrieve all Blocks of a specific Record
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param RealmId - Identifier of the Realm
 * @param StorageId - Identifier of the Storage
 * @param RecordId - Identifier of the Record
 * @param SupportedFeatures - Features required to be supported by the target NF

@return GetBlockListResponse
*/

// GetBlockListRequest
type GetBlockListRequest struct {
    RealmId *string
    StorageId *string
    RecordId *string
    SupportedFeatures *string
}

func (r *GetBlockListRequest) SetRealmId(RealmId string) {
    r.RealmId = &RealmId
}
func (r *GetBlockListRequest) SetStorageId(StorageId string) {
    r.StorageId = &StorageId
}
func (r *GetBlockListRequest) SetRecordId(RecordId string) {
    r.RecordId = &RecordId
}
func (r *GetBlockListRequest) SetSupportedFeatures(SupportedFeatures string) {
    r.SupportedFeatures = &SupportedFeatures
}

func (a *BlockCRUDApiService) GetBlockList(ctx context.Context, request *GetBlockListRequest) (*GetBlockListResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetBlockListResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/{realmId}/{storageId}/records/{recordId}/blocks"
	localVarPath = strings.Replace(localVarPath, "{"+"realmId"+"}", fmt.Sprintf("%v", request.RealmId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"storageId"+"}", fmt.Sprintf("%v", request.StorageId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"recordId"+"}", fmt.Sprintf("%v", request.RecordId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	
    if request.SupportedFeatures != nil {
        
        localVarQueryParams.Add("supported-features", openapi.ParameterToString(request.SupportedFeatures, ""))
    } 

    localVarHTTPContentTypes := []string{"application/json"}

    localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{ "multipart/parallel", "application/problem+json" }

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}




	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, err
	}

    apiError := openapi.GenericOpenAPIError{
        RawBody:     localVarBody,
        ErrorStatus: localVarHTTPResponse.Status,
    }

    switch localVarHTTPResponse.StatusCode {
        case 200:
            err = openapi.Deserialize(&localVarReturnValue.map[string]interface{}, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            localVarReturnValue.CacheControl = localVarHTTPResponse.Header.Get("Cache-Control")
            localVarReturnValue.ETag = localVarHTTPResponse.Header.Get("ETag")
            localVarReturnValue.LastModified = localVarHTTPResponse.Header.Get("Last-Modified")
            return &localVarReturnValue, nil
        case 204:
        return &localVarReturnValue, apiError
        case 400:
            var v GetBlockListError
            err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            apiError.ErrorModel = v
            return nil, apiError
        case 401:
            var v GetBlockListError
            err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            apiError.ErrorModel = v
            return nil, apiError
        case 403:
            var v GetBlockListError
            err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            apiError.ErrorModel = v
            return nil, apiError
        case 404:
            var v GetBlockListError
            err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            apiError.ErrorModel = v
            return nil, apiError
        case 500:
            var v GetBlockListError
            err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            apiError.ErrorModel = v
            return nil, apiError
        case 503:
            var v GetBlockListError
            err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                return nil, err
            }
            apiError.ErrorModel = v
            return nil, apiError
        default:
        return &localVarReturnValue, apiError
    }
}

type GetBlockListResponse struct {
    CacheControl string
ETag string
LastModified string
    map[string]interface{} map[string]interface{}
}

type GetBlockListError struct {
        map[string]interface{} models.ProblemDetails
}
