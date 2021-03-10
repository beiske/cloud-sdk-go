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

package platform_configuration_security

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

// NewEnableSecurityDeploymentParams creates a new EnableSecurityDeploymentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEnableSecurityDeploymentParams() *EnableSecurityDeploymentParams {
	return &EnableSecurityDeploymentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEnableSecurityDeploymentParamsWithTimeout creates a new EnableSecurityDeploymentParams object
// with the ability to set a timeout on a request.
func NewEnableSecurityDeploymentParamsWithTimeout(timeout time.Duration) *EnableSecurityDeploymentParams {
	return &EnableSecurityDeploymentParams{
		timeout: timeout,
	}
}

// NewEnableSecurityDeploymentParamsWithContext creates a new EnableSecurityDeploymentParams object
// with the ability to set a context for a request.
func NewEnableSecurityDeploymentParamsWithContext(ctx context.Context) *EnableSecurityDeploymentParams {
	return &EnableSecurityDeploymentParams{
		Context: ctx,
	}
}

// NewEnableSecurityDeploymentParamsWithHTTPClient creates a new EnableSecurityDeploymentParams object
// with the ability to set a custom HTTPClient for a request.
func NewEnableSecurityDeploymentParamsWithHTTPClient(client *http.Client) *EnableSecurityDeploymentParams {
	return &EnableSecurityDeploymentParams{
		HTTPClient: client,
	}
}

/* EnableSecurityDeploymentParams contains all the parameters to send to the API endpoint
   for the enable security deployment operation.

   Typically these are written to a http.Request.
*/
type EnableSecurityDeploymentParams struct {

	/* Version.

	   When specified, checks for conflicts against the version stored in the persistent store (returned in 'x-cloud-resource-version' of the GET request)
	*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the enable security deployment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnableSecurityDeploymentParams) WithDefaults() *EnableSecurityDeploymentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the enable security deployment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnableSecurityDeploymentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the enable security deployment params
func (o *EnableSecurityDeploymentParams) WithTimeout(timeout time.Duration) *EnableSecurityDeploymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enable security deployment params
func (o *EnableSecurityDeploymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enable security deployment params
func (o *EnableSecurityDeploymentParams) WithContext(ctx context.Context) *EnableSecurityDeploymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enable security deployment params
func (o *EnableSecurityDeploymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enable security deployment params
func (o *EnableSecurityDeploymentParams) WithHTTPClient(client *http.Client) *EnableSecurityDeploymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enable security deployment params
func (o *EnableSecurityDeploymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVersion adds the version to the enable security deployment params
func (o *EnableSecurityDeploymentParams) WithVersion(version *string) *EnableSecurityDeploymentParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the enable security deployment params
func (o *EnableSecurityDeploymentParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *EnableSecurityDeploymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Version != nil {

		// query param version
		var qrVersion string

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := qrVersion
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
