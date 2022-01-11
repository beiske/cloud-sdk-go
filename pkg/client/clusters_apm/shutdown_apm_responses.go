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

// ShutdownApmReader is a Reader for the ShutdownApm structure.
type ShutdownApmReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShutdownApmReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewShutdownApmAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewShutdownApmNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewShutdownApmRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShutdownApmAccepted creates a ShutdownApmAccepted with default headers values
func NewShutdownApmAccepted() *ShutdownApmAccepted {
	return &ShutdownApmAccepted{}
}

/* ShutdownApmAccepted describes a response with status code 202, with default header values.

The shutdown command was issued successfully. Use the "GET" command on the /{deployment_id} resource to monitor progress
*/
type ShutdownApmAccepted struct {
	Payload *models.ClusterCommandResponse
}

func (o *ShutdownApmAccepted) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/_shutdown][%d] shutdownApmAccepted  %+v", 202, o.Payload)
}
func (o *ShutdownApmAccepted) GetPayload() *models.ClusterCommandResponse {
	return o.Payload
}

func (o *ShutdownApmAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShutdownApmNotFound creates a ShutdownApmNotFound with default headers values
func NewShutdownApmNotFound() *ShutdownApmNotFound {
	return &ShutdownApmNotFound{}
}

/* ShutdownApmNotFound describes a response with status code 404, with default header values.

The cluster specified by {cluster_id} cannot be found. (code: `clusters.cluster_not_found`)
*/
type ShutdownApmNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *ShutdownApmNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/_shutdown][%d] shutdownApmNotFound  %+v", 404, o.Payload)
}
func (o *ShutdownApmNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *ShutdownApmNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewShutdownApmRetryWith creates a ShutdownApmRetryWith with default headers values
func NewShutdownApmRetryWith() *ShutdownApmRetryWith {
	return &ShutdownApmRetryWith{}
}

/* ShutdownApmRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type ShutdownApmRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *ShutdownApmRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/_shutdown][%d] shutdownApmRetryWith  %+v", 449, o.Payload)
}
func (o *ShutdownApmRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *ShutdownApmRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
