// Code generated by go-swagger; DO NOT EDIT.

package a_s_p_m

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

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// NewExecuteFunctionDataQueryParams creates a new ExecuteFunctionDataQueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExecuteFunctionDataQueryParams() *ExecuteFunctionDataQueryParams {
	return &ExecuteFunctionDataQueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExecuteFunctionDataQueryParamsWithTimeout creates a new ExecuteFunctionDataQueryParams object
// with the ability to set a timeout on a request.
func NewExecuteFunctionDataQueryParamsWithTimeout(timeout time.Duration) *ExecuteFunctionDataQueryParams {
	return &ExecuteFunctionDataQueryParams{
		timeout: timeout,
	}
}

// NewExecuteFunctionDataQueryParamsWithContext creates a new ExecuteFunctionDataQueryParams object
// with the ability to set a context for a request.
func NewExecuteFunctionDataQueryParamsWithContext(ctx context.Context) *ExecuteFunctionDataQueryParams {
	return &ExecuteFunctionDataQueryParams{
		Context: ctx,
	}
}

// NewExecuteFunctionDataQueryParamsWithHTTPClient creates a new ExecuteFunctionDataQueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewExecuteFunctionDataQueryParamsWithHTTPClient(client *http.Client) *ExecuteFunctionDataQueryParams {
	return &ExecuteFunctionDataQueryParams{
		HTTPClient: client,
	}
}

/*
ExecuteFunctionDataQueryParams contains all the parameters to send to the API endpoint

	for the execute function data query operation.

	Typically these are written to a http.Request.
*/
type ExecuteFunctionDataQueryParams struct {

	// Body.
	Body *models.TypesMSAQueryRequest

	// Field.
	Field string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the execute function data query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecuteFunctionDataQueryParams) WithDefaults() *ExecuteFunctionDataQueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the execute function data query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecuteFunctionDataQueryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the execute function data query params
func (o *ExecuteFunctionDataQueryParams) WithTimeout(timeout time.Duration) *ExecuteFunctionDataQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the execute function data query params
func (o *ExecuteFunctionDataQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the execute function data query params
func (o *ExecuteFunctionDataQueryParams) WithContext(ctx context.Context) *ExecuteFunctionDataQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the execute function data query params
func (o *ExecuteFunctionDataQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the execute function data query params
func (o *ExecuteFunctionDataQueryParams) WithHTTPClient(client *http.Client) *ExecuteFunctionDataQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the execute function data query params
func (o *ExecuteFunctionDataQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the execute function data query params
func (o *ExecuteFunctionDataQueryParams) WithBody(body *models.TypesMSAQueryRequest) *ExecuteFunctionDataQueryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the execute function data query params
func (o *ExecuteFunctionDataQueryParams) SetBody(body *models.TypesMSAQueryRequest) {
	o.Body = body
}

// WithField adds the field to the execute function data query params
func (o *ExecuteFunctionDataQueryParams) WithField(field string) *ExecuteFunctionDataQueryParams {
	o.SetField(field)
	return o
}

// SetField adds the field to the execute function data query params
func (o *ExecuteFunctionDataQueryParams) SetField(field string) {
	o.Field = field
}

// WriteToRequest writes these params to a swagger request
func (o *ExecuteFunctionDataQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// query param field
	qrField := o.Field
	qField := qrField
	if qField != "" {

		if err := r.SetQueryParam("field", qField); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
