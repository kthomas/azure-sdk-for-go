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

package alertsmanagement

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/alertsmanagement/mgmt/2019-03-01/alertsmanagement"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AlertModificationEvent = original.AlertModificationEvent

const (
	AlertCreated           AlertModificationEvent = original.AlertCreated
	MonitorConditionChange AlertModificationEvent = original.MonitorConditionChange
	StateChange            AlertModificationEvent = original.StateChange
)

type AlertRuleState = original.AlertRuleState

const (
	Disabled AlertRuleState = original.Disabled
	Enabled  AlertRuleState = original.Enabled
)

type AlertState = original.AlertState

const (
	AlertStateAcknowledged AlertState = original.AlertStateAcknowledged
	AlertStateClosed       AlertState = original.AlertStateClosed
	AlertStateNew          AlertState = original.AlertStateNew
)

type AlertsSortByFields = original.AlertsSortByFields

const (
	AlertsSortByFieldsAlertState           AlertsSortByFields = original.AlertsSortByFieldsAlertState
	AlertsSortByFieldsLastModifiedDateTime AlertsSortByFields = original.AlertsSortByFieldsLastModifiedDateTime
	AlertsSortByFieldsMonitorCondition     AlertsSortByFields = original.AlertsSortByFieldsMonitorCondition
	AlertsSortByFieldsName                 AlertsSortByFields = original.AlertsSortByFieldsName
	AlertsSortByFieldsSeverity             AlertsSortByFields = original.AlertsSortByFieldsSeverity
	AlertsSortByFieldsStartDateTime        AlertsSortByFields = original.AlertsSortByFieldsStartDateTime
	AlertsSortByFieldsTargetResource       AlertsSortByFields = original.AlertsSortByFieldsTargetResource
	AlertsSortByFieldsTargetResourceGroup  AlertsSortByFields = original.AlertsSortByFieldsTargetResourceGroup
	AlertsSortByFieldsTargetResourceName   AlertsSortByFields = original.AlertsSortByFieldsTargetResourceName
	AlertsSortByFieldsTargetResourceType   AlertsSortByFields = original.AlertsSortByFieldsTargetResourceType
)

type AlertsSummaryGroupByFields = original.AlertsSummaryGroupByFields

const (
	AlertsSummaryGroupByFieldsAlertRule        AlertsSummaryGroupByFields = original.AlertsSummaryGroupByFieldsAlertRule
	AlertsSummaryGroupByFieldsAlertState       AlertsSummaryGroupByFields = original.AlertsSummaryGroupByFieldsAlertState
	AlertsSummaryGroupByFieldsMonitorCondition AlertsSummaryGroupByFields = original.AlertsSummaryGroupByFieldsMonitorCondition
	AlertsSummaryGroupByFieldsMonitorService   AlertsSummaryGroupByFields = original.AlertsSummaryGroupByFieldsMonitorService
	AlertsSummaryGroupByFieldsSeverity         AlertsSummaryGroupByFields = original.AlertsSummaryGroupByFieldsSeverity
	AlertsSummaryGroupByFieldsSignalType       AlertsSummaryGroupByFields = original.AlertsSummaryGroupByFieldsSignalType
)

type MonitorCondition = original.MonitorCondition

const (
	Fired    MonitorCondition = original.Fired
	Resolved MonitorCondition = original.Resolved
)

type MonitorService = original.MonitorService

const (
	ActivityLogAdministrative MonitorService = original.ActivityLogAdministrative
	ActivityLogAutoscale      MonitorService = original.ActivityLogAutoscale
	ActivityLogPolicy         MonitorService = original.ActivityLogPolicy
	ActivityLogRecommendation MonitorService = original.ActivityLogRecommendation
	ActivityLogSecurity       MonitorService = original.ActivityLogSecurity
	ApplicationInsights       MonitorService = original.ApplicationInsights
	LogAnalytics              MonitorService = original.LogAnalytics
	Nagios                    MonitorService = original.Nagios
	Platform                  MonitorService = original.Platform
	ResourceHealth            MonitorService = original.ResourceHealth
	SCOM                      MonitorService = original.SCOM
	ServiceHealth             MonitorService = original.ServiceHealth
	SmartDetector             MonitorService = original.SmartDetector
	VMInsights                MonitorService = original.VMInsights
	Zabbix                    MonitorService = original.Zabbix
)

type Severity = original.Severity

const (
	Sev0 Severity = original.Sev0
	Sev1 Severity = original.Sev1
	Sev2 Severity = original.Sev2
	Sev3 Severity = original.Sev3
	Sev4 Severity = original.Sev4
)

type SignalType = original.SignalType

const (
	Log     SignalType = original.Log
	Metric  SignalType = original.Metric
	Unknown SignalType = original.Unknown
)

type SmartGroupModificationEvent = original.SmartGroupModificationEvent

const (
	SmartGroupModificationEventAlertAdded        SmartGroupModificationEvent = original.SmartGroupModificationEventAlertAdded
	SmartGroupModificationEventAlertRemoved      SmartGroupModificationEvent = original.SmartGroupModificationEventAlertRemoved
	SmartGroupModificationEventSmartGroupCreated SmartGroupModificationEvent = original.SmartGroupModificationEventSmartGroupCreated
	SmartGroupModificationEventStateChange       SmartGroupModificationEvent = original.SmartGroupModificationEventStateChange
)

type SmartGroupsSortByFields = original.SmartGroupsSortByFields

const (
	SmartGroupsSortByFieldsAlertsCount          SmartGroupsSortByFields = original.SmartGroupsSortByFieldsAlertsCount
	SmartGroupsSortByFieldsLastModifiedDateTime SmartGroupsSortByFields = original.SmartGroupsSortByFieldsLastModifiedDateTime
	SmartGroupsSortByFieldsSeverity             SmartGroupsSortByFields = original.SmartGroupsSortByFieldsSeverity
	SmartGroupsSortByFieldsStartDateTime        SmartGroupsSortByFields = original.SmartGroupsSortByFieldsStartDateTime
	SmartGroupsSortByFieldsState                SmartGroupsSortByFields = original.SmartGroupsSortByFieldsState
)

type State = original.State

const (
	StateAcknowledged State = original.StateAcknowledged
	StateClosed       State = original.StateClosed
	StateNew          State = original.StateNew
)

type TimeRange = original.TimeRange

const (
	Oned       TimeRange = original.Oned
	Oneh       TimeRange = original.Oneh
	Sevend     TimeRange = original.Sevend
	ThreeZerod TimeRange = original.ThreeZerod
)

type ActionGroupsInformation = original.ActionGroupsInformation
type Alert = original.Alert
type AlertModification = original.AlertModification
type AlertModificationItem = original.AlertModificationItem
type AlertModificationProperties = original.AlertModificationProperties
type AlertProperties = original.AlertProperties
type AlertRule = original.AlertRule
type AlertRuleProperties = original.AlertRuleProperties
type AlertRulesList = original.AlertRulesList
type AlertRulesListIterator = original.AlertRulesListIterator
type AlertRulesListPage = original.AlertRulesListPage
type AlertsClient = original.AlertsClient
type AlertsList = original.AlertsList
type AlertsListIterator = original.AlertsListIterator
type AlertsListPage = original.AlertsListPage
type AlertsSummary = original.AlertsSummary
type AlertsSummaryGroup = original.AlertsSummaryGroup
type AlertsSummaryGroupItem = original.AlertsSummaryGroupItem
type AzureResource = original.AzureResource
type BaseClient = original.BaseClient
type Detector = original.Detector
type ErrorResponse = original.ErrorResponse
type ErrorResponse1 = original.ErrorResponse1
type ErrorResponseBody = original.ErrorResponseBody
type Essentials = original.Essentials
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationsClient = original.OperationsClient
type OperationsList = original.OperationsList
type OperationsListIterator = original.OperationsListIterator
type OperationsListPage = original.OperationsListPage
type ProxyResource = original.ProxyResource
type SmartDetectorAlertRulesClient = original.SmartDetectorAlertRulesClient
type SmartGroup = original.SmartGroup
type SmartGroupAggregatedProperty = original.SmartGroupAggregatedProperty
type SmartGroupModification = original.SmartGroupModification
type SmartGroupModificationItem = original.SmartGroupModificationItem
type SmartGroupModificationProperties = original.SmartGroupModificationProperties
type SmartGroupProperties = original.SmartGroupProperties
type SmartGroupsClient = original.SmartGroupsClient
type SmartGroupsList = original.SmartGroupsList
type ThrottlingInformation = original.ThrottlingInformation

func New(scope string, subscriptionID string, subscriptionID1 string) BaseClient {
	return original.New(scope, subscriptionID, subscriptionID1)
}
func NewAlertRulesListIterator(page AlertRulesListPage) AlertRulesListIterator {
	return original.NewAlertRulesListIterator(page)
}
func NewAlertRulesListPage(getNextPage func(context.Context, AlertRulesList) (AlertRulesList, error)) AlertRulesListPage {
	return original.NewAlertRulesListPage(getNextPage)
}
func NewAlertsClient(scope string, subscriptionID string, subscriptionID1 string) AlertsClient {
	return original.NewAlertsClient(scope, subscriptionID, subscriptionID1)
}
func NewAlertsClientWithBaseURI(baseURI string, scope string, subscriptionID string, subscriptionID1 string) AlertsClient {
	return original.NewAlertsClientWithBaseURI(baseURI, scope, subscriptionID, subscriptionID1)
}
func NewAlertsListIterator(page AlertsListPage) AlertsListIterator {
	return original.NewAlertsListIterator(page)
}
func NewAlertsListPage(getNextPage func(context.Context, AlertsList) (AlertsList, error)) AlertsListPage {
	return original.NewAlertsListPage(getNextPage)
}
func NewOperationsClient(scope string, subscriptionID string, subscriptionID1 string) OperationsClient {
	return original.NewOperationsClient(scope, subscriptionID, subscriptionID1)
}
func NewOperationsClientWithBaseURI(baseURI string, scope string, subscriptionID string, subscriptionID1 string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, scope, subscriptionID, subscriptionID1)
}
func NewOperationsListIterator(page OperationsListPage) OperationsListIterator {
	return original.NewOperationsListIterator(page)
}
func NewOperationsListPage(getNextPage func(context.Context, OperationsList) (OperationsList, error)) OperationsListPage {
	return original.NewOperationsListPage(getNextPage)
}
func NewSmartDetectorAlertRulesClient(scope string, subscriptionID string, subscriptionID1 string) SmartDetectorAlertRulesClient {
	return original.NewSmartDetectorAlertRulesClient(scope, subscriptionID, subscriptionID1)
}
func NewSmartDetectorAlertRulesClientWithBaseURI(baseURI string, scope string, subscriptionID string, subscriptionID1 string) SmartDetectorAlertRulesClient {
	return original.NewSmartDetectorAlertRulesClientWithBaseURI(baseURI, scope, subscriptionID, subscriptionID1)
}
func NewSmartGroupsClient(scope string, subscriptionID string, subscriptionID1 string) SmartGroupsClient {
	return original.NewSmartGroupsClient(scope, subscriptionID, subscriptionID1)
}
func NewSmartGroupsClientWithBaseURI(baseURI string, scope string, subscriptionID string, subscriptionID1 string) SmartGroupsClient {
	return original.NewSmartGroupsClientWithBaseURI(baseURI, scope, subscriptionID, subscriptionID1)
}
func NewWithBaseURI(baseURI string, scope string, subscriptionID string, subscriptionID1 string) BaseClient {
	return original.NewWithBaseURI(baseURI, scope, subscriptionID, subscriptionID1)
}
func PossibleAlertModificationEventValues() []AlertModificationEvent {
	return original.PossibleAlertModificationEventValues()
}
func PossibleAlertRuleStateValues() []AlertRuleState {
	return original.PossibleAlertRuleStateValues()
}
func PossibleAlertStateValues() []AlertState {
	return original.PossibleAlertStateValues()
}
func PossibleAlertsSortByFieldsValues() []AlertsSortByFields {
	return original.PossibleAlertsSortByFieldsValues()
}
func PossibleAlertsSummaryGroupByFieldsValues() []AlertsSummaryGroupByFields {
	return original.PossibleAlertsSummaryGroupByFieldsValues()
}
func PossibleMonitorConditionValues() []MonitorCondition {
	return original.PossibleMonitorConditionValues()
}
func PossibleMonitorServiceValues() []MonitorService {
	return original.PossibleMonitorServiceValues()
}
func PossibleSeverityValues() []Severity {
	return original.PossibleSeverityValues()
}
func PossibleSignalTypeValues() []SignalType {
	return original.PossibleSignalTypeValues()
}
func PossibleSmartGroupModificationEventValues() []SmartGroupModificationEvent {
	return original.PossibleSmartGroupModificationEventValues()
}
func PossibleSmartGroupsSortByFieldsValues() []SmartGroupsSortByFields {
	return original.PossibleSmartGroupsSortByFieldsValues()
}
func PossibleStateValues() []State {
	return original.PossibleStateValues()
}
func PossibleTimeRangeValues() []TimeRange {
	return original.PossibleTimeRangeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
