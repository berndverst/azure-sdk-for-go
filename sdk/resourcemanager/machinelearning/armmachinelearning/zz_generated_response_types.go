//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearning

// BatchDeploymentsClientCreateOrUpdateResponse contains the response from method BatchDeploymentsClient.CreateOrUpdate.
type BatchDeploymentsClientCreateOrUpdateResponse struct {
	BatchDeployment
}

// BatchDeploymentsClientDeleteResponse contains the response from method BatchDeploymentsClient.Delete.
type BatchDeploymentsClientDeleteResponse struct {
	// placeholder for future response values
}

// BatchDeploymentsClientGetResponse contains the response from method BatchDeploymentsClient.Get.
type BatchDeploymentsClientGetResponse struct {
	BatchDeployment
}

// BatchDeploymentsClientListResponse contains the response from method BatchDeploymentsClient.List.
type BatchDeploymentsClientListResponse struct {
	BatchDeploymentTrackedResourceArmPaginatedResult
}

// BatchDeploymentsClientUpdateResponse contains the response from method BatchDeploymentsClient.Update.
type BatchDeploymentsClientUpdateResponse struct {
	BatchDeployment
}

// BatchEndpointsClientCreateOrUpdateResponse contains the response from method BatchEndpointsClient.CreateOrUpdate.
type BatchEndpointsClientCreateOrUpdateResponse struct {
	BatchEndpoint
}

// BatchEndpointsClientDeleteResponse contains the response from method BatchEndpointsClient.Delete.
type BatchEndpointsClientDeleteResponse struct {
	// placeholder for future response values
}

// BatchEndpointsClientGetResponse contains the response from method BatchEndpointsClient.Get.
type BatchEndpointsClientGetResponse struct {
	BatchEndpoint
}

// BatchEndpointsClientListKeysResponse contains the response from method BatchEndpointsClient.ListKeys.
type BatchEndpointsClientListKeysResponse struct {
	EndpointAuthKeys
}

// BatchEndpointsClientListResponse contains the response from method BatchEndpointsClient.List.
type BatchEndpointsClientListResponse struct {
	BatchEndpointTrackedResourceArmPaginatedResult
}

// BatchEndpointsClientUpdateResponse contains the response from method BatchEndpointsClient.Update.
type BatchEndpointsClientUpdateResponse struct {
	BatchEndpoint
}

// CodeContainersClientCreateOrUpdateResponse contains the response from method CodeContainersClient.CreateOrUpdate.
type CodeContainersClientCreateOrUpdateResponse struct {
	CodeContainer
}

// CodeContainersClientDeleteResponse contains the response from method CodeContainersClient.Delete.
type CodeContainersClientDeleteResponse struct {
	// placeholder for future response values
}

// CodeContainersClientGetResponse contains the response from method CodeContainersClient.Get.
type CodeContainersClientGetResponse struct {
	CodeContainer
}

// CodeContainersClientListResponse contains the response from method CodeContainersClient.List.
type CodeContainersClientListResponse struct {
	CodeContainerResourceArmPaginatedResult
}

// CodeVersionsClientCreateOrUpdateResponse contains the response from method CodeVersionsClient.CreateOrUpdate.
type CodeVersionsClientCreateOrUpdateResponse struct {
	CodeVersion
}

// CodeVersionsClientDeleteResponse contains the response from method CodeVersionsClient.Delete.
type CodeVersionsClientDeleteResponse struct {
	// placeholder for future response values
}

// CodeVersionsClientGetResponse contains the response from method CodeVersionsClient.Get.
type CodeVersionsClientGetResponse struct {
	CodeVersion
}

// CodeVersionsClientListResponse contains the response from method CodeVersionsClient.List.
type CodeVersionsClientListResponse struct {
	CodeVersionResourceArmPaginatedResult
}

// ComponentContainersClientCreateOrUpdateResponse contains the response from method ComponentContainersClient.CreateOrUpdate.
type ComponentContainersClientCreateOrUpdateResponse struct {
	ComponentContainer
}

// ComponentContainersClientDeleteResponse contains the response from method ComponentContainersClient.Delete.
type ComponentContainersClientDeleteResponse struct {
	// placeholder for future response values
}

// ComponentContainersClientGetResponse contains the response from method ComponentContainersClient.Get.
type ComponentContainersClientGetResponse struct {
	ComponentContainer
}

// ComponentContainersClientListResponse contains the response from method ComponentContainersClient.List.
type ComponentContainersClientListResponse struct {
	ComponentContainerResourceArmPaginatedResult
}

// ComponentVersionsClientCreateOrUpdateResponse contains the response from method ComponentVersionsClient.CreateOrUpdate.
type ComponentVersionsClientCreateOrUpdateResponse struct {
	ComponentVersion
}

// ComponentVersionsClientDeleteResponse contains the response from method ComponentVersionsClient.Delete.
type ComponentVersionsClientDeleteResponse struct {
	// placeholder for future response values
}

// ComponentVersionsClientGetResponse contains the response from method ComponentVersionsClient.Get.
type ComponentVersionsClientGetResponse struct {
	ComponentVersion
}

// ComponentVersionsClientListResponse contains the response from method ComponentVersionsClient.List.
type ComponentVersionsClientListResponse struct {
	ComponentVersionResourceArmPaginatedResult
}

// ComputeClientCreateOrUpdateResponse contains the response from method ComputeClient.CreateOrUpdate.
type ComputeClientCreateOrUpdateResponse struct {
	ComputeResource
}

// ComputeClientDeleteResponse contains the response from method ComputeClient.Delete.
type ComputeClientDeleteResponse struct {
	// placeholder for future response values
}

// ComputeClientGetResponse contains the response from method ComputeClient.Get.
type ComputeClientGetResponse struct {
	ComputeResource
}

// ComputeClientListKeysResponse contains the response from method ComputeClient.ListKeys.
type ComputeClientListKeysResponse struct {
	ComputeSecretsClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ComputeClientListKeysResponse.
func (c *ComputeClientListKeysResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalComputeSecretsClassification(data)
	if err != nil {
		return err
	}
	c.ComputeSecretsClassification = res
	return nil
}

// ComputeClientListNodesResponse contains the response from method ComputeClient.ListNodes.
type ComputeClientListNodesResponse struct {
	AmlComputeNodesInformation
}

// ComputeClientListResponse contains the response from method ComputeClient.List.
type ComputeClientListResponse struct {
	PaginatedComputeResourcesList
}

// ComputeClientRestartResponse contains the response from method ComputeClient.Restart.
type ComputeClientRestartResponse struct {
	// placeholder for future response values
}

// ComputeClientStartResponse contains the response from method ComputeClient.Start.
type ComputeClientStartResponse struct {
	// placeholder for future response values
}

// ComputeClientStopResponse contains the response from method ComputeClient.Stop.
type ComputeClientStopResponse struct {
	// placeholder for future response values
}

// ComputeClientUpdateResponse contains the response from method ComputeClient.Update.
type ComputeClientUpdateResponse struct {
	ComputeResource
}

// DataContainersClientCreateOrUpdateResponse contains the response from method DataContainersClient.CreateOrUpdate.
type DataContainersClientCreateOrUpdateResponse struct {
	DataContainer
}

// DataContainersClientDeleteResponse contains the response from method DataContainersClient.Delete.
type DataContainersClientDeleteResponse struct {
	// placeholder for future response values
}

// DataContainersClientGetResponse contains the response from method DataContainersClient.Get.
type DataContainersClientGetResponse struct {
	DataContainer
}

// DataContainersClientListResponse contains the response from method DataContainersClient.List.
type DataContainersClientListResponse struct {
	DataContainerResourceArmPaginatedResult
}

// DataVersionsClientCreateOrUpdateResponse contains the response from method DataVersionsClient.CreateOrUpdate.
type DataVersionsClientCreateOrUpdateResponse struct {
	DataVersionBase
}

// DataVersionsClientDeleteResponse contains the response from method DataVersionsClient.Delete.
type DataVersionsClientDeleteResponse struct {
	// placeholder for future response values
}

// DataVersionsClientGetResponse contains the response from method DataVersionsClient.Get.
type DataVersionsClientGetResponse struct {
	DataVersionBase
}

// DataVersionsClientListResponse contains the response from method DataVersionsClient.List.
type DataVersionsClientListResponse struct {
	DataVersionBaseResourceArmPaginatedResult
}

// DatastoresClientCreateOrUpdateResponse contains the response from method DatastoresClient.CreateOrUpdate.
type DatastoresClientCreateOrUpdateResponse struct {
	Datastore
}

// DatastoresClientDeleteResponse contains the response from method DatastoresClient.Delete.
type DatastoresClientDeleteResponse struct {
	// placeholder for future response values
}

// DatastoresClientGetResponse contains the response from method DatastoresClient.Get.
type DatastoresClientGetResponse struct {
	Datastore
}

// DatastoresClientListResponse contains the response from method DatastoresClient.List.
type DatastoresClientListResponse struct {
	DatastoreResourceArmPaginatedResult
}

// DatastoresClientListSecretsResponse contains the response from method DatastoresClient.ListSecrets.
type DatastoresClientListSecretsResponse struct {
	DatastoreSecretsClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DatastoresClientListSecretsResponse.
func (d *DatastoresClientListSecretsResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalDatastoreSecretsClassification(data)
	if err != nil {
		return err
	}
	d.DatastoreSecretsClassification = res
	return nil
}

// EnvironmentContainersClientCreateOrUpdateResponse contains the response from method EnvironmentContainersClient.CreateOrUpdate.
type EnvironmentContainersClientCreateOrUpdateResponse struct {
	EnvironmentContainer
}

// EnvironmentContainersClientDeleteResponse contains the response from method EnvironmentContainersClient.Delete.
type EnvironmentContainersClientDeleteResponse struct {
	// placeholder for future response values
}

// EnvironmentContainersClientGetResponse contains the response from method EnvironmentContainersClient.Get.
type EnvironmentContainersClientGetResponse struct {
	EnvironmentContainer
}

// EnvironmentContainersClientListResponse contains the response from method EnvironmentContainersClient.List.
type EnvironmentContainersClientListResponse struct {
	EnvironmentContainerResourceArmPaginatedResult
}

// EnvironmentVersionsClientCreateOrUpdateResponse contains the response from method EnvironmentVersionsClient.CreateOrUpdate.
type EnvironmentVersionsClientCreateOrUpdateResponse struct {
	EnvironmentVersion
}

// EnvironmentVersionsClientDeleteResponse contains the response from method EnvironmentVersionsClient.Delete.
type EnvironmentVersionsClientDeleteResponse struct {
	// placeholder for future response values
}

// EnvironmentVersionsClientGetResponse contains the response from method EnvironmentVersionsClient.Get.
type EnvironmentVersionsClientGetResponse struct {
	EnvironmentVersion
}

// EnvironmentVersionsClientListResponse contains the response from method EnvironmentVersionsClient.List.
type EnvironmentVersionsClientListResponse struct {
	EnvironmentVersionResourceArmPaginatedResult
}

// JobsClientCancelResponse contains the response from method JobsClient.Cancel.
type JobsClientCancelResponse struct {
	// placeholder for future response values
}

// JobsClientCreateOrUpdateResponse contains the response from method JobsClient.CreateOrUpdate.
type JobsClientCreateOrUpdateResponse struct {
	JobBase
}

// JobsClientDeleteResponse contains the response from method JobsClient.Delete.
type JobsClientDeleteResponse struct {
	// placeholder for future response values
}

// JobsClientGetResponse contains the response from method JobsClient.Get.
type JobsClientGetResponse struct {
	JobBase
}

// JobsClientListResponse contains the response from method JobsClient.List.
type JobsClientListResponse struct {
	JobBaseResourceArmPaginatedResult
}

// ModelContainersClientCreateOrUpdateResponse contains the response from method ModelContainersClient.CreateOrUpdate.
type ModelContainersClientCreateOrUpdateResponse struct {
	ModelContainer
}

// ModelContainersClientDeleteResponse contains the response from method ModelContainersClient.Delete.
type ModelContainersClientDeleteResponse struct {
	// placeholder for future response values
}

// ModelContainersClientGetResponse contains the response from method ModelContainersClient.Get.
type ModelContainersClientGetResponse struct {
	ModelContainer
}

// ModelContainersClientListResponse contains the response from method ModelContainersClient.List.
type ModelContainersClientListResponse struct {
	ModelContainerResourceArmPaginatedResult
}

// ModelVersionsClientCreateOrUpdateResponse contains the response from method ModelVersionsClient.CreateOrUpdate.
type ModelVersionsClientCreateOrUpdateResponse struct {
	ModelVersion
}

// ModelVersionsClientDeleteResponse contains the response from method ModelVersionsClient.Delete.
type ModelVersionsClientDeleteResponse struct {
	// placeholder for future response values
}

// ModelVersionsClientGetResponse contains the response from method ModelVersionsClient.Get.
type ModelVersionsClientGetResponse struct {
	ModelVersion
}

// ModelVersionsClientListResponse contains the response from method ModelVersionsClient.List.
type ModelVersionsClientListResponse struct {
	ModelVersionResourceArmPaginatedResult
}

// OnlineDeploymentsClientCreateOrUpdateResponse contains the response from method OnlineDeploymentsClient.CreateOrUpdate.
type OnlineDeploymentsClientCreateOrUpdateResponse struct {
	OnlineDeployment
}

// OnlineDeploymentsClientDeleteResponse contains the response from method OnlineDeploymentsClient.Delete.
type OnlineDeploymentsClientDeleteResponse struct {
	// placeholder for future response values
}

// OnlineDeploymentsClientGetLogsResponse contains the response from method OnlineDeploymentsClient.GetLogs.
type OnlineDeploymentsClientGetLogsResponse struct {
	DeploymentLogs
}

// OnlineDeploymentsClientGetResponse contains the response from method OnlineDeploymentsClient.Get.
type OnlineDeploymentsClientGetResponse struct {
	OnlineDeployment
}

// OnlineDeploymentsClientListResponse contains the response from method OnlineDeploymentsClient.List.
type OnlineDeploymentsClientListResponse struct {
	OnlineDeploymentTrackedResourceArmPaginatedResult
}

// OnlineDeploymentsClientListSKUsResponse contains the response from method OnlineDeploymentsClient.ListSKUs.
type OnlineDeploymentsClientListSKUsResponse struct {
	SKUResourceArmPaginatedResult
}

// OnlineDeploymentsClientUpdateResponse contains the response from method OnlineDeploymentsClient.Update.
type OnlineDeploymentsClientUpdateResponse struct {
	OnlineDeployment
}

// OnlineEndpointsClientCreateOrUpdateResponse contains the response from method OnlineEndpointsClient.CreateOrUpdate.
type OnlineEndpointsClientCreateOrUpdateResponse struct {
	OnlineEndpoint
}

// OnlineEndpointsClientDeleteResponse contains the response from method OnlineEndpointsClient.Delete.
type OnlineEndpointsClientDeleteResponse struct {
	// placeholder for future response values
}

// OnlineEndpointsClientGetResponse contains the response from method OnlineEndpointsClient.Get.
type OnlineEndpointsClientGetResponse struct {
	OnlineEndpoint
}

// OnlineEndpointsClientGetTokenResponse contains the response from method OnlineEndpointsClient.GetToken.
type OnlineEndpointsClientGetTokenResponse struct {
	EndpointAuthToken
}

// OnlineEndpointsClientListKeysResponse contains the response from method OnlineEndpointsClient.ListKeys.
type OnlineEndpointsClientListKeysResponse struct {
	EndpointAuthKeys
}

// OnlineEndpointsClientListResponse contains the response from method OnlineEndpointsClient.List.
type OnlineEndpointsClientListResponse struct {
	OnlineEndpointTrackedResourceArmPaginatedResult
}

// OnlineEndpointsClientRegenerateKeysResponse contains the response from method OnlineEndpointsClient.RegenerateKeys.
type OnlineEndpointsClientRegenerateKeysResponse struct {
	// placeholder for future response values
}

// OnlineEndpointsClientUpdateResponse contains the response from method OnlineEndpointsClient.Update.
type OnlineEndpointsClientUpdateResponse struct {
	OnlineEndpoint
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	AmlOperationListResult
}

// PrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionsClient.CreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.Delete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListResponse contains the response from method PrivateEndpointConnectionsClient.List.
type PrivateEndpointConnectionsClientListResponse struct {
	PrivateEndpointConnectionListResult
}

// PrivateLinkResourcesClientListResponse contains the response from method PrivateLinkResourcesClient.List.
type PrivateLinkResourcesClientListResponse struct {
	PrivateLinkResourceListResult
}

// QuotasClientListResponse contains the response from method QuotasClient.List.
type QuotasClientListResponse struct {
	ListWorkspaceQuotas
}

// QuotasClientUpdateResponse contains the response from method QuotasClient.Update.
type QuotasClientUpdateResponse struct {
	UpdateWorkspaceQuotasResult
}

// UsagesClientListResponse contains the response from method UsagesClient.List.
type UsagesClientListResponse struct {
	ListUsagesResult
}

// VirtualMachineSizesClientListResponse contains the response from method VirtualMachineSizesClient.List.
type VirtualMachineSizesClientListResponse struct {
	VirtualMachineSizeListResult
}

// WorkspaceConnectionsClientCreateResponse contains the response from method WorkspaceConnectionsClient.Create.
type WorkspaceConnectionsClientCreateResponse struct {
	WorkspaceConnectionPropertiesV2BasicResource
}

// WorkspaceConnectionsClientDeleteResponse contains the response from method WorkspaceConnectionsClient.Delete.
type WorkspaceConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// WorkspaceConnectionsClientGetResponse contains the response from method WorkspaceConnectionsClient.Get.
type WorkspaceConnectionsClientGetResponse struct {
	WorkspaceConnectionPropertiesV2BasicResource
}

// WorkspaceConnectionsClientListResponse contains the response from method WorkspaceConnectionsClient.List.
type WorkspaceConnectionsClientListResponse struct {
	WorkspaceConnectionPropertiesV2BasicResourceArmPaginatedResult
}

// WorkspaceFeaturesClientListResponse contains the response from method WorkspaceFeaturesClient.List.
type WorkspaceFeaturesClientListResponse struct {
	ListAmlUserFeatureResult
}

// WorkspacesClientCreateOrUpdateResponse contains the response from method WorkspacesClient.CreateOrUpdate.
type WorkspacesClientCreateOrUpdateResponse struct {
	Workspace
}

// WorkspacesClientDeleteResponse contains the response from method WorkspacesClient.Delete.
type WorkspacesClientDeleteResponse struct {
	// placeholder for future response values
}

// WorkspacesClientDiagnoseResponse contains the response from method WorkspacesClient.Diagnose.
type WorkspacesClientDiagnoseResponse struct {
	DiagnoseResponseResult
}

// WorkspacesClientGetResponse contains the response from method WorkspacesClient.Get.
type WorkspacesClientGetResponse struct {
	Workspace
}

// WorkspacesClientListByResourceGroupResponse contains the response from method WorkspacesClient.ListByResourceGroup.
type WorkspacesClientListByResourceGroupResponse struct {
	WorkspaceListResult
}

// WorkspacesClientListBySubscriptionResponse contains the response from method WorkspacesClient.ListBySubscription.
type WorkspacesClientListBySubscriptionResponse struct {
	WorkspaceListResult
}

// WorkspacesClientListKeysResponse contains the response from method WorkspacesClient.ListKeys.
type WorkspacesClientListKeysResponse struct {
	ListWorkspaceKeysResult
}

// WorkspacesClientListNotebookAccessTokenResponse contains the response from method WorkspacesClient.ListNotebookAccessToken.
type WorkspacesClientListNotebookAccessTokenResponse struct {
	NotebookAccessTokenResult
}

// WorkspacesClientListNotebookKeysResponse contains the response from method WorkspacesClient.ListNotebookKeys.
type WorkspacesClientListNotebookKeysResponse struct {
	ListNotebookKeysResult
}

// WorkspacesClientListOutboundNetworkDependenciesEndpointsResponse contains the response from method WorkspacesClient.ListOutboundNetworkDependenciesEndpoints.
type WorkspacesClientListOutboundNetworkDependenciesEndpointsResponse struct {
	ExternalFQDNResponse
}

// WorkspacesClientListStorageAccountKeysResponse contains the response from method WorkspacesClient.ListStorageAccountKeys.
type WorkspacesClientListStorageAccountKeysResponse struct {
	ListStorageAccountKeysResult
}

// WorkspacesClientPrepareNotebookResponse contains the response from method WorkspacesClient.PrepareNotebook.
type WorkspacesClientPrepareNotebookResponse struct {
	NotebookResourceInfo
}

// WorkspacesClientResyncKeysResponse contains the response from method WorkspacesClient.ResyncKeys.
type WorkspacesClientResyncKeysResponse struct {
	// placeholder for future response values
}

// WorkspacesClientUpdateResponse contains the response from method WorkspacesClient.Update.
type WorkspacesClientUpdateResponse struct {
	Workspace
}
