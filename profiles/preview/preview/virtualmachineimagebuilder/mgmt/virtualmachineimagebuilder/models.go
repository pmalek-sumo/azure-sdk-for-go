// +build go1.9

// Copyright 2019 Microsoft Corporation
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

package virtualmachineimagebuilder

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/virtualmachineimagebuilder/mgmt/2018-02-01-preview/virtualmachineimagebuilder"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ProvisioningErrorCode = original.ProvisioningErrorCode

const (
	BadCustomizerType       ProvisioningErrorCode = original.BadCustomizerType
	BadISOSource            ProvisioningErrorCode = original.BadISOSource
	BadPIRSource            ProvisioningErrorCode = original.BadPIRSource
	BadSourceType           ProvisioningErrorCode = original.BadSourceType
	NoCustomizerShellScript ProvisioningErrorCode = original.NoCustomizerShellScript
	Other                   ProvisioningErrorCode = original.Other
	ServerError             ProvisioningErrorCode = original.ServerError
)

type ProvisioningState = original.ProvisioningState

const (
	Creating  ProvisioningState = original.Creating
	Deleting  ProvisioningState = original.Deleting
	Failed    ProvisioningState = original.Failed
	Succeeded ProvisioningState = original.Succeeded
)

type ProvisioningState1 = original.ProvisioningState1

const (
	ProvisioningState1Creating  ProvisioningState1 = original.ProvisioningState1Creating
	ProvisioningState1Deleting  ProvisioningState1 = original.ProvisioningState1Deleting
	ProvisioningState1Failed    ProvisioningState1 = original.ProvisioningState1Failed
	ProvisioningState1Succeeded ProvisioningState1 = original.ProvisioningState1Succeeded
)

type RunState = original.RunState

const (
	RunStateFailed             RunState = original.RunStateFailed
	RunStatePartiallySucceeded RunState = original.RunStatePartiallySucceeded
	RunStateReady              RunState = original.RunStateReady
	RunStateRunning            RunState = original.RunStateRunning
	RunStateSucceeded          RunState = original.RunStateSucceeded
)

type RunSubState = original.RunSubState

const (
	Building     RunSubState = original.Building
	Customizing  RunSubState = original.Customizing
	Distributing RunSubState = original.Distributing
	Queued       RunSubState = original.Queued
)

type Type = original.Type

const (
	TypeImageTemplateSource Type = original.TypeImageTemplateSource
	TypeISO                 Type = original.TypeISO
	TypePlatformImage       Type = original.TypePlatformImage
)

type TypeBasicImageTemplateCustomizer = original.TypeBasicImageTemplateCustomizer

const (
	TypeImageTemplateCustomizer TypeBasicImageTemplateCustomizer = original.TypeImageTemplateCustomizer
	TypeShell                   TypeBasicImageTemplateCustomizer = original.TypeShell
)

type TypeBasicImageTemplateDistributor = original.TypeBasicImageTemplateDistributor

const (
	TypeImageTemplateDistributor TypeBasicImageTemplateDistributor = original.TypeImageTemplateDistributor
	TypeManagedImage             TypeBasicImageTemplateDistributor = original.TypeManagedImage
	TypeSharedImage              TypeBasicImageTemplateDistributor = original.TypeSharedImage
)

type APIError = original.APIError
type APIErrorBase = original.APIErrorBase
type BaseClient = original.BaseClient
type BasicImageTemplateCustomizer = original.BasicImageTemplateCustomizer
type BasicImageTemplateDistributor = original.BasicImageTemplateDistributor
type BasicImageTemplateSource = original.BasicImageTemplateSource
type ImageTemplate = original.ImageTemplate
type ImageTemplateCustomizer = original.ImageTemplateCustomizer
type ImageTemplateDistributor = original.ImageTemplateDistributor
type ImageTemplateIsoSource = original.ImageTemplateIsoSource
type ImageTemplateLastRunStatus = original.ImageTemplateLastRunStatus
type ImageTemplateListResult = original.ImageTemplateListResult
type ImageTemplateListResultIterator = original.ImageTemplateListResultIterator
type ImageTemplateListResultPage = original.ImageTemplateListResultPage
type ImageTemplateManagedImageDistributor = original.ImageTemplateManagedImageDistributor
type ImageTemplatePlatformImageSource = original.ImageTemplatePlatformImageSource
type ImageTemplateProperties = original.ImageTemplateProperties
type ImageTemplateSharedImageDistributor = original.ImageTemplateSharedImageDistributor
type ImageTemplateShellCustomizer = original.ImageTemplateShellCustomizer
type ImageTemplateSource = original.ImageTemplateSource
type ImageTemplateUpdateParameters = original.ImageTemplateUpdateParameters
type InnerError = original.InnerError
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type ProvisioningError = original.ProvisioningError
type Resource = original.Resource
type RunOutput = original.RunOutput
type RunOutputCollection = original.RunOutputCollection
type RunOutputCollectionIterator = original.RunOutputCollectionIterator
type RunOutputCollectionPage = original.RunOutputCollectionPage
type RunOutputProperties = original.RunOutputProperties
type SubResource = original.SubResource
type VirtualMachineImageTemplateClient = original.VirtualMachineImageTemplateClient
type VirtualMachineImageTemplateCreateOrUpdateFuture = original.VirtualMachineImageTemplateCreateOrUpdateFuture
type VirtualMachineImageTemplateDeleteFuture = original.VirtualMachineImageTemplateDeleteFuture
type VirtualMachineImageTemplateRunFuture = original.VirtualMachineImageTemplateRunFuture

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewImageTemplateListResultIterator(page ImageTemplateListResultPage) ImageTemplateListResultIterator {
	return original.NewImageTemplateListResultIterator(page)
}
func NewImageTemplateListResultPage(getNextPage func(context.Context, ImageTemplateListResult) (ImageTemplateListResult, error)) ImageTemplateListResultPage {
	return original.NewImageTemplateListResultPage(getNextPage)
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
func NewRunOutputCollectionIterator(page RunOutputCollectionPage) RunOutputCollectionIterator {
	return original.NewRunOutputCollectionIterator(page)
}
func NewRunOutputCollectionPage(getNextPage func(context.Context, RunOutputCollection) (RunOutputCollection, error)) RunOutputCollectionPage {
	return original.NewRunOutputCollectionPage(getNextPage)
}
func NewVirtualMachineImageTemplateClient(subscriptionID string) VirtualMachineImageTemplateClient {
	return original.NewVirtualMachineImageTemplateClient(subscriptionID)
}
func NewVirtualMachineImageTemplateClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineImageTemplateClient {
	return original.NewVirtualMachineImageTemplateClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleProvisioningErrorCodeValues() []ProvisioningErrorCode {
	return original.PossibleProvisioningErrorCodeValues()
}
func PossibleProvisioningState1Values() []ProvisioningState1 {
	return original.PossibleProvisioningState1Values()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleRunStateValues() []RunState {
	return original.PossibleRunStateValues()
}
func PossibleRunSubStateValues() []RunSubState {
	return original.PossibleRunSubStateValues()
}
func PossibleTypeBasicImageTemplateCustomizerValues() []TypeBasicImageTemplateCustomizer {
	return original.PossibleTypeBasicImageTemplateCustomizerValues()
}
func PossibleTypeBasicImageTemplateDistributorValues() []TypeBasicImageTemplateDistributor {
	return original.PossibleTypeBasicImageTemplateDistributorValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
