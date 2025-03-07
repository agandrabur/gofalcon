// Code generated by go-swagger; DO NOT EDIT.

package api_integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new api integrations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for api integrations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ExecuteCommand(params *ExecuteCommandParams, opts ...ClientOption) (*ExecuteCommandOK, error)

	ExecuteCommandProxy(params *ExecuteCommandProxyParams, opts ...ClientOption) (*ExecuteCommandProxyOK, error)

	GetCombinedPluginConfigs(params *GetCombinedPluginConfigsParams, opts ...ClientOption) (*GetCombinedPluginConfigsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ExecuteCommand executes a command
*/
func (a *Client) ExecuteCommand(params *ExecuteCommandParams, opts ...ClientOption) (*ExecuteCommandOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExecuteCommandParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExecuteCommand",
		Method:             "POST",
		PathPattern:        "/plugins/entities/execute/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExecuteCommandReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ExecuteCommandOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ExecuteCommand: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ExecuteCommandProxy executes a command and proxy the response directly
*/
func (a *Client) ExecuteCommandProxy(params *ExecuteCommandProxyParams, opts ...ClientOption) (*ExecuteCommandProxyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExecuteCommandProxyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExecuteCommandProxy",
		Method:             "POST",
		PathPattern:        "/plugins/entities/execute-proxy/v1",
		ProducesMediaTypes: []string{"*/*", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExecuteCommandProxyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ExecuteCommandProxyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ExecuteCommandProxy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCombinedPluginConfigs queries for config resources and returns details
*/
func (a *Client) GetCombinedPluginConfigs(params *GetCombinedPluginConfigsParams, opts ...ClientOption) (*GetCombinedPluginConfigsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCombinedPluginConfigsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCombinedPluginConfigs",
		Method:             "GET",
		PathPattern:        "/plugins/combined/configs/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCombinedPluginConfigsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCombinedPluginConfigsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetCombinedPluginConfigs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
