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

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// RestartDeploymentEsResourceReader is a Reader for the RestartDeploymentEsResource structure.
type RestartDeploymentEsResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestartDeploymentEsResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewRestartDeploymentEsResourceAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRestartDeploymentEsResourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewRestartDeploymentEsResourceUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewRestartDeploymentEsResourceRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRestartDeploymentEsResourceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRestartDeploymentEsResourceAccepted creates a RestartDeploymentEsResourceAccepted with default headers values
func NewRestartDeploymentEsResourceAccepted() *RestartDeploymentEsResourceAccepted {
	return &RestartDeploymentEsResourceAccepted{}
}

/* RestartDeploymentEsResourceAccepted describes a response with status code 202, with default header values.

The restart command was issued successfully.
*/
type RestartDeploymentEsResourceAccepted struct {
	Payload *models.DeploymentResourceCommandResponse
}

func (o *RestartDeploymentEsResourceAccepted) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/elasticsearch/{ref_id}/_restart][%d] restartDeploymentEsResourceAccepted  %+v", 202, o.Payload)
}
func (o *RestartDeploymentEsResourceAccepted) GetPayload() *models.DeploymentResourceCommandResponse {
	return o.Payload
}

func (o *RestartDeploymentEsResourceAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentResourceCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestartDeploymentEsResourceNotFound creates a RestartDeploymentEsResourceNotFound with default headers values
func NewRestartDeploymentEsResourceNotFound() *RestartDeploymentEsResourceNotFound {
	return &RestartDeploymentEsResourceNotFound{}
}

/* RestartDeploymentEsResourceNotFound describes a response with status code 404, with default header values.

 * The Deployment specified by {deployment_id} cannot be found. (code: `deployments.deployment_not_found`)
* The Resource specified by {ref_id} cannot be found. (code: `deployments.deployment_resource_not_found`)
*/
type RestartDeploymentEsResourceNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *RestartDeploymentEsResourceNotFound) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/elasticsearch/{ref_id}/_restart][%d] restartDeploymentEsResourceNotFound  %+v", 404, o.Payload)
}
func (o *RestartDeploymentEsResourceNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *RestartDeploymentEsResourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRestartDeploymentEsResourceUnprocessableEntity creates a RestartDeploymentEsResourceUnprocessableEntity with default headers values
func NewRestartDeploymentEsResourceUnprocessableEntity() *RestartDeploymentEsResourceUnprocessableEntity {
	return &RestartDeploymentEsResourceUnprocessableEntity{}
}

/* RestartDeploymentEsResourceUnprocessableEntity describes a response with status code 422, with default header values.

The command sent to a Resource found the Resource in an illegal state, the error message gives more details. (code: `deployments.deployment_resource_plan_change_error`)
*/
type RestartDeploymentEsResourceUnprocessableEntity struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *RestartDeploymentEsResourceUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/elasticsearch/{ref_id}/_restart][%d] restartDeploymentEsResourceUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *RestartDeploymentEsResourceUnprocessableEntity) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *RestartDeploymentEsResourceUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRestartDeploymentEsResourceRetryWith creates a RestartDeploymentEsResourceRetryWith with default headers values
func NewRestartDeploymentEsResourceRetryWith() *RestartDeploymentEsResourceRetryWith {
	return &RestartDeploymentEsResourceRetryWith{}
}

/* RestartDeploymentEsResourceRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type RestartDeploymentEsResourceRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *RestartDeploymentEsResourceRetryWith) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/elasticsearch/{ref_id}/_restart][%d] restartDeploymentEsResourceRetryWith  %+v", 449, o.Payload)
}
func (o *RestartDeploymentEsResourceRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *RestartDeploymentEsResourceRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRestartDeploymentEsResourceInternalServerError creates a RestartDeploymentEsResourceInternalServerError with default headers values
func NewRestartDeploymentEsResourceInternalServerError() *RestartDeploymentEsResourceInternalServerError {
	return &RestartDeploymentEsResourceInternalServerError{}
}

/* RestartDeploymentEsResourceInternalServerError describes a response with status code 500, with default header values.

A Resource that was previously stored no longer exists. (code: `deployments.deployment_resource_no_longer_exists`)
*/
type RestartDeploymentEsResourceInternalServerError struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *RestartDeploymentEsResourceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/elasticsearch/{ref_id}/_restart][%d] restartDeploymentEsResourceInternalServerError  %+v", 500, o.Payload)
}
func (o *RestartDeploymentEsResourceInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *RestartDeploymentEsResourceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
