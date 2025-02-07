package vmwarecloudsimple

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// DedicatedCloudServicesClient is the description of the new service
type DedicatedCloudServicesClient struct {
	BaseClient
}

// NewDedicatedCloudServicesClient creates an instance of the DedicatedCloudServicesClient client.
func NewDedicatedCloudServicesClient(subscriptionID string, referer string) DedicatedCloudServicesClient {
	return NewDedicatedCloudServicesClientWithBaseURI(DefaultBaseURI, subscriptionID, referer)
}

// NewDedicatedCloudServicesClientWithBaseURI creates an instance of the DedicatedCloudServicesClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewDedicatedCloudServicesClientWithBaseURI(baseURI string, subscriptionID string, referer string) DedicatedCloudServicesClient {
	return DedicatedCloudServicesClient{NewWithBaseURI(baseURI, subscriptionID, referer)}
}

// CreateOrUpdate create dedicate cloud service
// Parameters:
// resourceGroupName - the name of the resource group
// dedicatedCloudServiceName - dedicated cloud Service name
// dedicatedCloudServiceRequest - create Dedicated Cloud Service request
func (client DedicatedCloudServicesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, dedicatedCloudServiceName string, dedicatedCloudServiceRequest DedicatedCloudService) (result DedicatedCloudService, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedCloudServicesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dedicatedCloudServiceName,
			Constraints: []validation.Constraint{{Target: "dedicatedCloudServiceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]([-_.a-zA-Z0-9]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: dedicatedCloudServiceRequest,
			Constraints: []validation.Constraint{{Target: "dedicatedCloudServiceRequest.Location", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "dedicatedCloudServiceRequest.Name", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "dedicatedCloudServiceRequest.Name", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]([-_.a-zA-Z0-9]*[a-zA-Z0-9])?$`, Chain: nil}}},
				{Target: "dedicatedCloudServiceRequest.DedicatedCloudServiceProperties", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "dedicatedCloudServiceRequest.DedicatedCloudServiceProperties.GatewaySubnet", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("vmwarecloudsimple.DedicatedCloudServicesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, dedicatedCloudServiceName, dedicatedCloudServiceRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudServicesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudServicesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudServicesClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DedicatedCloudServicesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, dedicatedCloudServiceName string, dedicatedCloudServiceRequest DedicatedCloudService) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dedicatedCloudServiceName": autorest.Encode("path", dedicatedCloudServiceName),
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	dedicatedCloudServiceRequest.ID = nil
	dedicatedCloudServiceRequest.Name = nil
	dedicatedCloudServiceRequest.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices/{dedicatedCloudServiceName}", pathParameters),
		autorest.WithJSON(dedicatedCloudServiceRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedCloudServicesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DedicatedCloudServicesClient) CreateOrUpdateResponder(resp *http.Response) (result DedicatedCloudService, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete dedicate cloud service
// Parameters:
// resourceGroupName - the name of the resource group
// dedicatedCloudServiceName - dedicated cloud service name
func (client DedicatedCloudServicesClient) Delete(ctx context.Context, resourceGroupName string, dedicatedCloudServiceName string) (result DedicatedCloudServicesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedCloudServicesClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, dedicatedCloudServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudServicesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudServicesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DedicatedCloudServicesClient) DeletePreparer(ctx context.Context, resourceGroupName string, dedicatedCloudServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dedicatedCloudServiceName": autorest.Encode("path", dedicatedCloudServiceName),
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices/{dedicatedCloudServiceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedCloudServicesClient) DeleteSender(req *http.Request) (future DedicatedCloudServicesDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DedicatedCloudServicesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get returns Dedicate Cloud Service
// Parameters:
// resourceGroupName - the name of the resource group
// dedicatedCloudServiceName - dedicated cloud Service name
func (client DedicatedCloudServicesClient) Get(ctx context.Context, resourceGroupName string, dedicatedCloudServiceName string) (result DedicatedCloudService, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedCloudServicesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, dedicatedCloudServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudServicesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudServicesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudServicesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client DedicatedCloudServicesClient) GetPreparer(ctx context.Context, resourceGroupName string, dedicatedCloudServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dedicatedCloudServiceName": autorest.Encode("path", dedicatedCloudServiceName),
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices/{dedicatedCloudServiceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedCloudServicesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DedicatedCloudServicesClient) GetResponder(resp *http.Response) (result DedicatedCloudService, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup returns list of dedicated cloud services within a resource group
// Parameters:
// resourceGroupName - the name of the resource group
// filter - the filter to apply on the list operation
// top - the maximum number of record sets to return
// skipToken - to be used by nextLink implementation
func (client DedicatedCloudServicesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, top *int32, skipToken string) (result DedicatedCloudServiceListResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedCloudServicesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.dcslr.Response.Response != nil {
				sc = result.dcslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, filter, top, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudServicesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.dcslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudServicesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.dcslr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudServicesClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.dcslr.hasNextLink() && result.dcslr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client DedicatedCloudServicesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, filter string, top *int32, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedCloudServicesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client DedicatedCloudServicesClient) ListByResourceGroupResponder(resp *http.Response) (result DedicatedCloudServiceListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client DedicatedCloudServicesClient) listByResourceGroupNextResults(ctx context.Context, lastResults DedicatedCloudServiceListResponse) (result DedicatedCloudServiceListResponse, err error) {
	req, err := lastResults.dedicatedCloudServiceListResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudServicesClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudServicesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudServicesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client DedicatedCloudServicesClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, top *int32, skipToken string) (result DedicatedCloudServiceListResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedCloudServicesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, filter, top, skipToken)
	return
}

// ListBySubscription returns list of dedicated cloud services within a subscription
// Parameters:
// filter - the filter to apply on the list operation
// top - the maximum number of record sets to return
// skipToken - to be used by nextLink implementation
func (client DedicatedCloudServicesClient) ListBySubscription(ctx context.Context, filter string, top *int32, skipToken string) (result DedicatedCloudServiceListResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedCloudServicesClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.dcslr.Response.Response != nil {
				sc = result.dcslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx, filter, top, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudServicesClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.dcslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudServicesClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.dcslr, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudServicesClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}
	if result.dcslr.hasNextLink() && result.dcslr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client DedicatedCloudServicesClient) ListBySubscriptionPreparer(ctx context.Context, filter string, top *int32, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedCloudServicesClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client DedicatedCloudServicesClient) ListBySubscriptionResponder(resp *http.Response) (result DedicatedCloudServiceListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client DedicatedCloudServicesClient) listBySubscriptionNextResults(ctx context.Context, lastResults DedicatedCloudServiceListResponse) (result DedicatedCloudServiceListResponse, err error) {
	req, err := lastResults.dedicatedCloudServiceListResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudServicesClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudServicesClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudServicesClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client DedicatedCloudServicesClient) ListBySubscriptionComplete(ctx context.Context, filter string, top *int32, skipToken string) (result DedicatedCloudServiceListResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedCloudServicesClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx, filter, top, skipToken)
	return
}

// Update patch dedicated cloud service's properties
// Parameters:
// resourceGroupName - the name of the resource group
// dedicatedCloudServiceName - dedicated cloud service name
// dedicatedCloudServiceRequest - patch Dedicated Cloud Service request
func (client DedicatedCloudServicesClient) Update(ctx context.Context, resourceGroupName string, dedicatedCloudServiceName string, dedicatedCloudServiceRequest PatchPayload) (result DedicatedCloudService, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedCloudServicesClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, dedicatedCloudServiceName, dedicatedCloudServiceRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudServicesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudServicesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudServicesClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client DedicatedCloudServicesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, dedicatedCloudServiceName string, dedicatedCloudServiceRequest PatchPayload) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dedicatedCloudServiceName": autorest.Encode("path", dedicatedCloudServiceName),
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices/{dedicatedCloudServiceName}", pathParameters),
		autorest.WithJSON(dedicatedCloudServiceRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedCloudServicesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client DedicatedCloudServicesClient) UpdateResponder(resp *http.Response) (result DedicatedCloudService, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
