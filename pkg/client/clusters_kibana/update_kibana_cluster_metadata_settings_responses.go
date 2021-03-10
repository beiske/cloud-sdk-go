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

package clusters_kibana

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// UpdateKibanaClusterMetadataSettingsReader is a Reader for the UpdateKibanaClusterMetadataSettings structure.
type UpdateKibanaClusterMetadataSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateKibanaClusterMetadataSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateKibanaClusterMetadataSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateKibanaClusterMetadataSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateKibanaClusterMetadataSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewUpdateKibanaClusterMetadataSettingsRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateKibanaClusterMetadataSettingsOK creates a UpdateKibanaClusterMetadataSettingsOK with default headers values
func NewUpdateKibanaClusterMetadataSettingsOK() *UpdateKibanaClusterMetadataSettingsOK {
	return &UpdateKibanaClusterMetadataSettingsOK{}
}

/* UpdateKibanaClusterMetadataSettingsOK describes a response with status code 200, with default header values.

The cluster metadata was successfully updated
*/
type UpdateKibanaClusterMetadataSettingsOK struct {

	/* The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string

	/* The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string

	/* The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload *models.ClusterMetadataSettings
}

func (o *UpdateKibanaClusterMetadataSettingsOK) Error() string {
	return fmt.Sprintf("[PATCH /clusters/kibana/{cluster_id}/metadata/settings][%d] updateKibanaClusterMetadataSettingsOK  %+v", 200, o.Payload)
}
func (o *UpdateKibanaClusterMetadataSettingsOK) GetPayload() *models.ClusterMetadataSettings {
	return o.Payload
}

func (o *UpdateKibanaClusterMetadataSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-cloud-resource-created
	hdrXCloudResourceCreated := response.GetHeader("x-cloud-resource-created")

	if hdrXCloudResourceCreated != "" {
		o.XCloudResourceCreated = hdrXCloudResourceCreated
	}

	// hydrates response header x-cloud-resource-last-modified
	hdrXCloudResourceLastModified := response.GetHeader("x-cloud-resource-last-modified")

	if hdrXCloudResourceLastModified != "" {
		o.XCloudResourceLastModified = hdrXCloudResourceLastModified
	}

	// hydrates response header x-cloud-resource-version
	hdrXCloudResourceVersion := response.GetHeader("x-cloud-resource-version")

	if hdrXCloudResourceVersion != "" {
		o.XCloudResourceVersion = hdrXCloudResourceVersion
	}

	o.Payload = new(models.ClusterMetadataSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateKibanaClusterMetadataSettingsForbidden creates a UpdateKibanaClusterMetadataSettingsForbidden with default headers values
func NewUpdateKibanaClusterMetadataSettingsForbidden() *UpdateKibanaClusterMetadataSettingsForbidden {
	return &UpdateKibanaClusterMetadataSettingsForbidden{}
}

/* UpdateKibanaClusterMetadataSettingsForbidden describes a response with status code 403, with default header values.

The provided action was prohibited for the given cluster
*/
type UpdateKibanaClusterMetadataSettingsForbidden struct {
	Payload *models.BasicFailedReply
}

func (o *UpdateKibanaClusterMetadataSettingsForbidden) Error() string {
	return fmt.Sprintf("[PATCH /clusters/kibana/{cluster_id}/metadata/settings][%d] updateKibanaClusterMetadataSettingsForbidden  %+v", 403, o.Payload)
}
func (o *UpdateKibanaClusterMetadataSettingsForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateKibanaClusterMetadataSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateKibanaClusterMetadataSettingsNotFound creates a UpdateKibanaClusterMetadataSettingsNotFound with default headers values
func NewUpdateKibanaClusterMetadataSettingsNotFound() *UpdateKibanaClusterMetadataSettingsNotFound {
	return &UpdateKibanaClusterMetadataSettingsNotFound{}
}

/* UpdateKibanaClusterMetadataSettingsNotFound describes a response with status code 404, with default header values.

The cluster specified by {cluster_id} cannot be found (code: 'clusters.cluster_not_found')
*/
type UpdateKibanaClusterMetadataSettingsNotFound struct {
	Payload *models.BasicFailedReply
}

func (o *UpdateKibanaClusterMetadataSettingsNotFound) Error() string {
	return fmt.Sprintf("[PATCH /clusters/kibana/{cluster_id}/metadata/settings][%d] updateKibanaClusterMetadataSettingsNotFound  %+v", 404, o.Payload)
}
func (o *UpdateKibanaClusterMetadataSettingsNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateKibanaClusterMetadataSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateKibanaClusterMetadataSettingsRetryWith creates a UpdateKibanaClusterMetadataSettingsRetryWith with default headers values
func NewUpdateKibanaClusterMetadataSettingsRetryWith() *UpdateKibanaClusterMetadataSettingsRetryWith {
	return &UpdateKibanaClusterMetadataSettingsRetryWith{}
}

/* UpdateKibanaClusterMetadataSettingsRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: '"root.unauthorized.rbac.elevated_permissions_required"')
*/
type UpdateKibanaClusterMetadataSettingsRetryWith struct {
	Payload *models.BasicFailedReply
}

func (o *UpdateKibanaClusterMetadataSettingsRetryWith) Error() string {
	return fmt.Sprintf("[PATCH /clusters/kibana/{cluster_id}/metadata/settings][%d] updateKibanaClusterMetadataSettingsRetryWith  %+v", 449, o.Payload)
}
func (o *UpdateKibanaClusterMetadataSettingsRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateKibanaClusterMetadataSettingsRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
