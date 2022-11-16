//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdevcenter

const (
	moduleName    = "armdevcenter"
	moduleVersion = "v0.3.0"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// CatalogSyncState - The synchronization state of the catalog.
type CatalogSyncState string

const (
	CatalogSyncStateCanceled   CatalogSyncState = "Canceled"
	CatalogSyncStateFailed     CatalogSyncState = "Failed"
	CatalogSyncStateInProgress CatalogSyncState = "InProgress"
	CatalogSyncStateSucceeded  CatalogSyncState = "Succeeded"
)

// PossibleCatalogSyncStateValues returns the possible values for the CatalogSyncState const type.
func PossibleCatalogSyncStateValues() []CatalogSyncState {
	return []CatalogSyncState{
		CatalogSyncStateCanceled,
		CatalogSyncStateFailed,
		CatalogSyncStateInProgress,
		CatalogSyncStateSucceeded,
	}
}

// CheckNameAvailabilityReason - The reason why the given name is not available.
type CheckNameAvailabilityReason string

const (
	CheckNameAvailabilityReasonAlreadyExists CheckNameAvailabilityReason = "AlreadyExists"
	CheckNameAvailabilityReasonInvalid       CheckNameAvailabilityReason = "Invalid"
)

// PossibleCheckNameAvailabilityReasonValues returns the possible values for the CheckNameAvailabilityReason const type.
func PossibleCheckNameAvailabilityReasonValues() []CheckNameAvailabilityReason {
	return []CheckNameAvailabilityReason{
		CheckNameAvailabilityReasonAlreadyExists,
		CheckNameAvailabilityReasonInvalid,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// DomainJoinType - Active Directory join type
type DomainJoinType string

const (
	DomainJoinTypeAzureADJoin       DomainJoinType = "AzureADJoin"
	DomainJoinTypeHybridAzureADJoin DomainJoinType = "HybridAzureADJoin"
)

// PossibleDomainJoinTypeValues returns the possible values for the DomainJoinType const type.
func PossibleDomainJoinTypeValues() []DomainJoinType {
	return []DomainJoinType{
		DomainJoinTypeAzureADJoin,
		DomainJoinTypeHybridAzureADJoin,
	}
}

// EnableStatus - Enable or disable status. Indicates whether the property applied to is either enabled or disabled.
type EnableStatus string

const (
	EnableStatusDisabled EnableStatus = "Disabled"
	EnableStatusEnabled  EnableStatus = "Enabled"
)

// PossibleEnableStatusValues returns the possible values for the EnableStatus const type.
func PossibleEnableStatusValues() []EnableStatus {
	return []EnableStatus{
		EnableStatusDisabled,
		EnableStatusEnabled,
	}
}

// HealthCheckStatus - Health check status values
type HealthCheckStatus string

const (
	HealthCheckStatusFailed  HealthCheckStatus = "Failed"
	HealthCheckStatusPassed  HealthCheckStatus = "Passed"
	HealthCheckStatusPending HealthCheckStatus = "Pending"
	HealthCheckStatusRunning HealthCheckStatus = "Running"
	HealthCheckStatusUnknown HealthCheckStatus = "Unknown"
	HealthCheckStatusWarning HealthCheckStatus = "Warning"
)

// PossibleHealthCheckStatusValues returns the possible values for the HealthCheckStatus const type.
func PossibleHealthCheckStatusValues() []HealthCheckStatus {
	return []HealthCheckStatus{
		HealthCheckStatusFailed,
		HealthCheckStatusPassed,
		HealthCheckStatusPending,
		HealthCheckStatusRunning,
		HealthCheckStatusUnknown,
		HealthCheckStatusWarning,
	}
}

// HibernateSupport - Indicates whether hibernate is enabled/disabled.
type HibernateSupport string

const (
	HibernateSupportDisabled HibernateSupport = "Disabled"
	HibernateSupportEnabled  HibernateSupport = "Enabled"
)

// PossibleHibernateSupportValues returns the possible values for the HibernateSupport const type.
func PossibleHibernateSupportValues() []HibernateSupport {
	return []HibernateSupport{
		HibernateSupportDisabled,
		HibernateSupportEnabled,
	}
}

// ImageValidationStatus - Image validation status
type ImageValidationStatus string

const (
	ImageValidationStatusFailed    ImageValidationStatus = "Failed"
	ImageValidationStatusPending   ImageValidationStatus = "Pending"
	ImageValidationStatusSucceeded ImageValidationStatus = "Succeeded"
	ImageValidationStatusTimedOut  ImageValidationStatus = "TimedOut"
	ImageValidationStatusUnknown   ImageValidationStatus = "Unknown"
)

// PossibleImageValidationStatusValues returns the possible values for the ImageValidationStatus const type.
func PossibleImageValidationStatusValues() []ImageValidationStatus {
	return []ImageValidationStatus{
		ImageValidationStatusFailed,
		ImageValidationStatusPending,
		ImageValidationStatusSucceeded,
		ImageValidationStatusTimedOut,
		ImageValidationStatusUnknown,
	}
}

// LicenseType - License Types
type LicenseType string

const (
	LicenseTypeWindowsClient LicenseType = "Windows_Client"
)

// PossibleLicenseTypeValues returns the possible values for the LicenseType const type.
func PossibleLicenseTypeValues() []LicenseType {
	return []LicenseType{
		LicenseTypeWindowsClient,
	}
}

type LocalAdminStatus string

const (
	LocalAdminStatusDisabled LocalAdminStatus = "Disabled"
	LocalAdminStatusEnabled  LocalAdminStatus = "Enabled"
)

// PossibleLocalAdminStatusValues returns the possible values for the LocalAdminStatus const type.
func PossibleLocalAdminStatusValues() []LocalAdminStatus {
	return []LocalAdminStatus{
		LocalAdminStatusDisabled,
		LocalAdminStatusEnabled,
	}
}

// ManagedServiceIdentityType - Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                       ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeSystemAssigned             ManagedServiceIdentityType = "SystemAssigned"
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned, UserAssigned"
	ManagedServiceIdentityTypeUserAssigned               ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeNone,
		ManagedServiceIdentityTypeSystemAssigned,
		ManagedServiceIdentityTypeSystemAssignedUserAssigned,
		ManagedServiceIdentityTypeUserAssigned,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// ProvisioningState - Provisioning state of the resource.
type ProvisioningState string

const (
	ProvisioningStateAccepted                  ProvisioningState = "Accepted"
	ProvisioningStateCanceled                  ProvisioningState = "Canceled"
	ProvisioningStateCreated                   ProvisioningState = "Created"
	ProvisioningStateCreating                  ProvisioningState = "Creating"
	ProvisioningStateDeleted                   ProvisioningState = "Deleted"
	ProvisioningStateDeleting                  ProvisioningState = "Deleting"
	ProvisioningStateFailed                    ProvisioningState = "Failed"
	ProvisioningStateMovingResources           ProvisioningState = "MovingResources"
	ProvisioningStateNotSpecified              ProvisioningState = "NotSpecified"
	ProvisioningStateRolloutInProgress         ProvisioningState = "RolloutInProgress"
	ProvisioningStateRunning                   ProvisioningState = "Running"
	ProvisioningStateStorageProvisioningFailed ProvisioningState = "StorageProvisioningFailed"
	ProvisioningStateSucceeded                 ProvisioningState = "Succeeded"
	ProvisioningStateTransientFailure          ProvisioningState = "TransientFailure"
	ProvisioningStateUpdated                   ProvisioningState = "Updated"
	ProvisioningStateUpdating                  ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateCreated,
		ProvisioningStateCreating,
		ProvisioningStateDeleted,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateMovingResources,
		ProvisioningStateNotSpecified,
		ProvisioningStateRolloutInProgress,
		ProvisioningStateRunning,
		ProvisioningStateStorageProvisioningFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateTransientFailure,
		ProvisioningStateUpdated,
		ProvisioningStateUpdating,
	}
}

// SKUTier - This field is required to be implemented by the Resource Provider if the service has more than one tier, but
// is not required on a PUT.
type SKUTier string

const (
	SKUTierFree     SKUTier = "Free"
	SKUTierBasic    SKUTier = "Basic"
	SKUTierStandard SKUTier = "Standard"
	SKUTierPremium  SKUTier = "Premium"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierFree,
		SKUTierBasic,
		SKUTierStandard,
		SKUTierPremium,
	}
}

// ScheduledFrequency - The frequency of task execution.
type ScheduledFrequency string

const (
	ScheduledFrequencyDaily ScheduledFrequency = "Daily"
)

// PossibleScheduledFrequencyValues returns the possible values for the ScheduledFrequency const type.
func PossibleScheduledFrequencyValues() []ScheduledFrequency {
	return []ScheduledFrequency{
		ScheduledFrequencyDaily,
	}
}

// ScheduledType - The supported types for a scheduled task.
type ScheduledType string

const (
	ScheduledTypeStopDevBox ScheduledType = "StopDevBox"
)

// PossibleScheduledTypeValues returns the possible values for the ScheduledType const type.
func PossibleScheduledTypeValues() []ScheduledType {
	return []ScheduledType{
		ScheduledTypeStopDevBox,
	}
}

// UsageUnit - The unit details.
type UsageUnit string

const (
	UsageUnitCount UsageUnit = "Count"
)

// PossibleUsageUnitValues returns the possible values for the UsageUnit const type.
func PossibleUsageUnitValues() []UsageUnit {
	return []UsageUnit{
		UsageUnitCount,
	}
}
