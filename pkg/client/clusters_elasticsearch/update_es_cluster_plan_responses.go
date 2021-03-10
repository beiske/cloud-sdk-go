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

package clusters_elasticsearch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// UpdateEsClusterPlanReader is a Reader for the UpdateEsClusterPlan structure.
type UpdateEsClusterPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateEsClusterPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateEsClusterPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewUpdateEsClusterPlanAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateEsClusterPlanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateEsClusterPlanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewUpdateEsClusterPlanPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewUpdateEsClusterPlanRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateEsClusterPlanOK creates a UpdateEsClusterPlanOK with default headers values
func NewUpdateEsClusterPlanOK() *UpdateEsClusterPlanOK {
	return &UpdateEsClusterPlanOK{}
}

/* UpdateEsClusterPlanOK describes a response with status code 200, with default header values.

The cluster definition was valid - no further action was requested. The return object contains an internal representation of the plan, for use in debugging
*/
type UpdateEsClusterPlanOK struct {
	Payload *models.ClusterCrudResponse
}

func (o *UpdateEsClusterPlanOK) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/plan][%d] updateEsClusterPlanOK  %+v", 200, o.Payload)
}
func (o *UpdateEsClusterPlanOK) GetPayload() *models.ClusterCrudResponse {
	return o.Payload
}

func (o *UpdateEsClusterPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterCrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEsClusterPlanAccepted creates a UpdateEsClusterPlanAccepted with default headers values
func NewUpdateEsClusterPlanAccepted() *UpdateEsClusterPlanAccepted {
	return &UpdateEsClusterPlanAccepted{}
}

/* UpdateEsClusterPlanAccepted describes a response with status code 202, with default header values.

The plan definition was valid and the updated plan is in progress
*/
type UpdateEsClusterPlanAccepted struct {
	Payload *models.ClusterCrudResponse
}

func (o *UpdateEsClusterPlanAccepted) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/plan][%d] updateEsClusterPlanAccepted  %+v", 202, o.Payload)
}
func (o *UpdateEsClusterPlanAccepted) GetPayload() *models.ClusterCrudResponse {
	return o.Payload
}

func (o *UpdateEsClusterPlanAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterCrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEsClusterPlanBadRequest creates a UpdateEsClusterPlanBadRequest with default headers values
func NewUpdateEsClusterPlanBadRequest() *UpdateEsClusterPlanBadRequest {
	return &UpdateEsClusterPlanBadRequest{}
}

/* UpdateEsClusterPlanBadRequest describes a response with status code 400, with default header values.

 * The cluster definition contained errors. (code: `clusters.cluster_invalid_plan`)
* The features used in the cluster definition have not been implemented yet. (code: `clusters.plan_feature_not_implemented`)
*/
type UpdateEsClusterPlanBadRequest struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateEsClusterPlanBadRequest) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/plan][%d] updateEsClusterPlanBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateEsClusterPlanBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateEsClusterPlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateEsClusterPlanNotFound creates a UpdateEsClusterPlanNotFound with default headers values
func NewUpdateEsClusterPlanNotFound() *UpdateEsClusterPlanNotFound {
	return &UpdateEsClusterPlanNotFound{}
}

/* UpdateEsClusterPlanNotFound describes a response with status code 404, with default header values.

The cluster specified by {cluster_id} cannot be found. (code: `clusters.cluster_not_found`)
*/
type UpdateEsClusterPlanNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateEsClusterPlanNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/plan][%d] updateEsClusterPlanNotFound  %+v", 404, o.Payload)
}
func (o *UpdateEsClusterPlanNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateEsClusterPlanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateEsClusterPlanPreconditionFailed creates a UpdateEsClusterPlanPreconditionFailed with default headers values
func NewUpdateEsClusterPlanPreconditionFailed() *UpdateEsClusterPlanPreconditionFailed {
	return &UpdateEsClusterPlanPreconditionFailed{}
}

/* UpdateEsClusterPlanPreconditionFailed describes a response with status code 412, with default header values.

Potential risky settings have been specified. (code: `clusters.cluster_plan_state_error`)
*/
type UpdateEsClusterPlanPreconditionFailed struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateEsClusterPlanPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/plan][%d] updateEsClusterPlanPreconditionFailed  %+v", 412, o.Payload)
}
func (o *UpdateEsClusterPlanPreconditionFailed) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateEsClusterPlanPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateEsClusterPlanRetryWith creates a UpdateEsClusterPlanRetryWith with default headers values
func NewUpdateEsClusterPlanRetryWith() *UpdateEsClusterPlanRetryWith {
	return &UpdateEsClusterPlanRetryWith{}
}

/* UpdateEsClusterPlanRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type UpdateEsClusterPlanRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateEsClusterPlanRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/plan][%d] updateEsClusterPlanRetryWith  %+v", 449, o.Payload)
}
func (o *UpdateEsClusterPlanRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateEsClusterPlanRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
