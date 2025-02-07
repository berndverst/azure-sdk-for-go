//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package postgresqlflexibleservers

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/postgresql/mgmt/2020-11-05-preview/postgresqlflexibleservers"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ConfigurationDataType = original.ConfigurationDataType

const (
	Boolean     ConfigurationDataType = original.Boolean
	Enumeration ConfigurationDataType = original.Enumeration
	Integer     ConfigurationDataType = original.Integer
	Numeric     ConfigurationDataType = original.Numeric
)

type CreateMode = original.CreateMode

const (
	Default            CreateMode = original.Default
	PointInTimeRestore CreateMode = original.PointInTimeRestore
)

type HAEnabledEnum = original.HAEnabledEnum

const (
	Disabled HAEnabledEnum = original.Disabled
	Enabled  HAEnabledEnum = original.Enabled
)

type OperationOrigin = original.OperationOrigin

const (
	NotSpecified OperationOrigin = original.NotSpecified
	System       OperationOrigin = original.System
	User         OperationOrigin = original.User
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	SystemAssigned ResourceIdentityType = original.SystemAssigned
)

type ServerHAState = original.ServerHAState

const (
	CreatingStandby ServerHAState = original.CreatingStandby
	FailingOver     ServerHAState = original.FailingOver
	Healthy         ServerHAState = original.Healthy
	NotEnabled      ServerHAState = original.NotEnabled
	RemovingStandby ServerHAState = original.RemovingStandby
	ReplicatingData ServerHAState = original.ReplicatingData
)

type ServerPublicNetworkAccessState = original.ServerPublicNetworkAccessState

const (
	ServerPublicNetworkAccessStateDisabled ServerPublicNetworkAccessState = original.ServerPublicNetworkAccessStateDisabled
	ServerPublicNetworkAccessStateEnabled  ServerPublicNetworkAccessState = original.ServerPublicNetworkAccessStateEnabled
)

type ServerState = original.ServerState

const (
	ServerStateDisabled ServerState = original.ServerStateDisabled
	ServerStateDropping ServerState = original.ServerStateDropping
	ServerStateReady    ServerState = original.ServerStateReady
	ServerStateStarting ServerState = original.ServerStateStarting
	ServerStateStopped  ServerState = original.ServerStateStopped
	ServerStateStopping ServerState = original.ServerStateStopping
	ServerStateUpdating ServerState = original.ServerStateUpdating
)

type ServerVersion = original.ServerVersion

const (
	OneOne ServerVersion = original.OneOne
	OneTwo ServerVersion = original.OneTwo
)

type SkuTier = original.SkuTier

const (
	Burstable       SkuTier = original.Burstable
	GeneralPurpose  SkuTier = original.GeneralPurpose
	MemoryOptimized SkuTier = original.MemoryOptimized
)

type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type CapabilitiesListResult = original.CapabilitiesListResult
type CapabilitiesListResultIterator = original.CapabilitiesListResultIterator
type CapabilitiesListResultPage = original.CapabilitiesListResultPage
type CapabilityProperties = original.CapabilityProperties
type CheckNameAvailabilityClient = original.CheckNameAvailabilityClient
type CloudError = original.CloudError
type Configuration = original.Configuration
type ConfigurationListResult = original.ConfigurationListResult
type ConfigurationListResultIterator = original.ConfigurationListResultIterator
type ConfigurationListResultPage = original.ConfigurationListResultPage
type ConfigurationProperties = original.ConfigurationProperties
type ConfigurationsClient = original.ConfigurationsClient
type ConfigurationsUpdateFuture = original.ConfigurationsUpdateFuture
type Database = original.Database
type DatabaseListResult = original.DatabaseListResult
type DatabaseListResultIterator = original.DatabaseListResultIterator
type DatabaseListResultPage = original.DatabaseListResultPage
type DatabaseProperties = original.DatabaseProperties
type DatabasesClient = original.DatabasesClient
type DatabasesCreateFuture = original.DatabasesCreateFuture
type DatabasesDeleteFuture = original.DatabasesDeleteFuture
type DelegatedSubnetUsage = original.DelegatedSubnetUsage
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorResponse = original.ErrorResponse
type FirewallRule = original.FirewallRule
type FirewallRuleListResult = original.FirewallRuleListResult
type FirewallRuleListResultIterator = original.FirewallRuleListResultIterator
type FirewallRuleListResultPage = original.FirewallRuleListResultPage
type FirewallRuleProperties = original.FirewallRuleProperties
type FirewallRulesClient = original.FirewallRulesClient
type FirewallRulesCreateOrUpdateFuture = original.FirewallRulesCreateOrUpdateFuture
type FirewallRulesDeleteFuture = original.FirewallRulesDeleteFuture
type Identity = original.Identity
type LocationBasedCapabilitiesClient = original.LocationBasedCapabilitiesClient
type MaintenanceWindow = original.MaintenanceWindow
type NameAvailability = original.NameAvailability
type NameAvailabilityRequest = original.NameAvailabilityRequest
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationsClient = original.OperationsClient
type Plan = original.Plan
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type ResourceModelWithAllowedPropertySet = original.ResourceModelWithAllowedPropertySet
type ResourceModelWithAllowedPropertySetIdentity = original.ResourceModelWithAllowedPropertySetIdentity
type ResourceModelWithAllowedPropertySetPlan = original.ResourceModelWithAllowedPropertySetPlan
type ResourceModelWithAllowedPropertySetSku = original.ResourceModelWithAllowedPropertySetSku
type Server = original.Server
type ServerEditionCapability = original.ServerEditionCapability
type ServerForUpdate = original.ServerForUpdate
type ServerListResult = original.ServerListResult
type ServerListResultIterator = original.ServerListResultIterator
type ServerListResultPage = original.ServerListResultPage
type ServerProperties = original.ServerProperties
type ServerPropertiesDelegatedSubnetArguments = original.ServerPropertiesDelegatedSubnetArguments
type ServerPropertiesForUpdate = original.ServerPropertiesForUpdate
type ServerVersionCapability = original.ServerVersionCapability
type ServersClient = original.ServersClient
type ServersCreateFuture = original.ServersCreateFuture
type ServersDeleteFuture = original.ServersDeleteFuture
type ServersRestartFuture = original.ServersRestartFuture
type ServersStartFuture = original.ServersStartFuture
type ServersStopFuture = original.ServersStopFuture
type ServersUpdateFuture = original.ServersUpdateFuture
type Sku = original.Sku
type StorageEditionCapability = original.StorageEditionCapability
type StorageMBCapability = original.StorageMBCapability
type StorageProfile = original.StorageProfile
type TrackedResource = original.TrackedResource
type VcoreCapability = original.VcoreCapability
type VirtualNetworkSubnetUsageClient = original.VirtualNetworkSubnetUsageClient
type VirtualNetworkSubnetUsageParameter = original.VirtualNetworkSubnetUsageParameter
type VirtualNetworkSubnetUsageResult = original.VirtualNetworkSubnetUsageResult

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewCapabilitiesListResultIterator(page CapabilitiesListResultPage) CapabilitiesListResultIterator {
	return original.NewCapabilitiesListResultIterator(page)
}
func NewCapabilitiesListResultPage(cur CapabilitiesListResult, getNextPage func(context.Context, CapabilitiesListResult) (CapabilitiesListResult, error)) CapabilitiesListResultPage {
	return original.NewCapabilitiesListResultPage(cur, getNextPage)
}
func NewCheckNameAvailabilityClient(subscriptionID string) CheckNameAvailabilityClient {
	return original.NewCheckNameAvailabilityClient(subscriptionID)
}
func NewCheckNameAvailabilityClientWithBaseURI(baseURI string, subscriptionID string) CheckNameAvailabilityClient {
	return original.NewCheckNameAvailabilityClientWithBaseURI(baseURI, subscriptionID)
}
func NewConfigurationListResultIterator(page ConfigurationListResultPage) ConfigurationListResultIterator {
	return original.NewConfigurationListResultIterator(page)
}
func NewConfigurationListResultPage(cur ConfigurationListResult, getNextPage func(context.Context, ConfigurationListResult) (ConfigurationListResult, error)) ConfigurationListResultPage {
	return original.NewConfigurationListResultPage(cur, getNextPage)
}
func NewConfigurationsClient(subscriptionID string) ConfigurationsClient {
	return original.NewConfigurationsClient(subscriptionID)
}
func NewConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationsClient {
	return original.NewConfigurationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDatabaseListResultIterator(page DatabaseListResultPage) DatabaseListResultIterator {
	return original.NewDatabaseListResultIterator(page)
}
func NewDatabaseListResultPage(cur DatabaseListResult, getNextPage func(context.Context, DatabaseListResult) (DatabaseListResult, error)) DatabaseListResultPage {
	return original.NewDatabaseListResultPage(cur, getNextPage)
}
func NewDatabasesClient(subscriptionID string) DatabasesClient {
	return original.NewDatabasesClient(subscriptionID)
}
func NewDatabasesClientWithBaseURI(baseURI string, subscriptionID string) DatabasesClient {
	return original.NewDatabasesClientWithBaseURI(baseURI, subscriptionID)
}
func NewFirewallRuleListResultIterator(page FirewallRuleListResultPage) FirewallRuleListResultIterator {
	return original.NewFirewallRuleListResultIterator(page)
}
func NewFirewallRuleListResultPage(cur FirewallRuleListResult, getNextPage func(context.Context, FirewallRuleListResult) (FirewallRuleListResult, error)) FirewallRuleListResultPage {
	return original.NewFirewallRuleListResultPage(cur, getNextPage)
}
func NewFirewallRulesClient(subscriptionID string) FirewallRulesClient {
	return original.NewFirewallRulesClient(subscriptionID)
}
func NewFirewallRulesClientWithBaseURI(baseURI string, subscriptionID string) FirewallRulesClient {
	return original.NewFirewallRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewLocationBasedCapabilitiesClient(subscriptionID string) LocationBasedCapabilitiesClient {
	return original.NewLocationBasedCapabilitiesClient(subscriptionID)
}
func NewLocationBasedCapabilitiesClientWithBaseURI(baseURI string, subscriptionID string) LocationBasedCapabilitiesClient {
	return original.NewLocationBasedCapabilitiesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewServerListResultIterator(page ServerListResultPage) ServerListResultIterator {
	return original.NewServerListResultIterator(page)
}
func NewServerListResultPage(cur ServerListResult, getNextPage func(context.Context, ServerListResult) (ServerListResult, error)) ServerListResultPage {
	return original.NewServerListResultPage(cur, getNextPage)
}
func NewServersClient(subscriptionID string) ServersClient {
	return original.NewServersClient(subscriptionID)
}
func NewServersClientWithBaseURI(baseURI string, subscriptionID string) ServersClient {
	return original.NewServersClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualNetworkSubnetUsageClient(subscriptionID string) VirtualNetworkSubnetUsageClient {
	return original.NewVirtualNetworkSubnetUsageClient(subscriptionID)
}
func NewVirtualNetworkSubnetUsageClientWithBaseURI(baseURI string, subscriptionID string) VirtualNetworkSubnetUsageClient {
	return original.NewVirtualNetworkSubnetUsageClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleConfigurationDataTypeValues() []ConfigurationDataType {
	return original.PossibleConfigurationDataTypeValues()
}
func PossibleCreateModeValues() []CreateMode {
	return original.PossibleCreateModeValues()
}
func PossibleHAEnabledEnumValues() []HAEnabledEnum {
	return original.PossibleHAEnabledEnumValues()
}
func PossibleOperationOriginValues() []OperationOrigin {
	return original.PossibleOperationOriginValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleServerHAStateValues() []ServerHAState {
	return original.PossibleServerHAStateValues()
}
func PossibleServerPublicNetworkAccessStateValues() []ServerPublicNetworkAccessState {
	return original.PossibleServerPublicNetworkAccessStateValues()
}
func PossibleServerStateValues() []ServerState {
	return original.PossibleServerStateValues()
}
func PossibleServerVersionValues() []ServerVersion {
	return original.PossibleServerVersionValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
