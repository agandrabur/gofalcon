// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RestapiIndicatorFeedQueryResponse restapi indicator feed query response
//
// swagger:model restapi.IndicatorFeedQueryResponse
type RestapiIndicatorFeedQueryResponse struct {

	// errors
	Errors []*MsaAPIError `json:"errors"`

	// feed name
	FeedName string `json:"feed_name,omitempty"`

	// meta
	// Required: true
	Meta *MsaMetaInfo `json:"meta"`

	// resources
	Resources []*RestapiIndicatorFeedQueryItem `json:"resources"`
}

// Validate validates this restapi indicator feed query response
func (m *RestapiIndicatorFeedQueryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestapiIndicatorFeedQueryResponse) validateErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RestapiIndicatorFeedQueryResponse) validateMeta(formats strfmt.Registry) error {

	if err := validate.Required("meta", "body", m.Meta); err != nil {
		return err
	}

	if m.Meta != nil {
		if err := m.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

func (m *RestapiIndicatorFeedQueryResponse) validateResources(formats strfmt.Registry) error {
	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	for i := 0; i < len(m.Resources); i++ {
		if swag.IsZero(m.Resources[i]) { // not required
			continue
		}

		if m.Resources[i] != nil {
			if err := m.Resources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this restapi indicator feed query response based on the context it is used
func (m *RestapiIndicatorFeedQueryResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestapiIndicatorFeedQueryResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Errors); i++ {

		if m.Errors[i] != nil {

			if swag.IsZero(m.Errors[i]) { // not required
				return nil
			}

			if err := m.Errors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RestapiIndicatorFeedQueryResponse) contextValidateMeta(ctx context.Context, formats strfmt.Registry) error {

	if m.Meta != nil {

		if err := m.Meta.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

func (m *RestapiIndicatorFeedQueryResponse) contextValidateResources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Resources); i++ {

		if m.Resources[i] != nil {

			if swag.IsZero(m.Resources[i]) { // not required
				return nil
			}

			if err := m.Resources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestapiIndicatorFeedQueryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestapiIndicatorFeedQueryResponse) UnmarshalBinary(b []byte) error {
	var res RestapiIndicatorFeedQueryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
