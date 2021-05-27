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

// DomainCIDPolicyAssignmentsNistBenchmark domain c ID policy assignments nist benchmark
//
// swagger:model domain.CIDPolicyAssignments.nist_benchmark
type DomainCIDPolicyAssignmentsNistBenchmark struct {

	// benchmark short
	// Required: true
	BenchmarkShort *string `json:"benchmark_short"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// recommendation number
	// Required: true
	RecommendationNumber *string `json:"recommendation_number"`
}

// Validate validates this domain c ID policy assignments nist benchmark
func (m *DomainCIDPolicyAssignmentsNistBenchmark) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBenchmarkShort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecommendationNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainCIDPolicyAssignmentsNistBenchmark) validateBenchmarkShort(formats strfmt.Registry) error {

	if err := validate.Required("benchmark_short", "body", m.BenchmarkShort); err != nil {
		return err
	}

	return nil
}

func (m *DomainCIDPolicyAssignmentsNistBenchmark) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainCIDPolicyAssignmentsNistBenchmark) validateRecommendationNumber(formats strfmt.Registry) error {

	if err := validate.Required("recommendation_number", "body", m.RecommendationNumber); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain c ID policy assignments nist benchmark based on context it is used
func (m *DomainCIDPolicyAssignmentsNistBenchmark) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainCIDPolicyAssignmentsNistBenchmark) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainCIDPolicyAssignmentsNistBenchmark) UnmarshalBinary(b []byte) error {
	var res DomainCIDPolicyAssignmentsNistBenchmark
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
