// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TypesQueryResponse types query response
//
// swagger:model types.QueryResponse
type TypesQueryResponse struct {

	// result JSON
	// Required: true
	ResultJSON []TypesQueryResponseResultJSON `json:"resultJSON"`

	// result type
	// Required: true
	ResultType *string `json:"resultType"`
}

// Validate validates this types query response
func (m *TypesQueryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResultJSON(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResultType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TypesQueryResponse) validateResultJSON(formats strfmt.Registry) error {

	if err := validate.Required("resultJSON", "body", m.ResultJSON); err != nil {
		return err
	}

	return nil
}

func (m *TypesQueryResponse) validateResultType(formats strfmt.Registry) error {

	if err := validate.Required("resultType", "body", m.ResultType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this types query response based on context it is used
func (m *TypesQueryResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TypesQueryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TypesQueryResponse) UnmarshalBinary(b []byte) error {
	var res TypesQueryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
