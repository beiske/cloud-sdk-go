// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by go-swagger; DO NOT EDIT.

package clusters_enterprise_search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetEnterpriseSearchProxyRequestsParams creates a new GetEnterpriseSearchProxyRequestsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEnterpriseSearchProxyRequestsParams() *GetEnterpriseSearchProxyRequestsParams {
	return &GetEnterpriseSearchProxyRequestsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEnterpriseSearchProxyRequestsParamsWithTimeout creates a new GetEnterpriseSearchProxyRequestsParams object
// with the ability to set a timeout on a request.
func NewGetEnterpriseSearchProxyRequestsParamsWithTimeout(timeout time.Duration) *GetEnterpriseSearchProxyRequestsParams {
	return &GetEnterpriseSearchProxyRequestsParams{
		timeout: timeout,
	}
}

// NewGetEnterpriseSearchProxyRequestsParamsWithContext creates a new GetEnterpriseSearchProxyRequestsParams object
// with the ability to set a context for a request.
func NewGetEnterpriseSearchProxyRequestsParamsWithContext(ctx context.Context) *GetEnterpriseSearchProxyRequestsParams {
	return &GetEnterpriseSearchProxyRequestsParams{
		Context: ctx,
	}
}

// NewGetEnterpriseSearchProxyRequestsParamsWithHTTPClient creates a new GetEnterpriseSearchProxyRequestsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEnterpriseSearchProxyRequestsParamsWithHTTPClient(client *http.Client) *GetEnterpriseSearchProxyRequestsParams {
	return &GetEnterpriseSearchProxyRequestsParams{
		HTTPClient: client,
	}
}

/* GetEnterpriseSearchProxyRequestsParams contains all the parameters to send to the API endpoint
   for the get enterprise search proxy requests operation.

   Typically these are written to a http.Request.
*/
type GetEnterpriseSearchProxyRequestsParams struct {

	/* XManagementRequest.

	   When set to `true`, includes the X-Management-Request header value.
	*/
	XManagementRequest string

	/* ClusterID.

	   The Enterprise Search deployment identifier
	*/
	ClusterID string

	/* EnterpriseSearchPath.

	   The URL part to proxy to the Enterprise Search cluster. Example: /api/ent/v1/internal/read_only_mode or /api/ent/v1/internal/health
	*/
	EnterpriseSearchPath string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get enterprise search proxy requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEnterpriseSearchProxyRequestsParams) WithDefaults() *GetEnterpriseSearchProxyRequestsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get enterprise search proxy requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEnterpriseSearchProxyRequestsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get enterprise search proxy requests params
func (o *GetEnterpriseSearchProxyRequestsParams) WithTimeout(timeout time.Duration) *GetEnterpriseSearchProxyRequestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get enterprise search proxy requests params
func (o *GetEnterpriseSearchProxyRequestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get enterprise search proxy requests params
func (o *GetEnterpriseSearchProxyRequestsParams) WithContext(ctx context.Context) *GetEnterpriseSearchProxyRequestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get enterprise search proxy requests params
func (o *GetEnterpriseSearchProxyRequestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get enterprise search proxy requests params
func (o *GetEnterpriseSearchProxyRequestsParams) WithHTTPClient(client *http.Client) *GetEnterpriseSearchProxyRequestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get enterprise search proxy requests params
func (o *GetEnterpriseSearchProxyRequestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXManagementRequest adds the xManagementRequest to the get enterprise search proxy requests params
func (o *GetEnterpriseSearchProxyRequestsParams) WithXManagementRequest(xManagementRequest string) *GetEnterpriseSearchProxyRequestsParams {
	o.SetXManagementRequest(xManagementRequest)
	return o
}

// SetXManagementRequest adds the xManagementRequest to the get enterprise search proxy requests params
func (o *GetEnterpriseSearchProxyRequestsParams) SetXManagementRequest(xManagementRequest string) {
	o.XManagementRequest = xManagementRequest
}

// WithClusterID adds the clusterID to the get enterprise search proxy requests params
func (o *GetEnterpriseSearchProxyRequestsParams) WithClusterID(clusterID string) *GetEnterpriseSearchProxyRequestsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get enterprise search proxy requests params
func (o *GetEnterpriseSearchProxyRequestsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithEnterpriseSearchPath adds the enterpriseSearchPath to the get enterprise search proxy requests params
func (o *GetEnterpriseSearchProxyRequestsParams) WithEnterpriseSearchPath(enterpriseSearchPath string) *GetEnterpriseSearchProxyRequestsParams {
	o.SetEnterpriseSearchPath(enterpriseSearchPath)
	return o
}

// SetEnterpriseSearchPath adds the enterpriseSearchPath to the get enterprise search proxy requests params
func (o *GetEnterpriseSearchProxyRequestsParams) SetEnterpriseSearchPath(enterpriseSearchPath string) {
	o.EnterpriseSearchPath = enterpriseSearchPath
}

// WriteToRequest writes these params to a swagger request
func (o *GetEnterpriseSearchProxyRequestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Management-Request
	if err := r.SetHeaderParam("X-Management-Request", o.XManagementRequest); err != nil {
		return err
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param enterprise_search_path
	if err := r.SetPathParam("enterprise_search_path", o.EnterpriseSearchPath); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
