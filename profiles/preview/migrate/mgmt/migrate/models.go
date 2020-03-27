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

package migrate

import original "github.com/Azure/azure-sdk-for-go/services/migrate/mgmt/2018-02-02/migrate"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AssessmentSizingCriterion = original.AssessmentSizingCriterion

const (
	AsOnPremises     AssessmentSizingCriterion = original.AsOnPremises
	PerformanceBased AssessmentSizingCriterion = original.PerformanceBased
)

type AssessmentStage = original.AssessmentStage

const (
	Approved    AssessmentStage = original.Approved
	InProgress  AssessmentStage = original.InProgress
	UnderReview AssessmentStage = original.UnderReview
)

type AssessmentStatus = original.AssessmentStatus

const (
	Completed AssessmentStatus = original.Completed
	Created   AssessmentStatus = original.Created
	Invalid   AssessmentStatus = original.Invalid
	Running   AssessmentStatus = original.Running
	Updated   AssessmentStatus = original.Updated
)

type AzureDiskSize = original.AzureDiskSize

const (
	PremiumP10  AzureDiskSize = original.PremiumP10
	PremiumP20  AzureDiskSize = original.PremiumP20
	PremiumP30  AzureDiskSize = original.PremiumP30
	PremiumP4   AzureDiskSize = original.PremiumP4
	PremiumP40  AzureDiskSize = original.PremiumP40
	PremiumP50  AzureDiskSize = original.PremiumP50
	PremiumP6   AzureDiskSize = original.PremiumP6
	StandardS10 AzureDiskSize = original.StandardS10
	StandardS20 AzureDiskSize = original.StandardS20
	StandardS30 AzureDiskSize = original.StandardS30
	StandardS4  AzureDiskSize = original.StandardS4
	StandardS40 AzureDiskSize = original.StandardS40
	StandardS50 AzureDiskSize = original.StandardS50
	StandardS6  AzureDiskSize = original.StandardS6
	Unknown     AzureDiskSize = original.Unknown
)

type AzureDiskSuitabilityExplanation = original.AzureDiskSuitabilityExplanation

const (
	AzureDiskSuitabilityExplanationDiskSizeGreaterThanSupported           AzureDiskSuitabilityExplanation = original.AzureDiskSuitabilityExplanationDiskSizeGreaterThanSupported
	AzureDiskSuitabilityExplanationInternalErrorOccurredForDiskEvaluation AzureDiskSuitabilityExplanation = original.AzureDiskSuitabilityExplanationInternalErrorOccurredForDiskEvaluation
	AzureDiskSuitabilityExplanationNoDiskSizeFoundForSelectedRedundancy   AzureDiskSuitabilityExplanation = original.AzureDiskSuitabilityExplanationNoDiskSizeFoundForSelectedRedundancy
	AzureDiskSuitabilityExplanationNoDiskSizeFoundInSelectedLocation      AzureDiskSuitabilityExplanation = original.AzureDiskSuitabilityExplanationNoDiskSizeFoundInSelectedLocation
	AzureDiskSuitabilityExplanationNoSuitableDiskSizeForIops              AzureDiskSuitabilityExplanation = original.AzureDiskSuitabilityExplanationNoSuitableDiskSizeForIops
	AzureDiskSuitabilityExplanationNoSuitableDiskSizeForThroughput        AzureDiskSuitabilityExplanation = original.AzureDiskSuitabilityExplanationNoSuitableDiskSizeForThroughput
	AzureDiskSuitabilityExplanationNotApplicable                          AzureDiskSuitabilityExplanation = original.AzureDiskSuitabilityExplanationNotApplicable
	AzureDiskSuitabilityExplanationUnknown                                AzureDiskSuitabilityExplanation = original.AzureDiskSuitabilityExplanationUnknown
)

type AzureDiskType = original.AzureDiskType

const (
	AzureDiskTypePremium  AzureDiskType = original.AzureDiskTypePremium
	AzureDiskTypeStandard AzureDiskType = original.AzureDiskTypeStandard
	AzureDiskTypeUnknown  AzureDiskType = original.AzureDiskTypeUnknown
)

type AzureHybridUseBenefit = original.AzureHybridUseBenefit

const (
	AzureHybridUseBenefitNo      AzureHybridUseBenefit = original.AzureHybridUseBenefitNo
	AzureHybridUseBenefitUnknown AzureHybridUseBenefit = original.AzureHybridUseBenefitUnknown
	AzureHybridUseBenefitYes     AzureHybridUseBenefit = original.AzureHybridUseBenefitYes
)

type AzureLocation = original.AzureLocation

const (
	AzureLocationAustraliaEast      AzureLocation = original.AzureLocationAustraliaEast
	AzureLocationAustraliaSoutheast AzureLocation = original.AzureLocationAustraliaSoutheast
	AzureLocationBrazilSouth        AzureLocation = original.AzureLocationBrazilSouth
	AzureLocationCanadaCentral      AzureLocation = original.AzureLocationCanadaCentral
	AzureLocationCanadaEast         AzureLocation = original.AzureLocationCanadaEast
	AzureLocationCentralIndia       AzureLocation = original.AzureLocationCentralIndia
	AzureLocationCentralUs          AzureLocation = original.AzureLocationCentralUs
	AzureLocationChinaEast          AzureLocation = original.AzureLocationChinaEast
	AzureLocationChinaNorth         AzureLocation = original.AzureLocationChinaNorth
	AzureLocationEastAsia           AzureLocation = original.AzureLocationEastAsia
	AzureLocationEastUs             AzureLocation = original.AzureLocationEastUs
	AzureLocationEastUs2            AzureLocation = original.AzureLocationEastUs2
	AzureLocationGermanyCentral     AzureLocation = original.AzureLocationGermanyCentral
	AzureLocationGermanyNortheast   AzureLocation = original.AzureLocationGermanyNortheast
	AzureLocationJapanEast          AzureLocation = original.AzureLocationJapanEast
	AzureLocationJapanWest          AzureLocation = original.AzureLocationJapanWest
	AzureLocationKoreaCentral       AzureLocation = original.AzureLocationKoreaCentral
	AzureLocationKoreaSouth         AzureLocation = original.AzureLocationKoreaSouth
	AzureLocationNorthCentralUs     AzureLocation = original.AzureLocationNorthCentralUs
	AzureLocationNorthEurope        AzureLocation = original.AzureLocationNorthEurope
	AzureLocationSouthCentralUs     AzureLocation = original.AzureLocationSouthCentralUs
	AzureLocationSoutheastAsia      AzureLocation = original.AzureLocationSoutheastAsia
	AzureLocationSouthIndia         AzureLocation = original.AzureLocationSouthIndia
	AzureLocationUkSouth            AzureLocation = original.AzureLocationUkSouth
	AzureLocationUkWest             AzureLocation = original.AzureLocationUkWest
	AzureLocationUnknown            AzureLocation = original.AzureLocationUnknown
	AzureLocationWestCentralUs      AzureLocation = original.AzureLocationWestCentralUs
	AzureLocationWestEurope         AzureLocation = original.AzureLocationWestEurope
	AzureLocationWestIndia          AzureLocation = original.AzureLocationWestIndia
	AzureLocationWestUs             AzureLocation = original.AzureLocationWestUs
	AzureLocationWestUs2            AzureLocation = original.AzureLocationWestUs2
)

type AzureNetworkAdapterSuitabilityExplanation = original.AzureNetworkAdapterSuitabilityExplanation

const (
	AzureNetworkAdapterSuitabilityExplanationInternalErrorOccured AzureNetworkAdapterSuitabilityExplanation = original.AzureNetworkAdapterSuitabilityExplanationInternalErrorOccured
	AzureNetworkAdapterSuitabilityExplanationNotApplicable        AzureNetworkAdapterSuitabilityExplanation = original.AzureNetworkAdapterSuitabilityExplanationNotApplicable
	AzureNetworkAdapterSuitabilityExplanationUnknown              AzureNetworkAdapterSuitabilityExplanation = original.AzureNetworkAdapterSuitabilityExplanationUnknown
)

type AzureOfferCode = original.AzureOfferCode

const (
	AzureOfferCodeMSAZR0003P   AzureOfferCode = original.AzureOfferCodeMSAZR0003P
	AzureOfferCodeMSAZR0022P   AzureOfferCode = original.AzureOfferCodeMSAZR0022P
	AzureOfferCodeMSAZR0023P   AzureOfferCode = original.AzureOfferCodeMSAZR0023P
	AzureOfferCodeMSAZR0025P   AzureOfferCode = original.AzureOfferCodeMSAZR0025P
	AzureOfferCodeMSAZR0029P   AzureOfferCode = original.AzureOfferCodeMSAZR0029P
	AzureOfferCodeMSAZR0036P   AzureOfferCode = original.AzureOfferCodeMSAZR0036P
	AzureOfferCodeMSAZR0044P   AzureOfferCode = original.AzureOfferCodeMSAZR0044P
	AzureOfferCodeMSAZR0059P   AzureOfferCode = original.AzureOfferCodeMSAZR0059P
	AzureOfferCodeMSAZR0060P   AzureOfferCode = original.AzureOfferCodeMSAZR0060P
	AzureOfferCodeMSAZR0062P   AzureOfferCode = original.AzureOfferCodeMSAZR0062P
	AzureOfferCodeMSAZR0063P   AzureOfferCode = original.AzureOfferCodeMSAZR0063P
	AzureOfferCodeMSAZR0064P   AzureOfferCode = original.AzureOfferCodeMSAZR0064P
	AzureOfferCodeMSAZR0111P   AzureOfferCode = original.AzureOfferCodeMSAZR0111P
	AzureOfferCodeMSAZR0120P   AzureOfferCode = original.AzureOfferCodeMSAZR0120P
	AzureOfferCodeMSAZR0121P   AzureOfferCode = original.AzureOfferCodeMSAZR0121P
	AzureOfferCodeMSAZR0122P   AzureOfferCode = original.AzureOfferCodeMSAZR0122P
	AzureOfferCodeMSAZR0123P   AzureOfferCode = original.AzureOfferCodeMSAZR0123P
	AzureOfferCodeMSAZR0124P   AzureOfferCode = original.AzureOfferCodeMSAZR0124P
	AzureOfferCodeMSAZR0125P   AzureOfferCode = original.AzureOfferCodeMSAZR0125P
	AzureOfferCodeMSAZR0126P   AzureOfferCode = original.AzureOfferCodeMSAZR0126P
	AzureOfferCodeMSAZR0127P   AzureOfferCode = original.AzureOfferCodeMSAZR0127P
	AzureOfferCodeMSAZR0128P   AzureOfferCode = original.AzureOfferCodeMSAZR0128P
	AzureOfferCodeMSAZR0129P   AzureOfferCode = original.AzureOfferCodeMSAZR0129P
	AzureOfferCodeMSAZR0130P   AzureOfferCode = original.AzureOfferCodeMSAZR0130P
	AzureOfferCodeMSAZR0144P   AzureOfferCode = original.AzureOfferCodeMSAZR0144P
	AzureOfferCodeMSAZR0148P   AzureOfferCode = original.AzureOfferCodeMSAZR0148P
	AzureOfferCodeMSAZR0149P   AzureOfferCode = original.AzureOfferCodeMSAZR0149P
	AzureOfferCodeMSAZRDE0003P AzureOfferCode = original.AzureOfferCodeMSAZRDE0003P
	AzureOfferCodeMSAZRDE0044P AzureOfferCode = original.AzureOfferCodeMSAZRDE0044P
	AzureOfferCodeMSMCAZR0044P AzureOfferCode = original.AzureOfferCodeMSMCAZR0044P
	AzureOfferCodeMSMCAZR0059P AzureOfferCode = original.AzureOfferCodeMSMCAZR0059P
	AzureOfferCodeMSMCAZR0060P AzureOfferCode = original.AzureOfferCodeMSMCAZR0060P
	AzureOfferCodeMSMCAZR0063P AzureOfferCode = original.AzureOfferCodeMSMCAZR0063P
	AzureOfferCodeMSMCAZR0120P AzureOfferCode = original.AzureOfferCodeMSMCAZR0120P
	AzureOfferCodeMSMCAZR0121P AzureOfferCode = original.AzureOfferCodeMSMCAZR0121P
	AzureOfferCodeMSMCAZR0125P AzureOfferCode = original.AzureOfferCodeMSMCAZR0125P
	AzureOfferCodeMSMCAZR0128P AzureOfferCode = original.AzureOfferCodeMSMCAZR0128P
	AzureOfferCodeUnknown      AzureOfferCode = original.AzureOfferCodeUnknown
)

type AzurePricingTier = original.AzurePricingTier

const (
	Basic    AzurePricingTier = original.Basic
	Standard AzurePricingTier = original.Standard
)

type AzureStorageRedundancy = original.AzureStorageRedundancy

const (
	AzureStorageRedundancyGeoRedundant           AzureStorageRedundancy = original.AzureStorageRedundancyGeoRedundant
	AzureStorageRedundancyLocallyRedundant       AzureStorageRedundancy = original.AzureStorageRedundancyLocallyRedundant
	AzureStorageRedundancyReadAccessGeoRedundant AzureStorageRedundancy = original.AzureStorageRedundancyReadAccessGeoRedundant
	AzureStorageRedundancyUnknown                AzureStorageRedundancy = original.AzureStorageRedundancyUnknown
	AzureStorageRedundancyZoneRedundant          AzureStorageRedundancy = original.AzureStorageRedundancyZoneRedundant
)

type AzureVMSize = original.AzureVMSize

const (
	AzureVMSizeBasicA0        AzureVMSize = original.AzureVMSizeBasicA0
	AzureVMSizeBasicA1        AzureVMSize = original.AzureVMSizeBasicA1
	AzureVMSizeBasicA2        AzureVMSize = original.AzureVMSizeBasicA2
	AzureVMSizeBasicA3        AzureVMSize = original.AzureVMSizeBasicA3
	AzureVMSizeBasicA4        AzureVMSize = original.AzureVMSizeBasicA4
	AzureVMSizeStandardA0     AzureVMSize = original.AzureVMSizeStandardA0
	AzureVMSizeStandardA1     AzureVMSize = original.AzureVMSizeStandardA1
	AzureVMSizeStandardA10    AzureVMSize = original.AzureVMSizeStandardA10
	AzureVMSizeStandardA11    AzureVMSize = original.AzureVMSizeStandardA11
	AzureVMSizeStandardA1V2   AzureVMSize = original.AzureVMSizeStandardA1V2
	AzureVMSizeStandardA2     AzureVMSize = original.AzureVMSizeStandardA2
	AzureVMSizeStandardA2mV2  AzureVMSize = original.AzureVMSizeStandardA2mV2
	AzureVMSizeStandardA2V2   AzureVMSize = original.AzureVMSizeStandardA2V2
	AzureVMSizeStandardA3     AzureVMSize = original.AzureVMSizeStandardA3
	AzureVMSizeStandardA4     AzureVMSize = original.AzureVMSizeStandardA4
	AzureVMSizeStandardA4mV2  AzureVMSize = original.AzureVMSizeStandardA4mV2
	AzureVMSizeStandardA4V2   AzureVMSize = original.AzureVMSizeStandardA4V2
	AzureVMSizeStandardA5     AzureVMSize = original.AzureVMSizeStandardA5
	AzureVMSizeStandardA6     AzureVMSize = original.AzureVMSizeStandardA6
	AzureVMSizeStandardA7     AzureVMSize = original.AzureVMSizeStandardA7
	AzureVMSizeStandardA8     AzureVMSize = original.AzureVMSizeStandardA8
	AzureVMSizeStandardA8mV2  AzureVMSize = original.AzureVMSizeStandardA8mV2
	AzureVMSizeStandardA8V2   AzureVMSize = original.AzureVMSizeStandardA8V2
	AzureVMSizeStandardA9     AzureVMSize = original.AzureVMSizeStandardA9
	AzureVMSizeStandardD1     AzureVMSize = original.AzureVMSizeStandardD1
	AzureVMSizeStandardD11    AzureVMSize = original.AzureVMSizeStandardD11
	AzureVMSizeStandardD11V2  AzureVMSize = original.AzureVMSizeStandardD11V2
	AzureVMSizeStandardD12    AzureVMSize = original.AzureVMSizeStandardD12
	AzureVMSizeStandardD12V2  AzureVMSize = original.AzureVMSizeStandardD12V2
	AzureVMSizeStandardD13    AzureVMSize = original.AzureVMSizeStandardD13
	AzureVMSizeStandardD13V2  AzureVMSize = original.AzureVMSizeStandardD13V2
	AzureVMSizeStandardD14    AzureVMSize = original.AzureVMSizeStandardD14
	AzureVMSizeStandardD14V2  AzureVMSize = original.AzureVMSizeStandardD14V2
	AzureVMSizeStandardD15V2  AzureVMSize = original.AzureVMSizeStandardD15V2
	AzureVMSizeStandardD1V2   AzureVMSize = original.AzureVMSizeStandardD1V2
	AzureVMSizeStandardD2     AzureVMSize = original.AzureVMSizeStandardD2
	AzureVMSizeStandardD2V2   AzureVMSize = original.AzureVMSizeStandardD2V2
	AzureVMSizeStandardD3     AzureVMSize = original.AzureVMSizeStandardD3
	AzureVMSizeStandardD3V2   AzureVMSize = original.AzureVMSizeStandardD3V2
	AzureVMSizeStandardD4     AzureVMSize = original.AzureVMSizeStandardD4
	AzureVMSizeStandardD4V2   AzureVMSize = original.AzureVMSizeStandardD4V2
	AzureVMSizeStandardD5V2   AzureVMSize = original.AzureVMSizeStandardD5V2
	AzureVMSizeStandardDS1    AzureVMSize = original.AzureVMSizeStandardDS1
	AzureVMSizeStandardDS11   AzureVMSize = original.AzureVMSizeStandardDS11
	AzureVMSizeStandardDS11V2 AzureVMSize = original.AzureVMSizeStandardDS11V2
	AzureVMSizeStandardDS12   AzureVMSize = original.AzureVMSizeStandardDS12
	AzureVMSizeStandardDS12V2 AzureVMSize = original.AzureVMSizeStandardDS12V2
	AzureVMSizeStandardDS13   AzureVMSize = original.AzureVMSizeStandardDS13
	AzureVMSizeStandardDS13V2 AzureVMSize = original.AzureVMSizeStandardDS13V2
	AzureVMSizeStandardDS14   AzureVMSize = original.AzureVMSizeStandardDS14
	AzureVMSizeStandardDS14V2 AzureVMSize = original.AzureVMSizeStandardDS14V2
	AzureVMSizeStandardDS15V2 AzureVMSize = original.AzureVMSizeStandardDS15V2
	AzureVMSizeStandardDS1V2  AzureVMSize = original.AzureVMSizeStandardDS1V2
	AzureVMSizeStandardDS2    AzureVMSize = original.AzureVMSizeStandardDS2
	AzureVMSizeStandardDS2V2  AzureVMSize = original.AzureVMSizeStandardDS2V2
	AzureVMSizeStandardDS3    AzureVMSize = original.AzureVMSizeStandardDS3
	AzureVMSizeStandardDS3V2  AzureVMSize = original.AzureVMSizeStandardDS3V2
	AzureVMSizeStandardDS4    AzureVMSize = original.AzureVMSizeStandardDS4
	AzureVMSizeStandardDS4V2  AzureVMSize = original.AzureVMSizeStandardDS4V2
	AzureVMSizeStandardDS5V2  AzureVMSize = original.AzureVMSizeStandardDS5V2
	AzureVMSizeStandardF1     AzureVMSize = original.AzureVMSizeStandardF1
	AzureVMSizeStandardF16    AzureVMSize = original.AzureVMSizeStandardF16
	AzureVMSizeStandardF16s   AzureVMSize = original.AzureVMSizeStandardF16s
	AzureVMSizeStandardF1s    AzureVMSize = original.AzureVMSizeStandardF1s
	AzureVMSizeStandardF2     AzureVMSize = original.AzureVMSizeStandardF2
	AzureVMSizeStandardF2s    AzureVMSize = original.AzureVMSizeStandardF2s
	AzureVMSizeStandardF4     AzureVMSize = original.AzureVMSizeStandardF4
	AzureVMSizeStandardF4s    AzureVMSize = original.AzureVMSizeStandardF4s
	AzureVMSizeStandardF8     AzureVMSize = original.AzureVMSizeStandardF8
	AzureVMSizeStandardF8s    AzureVMSize = original.AzureVMSizeStandardF8s
	AzureVMSizeStandardG1     AzureVMSize = original.AzureVMSizeStandardG1
	AzureVMSizeStandardG2     AzureVMSize = original.AzureVMSizeStandardG2
	AzureVMSizeStandardG3     AzureVMSize = original.AzureVMSizeStandardG3
	AzureVMSizeStandardG4     AzureVMSize = original.AzureVMSizeStandardG4
	AzureVMSizeStandardG5     AzureVMSize = original.AzureVMSizeStandardG5
	AzureVMSizeStandardGS1    AzureVMSize = original.AzureVMSizeStandardGS1
	AzureVMSizeStandardGS2    AzureVMSize = original.AzureVMSizeStandardGS2
	AzureVMSizeStandardGS3    AzureVMSize = original.AzureVMSizeStandardGS3
	AzureVMSizeStandardGS4    AzureVMSize = original.AzureVMSizeStandardGS4
	AzureVMSizeStandardGS5    AzureVMSize = original.AzureVMSizeStandardGS5
	AzureVMSizeStandardH16    AzureVMSize = original.AzureVMSizeStandardH16
	AzureVMSizeStandardH16m   AzureVMSize = original.AzureVMSizeStandardH16m
	AzureVMSizeStandardH16mr  AzureVMSize = original.AzureVMSizeStandardH16mr
	AzureVMSizeStandardH16r   AzureVMSize = original.AzureVMSizeStandardH16r
	AzureVMSizeStandardH8     AzureVMSize = original.AzureVMSizeStandardH8
	AzureVMSizeStandardH8m    AzureVMSize = original.AzureVMSizeStandardH8m
	AzureVMSizeStandardL16s   AzureVMSize = original.AzureVMSizeStandardL16s
	AzureVMSizeStandardL32s   AzureVMSize = original.AzureVMSizeStandardL32s
	AzureVMSizeStandardL4s    AzureVMSize = original.AzureVMSizeStandardL4s
	AzureVMSizeStandardL8s    AzureVMSize = original.AzureVMSizeStandardL8s
	AzureVMSizeUnknown        AzureVMSize = original.AzureVMSizeUnknown
)

type AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanation

const (
	AzureVMSuitabilityExplanationBootTypeNotSupported                         AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationBootTypeNotSupported
	AzureVMSuitabilityExplanationBootTypeUnknown                              AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationBootTypeUnknown
	AzureVMSuitabilityExplanationCheckCentOsVersion                           AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationCheckCentOsVersion
	AzureVMSuitabilityExplanationCheckCoreOsLinuxVersion                      AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationCheckCoreOsLinuxVersion
	AzureVMSuitabilityExplanationCheckDebianLinuxVersion                      AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationCheckDebianLinuxVersion
	AzureVMSuitabilityExplanationCheckOpenSuseLinuxVersion                    AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationCheckOpenSuseLinuxVersion
	AzureVMSuitabilityExplanationCheckOracleLinuxVersion                      AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationCheckOracleLinuxVersion
	AzureVMSuitabilityExplanationCheckRedHatLinuxVersion                      AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationCheckRedHatLinuxVersion
	AzureVMSuitabilityExplanationCheckSuseLinuxVersion                        AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationCheckSuseLinuxVersion
	AzureVMSuitabilityExplanationCheckUbuntuLinuxVersion                      AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationCheckUbuntuLinuxVersion
	AzureVMSuitabilityExplanationCheckWindowsServer2008R2Version              AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationCheckWindowsServer2008R2Version
	AzureVMSuitabilityExplanationEndorsedWithConditionsLinuxDistributions     AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationEndorsedWithConditionsLinuxDistributions
	AzureVMSuitabilityExplanationGuestOperatingSystemArchitectureNotSupported AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationGuestOperatingSystemArchitectureNotSupported
	AzureVMSuitabilityExplanationGuestOperatingSystemNotSupported             AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationGuestOperatingSystemNotSupported
	AzureVMSuitabilityExplanationGuestOperatingSystemUnknown                  AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationGuestOperatingSystemUnknown
	AzureVMSuitabilityExplanationInternalErrorOccuredDuringComputeEvaluation  AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationInternalErrorOccuredDuringComputeEvaluation
	AzureVMSuitabilityExplanationInternalErrorOccuredDuringNetworkEvaluation  AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationInternalErrorOccuredDuringNetworkEvaluation
	AzureVMSuitabilityExplanationInternalErrorOccuredDuringStorageEvaluation  AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationInternalErrorOccuredDuringStorageEvaluation
	AzureVMSuitabilityExplanationMoreDisksThanSupported                       AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationMoreDisksThanSupported
	AzureVMSuitabilityExplanationNoGuestOperatingSystemConditionallySupported AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationNoGuestOperatingSystemConditionallySupported
	AzureVMSuitabilityExplanationNoSuitableVMSizeFound                        AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationNoSuitableVMSizeFound
	AzureVMSuitabilityExplanationNotApplicable                                AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationNotApplicable
	AzureVMSuitabilityExplanationNoVMSizeForBasicPricingTier                  AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationNoVMSizeForBasicPricingTier
	AzureVMSuitabilityExplanationNoVMSizeForSelectedAzureLocation             AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationNoVMSizeForSelectedAzureLocation
	AzureVMSuitabilityExplanationNoVMSizeForSelectedPricingTier               AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationNoVMSizeForSelectedPricingTier
	AzureVMSuitabilityExplanationNoVMSizeForStandardPricingTier               AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationNoVMSizeForStandardPricingTier
	AzureVMSuitabilityExplanationNoVMSizeSupportsNetworkPerformance           AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationNoVMSizeSupportsNetworkPerformance
	AzureVMSuitabilityExplanationNoVMSizeSupportsStoragePerformance           AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationNoVMSizeSupportsStoragePerformance
	AzureVMSuitabilityExplanationOneOrMoreAdaptersNotSuitable                 AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationOneOrMoreAdaptersNotSuitable
	AzureVMSuitabilityExplanationOneOrMoreDisksNotSuitable                    AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationOneOrMoreDisksNotSuitable
	AzureVMSuitabilityExplanationUnendorsedLinuxDistributions                 AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationUnendorsedLinuxDistributions
	AzureVMSuitabilityExplanationUnknown                                      AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationUnknown
	AzureVMSuitabilityExplanationWindowsClientVersionsConditionallySupported  AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationWindowsClientVersionsConditionallySupported
	AzureVMSuitabilityExplanationWindowsOSNoLongerUnderMSSupport              AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationWindowsOSNoLongerUnderMSSupport
	AzureVMSuitabilityExplanationWindowsServerVersionConditionallySupported   AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationWindowsServerVersionConditionallySupported
	AzureVMSuitabilityExplanationWindowsServerVersionsSupportedWithCaveat     AzureVMSuitabilityExplanation = original.AzureVMSuitabilityExplanationWindowsServerVersionsSupportedWithCaveat
)

type CloudSuitability = original.CloudSuitability

const (
	CloudSuitabilityConditionallySuitable CloudSuitability = original.CloudSuitabilityConditionallySuitable
	CloudSuitabilityNotSuitable           CloudSuitability = original.CloudSuitabilityNotSuitable
	CloudSuitabilityReadinessUnknown      CloudSuitability = original.CloudSuitabilityReadinessUnknown
	CloudSuitabilitySuitable              CloudSuitability = original.CloudSuitabilitySuitable
	CloudSuitabilityUnknown               CloudSuitability = original.CloudSuitabilityUnknown
)

type Currency = original.Currency

const (
	CurrencyARS     Currency = original.CurrencyARS
	CurrencyAUD     Currency = original.CurrencyAUD
	CurrencyBRL     Currency = original.CurrencyBRL
	CurrencyCAD     Currency = original.CurrencyCAD
	CurrencyCHF     Currency = original.CurrencyCHF
	CurrencyCNY     Currency = original.CurrencyCNY
	CurrencyDKK     Currency = original.CurrencyDKK
	CurrencyEUR     Currency = original.CurrencyEUR
	CurrencyGBP     Currency = original.CurrencyGBP
	CurrencyHKD     Currency = original.CurrencyHKD
	CurrencyIDR     Currency = original.CurrencyIDR
	CurrencyINR     Currency = original.CurrencyINR
	CurrencyJPY     Currency = original.CurrencyJPY
	CurrencyKRW     Currency = original.CurrencyKRW
	CurrencyMXN     Currency = original.CurrencyMXN
	CurrencyMYR     Currency = original.CurrencyMYR
	CurrencyNOK     Currency = original.CurrencyNOK
	CurrencyNZD     Currency = original.CurrencyNZD
	CurrencyRUB     Currency = original.CurrencyRUB
	CurrencySAR     Currency = original.CurrencySAR
	CurrencySEK     Currency = original.CurrencySEK
	CurrencyTRY     Currency = original.CurrencyTRY
	CurrencyTWD     Currency = original.CurrencyTWD
	CurrencyUnknown Currency = original.CurrencyUnknown
	CurrencyUSD     Currency = original.CurrencyUSD
	CurrencyZAR     Currency = original.CurrencyZAR
)

type DiscoveryStatus = original.DiscoveryStatus

const (
	DiscoveryStatusCompleted  DiscoveryStatus = original.DiscoveryStatusCompleted
	DiscoveryStatusInProgress DiscoveryStatus = original.DiscoveryStatusInProgress
	DiscoveryStatusNotStarted DiscoveryStatus = original.DiscoveryStatusNotStarted
	DiscoveryStatusUnknown    DiscoveryStatus = original.DiscoveryStatusUnknown
)

type MachineBootType = original.MachineBootType

const (
	MachineBootTypeBIOS    MachineBootType = original.MachineBootTypeBIOS
	MachineBootTypeEFI     MachineBootType = original.MachineBootTypeEFI
	MachineBootTypeUnknown MachineBootType = original.MachineBootTypeUnknown
)

type NameAvailabilityReason = original.NameAvailabilityReason

const (
	NameAvailabilityReasonAlreadyExists NameAvailabilityReason = original.NameAvailabilityReasonAlreadyExists
	NameAvailabilityReasonAvailable     NameAvailabilityReason = original.NameAvailabilityReasonAvailable
	NameAvailabilityReasonInvalid       NameAvailabilityReason = original.NameAvailabilityReasonInvalid
)

type Percentile = original.Percentile

const (
	Percentile50 Percentile = original.Percentile50
	Percentile90 Percentile = original.Percentile90
	Percentile95 Percentile = original.Percentile95
	Percentile99 Percentile = original.Percentile99
)

type ProvisioningState = original.ProvisioningState

const (
	Accepted  ProvisioningState = original.Accepted
	Creating  ProvisioningState = original.Creating
	Deleting  ProvisioningState = original.Deleting
	Failed    ProvisioningState = original.Failed
	Moving    ProvisioningState = original.Moving
	Succeeded ProvisioningState = original.Succeeded
)

type TimeRange = original.TimeRange

const (
	Day   TimeRange = original.Day
	Month TimeRange = original.Month
	Week  TimeRange = original.Week
)

type AssessedDisk = original.AssessedDisk
type AssessedMachine = original.AssessedMachine
type AssessedMachineProperties = original.AssessedMachineProperties
type AssessedMachineResultList = original.AssessedMachineResultList
type AssessedMachinesClient = original.AssessedMachinesClient
type AssessedNetworkAdapter = original.AssessedNetworkAdapter
type Assessment = original.Assessment
type AssessmentOptionsClient = original.AssessmentOptionsClient
type AssessmentOptionsResultList = original.AssessmentOptionsResultList
type AssessmentProperties = original.AssessmentProperties
type AssessmentResultList = original.AssessmentResultList
type AssessmentsClient = original.AssessmentsClient
type BaseClient = original.BaseClient
type CheckNameAvailabilityParameters = original.CheckNameAvailabilityParameters
type CheckNameAvailabilityResult = original.CheckNameAvailabilityResult
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type Disk = original.Disk
type DownloadURL = original.DownloadURL
type Group = original.Group
type GroupProperties = original.GroupProperties
type GroupResultList = original.GroupResultList
type GroupsClient = original.GroupsClient
type LocationClient = original.LocationClient
type Machine = original.Machine
type MachineProperties = original.MachineProperties
type MachineResultList = original.MachineResultList
type MachinesClient = original.MachinesClient
type NetworkAdapter = original.NetworkAdapter
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationResultList = original.OperationResultList
type OperationsClient = original.OperationsClient
type Project = original.Project
type ProjectKey = original.ProjectKey
type ProjectProperties = original.ProjectProperties
type ProjectResultList = original.ProjectResultList
type ProjectsClient = original.ProjectsClient
type VMFamily = original.VMFamily

func New(subscriptionID string, acceptLanguage string) BaseClient {
	return original.New(subscriptionID, acceptLanguage)
}
func NewAssessedMachinesClient(subscriptionID string, acceptLanguage string) AssessedMachinesClient {
	return original.NewAssessedMachinesClient(subscriptionID, acceptLanguage)
}
func NewAssessedMachinesClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) AssessedMachinesClient {
	return original.NewAssessedMachinesClientWithBaseURI(baseURI, subscriptionID, acceptLanguage)
}
func NewAssessmentOptionsClient(subscriptionID string, acceptLanguage string) AssessmentOptionsClient {
	return original.NewAssessmentOptionsClient(subscriptionID, acceptLanguage)
}
func NewAssessmentOptionsClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) AssessmentOptionsClient {
	return original.NewAssessmentOptionsClientWithBaseURI(baseURI, subscriptionID, acceptLanguage)
}
func NewAssessmentsClient(subscriptionID string, acceptLanguage string) AssessmentsClient {
	return original.NewAssessmentsClient(subscriptionID, acceptLanguage)
}
func NewAssessmentsClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) AssessmentsClient {
	return original.NewAssessmentsClientWithBaseURI(baseURI, subscriptionID, acceptLanguage)
}
func NewGroupsClient(subscriptionID string, acceptLanguage string) GroupsClient {
	return original.NewGroupsClient(subscriptionID, acceptLanguage)
}
func NewGroupsClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) GroupsClient {
	return original.NewGroupsClientWithBaseURI(baseURI, subscriptionID, acceptLanguage)
}
func NewLocationClient(subscriptionID string, acceptLanguage string) LocationClient {
	return original.NewLocationClient(subscriptionID, acceptLanguage)
}
func NewLocationClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) LocationClient {
	return original.NewLocationClientWithBaseURI(baseURI, subscriptionID, acceptLanguage)
}
func NewMachinesClient(subscriptionID string, acceptLanguage string) MachinesClient {
	return original.NewMachinesClient(subscriptionID, acceptLanguage)
}
func NewMachinesClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) MachinesClient {
	return original.NewMachinesClientWithBaseURI(baseURI, subscriptionID, acceptLanguage)
}
func NewOperationsClient(subscriptionID string, acceptLanguage string) OperationsClient {
	return original.NewOperationsClient(subscriptionID, acceptLanguage)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID, acceptLanguage)
}
func NewProjectsClient(subscriptionID string, acceptLanguage string) ProjectsClient {
	return original.NewProjectsClient(subscriptionID, acceptLanguage)
}
func NewProjectsClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) ProjectsClient {
	return original.NewProjectsClientWithBaseURI(baseURI, subscriptionID, acceptLanguage)
}
func NewWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID, acceptLanguage)
}
func PossibleAssessmentSizingCriterionValues() []AssessmentSizingCriterion {
	return original.PossibleAssessmentSizingCriterionValues()
}
func PossibleAssessmentStageValues() []AssessmentStage {
	return original.PossibleAssessmentStageValues()
}
func PossibleAssessmentStatusValues() []AssessmentStatus {
	return original.PossibleAssessmentStatusValues()
}
func PossibleAzureDiskSizeValues() []AzureDiskSize {
	return original.PossibleAzureDiskSizeValues()
}
func PossibleAzureDiskSuitabilityExplanationValues() []AzureDiskSuitabilityExplanation {
	return original.PossibleAzureDiskSuitabilityExplanationValues()
}
func PossibleAzureDiskTypeValues() []AzureDiskType {
	return original.PossibleAzureDiskTypeValues()
}
func PossibleAzureHybridUseBenefitValues() []AzureHybridUseBenefit {
	return original.PossibleAzureHybridUseBenefitValues()
}
func PossibleAzureLocationValues() []AzureLocation {
	return original.PossibleAzureLocationValues()
}
func PossibleAzureNetworkAdapterSuitabilityExplanationValues() []AzureNetworkAdapterSuitabilityExplanation {
	return original.PossibleAzureNetworkAdapterSuitabilityExplanationValues()
}
func PossibleAzureOfferCodeValues() []AzureOfferCode {
	return original.PossibleAzureOfferCodeValues()
}
func PossibleAzurePricingTierValues() []AzurePricingTier {
	return original.PossibleAzurePricingTierValues()
}
func PossibleAzureStorageRedundancyValues() []AzureStorageRedundancy {
	return original.PossibleAzureStorageRedundancyValues()
}
func PossibleAzureVMSizeValues() []AzureVMSize {
	return original.PossibleAzureVMSizeValues()
}
func PossibleAzureVMSuitabilityExplanationValues() []AzureVMSuitabilityExplanation {
	return original.PossibleAzureVMSuitabilityExplanationValues()
}
func PossibleCloudSuitabilityValues() []CloudSuitability {
	return original.PossibleCloudSuitabilityValues()
}
func PossibleCurrencyValues() []Currency {
	return original.PossibleCurrencyValues()
}
func PossibleDiscoveryStatusValues() []DiscoveryStatus {
	return original.PossibleDiscoveryStatusValues()
}
func PossibleMachineBootTypeValues() []MachineBootType {
	return original.PossibleMachineBootTypeValues()
}
func PossibleNameAvailabilityReasonValues() []NameAvailabilityReason {
	return original.PossibleNameAvailabilityReasonValues()
}
func PossiblePercentileValues() []Percentile {
	return original.PossiblePercentileValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleTimeRangeValues() []TimeRange {
	return original.PossibleTimeRangeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
