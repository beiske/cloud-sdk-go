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

package clusters_apm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// MoveApmInstancesAdvancedReader is a Reader for the MoveApmInstancesAdvanced structure.
type MoveApmInstancesAdvancedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MoveApmInstancesAdvancedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewMoveApmInstancesAdvancedAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMoveApmInstancesAdvancedBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewMoveApmInstancesAdvancedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMoveApmInstancesAdvancedNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewMoveApmInstancesAdvancedRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMoveApmInstancesAdvancedAccepted creates a MoveApmInstancesAdvancedAccepted with default headers values
func NewMoveApmInstancesAdvancedAccepted() *MoveApmInstancesAdvancedAccepted {
	return &MoveApmInstancesAdvancedAccepted{}
}

/* MoveApmInstancesAdvancedAccepted describes a response with status code 202, with default header values.

The move command was issued successfully. Use the "GET" command on the /{deployment_id} resource to monitor progress
*/
type MoveApmInstancesAdvancedAccepted struct {
	Payload *models.ClusterCommandResponse
}

func (o *MoveApmInstancesAdvancedAccepted) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/instances/_move][%d] moveApmInstancesAdvancedAccepted  %+v", 202, o.Payload)
}
func (o *MoveApmInstancesAdvancedAccepted) GetPayload() *models.ClusterCommandResponse {
	return o.Payload
}

func (o *MoveApmInstancesAdvancedAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMoveApmInstancesAdvancedBadRequest creates a MoveApmInstancesAdvancedBadRequest with default headers values
func NewMoveApmInstancesAdvancedBadRequest() *MoveApmInstancesAdvancedBadRequest {
	return &MoveApmInstancesAdvancedBadRequest{}
}

/* MoveApmInstancesAdvancedBadRequest describes a response with status code 400, with default header values.

 * The cluster definition contained errors. (code: `clusters.cluster_invalid_plan`)
* The cluster definition contained errors. (code: `clusters.plan_feature_not_implemented`)
*/
type MoveApmInstancesAdvancedBadRequest struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *MoveApmInstancesAdvancedBadRequest) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/instances/_move][%d] moveApmInstancesAdvancedBadRequest  %+v", 400, o.Payload)
}
func (o *MoveApmInstancesAdvancedBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *MoveApmInstancesAdvancedBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-cloud-error-codes
	hdrXCloudErrorCodes := response.GetHeader("x-cloud-error-codes")

	if hdrXCloudErrorCodes != "" {
		o.XCloudErrorCodes = hdrXCloudErrorCodes
	}

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMoveApmInstancesAdvancedForbidden creates a MoveApmInstancesAdvancedForbidden with default headers values
func NewMoveApmInstancesAdvancedForbidden() *MoveApmInstancesAdvancedForbidden {
	return &MoveApmInstancesAdvancedForbidden{}
}

/* MoveApmInstancesAdvancedForbidden describes a response with status code 403, with default header values.

The move command was prohibited for the given cluster. (code: `clusters.command_prohibited`)
*/
type MoveApmInstancesAdvancedForbidden struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *MoveApmInstancesAdvancedForbidden) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/instances/_move][%d] moveApmInstancesAdvancedForbidden  %+v", 403, o.Payload)
}
func (o *MoveApmInstancesAdvancedForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *MoveApmInstancesAdvancedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-cloud-error-codes
	hdrXCloudErrorCodes := response.GetHeader("x-cloud-error-codes")

	if hdrXCloudErrorCodes != "" {
		o.XCloudErrorCodes = hdrXCloudErrorCodes
	}

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMoveApmInstancesAdvancedNotFound creates a MoveApmInstancesAdvancedNotFound with default headers values
func NewMoveApmInstancesAdvancedNotFound() *MoveApmInstancesAdvancedNotFound {
	return &MoveApmInstancesAdvancedNotFound{}
}

/* MoveApmInstancesAdvancedNotFound describes a response with status code 404, with default header values.

The cluster specified by {cluster_id} cannot be found. (code: `clusters.cluster_not_found`)
*/
type MoveApmInstancesAdvancedNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *MoveApmInstancesAdvancedNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/instances/_move][%d] moveApmInstancesAdvancedNotFound  %+v", 404, o.Payload)
}
func (o *MoveApmInstancesAdvancedNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *MoveApmInstancesAdvancedNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-cloud-error-codes
	hdrXCloudErrorCodes := response.GetHeader("x-cloud-error-codes")

	if hdrXCloudErrorCodes != "" {
		o.XCloudErrorCodes = hdrXCloudErrorCodes
	}

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMoveApmInstancesAdvancedRetryWith creates a MoveApmInstancesAdvancedRetryWith with default headers values
func NewMoveApmInstancesAdvancedRetryWith() *MoveApmInstancesAdvancedRetryWith {
	return &MoveApmInstancesAdvancedRetryWith{}
}

/* MoveApmInstancesAdvancedRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type MoveApmInstancesAdvancedRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *MoveApmInstancesAdvancedRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/instances/_move][%d] moveApmInstancesAdvancedRetryWith  %+v", 449, o.Payload)
}
func (o *MoveApmInstancesAdvancedRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *MoveApmInstancesAdvancedRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-cloud-error-codes
	hdrXCloudErrorCodes := response.GetHeader("x-cloud-error-codes")

	if hdrXCloudErrorCodes != "" {
		o.XCloudErrorCodes = hdrXCloudErrorCodes
	}

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
