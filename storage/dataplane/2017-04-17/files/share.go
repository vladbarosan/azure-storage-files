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
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io/ioutil"
	"net/http"
	"net/url"
)

// ShareClient is the client for the Share methods of the Storage service.
type ShareClient struct {
	ManagementClient
}

// NewShareClient creates an instance of the ShareClient client.
func NewShareClient(url url.URL, p pipeline.Pipeline) ShareClient {
	return ShareClient{NewManagementClient(url, p)}
}

// Create creates a new share under the specified account. If the share with the same name already exists, the
// operation fails.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> metadata is a name-value pair to associate with a file storage object.
// Metadata names must adhere to the naming rules for C# identifiers. quota is specifies the maximum size of the share,
// in gigabytes.
func (client ShareClient) Create(ctx context.Context, timeout *int32, metadata map[string]string, quota *int32) (*ShareCreateResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}},
		{targetValue: quota,
			constraints: []constraint{{target: "quota", name: null, rule: false,
				chain: []constraint{{target: "quota", name: inclusiveMinimum, rule: 1, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.createPreparer(timeout, metadata, quota)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ShareCreateResponse), err
}

// createPreparer prepares the Create request.
func (client ShareClient) createPreparer(timeout *int32, metadata map[string]string, quota *int32) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	params.Set("restype", "share")
	req.URL.RawQuery = params.Encode()
	if metadata != nil {
		for k, v := range metadata {
			req.Header.Set("x-ms-meta-"+k, v)
		}
	}
	if quota != nil {
		req.Header.Set("x-ms-share-quota", fmt.Sprintf("%v", *quota))
	}
	req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// createResponder handles the response to the Create request.
func (client ShareClient) createResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	resp.Response().Body.Close()
	return &ShareCreateResponse{rawResponse: resp.Response()}, err
}

// CreateSnapshot creates a read-only snapshot of a share.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> metadata is a name-value pair to associate with a file storage object.
// Metadata names must adhere to the naming rules for C# identifiers.
func (client ShareClient) CreateSnapshot(ctx context.Context, timeout *int32, metadata map[string]string) (*ShareCreateSnapshotResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.createSnapshotPreparer(timeout, metadata)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createSnapshotResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ShareCreateSnapshotResponse), err
}

// createSnapshotPreparer prepares the CreateSnapshot request.
func (client ShareClient) createSnapshotPreparer(timeout *int32, metadata map[string]string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	params.Set("restype", "share")
	params.Set("comp", "snapshot")
	req.URL.RawQuery = params.Encode()
	if metadata != nil {
		for k, v := range metadata {
			req.Header.Set("x-ms-meta-"+k, v)
		}
	}
	req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// createSnapshotResponder handles the response to the CreateSnapshot request.
func (client ShareClient) createSnapshotResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	resp.Response().Body.Close()
	return &ShareCreateSnapshotResponse{rawResponse: resp.Response()}, err
}

// Delete operation marks the specified share or share snapshot for deletion. The share or share snapshot and any files
// contained within it are later deleted during garbage collection.
//
// sharesnapshot is the snapshot parameter is an opaque DateTime value that, when present, specifies the share snapshot
// to query. timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> deleteSnapshots is specifies the option include to delete the base share
// and all of its snapshots.
func (client ShareClient) Delete(ctx context.Context, sharesnapshot *string, timeout *int32, deleteSnapshots DeleteSnapshotsOptionType) (*ShareDeleteResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.deletePreparer(sharesnapshot, timeout, deleteSnapshots)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.deleteResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ShareDeleteResponse), err
}

// deletePreparer prepares the Delete request.
func (client ShareClient) deletePreparer(sharesnapshot *string, timeout *int32, deleteSnapshots DeleteSnapshotsOptionType) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("DELETE", client.url, nil)
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
	params.Set("restype", "share")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	if deleteSnapshots != DeleteSnapshotsOptionNone {
		req.Header.Set("x-ms-delete-snapshots", fmt.Sprintf("%v", deleteSnapshots))
	}
	return req, nil
}

// deleteResponder handles the response to the Delete request.
func (client ShareClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	resp.Response().Body.Close()
	return &ShareDeleteResponse{rawResponse: resp.Response()}, err
}

// GetAccessPolicy returns information about stored access policies specified on the share.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a>
func (client ShareClient) GetAccessPolicy(ctx context.Context, timeout *int32) (*SignedIdentifiers, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.getAccessPolicyPreparer(timeout)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getAccessPolicyResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*SignedIdentifiers), err
}

// getAccessPolicyPreparer prepares the GetAccessPolicy request.
func (client ShareClient) getAccessPolicyPreparer(timeout *int32) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("GET", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	params.Set("restype", "share")
	params.Set("comp", "acl")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// getAccessPolicyResponder handles the response to the GetAccessPolicy request.
func (client ShareClient) getAccessPolicyResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &SignedIdentifiers{rawResponse: resp.Response()}
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

// GetProperties returns all user-defined metadata and system properties for the specified share or share snapshot. The
// data returned does not include the share's list of files.
//
// sharesnapshot is the snapshot parameter is an opaque DateTime value that, when present, specifies the share snapshot
// to query. timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a>
func (client ShareClient) GetProperties(ctx context.Context, sharesnapshot *string, timeout *int32) (*ShareGetPropertiesResponse, error) {
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
	return resp.(*ShareGetPropertiesResponse), err
}

// getPropertiesPreparer prepares the GetProperties request.
func (client ShareClient) getPropertiesPreparer(sharesnapshot *string, timeout *int32) (pipeline.Request, error) {
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
	params.Set("restype", "share")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// getPropertiesResponder handles the response to the GetProperties request.
func (client ShareClient) getPropertiesResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	resp.Response().Body.Close()
	return &ShareGetPropertiesResponse{rawResponse: resp.Response()}, err
}

// GetStatistics retrieves statistics related to the share.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a>
func (client ShareClient) GetStatistics(ctx context.Context, timeout *int32) (*ShareStats, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.getStatisticsPreparer(timeout)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getStatisticsResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ShareStats), err
}

// getStatisticsPreparer prepares the GetStatistics request.
func (client ShareClient) getStatisticsPreparer(timeout *int32) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("GET", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	params.Set("restype", "share")
	params.Set("comp", "stats")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// getStatisticsResponder handles the response to the GetStatistics request.
func (client ShareClient) getStatisticsResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ShareStats{rawResponse: resp.Response()}
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

// SetAccessPolicy sets a stored access policy for use with shared access signatures.
//
// shareACL is the ACL for the share. timeout is the timeout parameter is expressed in seconds. For more information,
// see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a>
func (client ShareClient) SetAccessPolicy(ctx context.Context, shareACL []SignedIdentifier, timeout *int32) (*ShareSetAccessPolicyResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.setAccessPolicyPreparer(shareACL, timeout)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.setAccessPolicyResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ShareSetAccessPolicyResponse), err
}

// setAccessPolicyPreparer prepares the SetAccessPolicy request.
func (client ShareClient) setAccessPolicyPreparer(shareACL []SignedIdentifier, timeout *int32) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	params.Set("restype", "share")
	params.Set("comp", "acl")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	b, err := xml.Marshal(SignedIdentifiers{Value: shareACL})
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/xml")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// setAccessPolicyResponder handles the response to the SetAccessPolicy request.
func (client ShareClient) setAccessPolicyResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	resp.Response().Body.Close()
	return &ShareSetAccessPolicyResponse{rawResponse: resp.Response()}, err
}

// SetMetadata sets one or more user-defined name-value pairs for the specified share.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> metadata is a name-value pair to associate with a file storage object.
// Metadata names must adhere to the naming rules for C# identifiers.
func (client ShareClient) SetMetadata(ctx context.Context, timeout *int32, metadata map[string]string) (*ShareSetMetadataResponse, error) {
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
	return resp.(*ShareSetMetadataResponse), err
}

// setMetadataPreparer prepares the SetMetadata request.
func (client ShareClient) setMetadataPreparer(timeout *int32, metadata map[string]string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	params.Set("restype", "share")
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
func (client ShareClient) setMetadataResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	resp.Response().Body.Close()
	return &ShareSetMetadataResponse{rawResponse: resp.Response()}, err
}

// SetQuota sets quota for the specified share.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> quota is specifies the maximum size of the share, in gigabytes.
func (client ShareClient) SetQuota(ctx context.Context, timeout *int32, quota *int32) (*ShareSetQuotaResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}},
		{targetValue: quota,
			constraints: []constraint{{target: "quota", name: null, rule: false,
				chain: []constraint{{target: "quota", name: inclusiveMinimum, rule: 1, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.setQuotaPreparer(timeout, quota)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.setQuotaResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ShareSetQuotaResponse), err
}

// setQuotaPreparer prepares the SetQuota request.
func (client ShareClient) setQuotaPreparer(timeout *int32, quota *int32) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	params.Set("restype", "share")
	params.Set("comp", "properties")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	if quota != nil {
		req.Header.Set("x-ms-share-quota", fmt.Sprintf("%v", *quota))
	}
	return req, nil
}

// setQuotaResponder handles the response to the SetQuota request.
func (client ShareClient) setQuotaResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	resp.Response().Body.Close()
	return &ShareSetQuotaResponse{rawResponse: resp.Response()}, err
}
