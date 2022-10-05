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

// DomainExposedDataRecordSocialV1 domain exposed data record social v1
//
// swagger:model domain.ExposedDataRecordSocialV1
type DomainExposedDataRecordSocialV1 struct {

	// aim id
	// Required: true
	AimID *string `json:"aim_id"`

	// facebook id
	// Required: true
	FacebookID *string `json:"facebook_id"`

	// icq id
	// Required: true
	IcqID *string `json:"icq_id"`

	// instagram id
	// Required: true
	InstagramID *string `json:"instagram_id"`

	// msn id
	// Required: true
	MsnID *string `json:"msn_id"`

	// skype id
	// Required: true
	SkypeID *string `json:"skype_id"`

	// twitter id
	// Required: true
	TwitterID *string `json:"twitter_id"`

	// vk id
	// Required: true
	VkID *string `json:"vk_id"`
}

// Validate validates this domain exposed data record social v1
func (m *DomainExposedDataRecordSocialV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAimID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFacebookID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIcqID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstagramID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMsnID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSkypeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTwitterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVkID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainExposedDataRecordSocialV1) validateAimID(formats strfmt.Registry) error {

	if err := validate.Required("aim_id", "body", m.AimID); err != nil {
		return err
	}

	return nil
}

func (m *DomainExposedDataRecordSocialV1) validateFacebookID(formats strfmt.Registry) error {

	if err := validate.Required("facebook_id", "body", m.FacebookID); err != nil {
		return err
	}

	return nil
}

func (m *DomainExposedDataRecordSocialV1) validateIcqID(formats strfmt.Registry) error {

	if err := validate.Required("icq_id", "body", m.IcqID); err != nil {
		return err
	}

	return nil
}

func (m *DomainExposedDataRecordSocialV1) validateInstagramID(formats strfmt.Registry) error {

	if err := validate.Required("instagram_id", "body", m.InstagramID); err != nil {
		return err
	}

	return nil
}

func (m *DomainExposedDataRecordSocialV1) validateMsnID(formats strfmt.Registry) error {

	if err := validate.Required("msn_id", "body", m.MsnID); err != nil {
		return err
	}

	return nil
}

func (m *DomainExposedDataRecordSocialV1) validateSkypeID(formats strfmt.Registry) error {

	if err := validate.Required("skype_id", "body", m.SkypeID); err != nil {
		return err
	}

	return nil
}

func (m *DomainExposedDataRecordSocialV1) validateTwitterID(formats strfmt.Registry) error {

	if err := validate.Required("twitter_id", "body", m.TwitterID); err != nil {
		return err
	}

	return nil
}

func (m *DomainExposedDataRecordSocialV1) validateVkID(formats strfmt.Registry) error {

	if err := validate.Required("vk_id", "body", m.VkID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain exposed data record social v1 based on context it is used
func (m *DomainExposedDataRecordSocialV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainExposedDataRecordSocialV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainExposedDataRecordSocialV1) UnmarshalBinary(b []byte) error {
	var res DomainExposedDataRecordSocialV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
