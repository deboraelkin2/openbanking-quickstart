// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccountWithDetailsentity4 AccountWithDetailsentity4
//
// swagger:model AccountWithDetailsentity4
type AccountWithDetailsentity4 struct {

	// insurance account
	InsuranceAccount *InsuranceAccountentity2 `json:"insuranceAccount,omitempty"`
}

// Validate validates this account with detailsentity4
func (m *AccountWithDetailsentity4) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInsuranceAccount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountWithDetailsentity4) validateInsuranceAccount(formats strfmt.Registry) error {
	if swag.IsZero(m.InsuranceAccount) { // not required
		return nil
	}

	if m.InsuranceAccount != nil {
		if err := m.InsuranceAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("insuranceAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("insuranceAccount")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this account with detailsentity4 based on the context it is used
func (m *AccountWithDetailsentity4) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInsuranceAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountWithDetailsentity4) contextValidateInsuranceAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.InsuranceAccount != nil {
		if err := m.InsuranceAccount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("insuranceAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("insuranceAccount")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountWithDetailsentity4) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountWithDetailsentity4) UnmarshalBinary(b []byte) error {
	var res AccountWithDetailsentity4
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
