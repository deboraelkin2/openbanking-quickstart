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

// PaymentNetworkIdentifierType2 PaymentNetworkIdentifierType2
//
// Type of identifier
//
// swagger:model PaymentNetworkIdentifierType2
type PaymentNetworkIdentifierType2 string

func NewPaymentNetworkIdentifierType2(value PaymentNetworkIdentifierType2) *PaymentNetworkIdentifierType2 {
	v := value
	return &v
}

const (

	// PaymentNetworkIdentifierType2ACCOUNTNUMBER captures enum value "ACCOUNT_NUMBER"
	PaymentNetworkIdentifierType2ACCOUNTNUMBER PaymentNetworkIdentifierType2 = "ACCOUNT_NUMBER"

	// PaymentNetworkIdentifierType2TOKENIZEDACCOUNTNUMBER captures enum value "TOKENIZED_ACCOUNT_NUMBER"
	PaymentNetworkIdentifierType2TOKENIZEDACCOUNTNUMBER PaymentNetworkIdentifierType2 = "TOKENIZED_ACCOUNT_NUMBER"
)

// for schema
var paymentNetworkIdentifierType2Enum []interface{}

func init() {
	var res []PaymentNetworkIdentifierType2
	if err := json.Unmarshal([]byte(`["ACCOUNT_NUMBER","TOKENIZED_ACCOUNT_NUMBER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentNetworkIdentifierType2Enum = append(paymentNetworkIdentifierType2Enum, v)
	}
}

func (m PaymentNetworkIdentifierType2) validatePaymentNetworkIdentifierType2Enum(path, location string, value PaymentNetworkIdentifierType2) error {
	if err := validate.EnumCase(path, location, value, paymentNetworkIdentifierType2Enum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this payment network identifier type2
func (m PaymentNetworkIdentifierType2) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePaymentNetworkIdentifierType2Enum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this payment network identifier type2 based on context it is used
func (m PaymentNetworkIdentifierType2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
