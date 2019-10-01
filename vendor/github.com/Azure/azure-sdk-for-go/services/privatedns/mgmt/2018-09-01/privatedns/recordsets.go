package privatedns

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// RecordSetsClient is the the Private DNS Management Client.
type RecordSetsClient struct {
	BaseClient
}

// NewRecordSetsClient creates an instance of the RecordSetsClient client.
func NewRecordSetsClient(subscriptionID string) RecordSetsClient {
	return NewRecordSetsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRecordSetsClientWithBaseURI creates an instance of the RecordSetsClient client.
func NewRecordSetsClientWithBaseURI(baseURI string, subscriptionID string) RecordSetsClient {
	return RecordSetsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a record set within a Private DNS zone.
// Parameters:
// resourceGroupName - the name of the resource group.
// privateZoneName - the name of the Private DNS zone (without a terminating dot).
// recordType - the type of DNS record in this record set. Record sets of type SOA can be updated but not
// created (they are created when the Private DNS zone is created).
// relativeRecordSetName - the name of the record set, relative to the name of the zone.
// parameters - parameters supplied to the CreateOrUpdate operation.
// ifMatch - the ETag of the record set. Omit this value to always overwrite the current record set. Specify
// the last-seen ETag value to prevent accidentally overwriting any concurrent changes.
// ifNoneMatch - set to '*' to allow a new record set to be created, but to prevent updating an existing record
// set. Other values will be ignored.
func (client RecordSetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, relativeRecordSetName string, parameters RecordSet, ifMatch string, ifNoneMatch string) (result RecordSet, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecordSetsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, privateZoneName, recordType, relativeRecordSetName, parameters, ifMatch, ifNoneMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatedns.RecordSetsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "privatedns.RecordSetsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatedns.RecordSetsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RecordSetsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, relativeRecordSetName string, parameters RecordSet, ifMatch string, ifNoneMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateZoneName":       autorest.Encode("path", privateZoneName),
		"recordType":            autorest.Encode("path", recordType),
		"relativeRecordSetName": relativeRecordSetName,
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}/{relativeRecordSetName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	if len(ifNoneMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-None-Match", autorest.String(ifNoneMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client RecordSetsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client RecordSetsClient) CreateOrUpdateResponder(resp *http.Response) (result RecordSet, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a record set from a Private DNS zone. This operation cannot be undone.
// Parameters:
// resourceGroupName - the name of the resource group.
// privateZoneName - the name of the Private DNS zone (without a terminating dot).
// recordType - the type of DNS record in this record set. Record sets of type SOA cannot be deleted (they are
// deleted when the Private DNS zone is deleted).
// relativeRecordSetName - the name of the record set, relative to the name of the zone.
// ifMatch - the ETag of the record set. Omit this value to always delete the current record set. Specify the
// last-seen ETag value to prevent accidentally deleting any concurrent changes.
func (client RecordSetsClient) Delete(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, relativeRecordSetName string, ifMatch string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecordSetsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, privateZoneName, recordType, relativeRecordSetName, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatedns.RecordSetsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "privatedns.RecordSetsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatedns.RecordSetsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RecordSetsClient) DeletePreparer(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, relativeRecordSetName string, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateZoneName":       autorest.Encode("path", privateZoneName),
		"recordType":            autorest.Encode("path", recordType),
		"relativeRecordSetName": relativeRecordSetName,
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}/{relativeRecordSetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RecordSetsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RecordSetsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a record set.
// Parameters:
// resourceGroupName - the name of the resource group.
// privateZoneName - the name of the Private DNS zone (without a terminating dot).
// recordType - the type of DNS record in this record set.
// relativeRecordSetName - the name of the record set, relative to the name of the zone.
func (client RecordSetsClient) Get(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, relativeRecordSetName string) (result RecordSet, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecordSetsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, privateZoneName, recordType, relativeRecordSetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatedns.RecordSetsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "privatedns.RecordSetsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatedns.RecordSetsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RecordSetsClient) GetPreparer(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, relativeRecordSetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateZoneName":       autorest.Encode("path", privateZoneName),
		"recordType":            autorest.Encode("path", recordType),
		"relativeRecordSetName": relativeRecordSetName,
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}/{relativeRecordSetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RecordSetsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RecordSetsClient) GetResponder(resp *http.Response) (result RecordSet, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all record sets in a Private DNS zone.
// Parameters:
// resourceGroupName - the name of the resource group.
// privateZoneName - the name of the Private DNS zone (without a terminating dot).
// top - the maximum number of record sets to return. If not specified, returns up to 100 record sets.
// recordsetnamesuffix - the suffix label of the record set name to be used to filter the record set
// enumeration. If this parameter is specified, the returned enumeration will only contain records that end
// with ".<recordsetnamesuffix>".
func (client RecordSetsClient) List(ctx context.Context, resourceGroupName string, privateZoneName string, top *int32, recordsetnamesuffix string) (result RecordSetListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecordSetsClient.List")
		defer func() {
			sc := -1
			if result.rslr.Response.Response != nil {
				sc = result.rslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, privateZoneName, top, recordsetnamesuffix)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatedns.RecordSetsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "privatedns.RecordSetsClient", "List", resp, "Failure sending request")
		return
	}

	result.rslr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatedns.RecordSetsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client RecordSetsClient) ListPreparer(ctx context.Context, resourceGroupName string, privateZoneName string, top *int32, recordsetnamesuffix string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateZoneName":   autorest.Encode("path", privateZoneName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(recordsetnamesuffix) > 0 {
		queryParameters["$recordsetnamesuffix"] = autorest.Encode("query", recordsetnamesuffix)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/ALL", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RecordSetsClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RecordSetsClient) ListResponder(resp *http.Response) (result RecordSetListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client RecordSetsClient) listNextResults(ctx context.Context, lastResults RecordSetListResult) (result RecordSetListResult, err error) {
	req, err := lastResults.recordSetListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "privatedns.RecordSetsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "privatedns.RecordSetsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatedns.RecordSetsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client RecordSetsClient) ListComplete(ctx context.Context, resourceGroupName string, privateZoneName string, top *int32, recordsetnamesuffix string) (result RecordSetListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecordSetsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, privateZoneName, top, recordsetnamesuffix)
	return
}

// ListByType lists the record sets of a specified type in a Private DNS zone.
// Parameters:
// resourceGroupName - the name of the resource group.
// privateZoneName - the name of the Private DNS zone (without a terminating dot).
// recordType - the type of record sets to enumerate.
// top - the maximum number of record sets to return. If not specified, returns up to 100 record sets.
// recordsetnamesuffix - the suffix label of the record set name to be used to filter the record set
// enumeration. If this parameter is specified, the returned enumeration will only contain records that end
// with ".<recordsetnamesuffix>".
func (client RecordSetsClient) ListByType(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, top *int32, recordsetnamesuffix string) (result RecordSetListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecordSetsClient.ListByType")
		defer func() {
			sc := -1
			if result.rslr.Response.Response != nil {
				sc = result.rslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByTypeNextResults
	req, err := client.ListByTypePreparer(ctx, resourceGroupName, privateZoneName, recordType, top, recordsetnamesuffix)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatedns.RecordSetsClient", "ListByType", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByTypeSender(req)
	if err != nil {
		result.rslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "privatedns.RecordSetsClient", "ListByType", resp, "Failure sending request")
		return
	}

	result.rslr, err = client.ListByTypeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatedns.RecordSetsClient", "ListByType", resp, "Failure responding to request")
	}

	return
}

// ListByTypePreparer prepares the ListByType request.
func (client RecordSetsClient) ListByTypePreparer(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, top *int32, recordsetnamesuffix string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateZoneName":   autorest.Encode("path", privateZoneName),
		"recordType":        autorest.Encode("path", recordType),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(recordsetnamesuffix) > 0 {
		queryParameters["$recordsetnamesuffix"] = autorest.Encode("query", recordsetnamesuffix)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByTypeSender sends the ListByType request. The method will close the
// http.Response Body if it receives an error.
func (client RecordSetsClient) ListByTypeSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByTypeResponder handles the response to the ListByType request. The method always
// closes the http.Response Body.
func (client RecordSetsClient) ListByTypeResponder(resp *http.Response) (result RecordSetListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByTypeNextResults retrieves the next set of results, if any.
func (client RecordSetsClient) listByTypeNextResults(ctx context.Context, lastResults RecordSetListResult) (result RecordSetListResult, err error) {
	req, err := lastResults.recordSetListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "privatedns.RecordSetsClient", "listByTypeNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByTypeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "privatedns.RecordSetsClient", "listByTypeNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByTypeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatedns.RecordSetsClient", "listByTypeNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByTypeComplete enumerates all values, automatically crossing page boundaries as required.
func (client RecordSetsClient) ListByTypeComplete(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, top *int32, recordsetnamesuffix string) (result RecordSetListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecordSetsClient.ListByType")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByType(ctx, resourceGroupName, privateZoneName, recordType, top, recordsetnamesuffix)
	return
}

// Update updates a record set within a Private DNS zone.
// Parameters:
// resourceGroupName - the name of the resource group.
// privateZoneName - the name of the Private DNS zone (without a terminating dot).
// recordType - the type of DNS record in this record set.
// relativeRecordSetName - the name of the record set, relative to the name of the zone.
// parameters - parameters supplied to the Update operation.
// ifMatch - the ETag of the record set. Omit this value to always overwrite the current record set. Specify
// the last-seen ETag value to prevent accidentally overwriting concurrent changes.
func (client RecordSetsClient) Update(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, relativeRecordSetName string, parameters RecordSet, ifMatch string) (result RecordSet, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecordSetsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, privateZoneName, recordType, relativeRecordSetName, parameters, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatedns.RecordSetsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "privatedns.RecordSetsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatedns.RecordSetsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client RecordSetsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, relativeRecordSetName string, parameters RecordSet, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateZoneName":       autorest.Encode("path", privateZoneName),
		"recordType":            autorest.Encode("path", recordType),
		"relativeRecordSetName": relativeRecordSetName,
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}/{relativeRecordSetName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client RecordSetsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client RecordSetsClient) UpdateResponder(resp *http.Response) (result RecordSet, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
