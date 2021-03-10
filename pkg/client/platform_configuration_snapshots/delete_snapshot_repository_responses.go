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

package platform_configuration_snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// DeleteSnapshotRepositoryReader is a Reader for the DeleteSnapshotRepository structure.
type DeleteSnapshotRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSnapshotRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSnapshotRepositoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewDeleteSnapshotRepositoryAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 449:
		result := NewDeleteSnapshotRepositoryRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteSnapshotRepositoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteSnapshotRepositoryOK creates a DeleteSnapshotRepositoryOK with default headers values
func NewDeleteSnapshotRepositoryOK() *DeleteSnapshotRepositoryOK {
	return &DeleteSnapshotRepositoryOK{}
}

/* DeleteSnapshotRepositoryOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteSnapshotRepositoryOK struct {
	Payload models.EmptyResponse
}

func (o *DeleteSnapshotRepositoryOK) Error() string {
	return fmt.Sprintf("[DELETE /platform/configuration/snapshots/repositories/{repository_name}][%d] deleteSnapshotRepositoryOK  %+v", 200, o.Payload)
}
func (o *DeleteSnapshotRepositoryOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *DeleteSnapshotRepositoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSnapshotRepositoryAccepted creates a DeleteSnapshotRepositoryAccepted with default headers values
func NewDeleteSnapshotRepositoryAccepted() *DeleteSnapshotRepositoryAccepted {
	return &DeleteSnapshotRepositoryAccepted{}
}

/* DeleteSnapshotRepositoryAccepted describes a response with status code 202, with default header values.

Delete snapshot repository config
*/
type DeleteSnapshotRepositoryAccepted struct {
	Payload models.EmptyResponse
}

func (o *DeleteSnapshotRepositoryAccepted) Error() string {
	return fmt.Sprintf("[DELETE /platform/configuration/snapshots/repositories/{repository_name}][%d] deleteSnapshotRepositoryAccepted  %+v", 202, o.Payload)
}
func (o *DeleteSnapshotRepositoryAccepted) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *DeleteSnapshotRepositoryAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSnapshotRepositoryRetryWith creates a DeleteSnapshotRepositoryRetryWith with default headers values
func NewDeleteSnapshotRepositoryRetryWith() *DeleteSnapshotRepositoryRetryWith {
	return &DeleteSnapshotRepositoryRetryWith{}
}

/* DeleteSnapshotRepositoryRetryWith describes a response with status code 449, with default header values.

elevated permissions are required. (code: '"root.unauthorized.rbac.elevated_permissions_required"')
*/
type DeleteSnapshotRepositoryRetryWith struct {
	Payload *models.BasicFailedReply
}

func (o *DeleteSnapshotRepositoryRetryWith) Error() string {
	return fmt.Sprintf("[DELETE /platform/configuration/snapshots/repositories/{repository_name}][%d] deleteSnapshotRepositoryRetryWith  %+v", 449, o.Payload)
}
func (o *DeleteSnapshotRepositoryRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteSnapshotRepositoryRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSnapshotRepositoryInternalServerError creates a DeleteSnapshotRepositoryInternalServerError with default headers values
func NewDeleteSnapshotRepositoryInternalServerError() *DeleteSnapshotRepositoryInternalServerError {
	return &DeleteSnapshotRepositoryInternalServerError{}
}

/* DeleteSnapshotRepositoryInternalServerError describes a response with status code 500, with default header values.

Failed to delete references and disable snapshots in one or more referencing clusters.
*/
type DeleteSnapshotRepositoryInternalServerError struct {
	Payload *models.BasicFailedReply
}

func (o *DeleteSnapshotRepositoryInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /platform/configuration/snapshots/repositories/{repository_name}][%d] deleteSnapshotRepositoryInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteSnapshotRepositoryInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteSnapshotRepositoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
