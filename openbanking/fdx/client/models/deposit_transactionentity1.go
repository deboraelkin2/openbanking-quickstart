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

// DepositTransactionentity1 DepositTransactionentity1
//
// A deposit transaction
//
// swagger:model DepositTransactionentity1
type DepositTransactionentity1 struct {

	// Corresponds to AccountId in Account
	// Max Length: 256
	AccountID string `json:"accountId,omitempty"`

	// The amount of money in the account currency
	Amount float64 `json:"amount,omitempty"`

	// Transaction category, preferably MCC or SIC.
	Category string `json:"category,omitempty"`

	// Check number
	CheckNumber int32 `json:"checkNumber,omitempty"`

	// debit credit memo
	DebitCreditMemo DebitCreditMemo2 `json:"debitCreditMemo,omitempty"`

	// The description of the transaction
	Description string `json:"description,omitempty"`

	// Array of FI-specific attributes
	FiAttributes []*FIAttributeentity `json:"fiAttributes"`

	// The amount of money in the foreign currency
	ForeignAmount float64 `json:"foreignAmount,omitempty"`

	// foreign currency
	ForeignCurrency ISO4217Code3 `json:"foreignCurrency,omitempty"`

	// Array of Image Identifiers (unique to this transaction) used to retrieve Images of check or transaction receipt
	ImageIds []string `json:"imageIds"`

	// Breakdown of the transaction details
	LineItem []*LineItementity `json:"lineItem"`

	// Links (unique to this transaction) used to retrieve images of checks or transaction receipts
	Links []*HATEOASLink `json:"links"`

	// Secondary transaction description
	// Max Length: 255
	Memo string `json:"memo,omitempty"`

	// Payee name
	// Max Length: 255
	Payee string `json:"payee,omitempty"`

	// The date and time that the transaction was posted to the account. If not provided then TransactionTimestamp can be used as PostedTimeStamp
	// Example: 2021-07-15T14:46:41.375Z
	// Format: date-time
	PostedTimestamp strfmt.DateTime `json:"postedTimestamp,omitempty"`

	// A tracking reference identifier
	Reference string `json:"reference,omitempty"`

	// For reverse postings, the identity of the transaction being reversed. For the correction transaction, the identity of the reversing post. For credit card posting transactions, the identity of the authorization transaction
	// Max Length: 256
	ReferenceTransactionID string `json:"referenceTransactionId,omitempty"`

	// reward
	Reward *TransactionRewardentity2 `json:"reward,omitempty"`

	// status
	Status TransactionStatus2 `json:"status,omitempty"`

	// Transaction category detail
	SubCategory string `json:"subCategory,omitempty"`

	// Long term persistent identity of the transaction (unique to account)
	// Max Length: 256
	TransactionID string `json:"transactionId,omitempty"`

	// The date and time that the transaction was added to the server backend systems
	// Example: 2021-07-15T14:46:41.375Z
	// Format: date-time
	TransactionTimestamp strfmt.DateTime `json:"transactionTimestamp,omitempty"`

	// transaction type
	TransactionType DepositTransactionType2 `json:"transactionType,omitempty"`
}

// Validate validates this deposit transactionentity1
func (m *DepositTransactionentity1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDebitCreditMemo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFiAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateForeignCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineItem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferenceTransactionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReward(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DepositTransactionentity1) validateAccountID(formats strfmt.Registry) error {
	if swag.IsZero(m.AccountID) { // not required
		return nil
	}

	if err := validate.MaxLength("accountId", "body", m.AccountID, 256); err != nil {
		return err
	}

	return nil
}

func (m *DepositTransactionentity1) validateDebitCreditMemo(formats strfmt.Registry) error {
	if swag.IsZero(m.DebitCreditMemo) { // not required
		return nil
	}

	if err := m.DebitCreditMemo.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("debitCreditMemo")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("debitCreditMemo")
		}
		return err
	}

	return nil
}

func (m *DepositTransactionentity1) validateFiAttributes(formats strfmt.Registry) error {
	if swag.IsZero(m.FiAttributes) { // not required
		return nil
	}

	for i := 0; i < len(m.FiAttributes); i++ {
		if swag.IsZero(m.FiAttributes[i]) { // not required
			continue
		}

		if m.FiAttributes[i] != nil {
			if err := m.FiAttributes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fiAttributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fiAttributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DepositTransactionentity1) validateForeignCurrency(formats strfmt.Registry) error {
	if swag.IsZero(m.ForeignCurrency) { // not required
		return nil
	}

	if err := m.ForeignCurrency.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("foreignCurrency")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("foreignCurrency")
		}
		return err
	}

	return nil
}

func (m *DepositTransactionentity1) validateLineItem(formats strfmt.Registry) error {
	if swag.IsZero(m.LineItem) { // not required
		return nil
	}

	for i := 0; i < len(m.LineItem); i++ {
		if swag.IsZero(m.LineItem[i]) { // not required
			continue
		}

		if m.LineItem[i] != nil {
			if err := m.LineItem[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("lineItem" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("lineItem" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DepositTransactionentity1) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	for i := 0; i < len(m.Links); i++ {
		if swag.IsZero(m.Links[i]) { // not required
			continue
		}

		if m.Links[i] != nil {
			if err := m.Links[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DepositTransactionentity1) validateMemo(formats strfmt.Registry) error {
	if swag.IsZero(m.Memo) { // not required
		return nil
	}

	if err := validate.MaxLength("memo", "body", m.Memo, 255); err != nil {
		return err
	}

	return nil
}

func (m *DepositTransactionentity1) validatePayee(formats strfmt.Registry) error {
	if swag.IsZero(m.Payee) { // not required
		return nil
	}

	if err := validate.MaxLength("payee", "body", m.Payee, 255); err != nil {
		return err
	}

	return nil
}

func (m *DepositTransactionentity1) validatePostedTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.PostedTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("postedTimestamp", "body", "date-time", m.PostedTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DepositTransactionentity1) validateReferenceTransactionID(formats strfmt.Registry) error {
	if swag.IsZero(m.ReferenceTransactionID) { // not required
		return nil
	}

	if err := validate.MaxLength("referenceTransactionId", "body", m.ReferenceTransactionID, 256); err != nil {
		return err
	}

	return nil
}

func (m *DepositTransactionentity1) validateReward(formats strfmt.Registry) error {
	if swag.IsZero(m.Reward) { // not required
		return nil
	}

	if m.Reward != nil {
		if err := m.Reward.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reward")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reward")
			}
			return err
		}
	}

	return nil
}

func (m *DepositTransactionentity1) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *DepositTransactionentity1) validateTransactionID(formats strfmt.Registry) error {
	if swag.IsZero(m.TransactionID) { // not required
		return nil
	}

	if err := validate.MaxLength("transactionId", "body", m.TransactionID, 256); err != nil {
		return err
	}

	return nil
}

func (m *DepositTransactionentity1) validateTransactionTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.TransactionTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("transactionTimestamp", "body", "date-time", m.TransactionTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DepositTransactionentity1) validateTransactionType(formats strfmt.Registry) error {
	if swag.IsZero(m.TransactionType) { // not required
		return nil
	}

	if err := m.TransactionType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("transactionType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("transactionType")
		}
		return err
	}

	return nil
}

// ContextValidate validate this deposit transactionentity1 based on the context it is used
func (m *DepositTransactionentity1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDebitCreditMemo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFiAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateForeignCurrency(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLineItem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReward(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransactionType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DepositTransactionentity1) contextValidateDebitCreditMemo(ctx context.Context, formats strfmt.Registry) error {

	if err := m.DebitCreditMemo.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("debitCreditMemo")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("debitCreditMemo")
		}
		return err
	}

	return nil
}

func (m *DepositTransactionentity1) contextValidateFiAttributes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FiAttributes); i++ {

		if m.FiAttributes[i] != nil {
			if err := m.FiAttributes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fiAttributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fiAttributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DepositTransactionentity1) contextValidateForeignCurrency(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ForeignCurrency.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("foreignCurrency")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("foreignCurrency")
		}
		return err
	}

	return nil
}

func (m *DepositTransactionentity1) contextValidateLineItem(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LineItem); i++ {

		if m.LineItem[i] != nil {
			if err := m.LineItem[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("lineItem" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("lineItem" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DepositTransactionentity1) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Links); i++ {

		if m.Links[i] != nil {
			if err := m.Links[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DepositTransactionentity1) contextValidateReward(ctx context.Context, formats strfmt.Registry) error {

	if m.Reward != nil {
		if err := m.Reward.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reward")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reward")
			}
			return err
		}
	}

	return nil
}

func (m *DepositTransactionentity1) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Status.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *DepositTransactionentity1) contextValidateTransactionType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.TransactionType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("transactionType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("transactionType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DepositTransactionentity1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DepositTransactionentity1) UnmarshalBinary(b []byte) error {
	var res DepositTransactionentity1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
