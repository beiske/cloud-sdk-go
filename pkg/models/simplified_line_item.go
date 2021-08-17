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

// SimplifiedLineItem Line Item
//
// swagger:model SimplifiedLineItem
type SimplifiedLineItem struct {

	// Elastic Consumption Unit (ECU) Balance
	// Required: true
	EcuBalance *float64 `json:"ecu_balance"`

	// Original Elastic Consumption Unit (ECU) quantity
	// Required: true
	EcuQuantity *float64 `json:"ecu_quantity"`

	// Expiration of the line item
	// Required: true
	// Format: date-time
	End *strfmt.DateTime `json:"end"`

	// Line Item ID
	// Required: true
	ID *string `json:"id"`

	// Start of the line item's validity
	// Required: true
	// Format: date-time
	Start *strfmt.DateTime `json:"start"`
}

// Validate validates this simplified line item
func (m *SimplifiedLineItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEcuBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEcuQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SimplifiedLineItem) validateEcuBalance(formats strfmt.Registry) error {

	if err := validate.Required("ecu_balance", "body", m.EcuBalance); err != nil {
		return err
	}

	return nil
}

func (m *SimplifiedLineItem) validateEcuQuantity(formats strfmt.Registry) error {

	if err := validate.Required("ecu_quantity", "body", m.EcuQuantity); err != nil {
		return err
	}

	return nil
}

func (m *SimplifiedLineItem) validateEnd(formats strfmt.Registry) error {

	if err := validate.Required("end", "body", m.End); err != nil {
		return err
	}

	if err := validate.FormatOf("end", "body", "date-time", m.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SimplifiedLineItem) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *SimplifiedLineItem) validateStart(formats strfmt.Registry) error {

	if err := validate.Required("start", "body", m.Start); err != nil {
		return err
	}

	if err := validate.FormatOf("start", "body", "date-time", m.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this simplified line item based on context it is used
func (m *SimplifiedLineItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SimplifiedLineItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SimplifiedLineItem) UnmarshalBinary(b []byte) error {
	var res SimplifiedLineItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
