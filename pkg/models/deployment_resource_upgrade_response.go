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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeploymentResourceUpgradeResponse deployment resource upgrade response
//
// swagger:model DeploymentResourceUpgradeResponse
type DeploymentResourceUpgradeResponse struct {

	// resource id
	// Required: true
	ResourceID *string `json:"resource_id"`

	// stack version
	// Required: true
	StackVersion *string `json:"stack_version"`
}

// Validate validates this deployment resource upgrade response
func (m *DeploymentResourceUpgradeResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStackVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentResourceUpgradeResponse) validateResourceID(formats strfmt.Registry) error {

	if err := validate.Required("resource_id", "body", m.ResourceID); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentResourceUpgradeResponse) validateStackVersion(formats strfmt.Registry) error {

	if err := validate.Required("stack_version", "body", m.StackVersion); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this deployment resource upgrade response based on context it is used
func (m *DeploymentResourceUpgradeResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentResourceUpgradeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentResourceUpgradeResponse) UnmarshalBinary(b []byte) error {
	var res DeploymentResourceUpgradeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
