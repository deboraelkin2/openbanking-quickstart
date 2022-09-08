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

// RecurringPaymentforUpdateentity1 RecurringPaymentforUpdateentity1
//
// The data of the Recurring Payment to be scheduled
//
// swagger:model RecurringPaymentforUpdateentity1
type RecurringPaymentforUpdateentity1 struct {

	// Amount for the payment. Must be positive
	// Required: true
	// Minimum: 0
	Amount *float64 `json:"amount"`

	// Date that the funds are scheduled to be delivered
	// Example: 2021-07-15T00:00:00.000Z
	// Required: true
	// Format: date
	DueDate *strfmt.Date `json:"dueDate"`

	// duration
	Duration *RecurringPaymentDurationentity2 `json:"duration,omitempty"`

	// frequency
	// Required: true
	Frequency *RecurringPaymentFrequencyenum2 `json:"frequency"`

	// ID of the account used to source funds for payment
	// Required: true
	// Max Length: 256
	FromAccountID *string `json:"fromAccountId"`

	// User's account identifier with the payee
	MerchantAccountID string `json:"merchantAccountId,omitempty"`

	// ID of the payee to receive funds for the payment
	// Required: true
	// Max Length: 256
	ToPayeeID *string `json:"toPayeeId"`
}

// Validate validates this recurring paymentfor updateentity1
func (m *RecurringPaymentforUpdateentity1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDueDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFromAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToPayeeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecurringPaymentforUpdateentity1) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	if err := validate.Minimum("amount", "body", *m.Amount, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *RecurringPaymentforUpdateentity1) validateDueDate(formats strfmt.Registry) error {

	if err := validate.Required("dueDate", "body", m.DueDate); err != nil {
		return err
	}

	if err := validate.FormatOf("dueDate", "body", "date", m.DueDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecurringPaymentforUpdateentity1) validateDuration(formats strfmt.Registry) error {
	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	if m.Duration != nil {
		if err := m.Duration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("duration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("duration")
			}
			return err
		}
	}

	return nil
}

func (m *RecurringPaymentforUpdateentity1) validateFrequency(formats strfmt.Registry) error {

	if err := validate.Required("frequency", "body", m.Frequency); err != nil {
		return err
	}

	if err := validate.Required("frequency", "body", m.Frequency); err != nil {
		return err
	}

	if m.Frequency != nil {
		if err := m.Frequency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("frequency")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("frequency")
			}
			return err
		}
	}

	return nil
}

func (m *RecurringPaymentforUpdateentity1) validateFromAccountID(formats strfmt.Registry) error {

	if err := validate.Required("fromAccountId", "body", m.FromAccountID); err != nil {
		return err
	}

	if err := validate.MaxLength("fromAccountId", "body", *m.FromAccountID, 256); err != nil {
		return err
	}

	return nil
}

func (m *RecurringPaymentforUpdateentity1) validateToPayeeID(formats strfmt.Registry) error {

	if err := validate.Required("toPayeeId", "body", m.ToPayeeID); err != nil {
		return err
	}

	if err := validate.MaxLength("toPayeeId", "body", *m.ToPayeeID, 256); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this recurring paymentfor updateentity1 based on the context it is used
func (m *RecurringPaymentforUpdateentity1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDuration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFrequency(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecurringPaymentforUpdateentity1) contextValidateDuration(ctx context.Context, formats strfmt.Registry) error {

	if m.Duration != nil {
		if err := m.Duration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("duration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("duration")
			}
			return err
		}
	}

	return nil
}

func (m *RecurringPaymentforUpdateentity1) contextValidateFrequency(ctx context.Context, formats strfmt.Registry) error {

	if m.Frequency != nil {
		if err := m.Frequency.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("frequency")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("frequency")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecurringPaymentforUpdateentity1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecurringPaymentforUpdateentity1) UnmarshalBinary(b []byte) error {
	var res RecurringPaymentforUpdateentity1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
