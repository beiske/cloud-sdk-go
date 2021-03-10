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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StackVersionEnterpriseSearchConfig EnterpriseSearch related configuration of an Elastic Stack version
//
// swagger:model StackVersionEnterpriseSearchConfig
type StackVersionEnterpriseSearchConfig struct {

	// List of configuration options that cannot be overridden by user settings
	// Required: true
	Blacklist []string `json:"blacklist"`

	// capacity constraints
	CapacityConstraints *StackVersionInstanceCapacityConstraint `json:"capacity_constraints,omitempty"`

	// Node types that are compatible with this one
	CompatibleNodeTypes []string `json:"compatible_node_types"`

	// Docker image for the EnterpriseSearch
	// Required: true
	DockerImage *string `json:"docker_image"`

	// Node types that are supported by this stack version
	NodeTypes []*StackVersionNodeType `json:"node_types"`

	// Settings that are applied to all nodes of this type
	Settings interface{} `json:"settings,omitempty"`

	// Version of EnterpriseSearch
	Version string `json:"version,omitempty"`
}

// Validate validates this stack version enterprise search config
func (m *StackVersionEnterpriseSearchConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlacklist(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapacityConstraints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDockerImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackVersionEnterpriseSearchConfig) validateBlacklist(formats strfmt.Registry) error {

	if err := validate.Required("blacklist", "body", m.Blacklist); err != nil {
		return err
	}

	return nil
}

func (m *StackVersionEnterpriseSearchConfig) validateCapacityConstraints(formats strfmt.Registry) error {
	if swag.IsZero(m.CapacityConstraints) { // not required
		return nil
	}

	if m.CapacityConstraints != nil {
		if err := m.CapacityConstraints.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capacity_constraints")
			}
			return err
		}
	}

	return nil
}

func (m *StackVersionEnterpriseSearchConfig) validateDockerImage(formats strfmt.Registry) error {

	if err := validate.Required("docker_image", "body", m.DockerImage); err != nil {
		return err
	}

	return nil
}

func (m *StackVersionEnterpriseSearchConfig) validateNodeTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.NodeTypes); i++ {
		if swag.IsZero(m.NodeTypes[i]) { // not required
			continue
		}

		if m.NodeTypes[i] != nil {
			if err := m.NodeTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("node_types" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this stack version enterprise search config based on the context it is used
func (m *StackVersionEnterpriseSearchConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCapacityConstraints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodeTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackVersionEnterpriseSearchConfig) contextValidateCapacityConstraints(ctx context.Context, formats strfmt.Registry) error {

	if m.CapacityConstraints != nil {
		if err := m.CapacityConstraints.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capacity_constraints")
			}
			return err
		}
	}

	return nil
}

func (m *StackVersionEnterpriseSearchConfig) contextValidateNodeTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NodeTypes); i++ {

		if m.NodeTypes[i] != nil {
			if err := m.NodeTypes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("node_types" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StackVersionEnterpriseSearchConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackVersionEnterpriseSearchConfig) UnmarshalBinary(b []byte) error {
	var res StackVersionEnterpriseSearchConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
