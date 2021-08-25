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

// ResourceHeapDumps resource heap dumps
//
// swagger:model ResourceHeapDumps
type ResourceHeapDumps struct {

	// The heap dumps belonging to this resource
	// Required: true
	HeapDumps []*HeapDump `json:"heap_dumps"`

	// The randomly-generated id of a Resource
	// Required: true
	ID *string `json:"id"`

	// The locally-unique user-specified id of the resource that the heap dump was captured from
	// Required: true
	RefID *string `json:"ref_id"`
}

// Validate validates this resource heap dumps
func (m *ResourceHeapDumps) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHeapDumps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceHeapDumps) validateHeapDumps(formats strfmt.Registry) error {

	if err := validate.Required("heap_dumps", "body", m.HeapDumps); err != nil {
		return err
	}

	for i := 0; i < len(m.HeapDumps); i++ {
		if swag.IsZero(m.HeapDumps[i]) { // not required
			continue
		}

		if m.HeapDumps[i] != nil {
			if err := m.HeapDumps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("heap_dumps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResourceHeapDumps) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ResourceHeapDumps) validateRefID(formats strfmt.Registry) error {

	if err := validate.Required("ref_id", "body", m.RefID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this resource heap dumps based on the context it is used
func (m *ResourceHeapDumps) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHeapDumps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceHeapDumps) contextValidateHeapDumps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.HeapDumps); i++ {

		if m.HeapDumps[i] != nil {
			if err := m.HeapDumps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("heap_dumps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceHeapDumps) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceHeapDumps) UnmarshalBinary(b []byte) error {
	var res ResourceHeapDumps
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
