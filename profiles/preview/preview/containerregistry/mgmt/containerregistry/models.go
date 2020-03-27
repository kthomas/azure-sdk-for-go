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

package containerregistry

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/containerregistry/mgmt/2019-12-01-preview/containerregistry"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type Action = original.Action

const (
	Allow Action = original.Allow
)

type ActionsRequired = original.ActionsRequired

const (
	None     ActionsRequired = original.None
	Recreate ActionsRequired = original.Recreate
)

type DefaultAction = original.DefaultAction

const (
	DefaultActionAllow DefaultAction = original.DefaultActionAllow
	DefaultActionDeny  DefaultAction = original.DefaultActionDeny
)

type EncryptionStatus = original.EncryptionStatus

const (
	Disabled EncryptionStatus = original.Disabled
	Enabled  EncryptionStatus = original.Enabled
)

type ImportMode = original.ImportMode

const (
	Force   ImportMode = original.Force
	NoForce ImportMode = original.NoForce
)

type PasswordName = original.PasswordName

const (
	Password  PasswordName = original.Password
	Password2 PasswordName = original.Password2
)

type PolicyStatus = original.PolicyStatus

const (
	PolicyStatusDisabled PolicyStatus = original.PolicyStatusDisabled
	PolicyStatusEnabled  PolicyStatus = original.PolicyStatusEnabled
)

type ProvisioningState = original.ProvisioningState

const (
	Canceled  ProvisioningState = original.Canceled
	Creating  ProvisioningState = original.Creating
	Deleting  ProvisioningState = original.Deleting
	Failed    ProvisioningState = original.Failed
	Succeeded ProvisioningState = original.Succeeded
	Updating  ProvisioningState = original.Updating
)

type RegistryUsageUnit = original.RegistryUsageUnit

const (
	Bytes RegistryUsageUnit = original.Bytes
	Count RegistryUsageUnit = original.Count
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	ResourceIdentityTypeNone                       ResourceIdentityType = original.ResourceIdentityTypeNone
	ResourceIdentityTypeSystemAssigned             ResourceIdentityType = original.ResourceIdentityTypeSystemAssigned
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = original.ResourceIdentityTypeSystemAssignedUserAssigned
	ResourceIdentityTypeUserAssigned               ResourceIdentityType = original.ResourceIdentityTypeUserAssigned
)

type SkuName = original.SkuName

const (
	Basic    SkuName = original.Basic
	Classic  SkuName = original.Classic
	Premium  SkuName = original.Premium
	Standard SkuName = original.Standard
)

type SkuTier = original.SkuTier

const (
	SkuTierBasic    SkuTier = original.SkuTierBasic
	SkuTierClassic  SkuTier = original.SkuTierClassic
	SkuTierPremium  SkuTier = original.SkuTierPremium
	SkuTierStandard SkuTier = original.SkuTierStandard
)

type Status = original.Status

const (
	Approved     Status = original.Approved
	Disconnected Status = original.Disconnected
	Pending      Status = original.Pending
	Rejected     Status = original.Rejected
)

type TrustPolicyType = original.TrustPolicyType

const (
	Notary TrustPolicyType = original.Notary
)

type WebhookAction = original.WebhookAction

const (
	ChartDelete WebhookAction = original.ChartDelete
	ChartPush   WebhookAction = original.ChartPush
	Delete      WebhookAction = original.Delete
	Push        WebhookAction = original.Push
	Quarantine  WebhookAction = original.Quarantine
)

type WebhookStatus = original.WebhookStatus

const (
	WebhookStatusDisabled WebhookStatus = original.WebhookStatusDisabled
	WebhookStatusEnabled  WebhookStatus = original.WebhookStatusEnabled
)

type Actor = original.Actor
type BaseClient = original.BaseClient
type CallbackConfig = original.CallbackConfig
type EncryptionProperty = original.EncryptionProperty
type Event = original.Event
type EventContent = original.EventContent
type EventInfo = original.EventInfo
type EventListResult = original.EventListResult
type EventListResultIterator = original.EventListResultIterator
type EventListResultPage = original.EventListResultPage
type EventRequestMessage = original.EventRequestMessage
type EventResponseMessage = original.EventResponseMessage
type IPRule = original.IPRule
type IdentityProperties = original.IdentityProperties
type ImportImageParameters = original.ImportImageParameters
type ImportSource = original.ImportSource
type ImportSourceCredentials = original.ImportSourceCredentials
type KeyVaultProperties = original.KeyVaultProperties
type NetworkRuleSet = original.NetworkRuleSet
type OperationDefinition = original.OperationDefinition
type OperationDisplayDefinition = original.OperationDisplayDefinition
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationMetricSpecificationDefinition = original.OperationMetricSpecificationDefinition
type OperationPropertiesDefinition = original.OperationPropertiesDefinition
type OperationServiceSpecificationDefinition = original.OperationServiceSpecificationDefinition
type OperationsClient = original.OperationsClient
type Policies = original.Policies
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionListResult = original.PrivateEndpointConnectionListResult
type PrivateEndpointConnectionListResultIterator = original.PrivateEndpointConnectionListResultIterator
type PrivateEndpointConnectionListResultPage = original.PrivateEndpointConnectionListResultPage
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateEndpointConnectionsCreateOrUpdateFuture = original.PrivateEndpointConnectionsCreateOrUpdateFuture
type PrivateEndpointConnectionsDeleteFuture = original.PrivateEndpointConnectionsDeleteFuture
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourceListResult = original.PrivateLinkResourceListResult
type PrivateLinkResourceListResultIterator = original.PrivateLinkResourceListResultIterator
type PrivateLinkResourceListResultPage = original.PrivateLinkResourceListResultPage
type PrivateLinkResourceProperties = original.PrivateLinkResourceProperties
type PrivateLinkServiceConnectionState = original.PrivateLinkServiceConnectionState
type QuarantinePolicy = original.QuarantinePolicy
type RegenerateCredentialParameters = original.RegenerateCredentialParameters
type RegistriesClient = original.RegistriesClient
type RegistriesCreateFuture = original.RegistriesCreateFuture
type RegistriesDeleteFuture = original.RegistriesDeleteFuture
type RegistriesImportImageFuture = original.RegistriesImportImageFuture
type RegistriesUpdateFuture = original.RegistriesUpdateFuture
type Registry = original.Registry
type RegistryListCredentialsResult = original.RegistryListCredentialsResult
type RegistryListResult = original.RegistryListResult
type RegistryListResultIterator = original.RegistryListResultIterator
type RegistryListResultPage = original.RegistryListResultPage
type RegistryNameCheckRequest = original.RegistryNameCheckRequest
type RegistryNameStatus = original.RegistryNameStatus
type RegistryPassword = original.RegistryPassword
type RegistryProperties = original.RegistryProperties
type RegistryPropertiesUpdateParameters = original.RegistryPropertiesUpdateParameters
type RegistryUpdateParameters = original.RegistryUpdateParameters
type RegistryUsage = original.RegistryUsage
type RegistryUsageListResult = original.RegistryUsageListResult
type Replication = original.Replication
type ReplicationListResult = original.ReplicationListResult
type ReplicationListResultIterator = original.ReplicationListResultIterator
type ReplicationListResultPage = original.ReplicationListResultPage
type ReplicationProperties = original.ReplicationProperties
type ReplicationUpdateParameters = original.ReplicationUpdateParameters
type ReplicationsClient = original.ReplicationsClient
type ReplicationsCreateFuture = original.ReplicationsCreateFuture
type ReplicationsDeleteFuture = original.ReplicationsDeleteFuture
type ReplicationsUpdateFuture = original.ReplicationsUpdateFuture
type Request = original.Request
type Resource = original.Resource
type RetentionPolicy = original.RetentionPolicy
type Sku = original.Sku
type Source = original.Source
type Status1 = original.Status1
type StorageAccountProperties = original.StorageAccountProperties
type Target = original.Target
type TrustPolicy = original.TrustPolicy
type UserIdentityProperties = original.UserIdentityProperties
type VirtualNetworkRule = original.VirtualNetworkRule
type Webhook = original.Webhook
type WebhookCreateParameters = original.WebhookCreateParameters
type WebhookListResult = original.WebhookListResult
type WebhookListResultIterator = original.WebhookListResultIterator
type WebhookListResultPage = original.WebhookListResultPage
type WebhookProperties = original.WebhookProperties
type WebhookPropertiesCreateParameters = original.WebhookPropertiesCreateParameters
type WebhookPropertiesUpdateParameters = original.WebhookPropertiesUpdateParameters
type WebhookUpdateParameters = original.WebhookUpdateParameters
type WebhooksClient = original.WebhooksClient
type WebhooksCreateFuture = original.WebhooksCreateFuture
type WebhooksDeleteFuture = original.WebhooksDeleteFuture
type WebhooksUpdateFuture = original.WebhooksUpdateFuture

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewEventListResultIterator(page EventListResultPage) EventListResultIterator {
	return original.NewEventListResultIterator(page)
}
func NewEventListResultPage(getNextPage func(context.Context, EventListResult) (EventListResult, error)) EventListResultPage {
	return original.NewEventListResultPage(getNextPage)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateEndpointConnectionListResultIterator(page PrivateEndpointConnectionListResultPage) PrivateEndpointConnectionListResultIterator {
	return original.NewPrivateEndpointConnectionListResultIterator(page)
}
func NewPrivateEndpointConnectionListResultPage(getNextPage func(context.Context, PrivateEndpointConnectionListResult) (PrivateEndpointConnectionListResult, error)) PrivateEndpointConnectionListResultPage {
	return original.NewPrivateEndpointConnectionListResultPage(getNextPage)
}
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClient(subscriptionID)
}
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkResourceListResultIterator(page PrivateLinkResourceListResultPage) PrivateLinkResourceListResultIterator {
	return original.NewPrivateLinkResourceListResultIterator(page)
}
func NewPrivateLinkResourceListResultPage(getNextPage func(context.Context, PrivateLinkResourceListResult) (PrivateLinkResourceListResult, error)) PrivateLinkResourceListResultPage {
	return original.NewPrivateLinkResourceListResultPage(getNextPage)
}
func NewRegistriesClient(subscriptionID string) RegistriesClient {
	return original.NewRegistriesClient(subscriptionID)
}
func NewRegistriesClientWithBaseURI(baseURI string, subscriptionID string) RegistriesClient {
	return original.NewRegistriesClientWithBaseURI(baseURI, subscriptionID)
}
func NewRegistryListResultIterator(page RegistryListResultPage) RegistryListResultIterator {
	return original.NewRegistryListResultIterator(page)
}
func NewRegistryListResultPage(getNextPage func(context.Context, RegistryListResult) (RegistryListResult, error)) RegistryListResultPage {
	return original.NewRegistryListResultPage(getNextPage)
}
func NewReplicationListResultIterator(page ReplicationListResultPage) ReplicationListResultIterator {
	return original.NewReplicationListResultIterator(page)
}
func NewReplicationListResultPage(getNextPage func(context.Context, ReplicationListResult) (ReplicationListResult, error)) ReplicationListResultPage {
	return original.NewReplicationListResultPage(getNextPage)
}
func NewReplicationsClient(subscriptionID string) ReplicationsClient {
	return original.NewReplicationsClient(subscriptionID)
}
func NewReplicationsClientWithBaseURI(baseURI string, subscriptionID string) ReplicationsClient {
	return original.NewReplicationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWebhookListResultIterator(page WebhookListResultPage) WebhookListResultIterator {
	return original.NewWebhookListResultIterator(page)
}
func NewWebhookListResultPage(getNextPage func(context.Context, WebhookListResult) (WebhookListResult, error)) WebhookListResultPage {
	return original.NewWebhookListResultPage(getNextPage)
}
func NewWebhooksClient(subscriptionID string) WebhooksClient {
	return original.NewWebhooksClient(subscriptionID)
}
func NewWebhooksClientWithBaseURI(baseURI string, subscriptionID string) WebhooksClient {
	return original.NewWebhooksClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleActionValues() []Action {
	return original.PossibleActionValues()
}
func PossibleActionsRequiredValues() []ActionsRequired {
	return original.PossibleActionsRequiredValues()
}
func PossibleDefaultActionValues() []DefaultAction {
	return original.PossibleDefaultActionValues()
}
func PossibleEncryptionStatusValues() []EncryptionStatus {
	return original.PossibleEncryptionStatusValues()
}
func PossibleImportModeValues() []ImportMode {
	return original.PossibleImportModeValues()
}
func PossiblePasswordNameValues() []PasswordName {
	return original.PossiblePasswordNameValues()
}
func PossiblePolicyStatusValues() []PolicyStatus {
	return original.PossiblePolicyStatusValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleRegistryUsageUnitValues() []RegistryUsageUnit {
	return original.PossibleRegistryUsageUnitValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func PossibleTrustPolicyTypeValues() []TrustPolicyType {
	return original.PossibleTrustPolicyTypeValues()
}
func PossibleWebhookActionValues() []WebhookAction {
	return original.PossibleWebhookActionValues()
}
func PossibleWebhookStatusValues() []WebhookStatus {
	return original.PossibleWebhookStatusValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
