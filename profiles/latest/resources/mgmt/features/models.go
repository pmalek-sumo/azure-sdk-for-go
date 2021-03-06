// +build go1.9

// Copyright 2021 Microsoft Corporation
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

package features

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-07-01/features"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ChangeType = original.ChangeType

const (
	Create   ChangeType = original.Create
	Delete   ChangeType = original.Delete
	Deploy   ChangeType = original.Deploy
	Ignore   ChangeType = original.Ignore
	Modify   ChangeType = original.Modify
	NoChange ChangeType = original.NoChange
)

type DeploymentMode = original.DeploymentMode

const (
	Complete    DeploymentMode = original.Complete
	Incremental DeploymentMode = original.Incremental
)

type OnErrorDeploymentType = original.OnErrorDeploymentType

const (
	LastSuccessful     OnErrorDeploymentType = original.LastSuccessful
	SpecificDeployment OnErrorDeploymentType = original.SpecificDeployment
)

type PropertyChangeType = original.PropertyChangeType

const (
	PropertyChangeTypeArray  PropertyChangeType = original.PropertyChangeTypeArray
	PropertyChangeTypeCreate PropertyChangeType = original.PropertyChangeTypeCreate
	PropertyChangeTypeDelete PropertyChangeType = original.PropertyChangeTypeDelete
	PropertyChangeTypeModify PropertyChangeType = original.PropertyChangeTypeModify
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	None                       ResourceIdentityType = original.None
	SystemAssigned             ResourceIdentityType = original.SystemAssigned
	SystemAssignedUserAssigned ResourceIdentityType = original.SystemAssignedUserAssigned
	UserAssigned               ResourceIdentityType = original.UserAssigned
)

type WhatIfResultFormat = original.WhatIfResultFormat

const (
	FullResourcePayloads WhatIfResultFormat = original.FullResourcePayloads
	ResourceIDOnly       WhatIfResultFormat = original.ResourceIDOnly
)

type AliasPathType = original.AliasPathType
type AliasType = original.AliasType
type BaseClient = original.BaseClient
type BasicDependency = original.BasicDependency
type CloudError = original.CloudError
type DebugSetting = original.DebugSetting
type Dependency = original.Dependency
type Deployment = original.Deployment
type DeploymentExportResult = original.DeploymentExportResult
type DeploymentExtended = original.DeploymentExtended
type DeploymentExtendedFilter = original.DeploymentExtendedFilter
type DeploymentListResult = original.DeploymentListResult
type DeploymentListResultIterator = original.DeploymentListResultIterator
type DeploymentListResultPage = original.DeploymentListResultPage
type DeploymentOperation = original.DeploymentOperation
type DeploymentOperationProperties = original.DeploymentOperationProperties
type DeploymentOperationsClient = original.DeploymentOperationsClient
type DeploymentOperationsListResult = original.DeploymentOperationsListResult
type DeploymentOperationsListResultIterator = original.DeploymentOperationsListResultIterator
type DeploymentOperationsListResultPage = original.DeploymentOperationsListResultPage
type DeploymentProperties = original.DeploymentProperties
type DeploymentPropertiesExtended = original.DeploymentPropertiesExtended
type DeploymentValidateResult = original.DeploymentValidateResult
type DeploymentWhatIf = original.DeploymentWhatIf
type DeploymentWhatIfProperties = original.DeploymentWhatIfProperties
type DeploymentWhatIfSettings = original.DeploymentWhatIfSettings
type DeploymentsClient = original.DeploymentsClient
type DeploymentsCreateOrUpdateAtManagementGroupScopeFuture = original.DeploymentsCreateOrUpdateAtManagementGroupScopeFuture
type DeploymentsCreateOrUpdateAtScopeFuture = original.DeploymentsCreateOrUpdateAtScopeFuture
type DeploymentsCreateOrUpdateAtSubscriptionScopeFuture = original.DeploymentsCreateOrUpdateAtSubscriptionScopeFuture
type DeploymentsCreateOrUpdateAtTenantScopeFuture = original.DeploymentsCreateOrUpdateAtTenantScopeFuture
type DeploymentsCreateOrUpdateFuture = original.DeploymentsCreateOrUpdateFuture
type DeploymentsDeleteAtManagementGroupScopeFuture = original.DeploymentsDeleteAtManagementGroupScopeFuture
type DeploymentsDeleteAtScopeFuture = original.DeploymentsDeleteAtScopeFuture
type DeploymentsDeleteAtSubscriptionScopeFuture = original.DeploymentsDeleteAtSubscriptionScopeFuture
type DeploymentsDeleteAtTenantScopeFuture = original.DeploymentsDeleteAtTenantScopeFuture
type DeploymentsDeleteFuture = original.DeploymentsDeleteFuture
type DeploymentsWhatIfAtSubscriptionScopeFuture = original.DeploymentsWhatIfAtSubscriptionScopeFuture
type DeploymentsWhatIfFuture = original.DeploymentsWhatIfFuture
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorResponse = original.ErrorResponse
type ExportTemplateRequest = original.ExportTemplateRequest
type GenericResource = original.GenericResource
type GenericResourceExpanded = original.GenericResourceExpanded
type GenericResourceFilter = original.GenericResourceFilter
type HTTPMessage = original.HTTPMessage
type Identity = original.Identity
type IdentityUserAssignedIdentitiesValue = original.IdentityUserAssignedIdentitiesValue
type OnErrorDeployment = original.OnErrorDeployment
type OnErrorDeploymentExtended = original.OnErrorDeploymentExtended
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type ParametersLink = original.ParametersLink
type Plan = original.Plan
type Provider = original.Provider
type ProviderListResult = original.ProviderListResult
type ProviderListResultIterator = original.ProviderListResultIterator
type ProviderListResultPage = original.ProviderListResultPage
type ProviderResourceType = original.ProviderResourceType
type ProvidersClient = original.ProvidersClient
type Resource = original.Resource
type ResourceGroup = original.ResourceGroup
type ResourceGroupExportResult = original.ResourceGroupExportResult
type ResourceGroupFilter = original.ResourceGroupFilter
type ResourceGroupListResult = original.ResourceGroupListResult
type ResourceGroupListResultIterator = original.ResourceGroupListResultIterator
type ResourceGroupListResultPage = original.ResourceGroupListResultPage
type ResourceGroupPatchable = original.ResourceGroupPatchable
type ResourceGroupProperties = original.ResourceGroupProperties
type ResourceGroupsClient = original.ResourceGroupsClient
type ResourceGroupsDeleteFuture = original.ResourceGroupsDeleteFuture
type ResourceListResult = original.ResourceListResult
type ResourceListResultIterator = original.ResourceListResultIterator
type ResourceListResultPage = original.ResourceListResultPage
type ResourceProviderOperationDisplayProperties = original.ResourceProviderOperationDisplayProperties
type ResourcesClient = original.ResourcesClient
type ResourcesCreateOrUpdateByIDFuture = original.ResourcesCreateOrUpdateByIDFuture
type ResourcesCreateOrUpdateFuture = original.ResourcesCreateOrUpdateFuture
type ResourcesDeleteByIDFuture = original.ResourcesDeleteByIDFuture
type ResourcesDeleteFuture = original.ResourcesDeleteFuture
type ResourcesMoveInfo = original.ResourcesMoveInfo
type ResourcesMoveResourcesFuture = original.ResourcesMoveResourcesFuture
type ResourcesUpdateByIDFuture = original.ResourcesUpdateByIDFuture
type ResourcesUpdateFuture = original.ResourcesUpdateFuture
type ResourcesValidateMoveResourcesFuture = original.ResourcesValidateMoveResourcesFuture
type Sku = original.Sku
type SubResource = original.SubResource
type TagCount = original.TagCount
type TagDetails = original.TagDetails
type TagValue = original.TagValue
type TagsClient = original.TagsClient
type TagsListResult = original.TagsListResult
type TagsListResultIterator = original.TagsListResultIterator
type TagsListResultPage = original.TagsListResultPage
type TargetResource = original.TargetResource
type TemplateHashResult = original.TemplateHashResult
type TemplateLink = original.TemplateLink
type WhatIfChange = original.WhatIfChange
type WhatIfOperationProperties = original.WhatIfOperationProperties
type WhatIfOperationResult = original.WhatIfOperationResult
type WhatIfPropertyChange = original.WhatIfPropertyChange

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewDeploymentListResultIterator(page DeploymentListResultPage) DeploymentListResultIterator {
	return original.NewDeploymentListResultIterator(page)
}
func NewDeploymentListResultPage(cur DeploymentListResult, getNextPage func(context.Context, DeploymentListResult) (DeploymentListResult, error)) DeploymentListResultPage {
	return original.NewDeploymentListResultPage(cur, getNextPage)
}
func NewDeploymentOperationsClient(subscriptionID string) DeploymentOperationsClient {
	return original.NewDeploymentOperationsClient(subscriptionID)
}
func NewDeploymentOperationsClientWithBaseURI(baseURI string, subscriptionID string) DeploymentOperationsClient {
	return original.NewDeploymentOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDeploymentOperationsListResultIterator(page DeploymentOperationsListResultPage) DeploymentOperationsListResultIterator {
	return original.NewDeploymentOperationsListResultIterator(page)
}
func NewDeploymentOperationsListResultPage(cur DeploymentOperationsListResult, getNextPage func(context.Context, DeploymentOperationsListResult) (DeploymentOperationsListResult, error)) DeploymentOperationsListResultPage {
	return original.NewDeploymentOperationsListResultPage(cur, getNextPage)
}
func NewDeploymentsClient(subscriptionID string) DeploymentsClient {
	return original.NewDeploymentsClient(subscriptionID)
}
func NewDeploymentsClientWithBaseURI(baseURI string, subscriptionID string) DeploymentsClient {
	return original.NewDeploymentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewProviderListResultIterator(page ProviderListResultPage) ProviderListResultIterator {
	return original.NewProviderListResultIterator(page)
}
func NewProviderListResultPage(cur ProviderListResult, getNextPage func(context.Context, ProviderListResult) (ProviderListResult, error)) ProviderListResultPage {
	return original.NewProviderListResultPage(cur, getNextPage)
}
func NewProvidersClient(subscriptionID string) ProvidersClient {
	return original.NewProvidersClient(subscriptionID)
}
func NewProvidersClientWithBaseURI(baseURI string, subscriptionID string) ProvidersClient {
	return original.NewProvidersClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceGroupListResultIterator(page ResourceGroupListResultPage) ResourceGroupListResultIterator {
	return original.NewResourceGroupListResultIterator(page)
}
func NewResourceGroupListResultPage(cur ResourceGroupListResult, getNextPage func(context.Context, ResourceGroupListResult) (ResourceGroupListResult, error)) ResourceGroupListResultPage {
	return original.NewResourceGroupListResultPage(cur, getNextPage)
}
func NewResourceGroupsClient(subscriptionID string) ResourceGroupsClient {
	return original.NewResourceGroupsClient(subscriptionID)
}
func NewResourceGroupsClientWithBaseURI(baseURI string, subscriptionID string) ResourceGroupsClient {
	return original.NewResourceGroupsClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceListResultIterator(page ResourceListResultPage) ResourceListResultIterator {
	return original.NewResourceListResultIterator(page)
}
func NewResourceListResultPage(cur ResourceListResult, getNextPage func(context.Context, ResourceListResult) (ResourceListResult, error)) ResourceListResultPage {
	return original.NewResourceListResultPage(cur, getNextPage)
}
func NewResourcesClient(subscriptionID string) ResourcesClient {
	return original.NewResourcesClient(subscriptionID)
}
func NewResourcesClientWithBaseURI(baseURI string, subscriptionID string) ResourcesClient {
	return original.NewResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewTagsClient(subscriptionID string) TagsClient {
	return original.NewTagsClient(subscriptionID)
}
func NewTagsClientWithBaseURI(baseURI string, subscriptionID string) TagsClient {
	return original.NewTagsClientWithBaseURI(baseURI, subscriptionID)
}
func NewTagsListResultIterator(page TagsListResultPage) TagsListResultIterator {
	return original.NewTagsListResultIterator(page)
}
func NewTagsListResultPage(cur TagsListResult, getNextPage func(context.Context, TagsListResult) (TagsListResult, error)) TagsListResultPage {
	return original.NewTagsListResultPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleChangeTypeValues() []ChangeType {
	return original.PossibleChangeTypeValues()
}
func PossibleDeploymentModeValues() []DeploymentMode {
	return original.PossibleDeploymentModeValues()
}
func PossibleOnErrorDeploymentTypeValues() []OnErrorDeploymentType {
	return original.PossibleOnErrorDeploymentTypeValues()
}
func PossiblePropertyChangeTypeValues() []PropertyChangeType {
	return original.PossiblePropertyChangeTypeValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleWhatIfResultFormatValues() []WhatIfResultFormat {
	return original.PossibleWhatIfResultFormatValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
