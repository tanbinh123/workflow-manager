// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WorkflowQuery workflow query
//
// swagger:model WorkflowQuery
type WorkflowQuery struct {

	// limit
	// Maximum: 10000
	Limit int64 `json:"limit,omitempty"`

	// oldest first
	OldestFirst bool `json:"oldestFirst,omitempty"`

	// page token
	PageToken string `json:"pageToken,omitempty"`

	// Tracks whether the resolvedByUser query parameter was sent or omitted in the request.
	ResolvedByUserWrapper *ResolvedByUserWrapper `json:"resolvedByUserWrapper,omitempty"`

	// status
	Status WorkflowStatus `json:"status,omitempty"`

	// summary only
	SummaryOnly *bool `json:"summaryOnly,omitempty"`

	// workflow definition name
	// Required: true
	WorkflowDefinitionName *string `json:"workflowDefinitionName"`
}

// Validate validates this workflow query
func (m *WorkflowQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResolvedByUserWrapper(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkflowDefinitionName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowQuery) validateLimit(formats strfmt.Registry) error {

	if swag.IsZero(m.Limit) { // not required
		return nil
	}

	if err := validate.MaximumInt("limit", "body", int64(m.Limit), 10000, false); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowQuery) validateResolvedByUserWrapper(formats strfmt.Registry) error {

	if swag.IsZero(m.ResolvedByUserWrapper) { // not required
		return nil
	}

	if m.ResolvedByUserWrapper != nil {
		if err := m.ResolvedByUserWrapper.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resolvedByUserWrapper")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowQuery) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *WorkflowQuery) validateWorkflowDefinitionName(formats strfmt.Registry) error {

	if err := validate.Required("workflowDefinitionName", "body", m.WorkflowDefinitionName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowQuery) UnmarshalBinary(b []byte) error {
	var res WorkflowQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
