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

package platform_configuration_templates

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

// NewDeleteDeploymentTemplateParams creates a new DeleteDeploymentTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDeploymentTemplateParams() *DeleteDeploymentTemplateParams {
	return &DeleteDeploymentTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDeploymentTemplateParamsWithTimeout creates a new DeleteDeploymentTemplateParams object
// with the ability to set a timeout on a request.
func NewDeleteDeploymentTemplateParamsWithTimeout(timeout time.Duration) *DeleteDeploymentTemplateParams {
	return &DeleteDeploymentTemplateParams{
		timeout: timeout,
	}
}

// NewDeleteDeploymentTemplateParamsWithContext creates a new DeleteDeploymentTemplateParams object
// with the ability to set a context for a request.
func NewDeleteDeploymentTemplateParamsWithContext(ctx context.Context) *DeleteDeploymentTemplateParams {
	return &DeleteDeploymentTemplateParams{
		Context: ctx,
	}
}

// NewDeleteDeploymentTemplateParamsWithHTTPClient creates a new DeleteDeploymentTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDeploymentTemplateParamsWithHTTPClient(client *http.Client) *DeleteDeploymentTemplateParams {
	return &DeleteDeploymentTemplateParams{
		HTTPClient: client,
	}
}

/* DeleteDeploymentTemplateParams contains all the parameters to send to the API endpoint
   for the delete deployment template operation.

   Typically these are written to a http.Request.
*/
type DeleteDeploymentTemplateParams struct {

	/* TemplateID.

	   The identifier for the deployment template.
	*/
	TemplateID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete deployment template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDeploymentTemplateParams) WithDefaults() *DeleteDeploymentTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete deployment template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDeploymentTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete deployment template params
func (o *DeleteDeploymentTemplateParams) WithTimeout(timeout time.Duration) *DeleteDeploymentTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete deployment template params
func (o *DeleteDeploymentTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete deployment template params
func (o *DeleteDeploymentTemplateParams) WithContext(ctx context.Context) *DeleteDeploymentTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete deployment template params
func (o *DeleteDeploymentTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete deployment template params
func (o *DeleteDeploymentTemplateParams) WithHTTPClient(client *http.Client) *DeleteDeploymentTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete deployment template params
func (o *DeleteDeploymentTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTemplateID adds the templateID to the delete deployment template params
func (o *DeleteDeploymentTemplateParams) WithTemplateID(templateID string) *DeleteDeploymentTemplateParams {
	o.SetTemplateID(templateID)
	return o
}

// SetTemplateID adds the templateId to the delete deployment template params
func (o *DeleteDeploymentTemplateParams) SetTemplateID(templateID string) {
	o.TemplateID = templateID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDeploymentTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param template_id
	if err := r.SetPathParam("template_id", o.TemplateID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
