//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armaad

// ARMProxyResource - The resource model definition for a ARM proxy resource. It will have everything other than required
// location and tags
type ARMProxyResource struct {
	// READ-ONLY; The unique resource identifier of the Azure AD PrivateLink Policy.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the Azure AD PrivateLink Policy.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of Azure resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// AzureResourceBase - Common properties for all Azure resources.
type AzureResourceBase struct {
	// Name of this resource.
	Name *string `json:"name,omitempty"`

	// READ-ONLY; String Id used to locate any resource on Azure.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Type of this resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ErrorDefinition - Error definition.
type ErrorDefinition struct {
	// READ-ONLY; Service specific error code which serves as the substatus for the HTTP error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; Internal error details.
	Details []*ErrorDefinition `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; Description of the error.
	Message *string `json:"message,omitempty" azure:"ro"`
}

// ErrorResponse - Error response.
type ErrorResponse struct {
	// The error details.
	Error *ErrorDefinition `json:"error,omitempty"`
}

// PrivateEndpoint - Private endpoint object properties.
type PrivateEndpoint struct {
	// Full identifier of the private endpoint resource.
	ID *string `json:"id,omitempty"`
}

// PrivateEndpointConnection - Private endpoint connection resource.
type PrivateEndpointConnection struct {
	// Resource properties.
	Properties *PrivateEndpointConnectionProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// PrivateEndpointConnectionListResult - A list of private link resources
type PrivateEndpointConnectionListResult struct {
	// URL to next page of results
	NextLink *string `json:"nextLink,omitempty"`

	// Array of private link resources
	Value []*PrivateEndpointConnection `json:"value,omitempty"`
}

// PrivateEndpointConnectionProperties - Properties of the private endpoint connection resource.
type PrivateEndpointConnectionProperties struct {
	// Properties of the private endpoint object.
	PrivateEndpoint *PrivateEndpoint `json:"privateEndpoint,omitempty"`

	// Updated tag information to set into the PrivateLinkConnection instance.
	PrivateLinkConnectionTags *TagsResource `json:"privateLinkConnectionTags,omitempty"`

	// Approval state of the private link connection.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `json:"privateLinkServiceConnectionState,omitempty"`

	// READ-ONLY; Provisioning state of the private endpoint connection.
	ProvisioningState *PrivateEndpointConnectionProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// PrivateEndpointConnectionsClientBeginCreateOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginCreate
// method.
type PrivateEndpointConnectionsClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PrivateEndpointConnectionsClientBeginDeleteOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginDelete
// method.
type PrivateEndpointConnectionsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PrivateEndpointConnectionsClientGetOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Get
// method.
type PrivateEndpointConnectionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientListByPolicyNameOptions contains the optional parameters for the PrivateEndpointConnectionsClient.ListByPolicyName
// method.
type PrivateEndpointConnectionsClientListByPolicyNameOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkForAzureAdClientBeginCreateOptions contains the optional parameters for the PrivateLinkForAzureAdClient.BeginCreate
// method.
type PrivateLinkForAzureAdClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PrivateLinkForAzureAdClientDeleteOptions contains the optional parameters for the PrivateLinkForAzureAdClient.Delete method.
type PrivateLinkForAzureAdClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkForAzureAdClientGetOptions contains the optional parameters for the PrivateLinkForAzureAdClient.Get method.
type PrivateLinkForAzureAdClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkForAzureAdClientListBySubscriptionOptions contains the optional parameters for the PrivateLinkForAzureAdClient.ListBySubscription
// method.
type PrivateLinkForAzureAdClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkForAzureAdClientListOptions contains the optional parameters for the PrivateLinkForAzureAdClient.List method.
type PrivateLinkForAzureAdClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkForAzureAdClientUpdateOptions contains the optional parameters for the PrivateLinkForAzureAdClient.Update method.
type PrivateLinkForAzureAdClientUpdateOptions struct {
	// Private Link Policy resource with the tags to be updated.
	PrivateLinkPolicy *PrivateLinkPolicyUpdateParameter
}

// PrivateLinkPolicy - PrivateLink Policy configuration object.
type PrivateLinkPolicy struct {
	// Flag indicating whether all tenants are allowed
	AllTenants *bool `json:"allTenants,omitempty"`

	// Name of this resource.
	Name *string `json:"name,omitempty"`

	// Guid of the owner tenant
	OwnerTenantID *string `json:"ownerTenantId,omitempty"`

	// Name of the resource group
	ResourceGroup *string `json:"resourceGroup,omitempty"`

	// Name of the private link policy resource
	ResourceName *string `json:"resourceName,omitempty"`

	// Subscription Identifier
	SubscriptionID *string `json:"subscriptionId,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// The list of tenantIds.
	Tenants []*string `json:"tenants,omitempty"`

	// READ-ONLY; String Id used to locate any resource on Azure.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Type of this resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// PrivateLinkPolicyListResult - A list of private link policies
type PrivateLinkPolicyListResult struct {
	// The link used to get the next page of operations.
	NextLink *string `json:"nextLink,omitempty"`

	// Array of private link policies
	Value []*PrivateLinkPolicy `json:"value,omitempty"`
}

// PrivateLinkPolicyUpdateParameter - private Link policy parameters to be updated.
type PrivateLinkPolicyUpdateParameter struct {
	// Resource tags to be updated.
	Tags map[string]*string `json:"tags,omitempty"`
}

// PrivateLinkResource - A private link resource
type PrivateLinkResource struct {
	// Resource properties.
	Properties *PrivateLinkResourceProperties `json:"properties,omitempty"`

	// READ-ONLY; The unique resource identifier of the Azure AD PrivateLink Policy.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the Azure AD PrivateLink Policy.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of Azure resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// PrivateLinkResourceListResult - A list of private link resources
type PrivateLinkResourceListResult struct {
	// The link used to get the next page of operations.
	NextLink *string `json:"nextLink,omitempty"`

	// Array of private link resources
	Value []*PrivateLinkResource `json:"value,omitempty"`
}

// PrivateLinkResourceProperties - Properties of a private link resource.
type PrivateLinkResourceProperties struct {
	// READ-ONLY; The private link resource group id.
	GroupID *string `json:"groupId,omitempty" azure:"ro"`

	// READ-ONLY; The private link resource required member names.
	RequiredMembers []*string `json:"requiredMembers,omitempty" azure:"ro"`
}

// PrivateLinkResourcesClientGetOptions contains the optional parameters for the PrivateLinkResourcesClient.Get method.
type PrivateLinkResourcesClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResourcesClientListByPrivateLinkPolicyOptions contains the optional parameters for the PrivateLinkResourcesClient.ListByPrivateLinkPolicy
// method.
type PrivateLinkResourcesClientListByPrivateLinkPolicyOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkServiceConnectionState - An object that represents the approval state of the private link connection.
type PrivateLinkServiceConnectionState struct {
	// A message indicating if changes on the service provider require any updates on the consumer.
	ActionsRequired *string `json:"actionsRequired,omitempty"`

	// The reason for approval or rejection.
	Description *string `json:"description,omitempty"`

	// Indicates whether the connection has been approved, rejected or removed by the given policy owner.
	Status *PrivateEndpointServiceConnectionStatus `json:"status,omitempty"`
}

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a
// location
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// TagsResource - A container holding only the Tags for a resource, allowing the user to update the tags on a PrivateLinkConnection
// instance.
type TagsResource struct {
	// Resource tags
	Tags map[string]*string `json:"tags,omitempty"`
}
