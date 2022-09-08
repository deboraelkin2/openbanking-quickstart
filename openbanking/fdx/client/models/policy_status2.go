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

// PolicyStatus2 PolicyStatus2
//
// The status of an insurance policy account
//
// swagger:model PolicyStatus2
type PolicyStatus2 string

func NewPolicyStatus2(value PolicyStatus2) *PolicyStatus2 {
	v := value
	return &v
}

const (

	// PolicyStatus2ACTIVE captures enum value "ACTIVE"
	PolicyStatus2ACTIVE PolicyStatus2 = "ACTIVE"

	// PolicyStatus2DEATHCLAIMPAID captures enum value "DEATH_CLAIM_PAID"
	PolicyStatus2DEATHCLAIMPAID PolicyStatus2 = "DEATH_CLAIM_PAID"

	// PolicyStatus2DEATHCLAIMPENDING captures enum value "DEATH_CLAIM_PENDING"
	PolicyStatus2DEATHCLAIMPENDING PolicyStatus2 = "DEATH_CLAIM_PENDING"

	// PolicyStatus2EXPIRED captures enum value "EXPIRED"
	PolicyStatus2EXPIRED PolicyStatus2 = "EXPIRED"

	// PolicyStatus2GRACEPERIOD captures enum value "GRACE_PERIOD"
	PolicyStatus2GRACEPERIOD PolicyStatus2 = "GRACE_PERIOD"

	// PolicyStatus2LAPSEPENDING captures enum value "LAPSE_PENDING"
	PolicyStatus2LAPSEPENDING PolicyStatus2 = "LAPSE_PENDING"

	// PolicyStatus2TERMINATED captures enum value "TERMINATED"
	PolicyStatus2TERMINATED PolicyStatus2 = "TERMINATED"

	// PolicyStatus2WAIVER captures enum value "WAIVER"
	PolicyStatus2WAIVER PolicyStatus2 = "WAIVER"
)

// for schema
var policyStatus2Enum []interface{}

func init() {
	var res []PolicyStatus2
	if err := json.Unmarshal([]byte(`["ACTIVE","DEATH_CLAIM_PAID","DEATH_CLAIM_PENDING","EXPIRED","GRACE_PERIOD","LAPSE_PENDING","TERMINATED","WAIVER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		policyStatus2Enum = append(policyStatus2Enum, v)
	}
}

func (m PolicyStatus2) validatePolicyStatus2Enum(path, location string, value PolicyStatus2) error {
	if err := validate.EnumCase(path, location, value, policyStatus2Enum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this policy status2
func (m PolicyStatus2) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePolicyStatus2Enum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this policy status2 based on context it is used
func (m PolicyStatus2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
