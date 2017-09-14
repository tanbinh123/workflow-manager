package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ErrorEquals error equals
// swagger:model ErrorEquals
type ErrorEquals string

const (
	ErrorEqualsStatesALL ErrorEquals = "States.ALL"
)

// for schema
var errorEqualsEnum []interface{}

func init() {
	var res []ErrorEquals
	if err := json.Unmarshal([]byte(`["States.ALL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		errorEqualsEnum = append(errorEqualsEnum, v)
	}
}

func (m ErrorEquals) validateErrorEqualsEnum(path, location string, value ErrorEquals) error {
	if err := validate.Enum(path, location, value, errorEqualsEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this error equals
func (m ErrorEquals) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateErrorEqualsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}