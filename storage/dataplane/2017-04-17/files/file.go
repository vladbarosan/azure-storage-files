package storage

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// FileClient is the client for the File methods of the Storage service.
type FileClient struct {
	ManagementClient
}

// NewFileClient creates an instance of the FileClient client.
func NewFileClient(url url.URL, p pipeline.Pipeline) FileClient {
	return FileClient{NewManagementClient(url, p)}
}

// AbortCopy aborts a pending Copy File operation, and leaves a destination file with zero length and full metadata.
//
// copyID is the copy identifier provided in the x-ms-copy-id header of the original Copy File operation. timeout is
// the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a>
func (client FileClient) AbortCopy(ctx context.Context, copyID string, copyActionAbortConstant string, timeout *int32) (*FileAbortCopyResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.abortCopyPreparer(copyID, copyActionAbortConstant, timeout)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.abortCopyResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*FileAbortCopyResponse), err
}

// abortCopyPreparer prepares the AbortCopy request.
func (client FileClient) abortCopyPreparer(copyID string, copyActionAbortConstant string, timeout *int32) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("copyid", copyID)
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	params.Set("comp", "copy")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-copy-action", copyActionAbortConstant)
	req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// abortCopyResponder handles the response to the AbortCopy request.
func (client FileClient) abortCopyResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusNoContent)
	if resp == nil {
		return nil, err
	}
	resp.Response().Body.Close()
	return &FileAbortCopyResponse{rawResponse: resp.Response()}, err
}

// Create creates a new file or replaces a file. Note it only initializes the file with no content.
//
// fileContentLength is specifies the maximum size for the file, up to 1 TB. fileTypeConstant is dummy constant
// parameter, file type can only be file. timeout is the timeout parameter is expressed in seconds. For more
// information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> fileContentType is sets the MIME content type of the file. The default
// type is 'application/octet-stream'. fileContentEncoding is specifies which content encodings have been applied to
// the file. fileContentLanguage is specifies the natural languages used by this resource. fileCacheControl is sets the
// file's cache control. The File service stores this value but does not use or modify it. fileContentMD5 is sets the
// file's MD5 hash. fileContentDisposition is sets the file's Content-Disposition header. metadata is a name-value pair
// to associate with a file storage object. Metadata names must adhere to the naming rules for C# identifiers.
func (client FileClient) Create(ctx context.Context, fileContentLength int64, fileTypeConstant string, timeout *int32, fileContentType *string, fileContentEncoding *string, fileContentLanguage *string, fileCacheControl *string, fileContentMD5 *string, fileContentDisposition *string, metadata map[string]string) (*FileCreateResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.createPreparer(fileContentLength, fileTypeConstant, timeout, fileContentType, fileContentEncoding, fileContentLanguage, fileCacheControl, fileContentMD5, fileContentDisposition, metadata)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*FileCreateResponse), err
}

// createPreparer prepares the Create request.
func (client FileClient) createPreparer(fileContentLength int64, fileTypeConstant string, timeout *int32, fileContentType *string, fileContentEncoding *string, fileContentLanguage *string, fileCacheControl *string, fileContentMD5 *string, fileContentDisposition *string, metadata map[string]string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	req.Header.Set("x-ms-content-length", fmt.Sprintf("%v", fileContentLength))
	req.Header.Set("x-ms-type", fileTypeConstant)
	if fileContentType != nil {
		req.Header.Set("x-ms-content-type", *fileContentType)
	}
	if fileContentEncoding != nil {
		req.Header.Set("x-ms-content-encoding", *fileContentEncoding)
	}
	if fileContentLanguage != nil {
		req.Header.Set("x-ms-content-language", *fileContentLanguage)
	}
	if fileCacheControl != nil {
		req.Header.Set("x-ms-cache-control", *fileCacheControl)
	}
	if fileContentMD5 != nil {
		req.Header.Set("x-ms-content-md5", *fileContentMD5)
	}
	if fileContentDisposition != nil {
		req.Header.Set("x-ms-content-disposition", *fileContentDisposition)
	}
	if metadata != nil {
		for k, v := range metadata {
			req.Header.Set("x-ms-meta-"+k, v)
		}
	}
	return req, nil
}

// createResponder handles the response to the Create request.
func (client FileClient) createResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	resp.Response().Body.Close()
	return &FileCreateResponse{rawResponse: resp.Response()}, err
}

// Delete removes the file from the storage account.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a>
func (client FileClient) Delete(ctx context.Context, timeout *int32) (*FileDeleteResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.deletePreparer(timeout)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.deleteResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*FileDeleteResponse), err
}

// deletePreparer prepares the Delete request.
func (client FileClient) deletePreparer(timeout *int32) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("DELETE", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// deleteResponder handles the response to the Delete request.
func (client FileClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	resp.Response().Body.Close()
	return &FileDeleteResponse{rawResponse: resp.Response()}, err
}

// Download reads or downloads a file from the system, including its metadata and properties.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> rangeParameter is return file data only from the specified byte range.
// rangeGetContentMD5 is when this header is set to true and specified together with the Range header, the service
// returns the MD5 hash for the range, as long as the range is less than or equal to 4 MB in size.
func (client FileClient) Download(ctx context.Context, timeout *int32, rangeParameter *string, rangeGetContentMD5 *bool) (*DownloadResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.downloadPreparer(timeout, rangeParameter, rangeGetContentMD5)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.downloadResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*DownloadResponse), err
}

// downloadPreparer prepares the Download request.
func (client FileClient) downloadPreparer(timeout *int32, rangeParameter *string, rangeGetContentMD5 *bool) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("GET", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	if rangeParameter != nil {
		req.Header.Set("x-ms-range", *rangeParameter)
	}
	if rangeGetContentMD5 != nil {
		req.Header.Set("x-ms-range-get-content-md5", fmt.Sprintf("%v", *rangeGetContentMD5))
	}
	return req, nil
}

// downloadResponder handles the response to the Download request.
func (client FileClient) downloadResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusPartialContent)
	if resp == nil {
		return nil, err
	}
	return &DownloadResponse{rawResponse: resp.Response()}, err
}

// GetProperties returns all user-defined metadata, standard HTTP properties, and system properties for the file. It
// does not return the content of the file.
//
// sharesnapshot is the snapshot parameter is an opaque DateTime value that, when present, specifies the share snapshot
// to query. timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a>
func (client FileClient) GetProperties(ctx context.Context, sharesnapshot *string, timeout *int32) (*FileGetPropertiesResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.getPropertiesPreparer(sharesnapshot, timeout)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getPropertiesResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*FileGetPropertiesResponse), err
}

// getPropertiesPreparer prepares the GetProperties request.
func (client FileClient) getPropertiesPreparer(sharesnapshot *string, timeout *int32) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("HEAD", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if sharesnapshot != nil {
		params.Set("sharesnapshot", *sharesnapshot)
	}
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// getPropertiesResponder handles the response to the GetProperties request.
func (client FileClient) getPropertiesResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	resp.Response().Body.Close()
	return &FileGetPropertiesResponse{rawResponse: resp.Response()}, err
}

// GetRangeList returns the list of valid ranges for a file.
//
// sharesnapshot is the snapshot parameter is an opaque DateTime value that, when present, specifies the share snapshot
// to query. timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> rangeParameter is specifies the range of bytes over which to list ranges,
// inclusively.
func (client FileClient) GetRangeList(ctx context.Context, sharesnapshot *string, timeout *int32, rangeParameter *string) (*Ranges, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.getRangeListPreparer(sharesnapshot, timeout, rangeParameter)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getRangeListResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ranges), err
}

// getRangeListPreparer prepares the GetRangeList request.
func (client FileClient) getRangeListPreparer(sharesnapshot *string, timeout *int32, rangeParameter *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("GET", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if sharesnapshot != nil {
		params.Set("sharesnapshot", *sharesnapshot)
	}
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	params.Set("comp", "rangelist")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	if rangeParameter != nil {
		req.Header.Set("x-ms-range", *rangeParameter)
	}
	return req, nil
}

// getRangeListResponder handles the response to the GetRangeList request.
func (client FileClient) getRangeListResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &Ranges{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = xml.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// SetHTTPHeaders sets HTTP headers on the file.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> fileContentLength is resizes a file to the specified size. If the
// specified byte value is less than the current size of the file, then all ranges above the specified byte value are
// cleared. fileContentType is sets the MIME content type of the file. The default type is 'application/octet-stream'.
// fileContentEncoding is specifies which content encodings have been applied to the file. fileContentLanguage is
// specifies the natural languages used by this resource. fileCacheControl is sets the file's cache control. The File
// service stores this value but does not use or modify it. fileContentMD5 is sets the file's MD5 hash.
// fileContentDisposition is sets the file's Content-Disposition header.
func (client FileClient) SetHTTPHeaders(ctx context.Context, timeout *int32, fileContentLength *int64, fileContentType *string, fileContentEncoding *string, fileContentLanguage *string, fileCacheControl *string, fileContentMD5 *string, fileContentDisposition *string) (*FileSetHTTPHeadersResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.setHTTPHeadersPreparer(timeout, fileContentLength, fileContentType, fileContentEncoding, fileContentLanguage, fileCacheControl, fileContentMD5, fileContentDisposition)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.setHTTPHeadersResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*FileSetHTTPHeadersResponse), err
}

// setHTTPHeadersPreparer prepares the SetHTTPHeaders request.
func (client FileClient) setHTTPHeadersPreparer(timeout *int32, fileContentLength *int64, fileContentType *string, fileContentEncoding *string, fileContentLanguage *string, fileCacheControl *string, fileContentMD5 *string, fileContentDisposition *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	params.Set("comp", "properties")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	if fileContentLength != nil {
		req.Header.Set("x-ms-content-length", fmt.Sprintf("%v", *fileContentLength))
	}
	if fileContentType != nil {
		req.Header.Set("x-ms-content-type", *fileContentType)
	}
	if fileContentEncoding != nil {
		req.Header.Set("x-ms-content-encoding", *fileContentEncoding)
	}
	if fileContentLanguage != nil {
		req.Header.Set("x-ms-content-language", *fileContentLanguage)
	}
	if fileCacheControl != nil {
		req.Header.Set("x-ms-cache-control", *fileCacheControl)
	}
	if fileContentMD5 != nil {
		req.Header.Set("x-ms-content-md5", *fileContentMD5)
	}
	if fileContentDisposition != nil {
		req.Header.Set("x-ms-content-disposition", *fileContentDisposition)
	}
	return req, nil
}

// setHTTPHeadersResponder handles the response to the SetHTTPHeaders request.
func (client FileClient) setHTTPHeadersResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	resp.Response().Body.Close()
	return &FileSetHTTPHeadersResponse{rawResponse: resp.Response()}, err
}

// SetMetadata updates user-defined metadata for the specified file.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> metadata is a name-value pair to associate with a file storage object.
// Metadata names must adhere to the naming rules for C# identifiers.
func (client FileClient) SetMetadata(ctx context.Context, timeout *int32, metadata map[string]string) (*FileSetMetadataResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.setMetadataPreparer(timeout, metadata)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.setMetadataResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*FileSetMetadataResponse), err
}

// setMetadataPreparer prepares the SetMetadata request.
func (client FileClient) setMetadataPreparer(timeout *int32, metadata map[string]string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	params.Set("comp", "metadata")
	req.URL.RawQuery = params.Encode()
	if metadata != nil {
		for k, v := range metadata {
			req.Header.Set("x-ms-meta-"+k, v)
		}
	}
	req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// setMetadataResponder handles the response to the SetMetadata request.
func (client FileClient) setMetadataResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	resp.Response().Body.Close()
	return &FileSetMetadataResponse{rawResponse: resp.Response()}, err
}

// StartCopy copies a blob or file to a destination file within the storage account.
//
// copySource is specifies the URL of the source file or blob, up to 2 KB in length. To copy a file to another file
// within the same storage account, you may use Shared Key to authenticate the source file. If you are copying a file
// from another storage account, or if you are copying a blob from the same storage account or another storage account,
// then you must authenticate the source file or blob using a shared access signature. If the source is a public blob,
// no authentication is required to perform the copy operation. A file in a share snapshot can also be specified as a
// copy source. timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> metadata is a name-value pair to associate with a file storage object.
// Metadata names must adhere to the naming rules for C# identifiers.
func (client FileClient) StartCopy(ctx context.Context, copySource string, timeout *int32, metadata map[string]string) (*FileStartCopyResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.startCopyPreparer(copySource, timeout, metadata)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.startCopyResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*FileStartCopyResponse), err
}

// startCopyPreparer prepares the StartCopy request.
func (client FileClient) startCopyPreparer(copySource string, timeout *int32, metadata map[string]string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	if metadata != nil {
		for k, v := range metadata {
			req.Header.Set("x-ms-meta-"+k, v)
		}
	}
	req.Header.Set("x-ms-copy-source", copySource)
	return req, nil
}

// startCopyResponder handles the response to the StartCopy request.
func (client FileClient) startCopyResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	resp.Response().Body.Close()
	return &FileStartCopyResponse{rawResponse: resp.Response()}, err
}

// UploadRange upload a range of bytes to a file.
//
// rangeParameter is specifies the range of bytes to be written. Both the start and end of the range must be specified.
// For an update operation, the range can be up to 4 MB in size. For a clear operation, the range can be up to the
// value of the file's full size. The File service accepts only a single byte range for the Range and 'x-ms-range'
// headers, and the byte range must be specified in the following format: bytes=startByte-endByte. fileRangeWrite is
// specify one of the following options: - Update: Writes the bytes specified by the request body into the specified
// range. The Range and Content-Length headers must match to perform the update. - Clear: Clears the specified range
// and releases the space used in storage for that range. To clear a range, set the Content-Length header to zero, and
// set the Range header to a value that indicates the range to clear, up to maximum file size. contentLength is
// specifies the number of bytes being transmitted in the request body. When the x-ms-write header is set to clear, the
// value of this header must be set to zero. optionalbody is initial data. optionalbody will be closed upon successful
// return. Callers should ensure closure when receiving an error.timeout is the timeout parameter is expressed in
// seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> contentMD5 is an MD5 hash of the content. This hash is used to verify the
// integrity of the data during transport. When the Content-MD5 header is specified, the File service compares the hash
// of the content that has arrived with the header value that was sent. If the two hashes do not match, the operation
// will fail with error code 400 (Bad Request).
func (client FileClient) UploadRange(ctx context.Context, rangeParameter string, fileRangeWrite FileRangeWriteType, contentLength int64, body io.ReadSeeker, timeout *int32, contentMD5 *string) (*FileUploadRangeResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.uploadRangePreparer(rangeParameter, fileRangeWrite, contentLength, body, timeout, contentMD5)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.uploadRangeResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*FileUploadRangeResponse), err
}

// uploadRangePreparer prepares the UploadRange request.
func (client FileClient) uploadRangePreparer(rangeParameter string, fileRangeWrite FileRangeWriteType, contentLength int64, body io.ReadSeeker, timeout *int32, contentMD5 *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, body)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	params.Set("comp", "range")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-range", rangeParameter)
	req.Header.Set("x-ms-write", fmt.Sprintf("%v", fileRangeWrite))
	req.Header.Set("Content-Length", fmt.Sprintf("%v", contentLength))
	if contentMD5 != nil {
		req.Header.Set("Content-MD5", *contentMD5)
	}
	req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// uploadRangeResponder handles the response to the UploadRange request.
func (client FileClient) uploadRangeResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	resp.Response().Body.Close()
	return &FileUploadRangeResponse{rawResponse: resp.Response()}, err
}
