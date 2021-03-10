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

// User An API user
//
// swagger:model User
type User struct {

	// True if the user is a built-in read-only user
	Builtin *bool `json:"builtin,omitempty"`

	// The user's optional email address
	Email string `json:"email,omitempty"`

	// The user's optional full name
	FullName string `json:"full_name,omitempty"`

	// The users metadata
	Metadata *UserMetadata `json:"metadata,omitempty"`

	// The user's security information
	// Required: true
	Security *UserSecurity `json:"security"`

	// The user's identifier
	// Required: true
	UserName *string `json:"user_name"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateSecurity(formats strfmt.Registry) error {

	if err := validate.Required("security", "body", m.Security); err != nil {
		return err
	}

	if m.Security != nil {
		if err := m.Security.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateUserName(formats strfmt.Registry) error {

	if err := validate.Required("user_name", "body", m.UserName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this user based on the context it is used
func (m *User) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {
		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *User) contextValidateSecurity(ctx context.Context, formats strfmt.Registry) error {

	if m.Security != nil {
		if err := m.Security.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
