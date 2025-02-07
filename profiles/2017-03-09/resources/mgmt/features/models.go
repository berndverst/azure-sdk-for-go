//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package features

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2015-12-01/features"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type Client = original.Client
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsListResult = original.OperationsListResult
type OperationsListResultIterator = original.OperationsListResultIterator
type OperationsListResultPage = original.OperationsListResultPage
type Properties = original.Properties
type Result = original.Result

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewClient(subscriptionID string) Client {
	return original.NewClient(subscriptionID)
}
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return original.NewClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsListResultIterator(page OperationsListResultPage) OperationsListResultIterator {
	return original.NewOperationsListResultIterator(page)
}
func NewOperationsListResultPage(cur OperationsListResult, getNextPage func(context.Context, OperationsListResult) (OperationsListResult, error)) OperationsListResultPage {
	return original.NewOperationsListResultPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/2017-03-09"
}
func Version() string {
	return original.Version()
}
