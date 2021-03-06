package reservations

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// AutoQuotaIncreaseClient is the client for the AutoQuotaIncrease methods of the Reservations service.
type AutoQuotaIncreaseClient struct {
	BaseClient
}

// NewAutoQuotaIncreaseClient creates an instance of the AutoQuotaIncreaseClient client.
func NewAutoQuotaIncreaseClient() AutoQuotaIncreaseClient {
	return NewAutoQuotaIncreaseClientWithBaseURI(DefaultBaseURI)
}

// NewAutoQuotaIncreaseClientWithBaseURI creates an instance of the AutoQuotaIncreaseClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewAutoQuotaIncreaseClientWithBaseURI(baseURI string) AutoQuotaIncreaseClient {
	return AutoQuotaIncreaseClient{NewWithBaseURI(baseURI)}
}

// Create sets the Auto Quota Increase enrollment properties for the specified subscription.
// Parameters:
// subscriptionID - azure subscription id.
// autoQuotaIncreaseRequest - auto Quota increase request payload.
func (client AutoQuotaIncreaseClient) Create(ctx context.Context, subscriptionID string, autoQuotaIncreaseRequest AutoQuotaIncreaseDetail) (result AutoQuotaIncreaseDetail, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AutoQuotaIncreaseClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, subscriptionID, autoQuotaIncreaseRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.AutoQuotaIncreaseClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "reservations.AutoQuotaIncreaseClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.AutoQuotaIncreaseClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client AutoQuotaIncreaseClient) CreatePreparer(ctx context.Context, subscriptionID string, autoQuotaIncreaseRequest AutoQuotaIncreaseDetail) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2019-07-19-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	autoQuotaIncreaseRequest.ID = nil
	autoQuotaIncreaseRequest.Name = nil
	autoQuotaIncreaseRequest.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/autoQuotaIncrease", pathParameters),
		autorest.WithJSON(autoQuotaIncreaseRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client AutoQuotaIncreaseClient) CreateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client AutoQuotaIncreaseClient) CreateResponder(resp *http.Response) (result AutoQuotaIncreaseDetail, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetProperties gets the Auto Quota Increase enrollment details for the specified subscription.
// Parameters:
// subscriptionID - azure subscription id.
func (client AutoQuotaIncreaseClient) GetProperties(ctx context.Context, subscriptionID string) (result AutoQuotaIncreaseDetail, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AutoQuotaIncreaseClient.GetProperties")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPropertiesPreparer(ctx, subscriptionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.AutoQuotaIncreaseClient", "GetProperties", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetPropertiesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "reservations.AutoQuotaIncreaseClient", "GetProperties", resp, "Failure sending request")
		return
	}

	result, err = client.GetPropertiesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.AutoQuotaIncreaseClient", "GetProperties", resp, "Failure responding to request")
	}

	return
}

// GetPropertiesPreparer prepares the GetProperties request.
func (client AutoQuotaIncreaseClient) GetPropertiesPreparer(ctx context.Context, subscriptionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2019-07-19-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/autoQuotaIncrease", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetPropertiesSender sends the GetProperties request. The method will close the
// http.Response Body if it receives an error.
func (client AutoQuotaIncreaseClient) GetPropertiesSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetPropertiesResponder handles the response to the GetProperties request. The method always
// closes the http.Response Body.
func (client AutoQuotaIncreaseClient) GetPropertiesResponder(resp *http.Response) (result AutoQuotaIncreaseDetail, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
