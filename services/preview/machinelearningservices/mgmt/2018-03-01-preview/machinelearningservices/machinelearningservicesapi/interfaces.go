package machinelearningservicesapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/machinelearningservices/mgmt/2018-03-01-preview/machinelearningservices"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result machinelearningservices.OperationListResult, err error)
}

var _ OperationsClientAPI = (*machinelearningservices.OperationsClient)(nil)

// WorkspacesClientAPI contains the set of methods on the WorkspacesClient type.
type WorkspacesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, parameters machinelearningservices.Workspace) (result machinelearningservices.Workspace, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string) (result machinelearningservices.WorkspacesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string) (result machinelearningservices.Workspace, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, skiptoken string) (result machinelearningservices.WorkspaceListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, skiptoken string) (result machinelearningservices.WorkspaceListResultIterator, err error)
	ListBySubscription(ctx context.Context, skiptoken string) (result machinelearningservices.WorkspaceListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context, skiptoken string) (result machinelearningservices.WorkspaceListResultIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, workspaceName string) (result machinelearningservices.ListWorkspaceKeysResult, err error)
	ResyncKeys(ctx context.Context, resourceGroupName string, workspaceName string) (result autorest.Response, err error)
	Update(ctx context.Context, resourceGroupName string, workspaceName string, parameters machinelearningservices.WorkspaceUpdateParameters) (result machinelearningservices.Workspace, err error)
}

var _ WorkspacesClientAPI = (*machinelearningservices.WorkspacesClient)(nil)

// MachineLearningComputeClientAPI contains the set of methods on the MachineLearningComputeClient type.
type MachineLearningComputeClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters machinelearningservices.ComputeResource) (result machinelearningservices.MachineLearningComputeCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (result machinelearningservices.MachineLearningComputeDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (result machinelearningservices.ComputeResource, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, skiptoken string) (result machinelearningservices.PaginatedComputeResourcesListPage, err error)
	ListByWorkspaceComplete(ctx context.Context, resourceGroupName string, workspaceName string, skiptoken string) (result machinelearningservices.PaginatedComputeResourcesListIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (result machinelearningservices.ComputeSecretsModel, err error)
	SystemUpdate(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (result machinelearningservices.MachineLearningComputeSystemUpdateFuture, err error)
}

var _ MachineLearningComputeClientAPI = (*machinelearningservices.MachineLearningComputeClient)(nil)
