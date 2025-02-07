package virtualmachineimagebuilder

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// CreatedByType enumerates the values for created by type.
type CreatedByType string

const (
	// CreatedByTypeApplication ...
	CreatedByTypeApplication CreatedByType = "Application"
	// CreatedByTypeKey ...
	CreatedByTypeKey CreatedByType = "Key"
	// CreatedByTypeManagedIdentity ...
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	// CreatedByTypeUser ...
	CreatedByTypeUser CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns an array of possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{CreatedByTypeApplication, CreatedByTypeKey, CreatedByTypeManagedIdentity, CreatedByTypeUser}
}

// ProvisioningErrorCode enumerates the values for provisioning error code.
type ProvisioningErrorCode string

const (
	// ProvisioningErrorCodeBadCustomizerType ...
	ProvisioningErrorCodeBadCustomizerType ProvisioningErrorCode = "BadCustomizerType"
	// ProvisioningErrorCodeBadDistributeType ...
	ProvisioningErrorCodeBadDistributeType ProvisioningErrorCode = "BadDistributeType"
	// ProvisioningErrorCodeBadManagedImageSource ...
	ProvisioningErrorCodeBadManagedImageSource ProvisioningErrorCode = "BadManagedImageSource"
	// ProvisioningErrorCodeBadPIRSource ...
	ProvisioningErrorCodeBadPIRSource ProvisioningErrorCode = "BadPIRSource"
	// ProvisioningErrorCodeBadSharedImageDistribute ...
	ProvisioningErrorCodeBadSharedImageDistribute ProvisioningErrorCode = "BadSharedImageDistribute"
	// ProvisioningErrorCodeBadSharedImageVersionSource ...
	ProvisioningErrorCodeBadSharedImageVersionSource ProvisioningErrorCode = "BadSharedImageVersionSource"
	// ProvisioningErrorCodeBadSourceType ...
	ProvisioningErrorCodeBadSourceType ProvisioningErrorCode = "BadSourceType"
	// ProvisioningErrorCodeNoCustomizerScript ...
	ProvisioningErrorCodeNoCustomizerScript ProvisioningErrorCode = "NoCustomizerScript"
	// ProvisioningErrorCodeOther ...
	ProvisioningErrorCodeOther ProvisioningErrorCode = "Other"
	// ProvisioningErrorCodeServerError ...
	ProvisioningErrorCodeServerError ProvisioningErrorCode = "ServerError"
	// ProvisioningErrorCodeUnsupportedCustomizerType ...
	ProvisioningErrorCodeUnsupportedCustomizerType ProvisioningErrorCode = "UnsupportedCustomizerType"
)

// PossibleProvisioningErrorCodeValues returns an array of possible values for the ProvisioningErrorCode const type.
func PossibleProvisioningErrorCodeValues() []ProvisioningErrorCode {
	return []ProvisioningErrorCode{ProvisioningErrorCodeBadCustomizerType, ProvisioningErrorCodeBadDistributeType, ProvisioningErrorCodeBadManagedImageSource, ProvisioningErrorCodeBadPIRSource, ProvisioningErrorCodeBadSharedImageDistribute, ProvisioningErrorCodeBadSharedImageVersionSource, ProvisioningErrorCodeBadSourceType, ProvisioningErrorCodeNoCustomizerScript, ProvisioningErrorCodeOther, ProvisioningErrorCodeServerError, ProvisioningErrorCodeUnsupportedCustomizerType}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// ProvisioningStateCreating ...
	ProvisioningStateCreating ProvisioningState = "Creating"
	// ProvisioningStateDeleting ...
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed ...
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateSucceeded ...
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating ...
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{ProvisioningStateCreating, ProvisioningStateDeleting, ProvisioningStateFailed, ProvisioningStateSucceeded, ProvisioningStateUpdating}
}

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// ResourceIdentityTypeNone ...
	ResourceIdentityTypeNone ResourceIdentityType = "None"
	// ResourceIdentityTypeUserAssigned ...
	ResourceIdentityTypeUserAssigned ResourceIdentityType = "UserAssigned"
)

// PossibleResourceIdentityTypeValues returns an array of possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{ResourceIdentityTypeNone, ResourceIdentityTypeUserAssigned}
}

// RunState enumerates the values for run state.
type RunState string

const (
	// RunStateCanceled ...
	RunStateCanceled RunState = "Canceled"
	// RunStateCanceling ...
	RunStateCanceling RunState = "Canceling"
	// RunStateFailed ...
	RunStateFailed RunState = "Failed"
	// RunStatePartiallySucceeded ...
	RunStatePartiallySucceeded RunState = "PartiallySucceeded"
	// RunStateRunning ...
	RunStateRunning RunState = "Running"
	// RunStateSucceeded ...
	RunStateSucceeded RunState = "Succeeded"
)

// PossibleRunStateValues returns an array of possible values for the RunState const type.
func PossibleRunStateValues() []RunState {
	return []RunState{RunStateCanceled, RunStateCanceling, RunStateFailed, RunStatePartiallySucceeded, RunStateRunning, RunStateSucceeded}
}

// RunSubState enumerates the values for run sub state.
type RunSubState string

const (
	// RunSubStateBuilding ...
	RunSubStateBuilding RunSubState = "Building"
	// RunSubStateCustomizing ...
	RunSubStateCustomizing RunSubState = "Customizing"
	// RunSubStateDistributing ...
	RunSubStateDistributing RunSubState = "Distributing"
	// RunSubStateQueued ...
	RunSubStateQueued RunSubState = "Queued"
)

// PossibleRunSubStateValues returns an array of possible values for the RunSubState const type.
func PossibleRunSubStateValues() []RunSubState {
	return []RunSubState{RunSubStateBuilding, RunSubStateCustomizing, RunSubStateDistributing, RunSubStateQueued}
}

// SharedImageStorageAccountType enumerates the values for shared image storage account type.
type SharedImageStorageAccountType string

const (
	// SharedImageStorageAccountTypeStandardLRS ...
	SharedImageStorageAccountTypeStandardLRS SharedImageStorageAccountType = "Standard_LRS"
	// SharedImageStorageAccountTypeStandardZRS ...
	SharedImageStorageAccountTypeStandardZRS SharedImageStorageAccountType = "Standard_ZRS"
)

// PossibleSharedImageStorageAccountTypeValues returns an array of possible values for the SharedImageStorageAccountType const type.
func PossibleSharedImageStorageAccountTypeValues() []SharedImageStorageAccountType {
	return []SharedImageStorageAccountType{SharedImageStorageAccountTypeStandardLRS, SharedImageStorageAccountTypeStandardZRS}
}

// Type enumerates the values for type.
type Type string

const (
	// TypeImageTemplateSource ...
	TypeImageTemplateSource Type = "ImageTemplateSource"
	// TypeManagedImage ...
	TypeManagedImage Type = "ManagedImage"
	// TypePlatformImage ...
	TypePlatformImage Type = "PlatformImage"
	// TypeSharedImageVersion ...
	TypeSharedImageVersion Type = "SharedImageVersion"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{TypeImageTemplateSource, TypeManagedImage, TypePlatformImage, TypeSharedImageVersion}
}

// TypeBasicImageTemplateCustomizer enumerates the values for type basic image template customizer.
type TypeBasicImageTemplateCustomizer string

const (
	// TypeBasicImageTemplateCustomizerTypeFile ...
	TypeBasicImageTemplateCustomizerTypeFile TypeBasicImageTemplateCustomizer = "File"
	// TypeBasicImageTemplateCustomizerTypeImageTemplateCustomizer ...
	TypeBasicImageTemplateCustomizerTypeImageTemplateCustomizer TypeBasicImageTemplateCustomizer = "ImageTemplateCustomizer"
	// TypeBasicImageTemplateCustomizerTypePowerShell ...
	TypeBasicImageTemplateCustomizerTypePowerShell TypeBasicImageTemplateCustomizer = "PowerShell"
	// TypeBasicImageTemplateCustomizerTypeShell ...
	TypeBasicImageTemplateCustomizerTypeShell TypeBasicImageTemplateCustomizer = "Shell"
	// TypeBasicImageTemplateCustomizerTypeWindowsRestart ...
	TypeBasicImageTemplateCustomizerTypeWindowsRestart TypeBasicImageTemplateCustomizer = "WindowsRestart"
	// TypeBasicImageTemplateCustomizerTypeWindowsUpdate ...
	TypeBasicImageTemplateCustomizerTypeWindowsUpdate TypeBasicImageTemplateCustomizer = "WindowsUpdate"
)

// PossibleTypeBasicImageTemplateCustomizerValues returns an array of possible values for the TypeBasicImageTemplateCustomizer const type.
func PossibleTypeBasicImageTemplateCustomizerValues() []TypeBasicImageTemplateCustomizer {
	return []TypeBasicImageTemplateCustomizer{TypeBasicImageTemplateCustomizerTypeFile, TypeBasicImageTemplateCustomizerTypeImageTemplateCustomizer, TypeBasicImageTemplateCustomizerTypePowerShell, TypeBasicImageTemplateCustomizerTypeShell, TypeBasicImageTemplateCustomizerTypeWindowsRestart, TypeBasicImageTemplateCustomizerTypeWindowsUpdate}
}

// TypeBasicImageTemplateDistributor enumerates the values for type basic image template distributor.
type TypeBasicImageTemplateDistributor string

const (
	// TypeBasicImageTemplateDistributorTypeImageTemplateDistributor ...
	TypeBasicImageTemplateDistributorTypeImageTemplateDistributor TypeBasicImageTemplateDistributor = "ImageTemplateDistributor"
	// TypeBasicImageTemplateDistributorTypeManagedImage ...
	TypeBasicImageTemplateDistributorTypeManagedImage TypeBasicImageTemplateDistributor = "ManagedImage"
	// TypeBasicImageTemplateDistributorTypeSharedImage ...
	TypeBasicImageTemplateDistributorTypeSharedImage TypeBasicImageTemplateDistributor = "SharedImage"
	// TypeBasicImageTemplateDistributorTypeVHD ...
	TypeBasicImageTemplateDistributorTypeVHD TypeBasicImageTemplateDistributor = "VHD"
)

// PossibleTypeBasicImageTemplateDistributorValues returns an array of possible values for the TypeBasicImageTemplateDistributor const type.
func PossibleTypeBasicImageTemplateDistributorValues() []TypeBasicImageTemplateDistributor {
	return []TypeBasicImageTemplateDistributor{TypeBasicImageTemplateDistributorTypeImageTemplateDistributor, TypeBasicImageTemplateDistributorTypeManagedImage, TypeBasicImageTemplateDistributorTypeSharedImage, TypeBasicImageTemplateDistributorTypeVHD}
}
