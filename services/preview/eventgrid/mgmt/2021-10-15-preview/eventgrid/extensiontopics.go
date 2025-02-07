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

// ExtensionTopicsClient is the azure EventGrid Management Client
type ExtensionTopicsClient struct {
	BaseClient
}

// NewExtensionTopicsClient creates an instance of the ExtensionTopicsClient client.
func NewExtensionTopicsClient(subscriptionID string) ExtensionTopicsClient {
	return NewExtensionTopicsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExtensionTopicsClientWithBaseURI creates an instance of the ExtensionTopicsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewExtensionTopicsClientWithBaseURI(baseURI string, subscriptionID string) ExtensionTopicsClient {
	return ExtensionTopicsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get the properties of an extension topic.
// Parameters:
// scope - the identifier of the resource to which extension topic is queried. The scope can be a subscription,
// or a resource group, or a top level resource belonging to a resource provider namespace. For example, use
// '/subscriptions/{subscriptionId}/' for a subscription,
// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for a resource group, and
// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}'
// for Azure resource.
func (client ExtensionTopicsClient) Get(ctx context.Context, scope string) (result ExtensionTopic, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExtensionTopicsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, scope)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.ExtensionTopicsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.ExtensionTopicsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.ExtensionTopicsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ExtensionTopicsClient) GetPreparer(ctx context.Context, scope string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope": autorest.Encode("path", scope),
	}

	const APIVersion = "2021-10-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.EventGrid/extensionTopics/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ExtensionTopicsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ExtensionTopicsClient) GetResponder(resp *http.Response) (result ExtensionTopic, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
