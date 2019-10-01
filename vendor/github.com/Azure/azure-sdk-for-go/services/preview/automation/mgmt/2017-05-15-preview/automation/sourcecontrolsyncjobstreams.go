package automation

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"github.com/satori/go.uuid"
	"net/http"
)

// SourceControlSyncJobStreamsClient is the automation Client
type SourceControlSyncJobStreamsClient struct {
	BaseClient
}

// NewSourceControlSyncJobStreamsClient creates an instance of the SourceControlSyncJobStreamsClient client.
func NewSourceControlSyncJobStreamsClient(subscriptionID string) SourceControlSyncJobStreamsClient {
	return NewSourceControlSyncJobStreamsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSourceControlSyncJobStreamsClientWithBaseURI creates an instance of the SourceControlSyncJobStreamsClient client.
func NewSourceControlSyncJobStreamsClientWithBaseURI(baseURI string, subscriptionID string) SourceControlSyncJobStreamsClient {
	return SourceControlSyncJobStreamsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get retrieve a sync job stream identified by stream id.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// sourceControlName - the source control name.
// sourceControlSyncJobID - the source control sync job id.
// streamID - the id of the sync job stream.
func (client SourceControlSyncJobStreamsClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, sourceControlSyncJobID uuid.UUID, streamID string) (result SourceControlSyncJobStreamByID, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SourceControlSyncJobStreamsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.SourceControlSyncJobStreamsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, automationAccountName, sourceControlName, sourceControlSyncJobID, streamID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SourceControlSyncJobStreamsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.SourceControlSyncJobStreamsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SourceControlSyncJobStreamsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SourceControlSyncJobStreamsClient) GetPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, sourceControlSyncJobID uuid.UUID, streamID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName":  autorest.Encode("path", automationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"sourceControlName":      autorest.Encode("path", sourceControlName),
		"sourceControlSyncJobId": autorest.Encode("path", sourceControlSyncJobID),
		"streamId":               autorest.Encode("path", streamID),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-05-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/sourceControls/{sourceControlName}/sourceControlSyncJobs/{sourceControlSyncJobId}/streams/{streamId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SourceControlSyncJobStreamsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SourceControlSyncJobStreamsClient) GetResponder(resp *http.Response) (result SourceControlSyncJobStreamByID, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySyncJob retrieve a list of sync job streams identified by sync job id.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// sourceControlName - the source control name.
// sourceControlSyncJobID - the source control sync job id.
// filter - the filter to apply on the operation.
func (client SourceControlSyncJobStreamsClient) ListBySyncJob(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, sourceControlSyncJobID uuid.UUID, filter string) (result SourceControlSyncJobStreamsListBySyncJobPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SourceControlSyncJobStreamsClient.ListBySyncJob")
		defer func() {
			sc := -1
			if result.scsjslbsj.Response.Response != nil {
				sc = result.scsjslbsj.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.SourceControlSyncJobStreamsClient", "ListBySyncJob", err.Error())
	}

	result.fn = client.listBySyncJobNextResults
	req, err := client.ListBySyncJobPreparer(ctx, resourceGroupName, automationAccountName, sourceControlName, sourceControlSyncJobID, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SourceControlSyncJobStreamsClient", "ListBySyncJob", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySyncJobSender(req)
	if err != nil {
		result.scsjslbsj.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.SourceControlSyncJobStreamsClient", "ListBySyncJob", resp, "Failure sending request")
		return
	}

	result.scsjslbsj, err = client.ListBySyncJobResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SourceControlSyncJobStreamsClient", "ListBySyncJob", resp, "Failure responding to request")
	}

	return
}

// ListBySyncJobPreparer prepares the ListBySyncJob request.
func (client SourceControlSyncJobStreamsClient) ListBySyncJobPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, sourceControlSyncJobID uuid.UUID, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName":  autorest.Encode("path", automationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"sourceControlName":      autorest.Encode("path", sourceControlName),
		"sourceControlSyncJobId": autorest.Encode("path", sourceControlSyncJobID),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-05-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/sourceControls/{sourceControlName}/sourceControlSyncJobs/{sourceControlSyncJobId}/streams", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySyncJobSender sends the ListBySyncJob request. The method will close the
// http.Response Body if it receives an error.
func (client SourceControlSyncJobStreamsClient) ListBySyncJobSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListBySyncJobResponder handles the response to the ListBySyncJob request. The method always
// closes the http.Response Body.
func (client SourceControlSyncJobStreamsClient) ListBySyncJobResponder(resp *http.Response) (result SourceControlSyncJobStreamsListBySyncJob, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySyncJobNextResults retrieves the next set of results, if any.
func (client SourceControlSyncJobStreamsClient) listBySyncJobNextResults(ctx context.Context, lastResults SourceControlSyncJobStreamsListBySyncJob) (result SourceControlSyncJobStreamsListBySyncJob, err error) {
	req, err := lastResults.sourceControlSyncJobStreamsListBySyncJobPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "automation.SourceControlSyncJobStreamsClient", "listBySyncJobNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySyncJobSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "automation.SourceControlSyncJobStreamsClient", "listBySyncJobNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySyncJobResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SourceControlSyncJobStreamsClient", "listBySyncJobNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySyncJobComplete enumerates all values, automatically crossing page boundaries as required.
func (client SourceControlSyncJobStreamsClient) ListBySyncJobComplete(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, sourceControlSyncJobID uuid.UUID, filter string) (result SourceControlSyncJobStreamsListBySyncJobIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SourceControlSyncJobStreamsClient.ListBySyncJob")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySyncJob(ctx, resourceGroupName, automationAccountName, sourceControlName, sourceControlSyncJobID, filter)
	return
}
