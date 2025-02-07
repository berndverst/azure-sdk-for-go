package eventgrid

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

// VerifiedPartnersClient is the azure EventGrid Management Client
type VerifiedPartnersClient struct {
	BaseClient
}

// NewVerifiedPartnersClient creates an instance of the VerifiedPartnersClient client.
func NewVerifiedPartnersClient(subscriptionID string) VerifiedPartnersClient {
	return NewVerifiedPartnersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVerifiedPartnersClientWithBaseURI creates an instance of the VerifiedPartnersClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewVerifiedPartnersClientWithBaseURI(baseURI string, subscriptionID string) VerifiedPartnersClient {
	return VerifiedPartnersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get properties of a verified partner.
// Parameters:
// verifiedPartnerName - name of the verified partner.
func (client VerifiedPartnersClient) Get(ctx context.Context, verifiedPartnerName string) (result VerifiedPartner, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VerifiedPartnersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, verifiedPartnerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.VerifiedPartnersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.VerifiedPartnersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.VerifiedPartnersClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client VerifiedPartnersClient) GetPreparer(ctx context.Context, verifiedPartnerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"verifiedPartnerName": autorest.Encode("path", verifiedPartnerName),
	}

	const APIVersion = "2021-10-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.EventGrid/verifiedPartners/{verifiedPartnerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VerifiedPartnersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VerifiedPartnersClient) GetResponder(resp *http.Response) (result VerifiedPartner, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get a list of all verified partners.
// Parameters:
// filter - the query used to filter the search results using OData syntax. Filtering is permitted on the
// 'name' property only and with limited number of OData operations. These operations are: the 'contains'
// function as well as the following logical operations: not, and, or, eq (for equal), and ne (for not equal).
// No arithmetic operations are supported. The following is a valid filter example: $filter=contains(namE,
// 'PATTERN') and name ne 'PATTERN-1'. The following is not a valid filter example: $filter=location eq
// 'westus'.
// top - the number of results to return per page for the list operation. Valid range for top parameter is 1 to
// 100. If not specified, the default number of results to be returned is 20 items per page.
func (client VerifiedPartnersClient) List(ctx context.Context, filter string, top *int32) (result VerifiedPartnersListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VerifiedPartnersClient.List")
		defer func() {
			sc := -1
			if result.vplr.Response.Response != nil {
				sc = result.vplr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, filter, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.VerifiedPartnersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.vplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.VerifiedPartnersClient", "List", resp, "Failure sending request")
		return
	}

	result.vplr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.VerifiedPartnersClient", "List", resp, "Failure responding to request")
		return
	}
	if result.vplr.hasNextLink() && result.vplr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client VerifiedPartnersClient) ListPreparer(ctx context.Context, filter string, top *int32) (*http.Request, error) {
	const APIVersion = "2021-10-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.EventGrid/verifiedPartners"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VerifiedPartnersClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VerifiedPartnersClient) ListResponder(resp *http.Response) (result VerifiedPartnersListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client VerifiedPartnersClient) listNextResults(ctx context.Context, lastResults VerifiedPartnersListResult) (result VerifiedPartnersListResult, err error) {
	req, err := lastResults.verifiedPartnersListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "eventgrid.VerifiedPartnersClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "eventgrid.VerifiedPartnersClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.VerifiedPartnersClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client VerifiedPartnersClient) ListComplete(ctx context.Context, filter string, top *int32) (result VerifiedPartnersListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VerifiedPartnersClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, filter, top)
	return
}
