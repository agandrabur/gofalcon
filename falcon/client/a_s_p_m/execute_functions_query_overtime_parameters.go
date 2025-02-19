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

// NewExecuteFunctionsQueryOvertimeParams creates a new ExecuteFunctionsQueryOvertimeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExecuteFunctionsQueryOvertimeParams() *ExecuteFunctionsQueryOvertimeParams {
	return &ExecuteFunctionsQueryOvertimeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExecuteFunctionsQueryOvertimeParamsWithTimeout creates a new ExecuteFunctionsQueryOvertimeParams object
// with the ability to set a timeout on a request.
func NewExecuteFunctionsQueryOvertimeParamsWithTimeout(timeout time.Duration) *ExecuteFunctionsQueryOvertimeParams {
	return &ExecuteFunctionsQueryOvertimeParams{
		timeout: timeout,
	}
}

// NewExecuteFunctionsQueryOvertimeParamsWithContext creates a new ExecuteFunctionsQueryOvertimeParams object
// with the ability to set a context for a request.
func NewExecuteFunctionsQueryOvertimeParamsWithContext(ctx context.Context) *ExecuteFunctionsQueryOvertimeParams {
	return &ExecuteFunctionsQueryOvertimeParams{
		Context: ctx,
	}
}

// NewExecuteFunctionsQueryOvertimeParamsWithHTTPClient creates a new ExecuteFunctionsQueryOvertimeParams object
// with the ability to set a custom HTTPClient for a request.
func NewExecuteFunctionsQueryOvertimeParamsWithHTTPClient(client *http.Client) *ExecuteFunctionsQueryOvertimeParams {
	return &ExecuteFunctionsQueryOvertimeParams{
		HTTPClient: client,
	}
}

/*
ExecuteFunctionsQueryOvertimeParams contains all the parameters to send to the API endpoint

	for the execute functions query overtime operation.

	Typically these are written to a http.Request.
*/
type ExecuteFunctionsQueryOvertimeParams struct {

	// Body.
	Body *models.TypesMSAQueryRequest

	// Field.
	Field string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the execute functions query overtime params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecuteFunctionsQueryOvertimeParams) WithDefaults() *ExecuteFunctionsQueryOvertimeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the execute functions query overtime params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecuteFunctionsQueryOvertimeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the execute functions query overtime params
func (o *ExecuteFunctionsQueryOvertimeParams) WithTimeout(timeout time.Duration) *ExecuteFunctionsQueryOvertimeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the execute functions query overtime params
func (o *ExecuteFunctionsQueryOvertimeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the execute functions query overtime params
func (o *ExecuteFunctionsQueryOvertimeParams) WithContext(ctx context.Context) *ExecuteFunctionsQueryOvertimeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the execute functions query overtime params
func (o *ExecuteFunctionsQueryOvertimeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the execute functions query overtime params
func (o *ExecuteFunctionsQueryOvertimeParams) WithHTTPClient(client *http.Client) *ExecuteFunctionsQueryOvertimeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the execute functions query overtime params
func (o *ExecuteFunctionsQueryOvertimeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the execute functions query overtime params
func (o *ExecuteFunctionsQueryOvertimeParams) WithBody(body *models.TypesMSAQueryRequest) *ExecuteFunctionsQueryOvertimeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the execute functions query overtime params
func (o *ExecuteFunctionsQueryOvertimeParams) SetBody(body *models.TypesMSAQueryRequest) {
	o.Body = body
}

// WithField adds the field to the execute functions query overtime params
func (o *ExecuteFunctionsQueryOvertimeParams) WithField(field string) *ExecuteFunctionsQueryOvertimeParams {
	o.SetField(field)
	return o
}

// SetField adds the field to the execute functions query overtime params
func (o *ExecuteFunctionsQueryOvertimeParams) SetField(field string) {
	o.Field = field
}

// WriteToRequest writes these params to a swagger request
func (o *ExecuteFunctionsQueryOvertimeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
