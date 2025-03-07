// Code generated by go-swagger; DO NOT EDIT.

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetDeploymentsExternalV1Params creates a new GetDeploymentsExternalV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeploymentsExternalV1Params() *GetDeploymentsExternalV1Params {
	return &GetDeploymentsExternalV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentsExternalV1ParamsWithTimeout creates a new GetDeploymentsExternalV1Params object
// with the ability to set a timeout on a request.
func NewGetDeploymentsExternalV1ParamsWithTimeout(timeout time.Duration) *GetDeploymentsExternalV1Params {
	return &GetDeploymentsExternalV1Params{
		timeout: timeout,
	}
}

// NewGetDeploymentsExternalV1ParamsWithContext creates a new GetDeploymentsExternalV1Params object
// with the ability to set a context for a request.
func NewGetDeploymentsExternalV1ParamsWithContext(ctx context.Context) *GetDeploymentsExternalV1Params {
	return &GetDeploymentsExternalV1Params{
		Context: ctx,
	}
}

// NewGetDeploymentsExternalV1ParamsWithHTTPClient creates a new GetDeploymentsExternalV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeploymentsExternalV1ParamsWithHTTPClient(client *http.Client) *GetDeploymentsExternalV1Params {
	return &GetDeploymentsExternalV1Params{
		HTTPClient: client,
	}
}

/*
GetDeploymentsExternalV1Params contains all the parameters to send to the API endpoint

	for the get deployments external v1 operation.

	Typically these are written to a http.Request.
*/
type GetDeploymentsExternalV1Params struct {

	/* Authorization.

	   authorization header
	*/
	Authorization string

	/* XCSUSERNAME.

	   user name
	*/
	XCSUSERNAME *string

	/* Ids.

	   release version ids to retrieve deployment details
	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get deployments external v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploymentsExternalV1Params) WithDefaults() *GetDeploymentsExternalV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get deployments external v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploymentsExternalV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get deployments external v1 params
func (o *GetDeploymentsExternalV1Params) WithTimeout(timeout time.Duration) *GetDeploymentsExternalV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deployments external v1 params
func (o *GetDeploymentsExternalV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deployments external v1 params
func (o *GetDeploymentsExternalV1Params) WithContext(ctx context.Context) *GetDeploymentsExternalV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deployments external v1 params
func (o *GetDeploymentsExternalV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deployments external v1 params
func (o *GetDeploymentsExternalV1Params) WithHTTPClient(client *http.Client) *GetDeploymentsExternalV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deployments external v1 params
func (o *GetDeploymentsExternalV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get deployments external v1 params
func (o *GetDeploymentsExternalV1Params) WithAuthorization(authorization string) *GetDeploymentsExternalV1Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get deployments external v1 params
func (o *GetDeploymentsExternalV1Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXCSUSERNAME adds the xCSUSERNAME to the get deployments external v1 params
func (o *GetDeploymentsExternalV1Params) WithXCSUSERNAME(xCSUSERNAME *string) *GetDeploymentsExternalV1Params {
	o.SetXCSUSERNAME(xCSUSERNAME)
	return o
}

// SetXCSUSERNAME adds the xCSUSERNAME to the get deployments external v1 params
func (o *GetDeploymentsExternalV1Params) SetXCSUSERNAME(xCSUSERNAME *string) {
	o.XCSUSERNAME = xCSUSERNAME
}

// WithIds adds the ids to the get deployments external v1 params
func (o *GetDeploymentsExternalV1Params) WithIds(ids []string) *GetDeploymentsExternalV1Params {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get deployments external v1 params
func (o *GetDeploymentsExternalV1Params) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentsExternalV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.XCSUSERNAME != nil {

		// header param X-CS-USERNAME
		if err := r.SetHeaderParam("X-CS-USERNAME", *o.XCSUSERNAME); err != nil {
			return err
		}
	}

	if o.Ids != nil {

		// binding items for ids
		joinedIds := o.bindParamIds(reg)

		// query array param ids
		if err := r.SetQueryParam("ids", joinedIds...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetDeploymentsExternalV1 binds the parameter ids
func (o *GetDeploymentsExternalV1Params) bindParamIds(formats strfmt.Registry) []string {
	idsIR := o.Ids

	var idsIC []string
	for _, idsIIR := range idsIR { // explode []string

		idsIIV := idsIIR // string as string
		idsIC = append(idsIC, idsIIV)
	}

	// items.CollectionFormat: "csv"
	idsIS := swag.JoinByFormat(idsIC, "csv")

	return idsIS
}
