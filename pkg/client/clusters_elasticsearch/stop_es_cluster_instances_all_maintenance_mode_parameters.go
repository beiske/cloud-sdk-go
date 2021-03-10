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

package clusters_elasticsearch

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

// NewStopEsClusterInstancesAllMaintenanceModeParams creates a new StopEsClusterInstancesAllMaintenanceModeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStopEsClusterInstancesAllMaintenanceModeParams() *StopEsClusterInstancesAllMaintenanceModeParams {
	return &StopEsClusterInstancesAllMaintenanceModeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStopEsClusterInstancesAllMaintenanceModeParamsWithTimeout creates a new StopEsClusterInstancesAllMaintenanceModeParams object
// with the ability to set a timeout on a request.
func NewStopEsClusterInstancesAllMaintenanceModeParamsWithTimeout(timeout time.Duration) *StopEsClusterInstancesAllMaintenanceModeParams {
	return &StopEsClusterInstancesAllMaintenanceModeParams{
		timeout: timeout,
	}
}

// NewStopEsClusterInstancesAllMaintenanceModeParamsWithContext creates a new StopEsClusterInstancesAllMaintenanceModeParams object
// with the ability to set a context for a request.
func NewStopEsClusterInstancesAllMaintenanceModeParamsWithContext(ctx context.Context) *StopEsClusterInstancesAllMaintenanceModeParams {
	return &StopEsClusterInstancesAllMaintenanceModeParams{
		Context: ctx,
	}
}

// NewStopEsClusterInstancesAllMaintenanceModeParamsWithHTTPClient creates a new StopEsClusterInstancesAllMaintenanceModeParams object
// with the ability to set a custom HTTPClient for a request.
func NewStopEsClusterInstancesAllMaintenanceModeParamsWithHTTPClient(client *http.Client) *StopEsClusterInstancesAllMaintenanceModeParams {
	return &StopEsClusterInstancesAllMaintenanceModeParams{
		HTTPClient: client,
	}
}

/* StopEsClusterInstancesAllMaintenanceModeParams contains all the parameters to send to the API endpoint
   for the stop es cluster instances all maintenance mode operation.

   Typically these are written to a http.Request.
*/
type StopEsClusterInstancesAllMaintenanceModeParams struct {

	/* ClusterID.

	   The Elasticsearch cluster identifier.
	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stop es cluster instances all maintenance mode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopEsClusterInstancesAllMaintenanceModeParams) WithDefaults() *StopEsClusterInstancesAllMaintenanceModeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stop es cluster instances all maintenance mode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopEsClusterInstancesAllMaintenanceModeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stop es cluster instances all maintenance mode params
func (o *StopEsClusterInstancesAllMaintenanceModeParams) WithTimeout(timeout time.Duration) *StopEsClusterInstancesAllMaintenanceModeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop es cluster instances all maintenance mode params
func (o *StopEsClusterInstancesAllMaintenanceModeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop es cluster instances all maintenance mode params
func (o *StopEsClusterInstancesAllMaintenanceModeParams) WithContext(ctx context.Context) *StopEsClusterInstancesAllMaintenanceModeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop es cluster instances all maintenance mode params
func (o *StopEsClusterInstancesAllMaintenanceModeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop es cluster instances all maintenance mode params
func (o *StopEsClusterInstancesAllMaintenanceModeParams) WithHTTPClient(client *http.Client) *StopEsClusterInstancesAllMaintenanceModeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop es cluster instances all maintenance mode params
func (o *StopEsClusterInstancesAllMaintenanceModeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the stop es cluster instances all maintenance mode params
func (o *StopEsClusterInstancesAllMaintenanceModeParams) WithClusterID(clusterID string) *StopEsClusterInstancesAllMaintenanceModeParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the stop es cluster instances all maintenance mode params
func (o *StopEsClusterInstancesAllMaintenanceModeParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *StopEsClusterInstancesAllMaintenanceModeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
