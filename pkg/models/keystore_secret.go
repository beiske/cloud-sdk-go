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

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KeystoreSecret The value that you configure for the Elasticsearch keystore secret.
//
// swagger:model KeystoreSecret
type KeystoreSecret struct {

	// Stores the keystore secret as a file. The default is false, which stores the keystore secret as string when value is a plain string, or true when value is an object.
	AsFile *bool `json:"as_file,omitempty"`

	// Value of this setting. This can either be a string or a JSON object that is stored as a JSON string in the keystore. NOTE: When the keystore secret is unspecified, it is removed.
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this keystore secret
func (m *KeystoreSecret) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this keystore secret based on context it is used
func (m *KeystoreSecret) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KeystoreSecret) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeystoreSecret) UnmarshalBinary(b []byte) error {
	var res KeystoreSecret
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
