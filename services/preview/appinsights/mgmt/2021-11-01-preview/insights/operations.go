package insights

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
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

// OperationsClient is the composite Swagger for Application Insights Management Client
type OperationsClient struct {
	BaseClient
}

// NewOperationsClient creates an instance of the OperationsClient client.
func NewOperationsClient(subscriptionID string) OperationsClient {
	return NewOperationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewOperationsClientWithBaseURI creates an instance of the OperationsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return OperationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists all of the available insights REST API operations.
func (client OperationsClient) List(ctx context.Context) (result OperationListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationsClient.List")
		defer func() {
			sc := -1
			if result.olr.Response.Response != nil {
				sc = result.olr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.OperationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.olr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.OperationsClient", "List", resp, "Failure sending request")
		return
	}

	result.olr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.OperationsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.olr.hasNextLink() && result.olr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client OperationsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2015-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Insights/operations"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client OperationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client OperationsClient) ListResponder(resp *http.Response) (result OperationListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client OperationsClient) listNextResults(ctx context.Context, lastResults OperationListResult) (result OperationListResult, err error) {
	req, err := lastResults.operationListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "insights.OperationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "insights.OperationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.OperationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client OperationsClient) ListComplete(ctx context.Context) (result OperationListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// List1 list the available operations supported by the resource provider.
func (client OperationsClient) List1(ctx context.Context) (result OperationsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationsClient.List1")
		defer func() {
			sc := -1
			if result.olr.Response.Response != nil {
				sc = result.olr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.list1NextResults
	req, err := client.List1Preparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.OperationsClient", "List1", nil, "Failure preparing request")
		return
	}

	resp, err := client.List1Sender(req)
	if err != nil {
		result.olr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.OperationsClient", "List1", resp, "Failure sending request")
		return
	}

	result.olr, err = client.List1Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.OperationsClient", "List1", resp, "Failure responding to request")
		return
	}
	if result.olr.hasNextLink() && result.olr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// List1Preparer prepares the List1 request.
func (client OperationsClient) List1Preparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2020-06-02-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/microsoft.insights/operations"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// List1Sender sends the List1 request. The method will close the
// http.Response Body if it receives an error.
func (client OperationsClient) List1Sender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// List1Responder handles the response to the List1 request. The method always
// closes the http.Response Body.
func (client OperationsClient) List1Responder(resp *http.Response) (result OperationsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// list1NextResults retrieves the next set of results, if any.
func (client OperationsClient) list1NextResults(ctx context.Context, lastResults OperationsListResult) (result OperationsListResult, err error) {
	req, err := lastResults.operationsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "insights.OperationsClient", "list1NextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.List1Sender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "insights.OperationsClient", "list1NextResults", resp, "Failure sending next results request")
	}
	result, err = client.List1Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.OperationsClient", "list1NextResults", resp, "Failure responding to next results request")
	}
	return
}

// List1Complete enumerates all values, automatically crossing page boundaries as required.
func (client OperationsClient) List1Complete(ctx context.Context) (result OperationsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationsClient.List1")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List1(ctx)
	return
}
