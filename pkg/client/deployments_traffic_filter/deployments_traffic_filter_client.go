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

package deployments_traffic_filter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new deployments traffic filter API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for deployments traffic filter API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateTrafficFilterRuleset(params *CreateTrafficFilterRulesetParams, authInfo runtime.ClientAuthInfoWriter) (*CreateTrafficFilterRulesetCreated, error)

	CreateTrafficFilterRulesetAssociation(params *CreateTrafficFilterRulesetAssociationParams, authInfo runtime.ClientAuthInfoWriter) (*CreateTrafficFilterRulesetAssociationOK, *CreateTrafficFilterRulesetAssociationCreated, error)

	DeleteTrafficFilterRuleset(params *DeleteTrafficFilterRulesetParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteTrafficFilterRulesetOK, error)

	DeleteTrafficFilterRulesetAssociation(params *DeleteTrafficFilterRulesetAssociationParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteTrafficFilterRulesetAssociationOK, error)

	GetTrafficFilterDeploymentRulesetAssociations(params *GetTrafficFilterDeploymentRulesetAssociationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetTrafficFilterDeploymentRulesetAssociationsOK, error)

	GetTrafficFilterRuleset(params *GetTrafficFilterRulesetParams, authInfo runtime.ClientAuthInfoWriter) (*GetTrafficFilterRulesetOK, error)

	GetTrafficFilterRulesetDeploymentAssociations(params *GetTrafficFilterRulesetDeploymentAssociationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetTrafficFilterRulesetDeploymentAssociationsOK, error)

	GetTrafficFilterRulesets(params *GetTrafficFilterRulesetsParams, authInfo runtime.ClientAuthInfoWriter) (*GetTrafficFilterRulesetsOK, error)

	UpdateTrafficFilterRuleset(params *UpdateTrafficFilterRulesetParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateTrafficFilterRulesetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateTrafficFilterRuleset creates a ruleset

  Creates a ruleset that consists of a set of rules.
*/
func (a *Client) CreateTrafficFilterRuleset(params *CreateTrafficFilterRulesetParams, authInfo runtime.ClientAuthInfoWriter) (*CreateTrafficFilterRulesetCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTrafficFilterRulesetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "create-traffic-filter-ruleset",
		Method:             "POST",
		PathPattern:        "/deployments/traffic-filter/rulesets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateTrafficFilterRulesetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateTrafficFilterRulesetCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for create-traffic-filter-ruleset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateTrafficFilterRulesetAssociation creates ruleset association

  Applies the ruleset to the specified deployment.
*/
func (a *Client) CreateTrafficFilterRulesetAssociation(params *CreateTrafficFilterRulesetAssociationParams, authInfo runtime.ClientAuthInfoWriter) (*CreateTrafficFilterRulesetAssociationOK, *CreateTrafficFilterRulesetAssociationCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTrafficFilterRulesetAssociationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "create-traffic-filter-ruleset-association",
		Method:             "POST",
		PathPattern:        "/deployments/traffic-filter/rulesets/{ruleset_id}/associations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateTrafficFilterRulesetAssociationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateTrafficFilterRulesetAssociationOK:
		return value, nil, nil
	case *CreateTrafficFilterRulesetAssociationCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deployments_traffic_filter: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteTrafficFilterRuleset deletes a ruleset

  Deletes the ruleset by ID.
*/
func (a *Client) DeleteTrafficFilterRuleset(params *DeleteTrafficFilterRulesetParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteTrafficFilterRulesetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTrafficFilterRulesetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "delete-traffic-filter-ruleset",
		Method:             "DELETE",
		PathPattern:        "/deployments/traffic-filter/rulesets/{ruleset_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteTrafficFilterRulesetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteTrafficFilterRulesetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-traffic-filter-ruleset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteTrafficFilterRulesetAssociation deletes ruleset association

  Deletes the traffic rules in the ruleset from the deployment.
*/
func (a *Client) DeleteTrafficFilterRulesetAssociation(params *DeleteTrafficFilterRulesetAssociationParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteTrafficFilterRulesetAssociationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTrafficFilterRulesetAssociationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "delete-traffic-filter-ruleset-association",
		Method:             "DELETE",
		PathPattern:        "/deployments/traffic-filter/rulesets/{ruleset_id}/associations/{association_type}/{associated_entity_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteTrafficFilterRulesetAssociationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteTrafficFilterRulesetAssociationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-traffic-filter-ruleset-association: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTrafficFilterDeploymentRulesetAssociations gets associated rulesets

  Retrieves the rulesets associated with a deployment.
*/
func (a *Client) GetTrafficFilterDeploymentRulesetAssociations(params *GetTrafficFilterDeploymentRulesetAssociationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetTrafficFilterDeploymentRulesetAssociationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTrafficFilterDeploymentRulesetAssociationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-traffic-filter-deployment-ruleset-associations",
		Method:             "GET",
		PathPattern:        "/deployments/traffic-filter/associations/{association_type}/{associated_entity_id}/rulesets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTrafficFilterDeploymentRulesetAssociationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTrafficFilterDeploymentRulesetAssociationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-traffic-filter-deployment-ruleset-associations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTrafficFilterRuleset retrieves the ruleset by ID

  Retrieves a list of resources that are associated to the specified ruleset.
*/
func (a *Client) GetTrafficFilterRuleset(params *GetTrafficFilterRulesetParams, authInfo runtime.ClientAuthInfoWriter) (*GetTrafficFilterRulesetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTrafficFilterRulesetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-traffic-filter-ruleset",
		Method:             "GET",
		PathPattern:        "/deployments/traffic-filter/rulesets/{ruleset_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTrafficFilterRulesetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTrafficFilterRulesetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-traffic-filter-ruleset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTrafficFilterRulesetDeploymentAssociations gets associated deployments

  Retrieves a list of deployments that are associated to the specified ruleset.
*/
func (a *Client) GetTrafficFilterRulesetDeploymentAssociations(params *GetTrafficFilterRulesetDeploymentAssociationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetTrafficFilterRulesetDeploymentAssociationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTrafficFilterRulesetDeploymentAssociationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-traffic-filter-ruleset-deployment-associations",
		Method:             "GET",
		PathPattern:        "/deployments/traffic-filter/rulesets/{ruleset_id}/associations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTrafficFilterRulesetDeploymentAssociationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTrafficFilterRulesetDeploymentAssociationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-traffic-filter-ruleset-deployment-associations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTrafficFilterRulesets lists traffic filter rulesets

  List all of the traffic filter rulesets.
*/
func (a *Client) GetTrafficFilterRulesets(params *GetTrafficFilterRulesetsParams, authInfo runtime.ClientAuthInfoWriter) (*GetTrafficFilterRulesetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTrafficFilterRulesetsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-traffic-filter-rulesets",
		Method:             "GET",
		PathPattern:        "/deployments/traffic-filter/rulesets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTrafficFilterRulesetsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTrafficFilterRulesetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-traffic-filter-rulesets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateTrafficFilterRuleset updates a ruleset

  Updates the ruleset with the definition.
*/
func (a *Client) UpdateTrafficFilterRuleset(params *UpdateTrafficFilterRulesetParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateTrafficFilterRulesetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateTrafficFilterRulesetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "update-traffic-filter-ruleset",
		Method:             "PUT",
		PathPattern:        "/deployments/traffic-filter/rulesets/{ruleset_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateTrafficFilterRulesetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateTrafficFilterRulesetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for update-traffic-filter-ruleset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
