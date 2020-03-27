// +build go1.9

// Copyright 2020 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package databox

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/databox/mgmt/2019-09-01/databox"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AccessProtocol = original.AccessProtocol

const (
	NFS AccessProtocol = original.NFS
	SMB AccessProtocol = original.SMB
)

type AddressType = original.AddressType

const (
	Commercial  AddressType = original.Commercial
	None        AddressType = original.None
	Residential AddressType = original.Residential
)

type AddressValidationStatus = original.AddressValidationStatus

const (
	Ambiguous AddressValidationStatus = original.Ambiguous
	Invalid   AddressValidationStatus = original.Invalid
	Valid     AddressValidationStatus = original.Valid
)

type CopyLogDetailsType = original.CopyLogDetailsType

const (
	CopyLogDetailsTypeCopyLogDetails CopyLogDetailsType = original.CopyLogDetailsTypeCopyLogDetails
	CopyLogDetailsTypeDataBox        CopyLogDetailsType = original.CopyLogDetailsTypeDataBox
	CopyLogDetailsTypeDataBoxDisk    CopyLogDetailsType = original.CopyLogDetailsTypeDataBoxDisk
	CopyLogDetailsTypeDataBoxHeavy   CopyLogDetailsType = original.CopyLogDetailsTypeDataBoxHeavy
)

type CopyStatus = original.CopyStatus

const (
	Completed                   CopyStatus = original.Completed
	CompletedWithErrors         CopyStatus = original.CompletedWithErrors
	DeviceFormatted             CopyStatus = original.DeviceFormatted
	DeviceMetadataModified      CopyStatus = original.DeviceMetadataModified
	Failed                      CopyStatus = original.Failed
	HardwareError               CopyStatus = original.HardwareError
	InProgress                  CopyStatus = original.InProgress
	NotReturned                 CopyStatus = original.NotReturned
	NotStarted                  CopyStatus = original.NotStarted
	StorageAccountNotAccessible CopyStatus = original.StorageAccountNotAccessible
	UnsupportedData             CopyStatus = original.UnsupportedData
)

type DataDestinationType = original.DataDestinationType

const (
	ManagedDisk    DataDestinationType = original.ManagedDisk
	StorageAccount DataDestinationType = original.StorageAccount
)

type DataDestinationTypeBasicDestinationAccountDetails = original.DataDestinationTypeBasicDestinationAccountDetails

const (
	DataDestinationTypeDestinationAccountDetails DataDestinationTypeBasicDestinationAccountDetails = original.DataDestinationTypeDestinationAccountDetails
	DataDestinationTypeManagedDisk               DataDestinationTypeBasicDestinationAccountDetails = original.DataDestinationTypeManagedDisk
	DataDestinationTypeStorageAccount            DataDestinationTypeBasicDestinationAccountDetails = original.DataDestinationTypeStorageAccount
)

type JobDeliveryType = original.JobDeliveryType

const (
	NonScheduled JobDeliveryType = original.NonScheduled
	Scheduled    JobDeliveryType = original.Scheduled
)

type JobDetailsTypeEnum = original.JobDetailsTypeEnum

const (
	JobDetailsTypeDataBox      JobDetailsTypeEnum = original.JobDetailsTypeDataBox
	JobDetailsTypeDataBoxDisk  JobDetailsTypeEnum = original.JobDetailsTypeDataBoxDisk
	JobDetailsTypeDataBoxHeavy JobDetailsTypeEnum = original.JobDetailsTypeDataBoxHeavy
	JobDetailsTypeJobDetails   JobDetailsTypeEnum = original.JobDetailsTypeJobDetails
)

type JobSecretsTypeEnum = original.JobSecretsTypeEnum

const (
	JobSecretsTypeDataBox      JobSecretsTypeEnum = original.JobSecretsTypeDataBox
	JobSecretsTypeDataBoxDisk  JobSecretsTypeEnum = original.JobSecretsTypeDataBoxDisk
	JobSecretsTypeDataBoxHeavy JobSecretsTypeEnum = original.JobSecretsTypeDataBoxHeavy
	JobSecretsTypeJobSecrets   JobSecretsTypeEnum = original.JobSecretsTypeJobSecrets
)

type NotificationStageName = original.NotificationStageName

const (
	AtAzureDC      NotificationStageName = original.AtAzureDC
	DataCopy       NotificationStageName = original.DataCopy
	Delivered      NotificationStageName = original.Delivered
	DevicePrepared NotificationStageName = original.DevicePrepared
	Dispatched     NotificationStageName = original.Dispatched
	PickedUp       NotificationStageName = original.PickedUp
)

type OverallValidationStatus = original.OverallValidationStatus

const (
	AllValidToProceed              OverallValidationStatus = original.AllValidToProceed
	CertainInputValidationsSkipped OverallValidationStatus = original.CertainInputValidationsSkipped
	InputsRevisitRequired          OverallValidationStatus = original.InputsRevisitRequired
)

type ShareDestinationFormatType = original.ShareDestinationFormatType

const (
	ShareDestinationFormatTypeAzureFile   ShareDestinationFormatType = original.ShareDestinationFormatTypeAzureFile
	ShareDestinationFormatTypeBlockBlob   ShareDestinationFormatType = original.ShareDestinationFormatTypeBlockBlob
	ShareDestinationFormatTypeHCS         ShareDestinationFormatType = original.ShareDestinationFormatTypeHCS
	ShareDestinationFormatTypeManagedDisk ShareDestinationFormatType = original.ShareDestinationFormatTypeManagedDisk
	ShareDestinationFormatTypePageBlob    ShareDestinationFormatType = original.ShareDestinationFormatTypePageBlob
	ShareDestinationFormatTypeUnknownType ShareDestinationFormatType = original.ShareDestinationFormatTypeUnknownType
)

type SkuDisabledReason = original.SkuDisabledReason

const (
	SkuDisabledReasonCountry            SkuDisabledReason = original.SkuDisabledReasonCountry
	SkuDisabledReasonFeature            SkuDisabledReason = original.SkuDisabledReasonFeature
	SkuDisabledReasonNone               SkuDisabledReason = original.SkuDisabledReasonNone
	SkuDisabledReasonNoSubscriptionInfo SkuDisabledReason = original.SkuDisabledReasonNoSubscriptionInfo
	SkuDisabledReasonOfferType          SkuDisabledReason = original.SkuDisabledReasonOfferType
	SkuDisabledReasonRegion             SkuDisabledReason = original.SkuDisabledReasonRegion
)

type SkuName = original.SkuName

const (
	DataBox      SkuName = original.DataBox
	DataBoxDisk  SkuName = original.DataBoxDisk
	DataBoxHeavy SkuName = original.DataBoxHeavy
)

type SkuNameBasicScheduleAvailabilityRequest = original.SkuNameBasicScheduleAvailabilityRequest

const (
	SkuNameDataBox                     SkuNameBasicScheduleAvailabilityRequest = original.SkuNameDataBox
	SkuNameDataBoxDisk                 SkuNameBasicScheduleAvailabilityRequest = original.SkuNameDataBoxDisk
	SkuNameDataBoxHeavy                SkuNameBasicScheduleAvailabilityRequest = original.SkuNameDataBoxHeavy
	SkuNameScheduleAvailabilityRequest SkuNameBasicScheduleAvailabilityRequest = original.SkuNameScheduleAvailabilityRequest
)

type StageName = original.StageName

const (
	StageNameAborted                       StageName = original.StageNameAborted
	StageNameAtAzureDC                     StageName = original.StageNameAtAzureDC
	StageNameCancelled                     StageName = original.StageNameCancelled
	StageNameCompleted                     StageName = original.StageNameCompleted
	StageNameCompletedWithErrors           StageName = original.StageNameCompletedWithErrors
	StageNameCompletedWithWarnings         StageName = original.StageNameCompletedWithWarnings
	StageNameDataCopy                      StageName = original.StageNameDataCopy
	StageNameDelivered                     StageName = original.StageNameDelivered
	StageNameDeviceOrdered                 StageName = original.StageNameDeviceOrdered
	StageNameDevicePrepared                StageName = original.StageNameDevicePrepared
	StageNameDispatched                    StageName = original.StageNameDispatched
	StageNameFailedIssueDetectedAtAzureDC  StageName = original.StageNameFailedIssueDetectedAtAzureDC
	StageNameFailedIssueReportedAtCustomer StageName = original.StageNameFailedIssueReportedAtCustomer
	StageNamePickedUp                      StageName = original.StageNamePickedUp
	StageNameReadyToDispatchFromAzureDC    StageName = original.StageNameReadyToDispatchFromAzureDC
	StageNameReadyToReceiveAtAzureDC       StageName = original.StageNameReadyToReceiveAtAzureDC
)

type StageStatus = original.StageStatus

const (
	StageStatusCancelled           StageStatus = original.StageStatusCancelled
	StageStatusCancelling          StageStatus = original.StageStatusCancelling
	StageStatusFailed              StageStatus = original.StageStatusFailed
	StageStatusInProgress          StageStatus = original.StageStatusInProgress
	StageStatusNone                StageStatus = original.StageStatusNone
	StageStatusSucceeded           StageStatus = original.StageStatusSucceeded
	StageStatusSucceededWithErrors StageStatus = original.StageStatusSucceededWithErrors
)

type TransportShipmentTypes = original.TransportShipmentTypes

const (
	CustomerManaged  TransportShipmentTypes = original.CustomerManaged
	MicrosoftManaged TransportShipmentTypes = original.MicrosoftManaged
)

type ValidationCategory = original.ValidationCategory

const (
	ValidationCategoryJobCreationValidation ValidationCategory = original.ValidationCategoryJobCreationValidation
	ValidationCategoryValidationRequest     ValidationCategory = original.ValidationCategoryValidationRequest
)

type ValidationStatus = original.ValidationStatus

const (
	ValidationStatusInvalid ValidationStatus = original.ValidationStatusInvalid
	ValidationStatusSkipped ValidationStatus = original.ValidationStatusSkipped
	ValidationStatusValid   ValidationStatus = original.ValidationStatusValid
)

type ValidationType = original.ValidationType

const (
	ValidationTypeValidateAddress                          ValidationType = original.ValidationTypeValidateAddress
	ValidationTypeValidateCreateOrderLimit                 ValidationType = original.ValidationTypeValidateCreateOrderLimit
	ValidationTypeValidateDataDestinationDetails           ValidationType = original.ValidationTypeValidateDataDestinationDetails
	ValidationTypeValidatePreferences                      ValidationType = original.ValidationTypeValidatePreferences
	ValidationTypeValidateSkuAvailability                  ValidationType = original.ValidationTypeValidateSkuAvailability
	ValidationTypeValidateSubscriptionIsAllowedToCreateJob ValidationType = original.ValidationTypeValidateSubscriptionIsAllowedToCreateJob
	ValidationTypeValidationInputRequest                   ValidationType = original.ValidationTypeValidationInputRequest
)

type ValidationTypeBasicValidationInputResponse = original.ValidationTypeBasicValidationInputResponse

const (
	ValidationTypeBasicValidationInputResponseValidationTypeValidateAddress                          ValidationTypeBasicValidationInputResponse = original.ValidationTypeBasicValidationInputResponseValidationTypeValidateAddress
	ValidationTypeBasicValidationInputResponseValidationTypeValidateCreateOrderLimit                 ValidationTypeBasicValidationInputResponse = original.ValidationTypeBasicValidationInputResponseValidationTypeValidateCreateOrderLimit
	ValidationTypeBasicValidationInputResponseValidationTypeValidateDataDestinationDetails           ValidationTypeBasicValidationInputResponse = original.ValidationTypeBasicValidationInputResponseValidationTypeValidateDataDestinationDetails
	ValidationTypeBasicValidationInputResponseValidationTypeValidatePreferences                      ValidationTypeBasicValidationInputResponse = original.ValidationTypeBasicValidationInputResponseValidationTypeValidatePreferences
	ValidationTypeBasicValidationInputResponseValidationTypeValidateSkuAvailability                  ValidationTypeBasicValidationInputResponse = original.ValidationTypeBasicValidationInputResponseValidationTypeValidateSkuAvailability
	ValidationTypeBasicValidationInputResponseValidationTypeValidateSubscriptionIsAllowedToCreateJob ValidationTypeBasicValidationInputResponse = original.ValidationTypeBasicValidationInputResponseValidationTypeValidateSubscriptionIsAllowedToCreateJob
	ValidationTypeBasicValidationInputResponseValidationTypeValidationInputResponse                  ValidationTypeBasicValidationInputResponse = original.ValidationTypeBasicValidationInputResponseValidationTypeValidationInputResponse
)

type AccountCopyLogDetails = original.AccountCopyLogDetails
type AccountCredentialDetails = original.AccountCredentialDetails
type AddressValidationOutput = original.AddressValidationOutput
type AddressValidationProperties = original.AddressValidationProperties
type ApplianceNetworkConfiguration = original.ApplianceNetworkConfiguration
type ArmBaseObject = original.ArmBaseObject
type AvailableSkuRequest = original.AvailableSkuRequest
type AvailableSkusResult = original.AvailableSkusResult
type AvailableSkusResultIterator = original.AvailableSkusResultIterator
type AvailableSkusResultPage = original.AvailableSkusResultPage
type BaseClient = original.BaseClient
type BasicCopyLogDetails = original.BasicCopyLogDetails
type BasicDestinationAccountDetails = original.BasicDestinationAccountDetails
type BasicJobDetails = original.BasicJobDetails
type BasicJobSecrets = original.BasicJobSecrets
type BasicScheduleAvailabilityRequest = original.BasicScheduleAvailabilityRequest
type BasicValidationInputRequest = original.BasicValidationInputRequest
type BasicValidationInputResponse = original.BasicValidationInputResponse
type BasicValidationRequest = original.BasicValidationRequest
type CancellationReason = original.CancellationReason
type CloudError = original.CloudError
type ContactDetails = original.ContactDetails
type CopyLogDetails = original.CopyLogDetails
type CopyProgress = original.CopyProgress
type CreateJobValidations = original.CreateJobValidations
type CreateOrderLimitForSubscriptionValidationRequest = original.CreateOrderLimitForSubscriptionValidationRequest
type CreateOrderLimitForSubscriptionValidationResponseProperties = original.CreateOrderLimitForSubscriptionValidationResponseProperties
type DataDestinationDetailsValidationRequest = original.DataDestinationDetailsValidationRequest
type DataDestinationDetailsValidationResponseProperties = original.DataDestinationDetailsValidationResponseProperties
type DcAccessSecurityCode = original.DcAccessSecurityCode
type DestinationAccountDetails = original.DestinationAccountDetails
type DestinationManagedDiskDetails = original.DestinationManagedDiskDetails
type DestinationStorageAccountDetails = original.DestinationStorageAccountDetails
type DestinationToServiceLocationMap = original.DestinationToServiceLocationMap
type DiskCopyLogDetails = original.DiskCopyLogDetails
type DiskCopyProgress = original.DiskCopyProgress
type DiskJobDetails = original.DiskJobDetails
type DiskJobSecrets = original.DiskJobSecrets
type DiskScheduleAvailabilityRequest = original.DiskScheduleAvailabilityRequest
type DiskSecret = original.DiskSecret
type Error = original.Error
type HeavyAccountCopyLogDetails = original.HeavyAccountCopyLogDetails
type HeavyJobDetails = original.HeavyJobDetails
type HeavyJobSecrets = original.HeavyJobSecrets
type HeavyScheduleAvailabilityRequest = original.HeavyScheduleAvailabilityRequest
type HeavySecret = original.HeavySecret
type JobDeliveryInfo = original.JobDeliveryInfo
type JobDetails = original.JobDetails
type JobDetailsType = original.JobDetailsType
type JobErrorDetails = original.JobErrorDetails
type JobProperties = original.JobProperties
type JobResource = original.JobResource
type JobResourceList = original.JobResourceList
type JobResourceListIterator = original.JobResourceListIterator
type JobResourceListPage = original.JobResourceListPage
type JobResourceUpdateParameter = original.JobResourceUpdateParameter
type JobSecrets = original.JobSecrets
type JobSecretsType = original.JobSecretsType
type JobStages = original.JobStages
type JobsClient = original.JobsClient
type JobsCreateFuture = original.JobsCreateFuture
type JobsDeleteFuture = original.JobsDeleteFuture
type JobsUpdateFuture = original.JobsUpdateFuture
type NotificationPreference = original.NotificationPreference
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationList = original.OperationList
type OperationListIterator = original.OperationListIterator
type OperationListPage = original.OperationListPage
type OperationsClient = original.OperationsClient
type PackageShippingDetails = original.PackageShippingDetails
type Preferences = original.Preferences
type PreferencesValidationRequest = original.PreferencesValidationRequest
type PreferencesValidationResponseProperties = original.PreferencesValidationResponseProperties
type RegionConfigurationRequest = original.RegionConfigurationRequest
type RegionConfigurationResponse = original.RegionConfigurationResponse
type Resource = original.Resource
type ScheduleAvailabilityRequest = original.ScheduleAvailabilityRequest
type ScheduleAvailabilityRequestType = original.ScheduleAvailabilityRequestType
type ScheduleAvailabilityResponse = original.ScheduleAvailabilityResponse
type Secret = original.Secret
type ServiceClient = original.ServiceClient
type ShareCredentialDetails = original.ShareCredentialDetails
type ShipmentPickUpRequest = original.ShipmentPickUpRequest
type ShipmentPickUpResponse = original.ShipmentPickUpResponse
type ShippingAddress = original.ShippingAddress
type Sku = original.Sku
type SkuAvailabilityValidationRequest = original.SkuAvailabilityValidationRequest
type SkuAvailabilityValidationResponseProperties = original.SkuAvailabilityValidationResponseProperties
type SkuCapacity = original.SkuCapacity
type SkuCost = original.SkuCost
type SkuInformation = original.SkuInformation
type SkuProperties = original.SkuProperties
type SubscriptionIsAllowedToCreateJobValidationRequest = original.SubscriptionIsAllowedToCreateJobValidationRequest
type SubscriptionIsAllowedToCreateJobValidationResponseProperties = original.SubscriptionIsAllowedToCreateJobValidationResponseProperties
type TransportAvailabilityDetails = original.TransportAvailabilityDetails
type TransportAvailabilityRequest = original.TransportAvailabilityRequest
type TransportAvailabilityResponse = original.TransportAvailabilityResponse
type TransportPreferences = original.TransportPreferences
type UnencryptedCredentials = original.UnencryptedCredentials
type UnencryptedCredentialsList = original.UnencryptedCredentialsList
type UpdateJobDetails = original.UpdateJobDetails
type UpdateJobProperties = original.UpdateJobProperties
type ValidateAddress = original.ValidateAddress
type ValidationInputRequest = original.ValidationInputRequest
type ValidationInputResponse = original.ValidationInputResponse
type ValidationRequest = original.ValidationRequest
type ValidationResponse = original.ValidationResponse
type ValidationResponseProperties = original.ValidationResponseProperties

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAvailableSkusResultIterator(page AvailableSkusResultPage) AvailableSkusResultIterator {
	return original.NewAvailableSkusResultIterator(page)
}
func NewAvailableSkusResultPage(getNextPage func(context.Context, AvailableSkusResult) (AvailableSkusResult, error)) AvailableSkusResultPage {
	return original.NewAvailableSkusResultPage(getNextPage)
}
func NewJobResourceListIterator(page JobResourceListPage) JobResourceListIterator {
	return original.NewJobResourceListIterator(page)
}
func NewJobResourceListPage(getNextPage func(context.Context, JobResourceList) (JobResourceList, error)) JobResourceListPage {
	return original.NewJobResourceListPage(getNextPage)
}
func NewJobsClient(subscriptionID string) JobsClient {
	return original.NewJobsClient(subscriptionID)
}
func NewJobsClientWithBaseURI(baseURI string, subscriptionID string) JobsClient {
	return original.NewJobsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListIterator(page OperationListPage) OperationListIterator {
	return original.NewOperationListIterator(page)
}
func NewOperationListPage(getNextPage func(context.Context, OperationList) (OperationList, error)) OperationListPage {
	return original.NewOperationListPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewServiceClient(subscriptionID string) ServiceClient {
	return original.NewServiceClient(subscriptionID)
}
func NewServiceClientWithBaseURI(baseURI string, subscriptionID string) ServiceClient {
	return original.NewServiceClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccessProtocolValues() []AccessProtocol {
	return original.PossibleAccessProtocolValues()
}
func PossibleAddressTypeValues() []AddressType {
	return original.PossibleAddressTypeValues()
}
func PossibleAddressValidationStatusValues() []AddressValidationStatus {
	return original.PossibleAddressValidationStatusValues()
}
func PossibleCopyLogDetailsTypeValues() []CopyLogDetailsType {
	return original.PossibleCopyLogDetailsTypeValues()
}
func PossibleCopyStatusValues() []CopyStatus {
	return original.PossibleCopyStatusValues()
}
func PossibleDataDestinationTypeBasicDestinationAccountDetailsValues() []DataDestinationTypeBasicDestinationAccountDetails {
	return original.PossibleDataDestinationTypeBasicDestinationAccountDetailsValues()
}
func PossibleDataDestinationTypeValues() []DataDestinationType {
	return original.PossibleDataDestinationTypeValues()
}
func PossibleJobDeliveryTypeValues() []JobDeliveryType {
	return original.PossibleJobDeliveryTypeValues()
}
func PossibleJobDetailsTypeEnumValues() []JobDetailsTypeEnum {
	return original.PossibleJobDetailsTypeEnumValues()
}
func PossibleJobSecretsTypeEnumValues() []JobSecretsTypeEnum {
	return original.PossibleJobSecretsTypeEnumValues()
}
func PossibleNotificationStageNameValues() []NotificationStageName {
	return original.PossibleNotificationStageNameValues()
}
func PossibleOverallValidationStatusValues() []OverallValidationStatus {
	return original.PossibleOverallValidationStatusValues()
}
func PossibleShareDestinationFormatTypeValues() []ShareDestinationFormatType {
	return original.PossibleShareDestinationFormatTypeValues()
}
func PossibleSkuDisabledReasonValues() []SkuDisabledReason {
	return original.PossibleSkuDisabledReasonValues()
}
func PossibleSkuNameBasicScheduleAvailabilityRequestValues() []SkuNameBasicScheduleAvailabilityRequest {
	return original.PossibleSkuNameBasicScheduleAvailabilityRequestValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleStageNameValues() []StageName {
	return original.PossibleStageNameValues()
}
func PossibleStageStatusValues() []StageStatus {
	return original.PossibleStageStatusValues()
}
func PossibleTransportShipmentTypesValues() []TransportShipmentTypes {
	return original.PossibleTransportShipmentTypesValues()
}
func PossibleValidationCategoryValues() []ValidationCategory {
	return original.PossibleValidationCategoryValues()
}
func PossibleValidationStatusValues() []ValidationStatus {
	return original.PossibleValidationStatusValues()
}
func PossibleValidationTypeBasicValidationInputResponseValues() []ValidationTypeBasicValidationInputResponse {
	return original.PossibleValidationTypeBasicValidationInputResponseValues()
}
func PossibleValidationTypeValues() []ValidationType {
	return original.PossibleValidationTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
