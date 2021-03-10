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

package clusters_kibana

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
	"github.com/go-openapi/swag"
)

// NewSetKibanaClusterMetadataRawParams creates a new SetKibanaClusterMetadataRawParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetKibanaClusterMetadataRawParams() *SetKibanaClusterMetadataRawParams {
	return &SetKibanaClusterMetadataRawParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetKibanaClusterMetadataRawParamsWithTimeout creates a new SetKibanaClusterMetadataRawParams object
// with the ability to set a timeout on a request.
func NewSetKibanaClusterMetadataRawParamsWithTimeout(timeout time.Duration) *SetKibanaClusterMetadataRawParams {
	return &SetKibanaClusterMetadataRawParams{
		timeout: timeout,
	}
}

// NewSetKibanaClusterMetadataRawParamsWithContext creates a new SetKibanaClusterMetadataRawParams object
// with the ability to set a context for a request.
func NewSetKibanaClusterMetadataRawParamsWithContext(ctx context.Context) *SetKibanaClusterMetadataRawParams {
	return &SetKibanaClusterMetadataRawParams{
		Context: ctx,
	}
}

// NewSetKibanaClusterMetadataRawParamsWithHTTPClient creates a new SetKibanaClusterMetadataRawParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetKibanaClusterMetadataRawParamsWithHTTPClient(client *http.Client) *SetKibanaClusterMetadataRawParams {
	return &SetKibanaClusterMetadataRawParams{
		HTTPClient: client,
	}
}

/* SetKibanaClusterMetadataRawParams contains all the parameters to send to the API endpoint
   for the set kibana cluster metadata raw operation.

   Typically these are written to a http.Request.
*/
type SetKibanaClusterMetadataRawParams struct {

	/* Body.

	   The freeform JSON for the cluster (should always be based on the current version retrieved from the GET)
	*/
	Body string

	/* ClusterID.

	   The Kibana deployment identifier.
	*/
	ClusterID string

	/* Version.

	   Checks for conflicts against the metadata version, then returns the value in the `x-cloud-resource-version` header.
	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set kibana cluster metadata raw params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetKibanaClusterMetadataRawParams) WithDefaults() *SetKibanaClusterMetadataRawParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set kibana cluster metadata raw params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetKibanaClusterMetadataRawParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set kibana cluster metadata raw params
func (o *SetKibanaClusterMetadataRawParams) WithTimeout(timeout time.Duration) *SetKibanaClusterMetadataRawParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set kibana cluster metadata raw params
func (o *SetKibanaClusterMetadataRawParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set kibana cluster metadata raw params
func (o *SetKibanaClusterMetadataRawParams) WithContext(ctx context.Context) *SetKibanaClusterMetadataRawParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set kibana cluster metadata raw params
func (o *SetKibanaClusterMetadataRawParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set kibana cluster metadata raw params
func (o *SetKibanaClusterMetadataRawParams) WithHTTPClient(client *http.Client) *SetKibanaClusterMetadataRawParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set kibana cluster metadata raw params
func (o *SetKibanaClusterMetadataRawParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set kibana cluster metadata raw params
func (o *SetKibanaClusterMetadataRawParams) WithBody(body string) *SetKibanaClusterMetadataRawParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set kibana cluster metadata raw params
func (o *SetKibanaClusterMetadataRawParams) SetBody(body string) {
	o.Body = body
}

// WithClusterID adds the clusterID to the set kibana cluster metadata raw params
func (o *SetKibanaClusterMetadataRawParams) WithClusterID(clusterID string) *SetKibanaClusterMetadataRawParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the set kibana cluster metadata raw params
func (o *SetKibanaClusterMetadataRawParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithVersion adds the version to the set kibana cluster metadata raw params
func (o *SetKibanaClusterMetadataRawParams) WithVersion(version *int64) *SetKibanaClusterMetadataRawParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the set kibana cluster metadata raw params
func (o *SetKibanaClusterMetadataRawParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *SetKibanaClusterMetadataRawParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if o.Version != nil {

		// query param version
		var qrVersion int64

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := swag.FormatInt64(qrVersion)
		if qVersion != "" {

			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
