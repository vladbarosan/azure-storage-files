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

// ServiceClient is the client for the Service methods of the Storage service.
type ServiceClient struct {
	ManagementClient
}

// NewServiceClient creates an instance of the ServiceClient client.
func NewServiceClient(url url.URL, p pipeline.Pipeline) ServiceClient {
	return ServiceClient{NewManagementClient(url, p)}
}

// GetProperties gets the properties of a storage account's File service, including properties for Storage Analytics
// metrics and CORS (Cross-Origin Resource Sharing) rules.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a>
func (client ServiceClient) GetProperties(ctx context.Context, timeout *int32) (*ServiceProperties, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.getPropertiesPreparer(timeout)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getPropertiesResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ServiceProperties), err
}

// getPropertiesPreparer prepares the GetProperties request.
func (client ServiceClient) getPropertiesPreparer(timeout *int32) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("GET", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	params.Set("restype", "service")
	params.Set("comp", "properties")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// getPropertiesResponder handles the response to the GetProperties request.
func (client ServiceClient) getPropertiesResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ServiceProperties{rawResponse: resp.Response()}
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

// ListSharesSegment the List Shares Segment operation returns a list of the shares and share snapshots under the
// specified account.
//
// prefix is filters the results to return only entries whose name begins with the specified prefix. marker is a string
// value that identifies the portion of the list to be returned with the next list operation. The operation returns a
// marker value within the response body if the list returned was not complete. The marker value may then be used in a
// subsequent call to request the next set of list items. The marker value is opaque to the client. maxresults is
// specifies the maximum number of entries to return. If the request does not specify maxresults, or specifies a value
// greater than 5,000, the server will return up to 5,000 items. include is include this parameter to specify one or
// more datasets to include in the response. timeout is the timeout parameter is expressed in seconds. For more
// information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a>
func (client ServiceClient) ListSharesSegment(ctx context.Context, prefix *string, marker *string, maxresults *int32, include []ListSharesIncludeType, timeout *int32) (*ListSharesResponse, error) {
	if err := validate([]validation{
		{targetValue: maxresults,
			constraints: []constraint{{target: "maxresults", name: null, rule: false,
				chain: []constraint{{target: "maxresults", name: inclusiveMinimum, rule: 1, chain: nil}}}}},
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.listSharesSegmentPreparer(prefix, marker, maxresults, include, timeout)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listSharesSegmentResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ListSharesResponse), err
}

// listSharesSegmentPreparer prepares the ListSharesSegment request.
func (client ServiceClient) listSharesSegmentPreparer(prefix *string, marker *string, maxresults *int32, include []ListSharesIncludeType, timeout *int32) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("GET", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if prefix != nil {
		params.Set("prefix", *prefix)
	}
	if marker != nil {
		params.Set("marker", *marker)
	}
	if maxresults != nil {
		params.Set("maxresults", fmt.Sprintf("%v", *maxresults))
	}
	if include != nil {
		params.Set("include", fmt.Sprintf("%v", joinConst(include, ",")))
	}
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	params.Set("comp", "list")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// listSharesSegmentResponder handles the response to the ListSharesSegment request.
func (client ServiceClient) listSharesSegmentResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ListSharesResponse{rawResponse: resp.Response()}
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

// SetProperties sets properties for a storage account's File service endpoint, including properties for Storage
// Analytics metrics and CORS (Cross-Origin Resource Sharing) rules.
//
// storageServiceProperties is the StorageService properties. timeout is the timeout parameter is expressed in seconds.
// For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a>
func (client ServiceClient) SetProperties(ctx context.Context, storageServiceProperties ServiceProperties, timeout *int32) (*ServiceSetPropertiesResponse, error) {
	if err := validate([]validation{
		{targetValue: storageServiceProperties,
			constraints: []constraint{{target: "storageServiceProperties.HourMetrics", name: null, rule: false,
				chain: []constraint{{target: "storageServiceProperties.HourMetrics.Enabled", name: null, rule: true, chain: nil},
					{target: "storageServiceProperties.HourMetrics.RetentionPolicy", name: null, rule: false,
						chain: []constraint{{target: "storageServiceProperties.HourMetrics.RetentionPolicy.Enabled", name: null, rule: true, chain: nil},
							{target: "storageServiceProperties.HourMetrics.RetentionPolicy.Days", name: null, rule: false,
								chain: []constraint{{target: "storageServiceProperties.HourMetrics.RetentionPolicy.Days", name: inclusiveMaximum, rule: 365, chain: nil},
									{target: "storageServiceProperties.HourMetrics.RetentionPolicy.Days", name: inclusiveMinimum, rule: 1, chain: nil},
								}},
						}},
				}},
				{target: "storageServiceProperties.MinuteMetrics", name: null, rule: false,
					chain: []constraint{{target: "storageServiceProperties.MinuteMetrics.Enabled", name: null, rule: true, chain: nil},
						{target: "storageServiceProperties.MinuteMetrics.RetentionPolicy", name: null, rule: false,
							chain: []constraint{{target: "storageServiceProperties.MinuteMetrics.RetentionPolicy.Enabled", name: null, rule: true, chain: nil},
								{target: "storageServiceProperties.MinuteMetrics.RetentionPolicy.Days", name: null, rule: false,
									chain: []constraint{{target: "storageServiceProperties.MinuteMetrics.RetentionPolicy.Days", name: inclusiveMaximum, rule: 365, chain: nil},
										{target: "storageServiceProperties.MinuteMetrics.RetentionPolicy.Days", name: inclusiveMinimum, rule: 1, chain: nil},
									}},
							}},
					}}}},
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.setPropertiesPreparer(storageServiceProperties, timeout)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.setPropertiesResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ServiceSetPropertiesResponse), err
}

// setPropertiesPreparer prepares the SetProperties request.
func (client ServiceClient) setPropertiesPreparer(storageServiceProperties ServiceProperties, timeout *int32) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	params.Set("restype", "service")
	params.Set("comp", "properties")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	b, err := xml.Marshal(storageServiceProperties)
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

// setPropertiesResponder handles the response to the SetProperties request.
func (client ServiceClient) setPropertiesResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	resp.Response().Body.Close()
	return &ServiceSetPropertiesResponse{rawResponse: resp.Response()}, err
}
