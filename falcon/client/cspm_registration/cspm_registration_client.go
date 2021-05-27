// Code generated by go-swagger; DO NOT EDIT.

package cspm_registration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new cspm registration API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cspm registration API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateCSPMAwsAccount(params *CreateCSPMAwsAccountParams, opts ...ClientOption) (*CreateCSPMAwsAccountCreated, *CreateCSPMAwsAccountMultiStatus, error)

	CreateCSPMAzureAccount(params *CreateCSPMAzureAccountParams, opts ...ClientOption) (*CreateCSPMAzureAccountCreated, *CreateCSPMAzureAccountMultiStatus, error)

	DeleteCSPMAwsAccount(params *DeleteCSPMAwsAccountParams, opts ...ClientOption) (*DeleteCSPMAwsAccountOK, *DeleteCSPMAwsAccountMultiStatus, error)

	DeleteCSPMAzureAccount(params *DeleteCSPMAzureAccountParams, opts ...ClientOption) (*DeleteCSPMAzureAccountOK, *DeleteCSPMAzureAccountMultiStatus, error)

	GetCSPMAwsAccount(params *GetCSPMAwsAccountParams, opts ...ClientOption) (*GetCSPMAwsAccountOK, *GetCSPMAwsAccountMultiStatus, error)

	GetCSPMAwsAccountScriptsAttachment(params *GetCSPMAwsAccountScriptsAttachmentParams, opts ...ClientOption) (*GetCSPMAwsAccountScriptsAttachmentOK, error)

	GetCSPMAwsConsoleSetupURLs(params *GetCSPMAwsConsoleSetupURLsParams, opts ...ClientOption) (*GetCSPMAwsConsoleSetupURLsOK, *GetCSPMAwsConsoleSetupURLsMultiStatus, error)

	GetCSPMAzureAccount(params *GetCSPMAzureAccountParams, opts ...ClientOption) (*GetCSPMAzureAccountOK, *GetCSPMAzureAccountMultiStatus, error)

	GetCSPMAzureUserScriptsAttachment(params *GetCSPMAzureUserScriptsAttachmentParams, opts ...ClientOption) (*GetCSPMAzureUserScriptsAttachmentOK, error)

	GetCSPMPolicy(params *GetCSPMPolicyParams, opts ...ClientOption) (*GetCSPMPolicyOK, *GetCSPMPolicyMultiStatus, error)

	GetCSPMPolicySettings(params *GetCSPMPolicySettingsParams, opts ...ClientOption) (*GetCSPMPolicySettingsOK, *GetCSPMPolicySettingsMultiStatus, error)

	GetCSPMScanSchedule(params *GetCSPMScanScheduleParams, opts ...ClientOption) (*GetCSPMScanScheduleOK, error)

	GetIOAEvents(params *GetIOAEventsParams, opts ...ClientOption) (*GetIOAEventsOK, error)

	GetIOAUsers(params *GetIOAUsersParams, opts ...ClientOption) (*GetIOAUsersOK, error)

	PatchCSPMAwsAccount(params *PatchCSPMAwsAccountParams, opts ...ClientOption) (*PatchCSPMAwsAccountCreated, *PatchCSPMAwsAccountMultiStatus, error)

	UpdateCSPMAzureAccountClientID(params *UpdateCSPMAzureAccountClientIDParams, opts ...ClientOption) (*UpdateCSPMAzureAccountClientIDCreated, error)

	UpdateCSPMAzureTenantDefaultSubscriptionID(params *UpdateCSPMAzureTenantDefaultSubscriptionIDParams, opts ...ClientOption) (*UpdateCSPMAzureTenantDefaultSubscriptionIDCreated, error)

	UpdateCSPMPolicySettings(params *UpdateCSPMPolicySettingsParams, opts ...ClientOption) (*UpdateCSPMPolicySettingsOK, *UpdateCSPMPolicySettingsMultiStatus, error)

	UpdateCSPMScanSchedule(params *UpdateCSPMScanScheduleParams, opts ...ClientOption) (*UpdateCSPMScanScheduleOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateCSPMAwsAccount creates a new account in our system for a customer and generates a script for them to run in their a w s cloud environment to grant us access
*/
func (a *Client) CreateCSPMAwsAccount(params *CreateCSPMAwsAccountParams, opts ...ClientOption) (*CreateCSPMAwsAccountCreated, *CreateCSPMAwsAccountMultiStatus, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCSPMAwsAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateCSPMAwsAccount",
		Method:             "POST",
		PathPattern:        "/cloud-connect-cspm-aws/entities/account/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateCSPMAwsAccountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateCSPMAwsAccountCreated:
		return value, nil, nil
	case *CreateCSPMAwsAccountMultiStatus:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cspm_registration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateCSPMAzureAccount creates a new account in our system for a customer and generates a script for them to run in their cloud environment to grant us access
*/
func (a *Client) CreateCSPMAzureAccount(params *CreateCSPMAzureAccountParams, opts ...ClientOption) (*CreateCSPMAzureAccountCreated, *CreateCSPMAzureAccountMultiStatus, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCSPMAzureAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateCSPMAzureAccount",
		Method:             "POST",
		PathPattern:        "/cloud-connect-cspm-azure/entities/account/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateCSPMAzureAccountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateCSPMAzureAccountCreated:
		return value, nil, nil
	case *CreateCSPMAzureAccountMultiStatus:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cspm_registration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteCSPMAwsAccount deletes an existing a w s account or organization in our system
*/
func (a *Client) DeleteCSPMAwsAccount(params *DeleteCSPMAwsAccountParams, opts ...ClientOption) (*DeleteCSPMAwsAccountOK, *DeleteCSPMAwsAccountMultiStatus, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCSPMAwsAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteCSPMAwsAccount",
		Method:             "DELETE",
		PathPattern:        "/cloud-connect-cspm-aws/entities/account/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCSPMAwsAccountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteCSPMAwsAccountOK:
		return value, nil, nil
	case *DeleteCSPMAwsAccountMultiStatus:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteCSPMAwsAccountDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteCSPMAzureAccount deletes an azure subscription from the system
*/
func (a *Client) DeleteCSPMAzureAccount(params *DeleteCSPMAzureAccountParams, opts ...ClientOption) (*DeleteCSPMAzureAccountOK, *DeleteCSPMAzureAccountMultiStatus, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCSPMAzureAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteCSPMAzureAccount",
		Method:             "DELETE",
		PathPattern:        "/cloud-connect-cspm-azure/entities/account/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCSPMAzureAccountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteCSPMAzureAccountOK:
		return value, nil, nil
	case *DeleteCSPMAzureAccountMultiStatus:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteCSPMAzureAccountDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetCSPMAwsAccount returns information about the current status of an a w s account
*/
func (a *Client) GetCSPMAwsAccount(params *GetCSPMAwsAccountParams, opts ...ClientOption) (*GetCSPMAwsAccountOK, *GetCSPMAwsAccountMultiStatus, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCSPMAwsAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCSPMAwsAccount",
		Method:             "GET",
		PathPattern:        "/cloud-connect-cspm-aws/entities/account/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCSPMAwsAccountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetCSPMAwsAccountOK:
		return value, nil, nil
	case *GetCSPMAwsAccountMultiStatus:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCSPMAwsAccountDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetCSPMAwsAccountScriptsAttachment returns a script for customer to run in their cloud environment to grant us access to their a w s environment as a downloadable attachment
*/
func (a *Client) GetCSPMAwsAccountScriptsAttachment(params *GetCSPMAwsAccountScriptsAttachmentParams, opts ...ClientOption) (*GetCSPMAwsAccountScriptsAttachmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCSPMAwsAccountScriptsAttachmentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCSPMAwsAccountScriptsAttachment",
		Method:             "GET",
		PathPattern:        "/cloud-connect-cspm-aws/entities/user-scripts-download/v1",
		ProducesMediaTypes: []string{"application/json", "application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCSPMAwsAccountScriptsAttachmentReader{formats: a.formats},
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
	success, ok := result.(*GetCSPMAwsAccountScriptsAttachmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCSPMAwsAccountScriptsAttachmentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetCSPMAwsConsoleSetupURLs returns a URL for customer to visit in their cloud environment to grant us access to their a w s environment
*/
func (a *Client) GetCSPMAwsConsoleSetupURLs(params *GetCSPMAwsConsoleSetupURLsParams, opts ...ClientOption) (*GetCSPMAwsConsoleSetupURLsOK, *GetCSPMAwsConsoleSetupURLsMultiStatus, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCSPMAwsConsoleSetupURLsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCSPMAwsConsoleSetupURLs",
		Method:             "GET",
		PathPattern:        "/cloud-connect-cspm-aws/entities/console-setup-urls/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCSPMAwsConsoleSetupURLsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetCSPMAwsConsoleSetupURLsOK:
		return value, nil, nil
	case *GetCSPMAwsConsoleSetupURLsMultiStatus:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCSPMAwsConsoleSetupURLsDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetCSPMAzureAccount returns information about azure account registration
*/
func (a *Client) GetCSPMAzureAccount(params *GetCSPMAzureAccountParams, opts ...ClientOption) (*GetCSPMAzureAccountOK, *GetCSPMAzureAccountMultiStatus, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCSPMAzureAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCSPMAzureAccount",
		Method:             "GET",
		PathPattern:        "/cloud-connect-cspm-azure/entities/account/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCSPMAzureAccountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetCSPMAzureAccountOK:
		return value, nil, nil
	case *GetCSPMAzureAccountMultiStatus:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCSPMAzureAccountDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetCSPMAzureUserScriptsAttachment returns a script for customer to run in their cloud environment to grant us access to their azure environment as a downloadable attachment
*/
func (a *Client) GetCSPMAzureUserScriptsAttachment(params *GetCSPMAzureUserScriptsAttachmentParams, opts ...ClientOption) (*GetCSPMAzureUserScriptsAttachmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCSPMAzureUserScriptsAttachmentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCSPMAzureUserScriptsAttachment",
		Method:             "GET",
		PathPattern:        "/cloud-connect-cspm-azure/entities/user-scripts-download/v1",
		ProducesMediaTypes: []string{"application/json", "application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCSPMAzureUserScriptsAttachmentReader{formats: a.formats},
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
	success, ok := result.(*GetCSPMAzureUserScriptsAttachmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCSPMAzureUserScriptsAttachmentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetCSPMPolicy givens a policy ID returns detailed policy information
*/
func (a *Client) GetCSPMPolicy(params *GetCSPMPolicyParams, opts ...ClientOption) (*GetCSPMPolicyOK, *GetCSPMPolicyMultiStatus, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCSPMPolicyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCSPMPolicy",
		Method:             "GET",
		PathPattern:        "/settings/entities/policy-details/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCSPMPolicyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetCSPMPolicyOK:
		return value, nil, nil
	case *GetCSPMPolicyMultiStatus:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCSPMPolicyDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetCSPMPolicySettings returns information about current policy settings
*/
func (a *Client) GetCSPMPolicySettings(params *GetCSPMPolicySettingsParams, opts ...ClientOption) (*GetCSPMPolicySettingsOK, *GetCSPMPolicySettingsMultiStatus, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCSPMPolicySettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCSPMPolicySettings",
		Method:             "GET",
		PathPattern:        "/settings/entities/policy/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCSPMPolicySettingsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetCSPMPolicySettingsOK:
		return value, nil, nil
	case *GetCSPMPolicySettingsMultiStatus:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCSPMPolicySettingsDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetCSPMScanSchedule returns scan schedule configuration for one or more cloud platforms
*/
func (a *Client) GetCSPMScanSchedule(params *GetCSPMScanScheduleParams, opts ...ClientOption) (*GetCSPMScanScheduleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCSPMScanScheduleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCSPMScanSchedule",
		Method:             "GET",
		PathPattern:        "/settings/scan-schedule/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCSPMScanScheduleReader{formats: a.formats},
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
	success, ok := result.(*GetCSPMScanScheduleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCSPMScanScheduleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetIOAEvents fors c s p m i o a events gets list of i o a events
*/
func (a *Client) GetIOAEvents(params *GetIOAEventsParams, opts ...ClientOption) (*GetIOAEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIOAEventsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetIOAEvents",
		Method:             "GET",
		PathPattern:        "/ioa/entities/events/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIOAEventsReader{formats: a.formats},
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
	success, ok := result.(*GetIOAEventsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIOAEventsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetIOAUsers fors c s p m i o a users gets list of i o a users
*/
func (a *Client) GetIOAUsers(params *GetIOAUsersParams, opts ...ClientOption) (*GetIOAUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIOAUsersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetIOAUsers",
		Method:             "GET",
		PathPattern:        "/ioa/entities/users/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIOAUsersReader{formats: a.formats},
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
	success, ok := result.(*GetIOAUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIOAUsersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PatchCSPMAwsAccount patches a existing account in our system for a customer
*/
func (a *Client) PatchCSPMAwsAccount(params *PatchCSPMAwsAccountParams, opts ...ClientOption) (*PatchCSPMAwsAccountCreated, *PatchCSPMAwsAccountMultiStatus, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchCSPMAwsAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchCSPMAwsAccount",
		Method:             "PATCH",
		PathPattern:        "/cloud-connect-cspm-aws/entities/account/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchCSPMAwsAccountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PatchCSPMAwsAccountCreated:
		return value, nil, nil
	case *PatchCSPMAwsAccountMultiStatus:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cspm_registration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateCSPMAzureAccountClientID updates an azure service account in our system by with the user created client id created with the public key we ve provided
*/
func (a *Client) UpdateCSPMAzureAccountClientID(params *UpdateCSPMAzureAccountClientIDParams, opts ...ClientOption) (*UpdateCSPMAzureAccountClientIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCSPMAzureAccountClientIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateCSPMAzureAccountClientID",
		Method:             "PATCH",
		PathPattern:        "/cloud-connect-cspm-azure/entities/client-id/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateCSPMAzureAccountClientIDReader{formats: a.formats},
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
	success, ok := result.(*UpdateCSPMAzureAccountClientIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateCSPMAzureAccountClientID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateCSPMAzureTenantDefaultSubscriptionID updates an azure default subscription id in our system for given tenant id
*/
func (a *Client) UpdateCSPMAzureTenantDefaultSubscriptionID(params *UpdateCSPMAzureTenantDefaultSubscriptionIDParams, opts ...ClientOption) (*UpdateCSPMAzureTenantDefaultSubscriptionIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCSPMAzureTenantDefaultSubscriptionIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateCSPMAzureTenantDefaultSubscriptionID",
		Method:             "PATCH",
		PathPattern:        "/cloud-connect-cspm-azure/entities/default-subscription-id/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateCSPMAzureTenantDefaultSubscriptionIDReader{formats: a.formats},
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
	success, ok := result.(*UpdateCSPMAzureTenantDefaultSubscriptionIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateCSPMAzureTenantDefaultSubscriptionID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateCSPMPolicySettings updates a policy setting can be used to override policy severity or to disable a policy entirely
*/
func (a *Client) UpdateCSPMPolicySettings(params *UpdateCSPMPolicySettingsParams, opts ...ClientOption) (*UpdateCSPMPolicySettingsOK, *UpdateCSPMPolicySettingsMultiStatus, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCSPMPolicySettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateCSPMPolicySettings",
		Method:             "PATCH",
		PathPattern:        "/settings/entities/policy/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateCSPMPolicySettingsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateCSPMPolicySettingsOK:
		return value, nil, nil
	case *UpdateCSPMPolicySettingsMultiStatus:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateCSPMPolicySettingsDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateCSPMScanSchedule updates scan schedule configuration for one or more cloud platforms
*/
func (a *Client) UpdateCSPMScanSchedule(params *UpdateCSPMScanScheduleParams, opts ...ClientOption) (*UpdateCSPMScanScheduleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCSPMScanScheduleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateCSPMScanSchedule",
		Method:             "POST",
		PathPattern:        "/settings/scan-schedule/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateCSPMScanScheduleReader{formats: a.formats},
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
	success, ok := result.(*UpdateCSPMScanScheduleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateCSPMScanScheduleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
