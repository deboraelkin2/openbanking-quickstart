// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// MutualFundType2 MutualFundType2
//
// Mutual fund type. One of OPENEND, CLOSEEND, OTHER
//
// swagger:model MutualFundType2
type MutualFundType2 string

func NewMutualFundType2(value MutualFundType2) *MutualFundType2 {
	v := value
	return &v
}

const (

	// MutualFundType2OPENEND captures enum value "OPENEND"
	MutualFundType2OPENEND MutualFundType2 = "OPENEND"

	// MutualFundType2CLOSEEND captures enum value "CLOSEEND"
	MutualFundType2CLOSEEND MutualFundType2 = "CLOSEEND"

	// MutualFundType2OTHER captures enum value "OTHER"
	MutualFundType2OTHER MutualFundType2 = "OTHER"
)

// for schema
var mutualFundType2Enum []interface{}

func init() {
	var res []MutualFundType2
	if err := json.Unmarshal([]byte(`["OPENEND","CLOSEEND","OTHER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mutualFundType2Enum = append(mutualFundType2Enum, v)
	}
}

func (m MutualFundType2) validateMutualFundType2Enum(path, location string, value MutualFundType2) error {
	if err := validate.EnumCase(path, location, value, mutualFundType2Enum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this mutual fund type2
func (m MutualFundType2) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMutualFundType2Enum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this mutual fund type2 based on context it is used
func (m MutualFundType2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
