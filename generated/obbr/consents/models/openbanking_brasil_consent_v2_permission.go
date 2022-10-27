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

// OpenbankingBrasilConsentV2Permission Permission
//
// Especifica os tipos de permisses de acesso s APIs no escopo do Open Banking Brasil - Fase 2, de acordo com os blocos de consentimento fornecidos pelo usurio e necessrios ao acesso a cada endpoint das APIs.
//
// swagger:model OpenbankingBrasilConsentV2Permission
type OpenbankingBrasilConsentV2Permission string

func NewOpenbankingBrasilConsentV2Permission(value OpenbankingBrasilConsentV2Permission) *OpenbankingBrasilConsentV2Permission {
	v := value
	return &v
}

const (

	// OpenbankingBrasilConsentV2PermissionACCOUNTSREAD captures enum value "ACCOUNTS_READ"
	OpenbankingBrasilConsentV2PermissionACCOUNTSREAD OpenbankingBrasilConsentV2Permission = "ACCOUNTS_READ"

	// OpenbankingBrasilConsentV2PermissionACCOUNTSBALANCESREAD captures enum value "ACCOUNTS_BALANCES_READ"
	OpenbankingBrasilConsentV2PermissionACCOUNTSBALANCESREAD OpenbankingBrasilConsentV2Permission = "ACCOUNTS_BALANCES_READ"

	// OpenbankingBrasilConsentV2PermissionACCOUNTSTRANSACTIONSREAD captures enum value "ACCOUNTS_TRANSACTIONS_READ"
	OpenbankingBrasilConsentV2PermissionACCOUNTSTRANSACTIONSREAD OpenbankingBrasilConsentV2Permission = "ACCOUNTS_TRANSACTIONS_READ"

	// OpenbankingBrasilConsentV2PermissionACCOUNTSOVERDRAFTLIMITSREAD captures enum value "ACCOUNTS_OVERDRAFT_LIMITS_READ"
	OpenbankingBrasilConsentV2PermissionACCOUNTSOVERDRAFTLIMITSREAD OpenbankingBrasilConsentV2Permission = "ACCOUNTS_OVERDRAFT_LIMITS_READ"

	// OpenbankingBrasilConsentV2PermissionCREDITCARDSACCOUNTSREAD captures enum value "CREDIT_CARDS_ACCOUNTS_READ"
	OpenbankingBrasilConsentV2PermissionCREDITCARDSACCOUNTSREAD OpenbankingBrasilConsentV2Permission = "CREDIT_CARDS_ACCOUNTS_READ"

	// OpenbankingBrasilConsentV2PermissionCREDITCARDSACCOUNTSBILLSREAD captures enum value "CREDIT_CARDS_ACCOUNTS_BILLS_READ"
	OpenbankingBrasilConsentV2PermissionCREDITCARDSACCOUNTSBILLSREAD OpenbankingBrasilConsentV2Permission = "CREDIT_CARDS_ACCOUNTS_BILLS_READ"

	// OpenbankingBrasilConsentV2PermissionCREDITCARDSACCOUNTSBILLSTRANSACTIONSREAD captures enum value "CREDIT_CARDS_ACCOUNTS_BILLS_TRANSACTIONS_READ"
	OpenbankingBrasilConsentV2PermissionCREDITCARDSACCOUNTSBILLSTRANSACTIONSREAD OpenbankingBrasilConsentV2Permission = "CREDIT_CARDS_ACCOUNTS_BILLS_TRANSACTIONS_READ"

	// OpenbankingBrasilConsentV2PermissionCREDITCARDSACCOUNTSLIMITSREAD captures enum value "CREDIT_CARDS_ACCOUNTS_LIMITS_READ"
	OpenbankingBrasilConsentV2PermissionCREDITCARDSACCOUNTSLIMITSREAD OpenbankingBrasilConsentV2Permission = "CREDIT_CARDS_ACCOUNTS_LIMITS_READ"

	// OpenbankingBrasilConsentV2PermissionCREDITCARDSACCOUNTSTRANSACTIONSREAD captures enum value "CREDIT_CARDS_ACCOUNTS_TRANSACTIONS_READ"
	OpenbankingBrasilConsentV2PermissionCREDITCARDSACCOUNTSTRANSACTIONSREAD OpenbankingBrasilConsentV2Permission = "CREDIT_CARDS_ACCOUNTS_TRANSACTIONS_READ"

	// OpenbankingBrasilConsentV2PermissionCUSTOMERSPERSONALIDENTIFICATIONSREAD captures enum value "CUSTOMERS_PERSONAL_IDENTIFICATIONS_READ"
	OpenbankingBrasilConsentV2PermissionCUSTOMERSPERSONALIDENTIFICATIONSREAD OpenbankingBrasilConsentV2Permission = "CUSTOMERS_PERSONAL_IDENTIFICATIONS_READ"

	// OpenbankingBrasilConsentV2PermissionCUSTOMERSPERSONALADITTIONALINFOREAD captures enum value "CUSTOMERS_PERSONAL_ADITTIONALINFO_READ"
	OpenbankingBrasilConsentV2PermissionCUSTOMERSPERSONALADITTIONALINFOREAD OpenbankingBrasilConsentV2Permission = "CUSTOMERS_PERSONAL_ADITTIONALINFO_READ"

	// OpenbankingBrasilConsentV2PermissionCUSTOMERSBUSINESSIDENTIFICATIONSREAD captures enum value "CUSTOMERS_BUSINESS_IDENTIFICATIONS_READ"
	OpenbankingBrasilConsentV2PermissionCUSTOMERSBUSINESSIDENTIFICATIONSREAD OpenbankingBrasilConsentV2Permission = "CUSTOMERS_BUSINESS_IDENTIFICATIONS_READ"

	// OpenbankingBrasilConsentV2PermissionCUSTOMERSBUSINESSADITTIONALINFOREAD captures enum value "CUSTOMERS_BUSINESS_ADITTIONALINFO_READ"
	OpenbankingBrasilConsentV2PermissionCUSTOMERSBUSINESSADITTIONALINFOREAD OpenbankingBrasilConsentV2Permission = "CUSTOMERS_BUSINESS_ADITTIONALINFO_READ"

	// OpenbankingBrasilConsentV2PermissionFINANCINGSREAD captures enum value "FINANCINGS_READ"
	OpenbankingBrasilConsentV2PermissionFINANCINGSREAD OpenbankingBrasilConsentV2Permission = "FINANCINGS_READ"

	// OpenbankingBrasilConsentV2PermissionFINANCINGSSCHEDULEDINSTALMENTSREAD captures enum value "FINANCINGS_SCHEDULED_INSTALMENTS_READ"
	OpenbankingBrasilConsentV2PermissionFINANCINGSSCHEDULEDINSTALMENTSREAD OpenbankingBrasilConsentV2Permission = "FINANCINGS_SCHEDULED_INSTALMENTS_READ"

	// OpenbankingBrasilConsentV2PermissionFINANCINGSPAYMENTSREAD captures enum value "FINANCINGS_PAYMENTS_READ"
	OpenbankingBrasilConsentV2PermissionFINANCINGSPAYMENTSREAD OpenbankingBrasilConsentV2Permission = "FINANCINGS_PAYMENTS_READ"

	// OpenbankingBrasilConsentV2PermissionFINANCINGSWARRANTIESREAD captures enum value "FINANCINGS_WARRANTIES_READ"
	OpenbankingBrasilConsentV2PermissionFINANCINGSWARRANTIESREAD OpenbankingBrasilConsentV2Permission = "FINANCINGS_WARRANTIES_READ"

	// OpenbankingBrasilConsentV2PermissionINVOICEFINANCINGSREAD captures enum value "INVOICE_FINANCINGS_READ"
	OpenbankingBrasilConsentV2PermissionINVOICEFINANCINGSREAD OpenbankingBrasilConsentV2Permission = "INVOICE_FINANCINGS_READ"

	// OpenbankingBrasilConsentV2PermissionINVOICEFINANCINGSSCHEDULEDINSTALMENTSREAD captures enum value "INVOICE_FINANCINGS_SCHEDULED_INSTALMENTS_READ"
	OpenbankingBrasilConsentV2PermissionINVOICEFINANCINGSSCHEDULEDINSTALMENTSREAD OpenbankingBrasilConsentV2Permission = "INVOICE_FINANCINGS_SCHEDULED_INSTALMENTS_READ"

	// OpenbankingBrasilConsentV2PermissionINVOICEFINANCINGSPAYMENTSREAD captures enum value "INVOICE_FINANCINGS_PAYMENTS_READ"
	OpenbankingBrasilConsentV2PermissionINVOICEFINANCINGSPAYMENTSREAD OpenbankingBrasilConsentV2Permission = "INVOICE_FINANCINGS_PAYMENTS_READ"

	// OpenbankingBrasilConsentV2PermissionINVOICEFINANCINGSWARRANTIESREAD captures enum value "INVOICE_FINANCINGS_WARRANTIES_READ"
	OpenbankingBrasilConsentV2PermissionINVOICEFINANCINGSWARRANTIESREAD OpenbankingBrasilConsentV2Permission = "INVOICE_FINANCINGS_WARRANTIES_READ"

	// OpenbankingBrasilConsentV2PermissionLOANSREAD captures enum value "LOANS_READ"
	OpenbankingBrasilConsentV2PermissionLOANSREAD OpenbankingBrasilConsentV2Permission = "LOANS_READ"

	// OpenbankingBrasilConsentV2PermissionLOANSSCHEDULEDINSTALMENTSREAD captures enum value "LOANS_SCHEDULED_INSTALMENTS_READ"
	OpenbankingBrasilConsentV2PermissionLOANSSCHEDULEDINSTALMENTSREAD OpenbankingBrasilConsentV2Permission = "LOANS_SCHEDULED_INSTALMENTS_READ"

	// OpenbankingBrasilConsentV2PermissionLOANSPAYMENTSREAD captures enum value "LOANS_PAYMENTS_READ"
	OpenbankingBrasilConsentV2PermissionLOANSPAYMENTSREAD OpenbankingBrasilConsentV2Permission = "LOANS_PAYMENTS_READ"

	// OpenbankingBrasilConsentV2PermissionLOANSWARRANTIESREAD captures enum value "LOANS_WARRANTIES_READ"
	OpenbankingBrasilConsentV2PermissionLOANSWARRANTIESREAD OpenbankingBrasilConsentV2Permission = "LOANS_WARRANTIES_READ"

	// OpenbankingBrasilConsentV2PermissionUNARRANGEDACCOUNTSOVERDRAFTREAD captures enum value "UNARRANGED_ACCOUNTS_OVERDRAFT_READ"
	OpenbankingBrasilConsentV2PermissionUNARRANGEDACCOUNTSOVERDRAFTREAD OpenbankingBrasilConsentV2Permission = "UNARRANGED_ACCOUNTS_OVERDRAFT_READ"

	// OpenbankingBrasilConsentV2PermissionUNARRANGEDACCOUNTSOVERDRAFTSCHEDULEDINSTALMENTSREAD captures enum value "UNARRANGED_ACCOUNTS_OVERDRAFT_SCHEDULED_INSTALMENTS_READ"
	OpenbankingBrasilConsentV2PermissionUNARRANGEDACCOUNTSOVERDRAFTSCHEDULEDINSTALMENTSREAD OpenbankingBrasilConsentV2Permission = "UNARRANGED_ACCOUNTS_OVERDRAFT_SCHEDULED_INSTALMENTS_READ"

	// OpenbankingBrasilConsentV2PermissionUNARRANGEDACCOUNTSOVERDRAFTPAYMENTSREAD captures enum value "UNARRANGED_ACCOUNTS_OVERDRAFT_PAYMENTS_READ"
	OpenbankingBrasilConsentV2PermissionUNARRANGEDACCOUNTSOVERDRAFTPAYMENTSREAD OpenbankingBrasilConsentV2Permission = "UNARRANGED_ACCOUNTS_OVERDRAFT_PAYMENTS_READ"

	// OpenbankingBrasilConsentV2PermissionUNARRANGEDACCOUNTSOVERDRAFTWARRANTIESREAD captures enum value "UNARRANGED_ACCOUNTS_OVERDRAFT_WARRANTIES_READ"
	OpenbankingBrasilConsentV2PermissionUNARRANGEDACCOUNTSOVERDRAFTWARRANTIESREAD OpenbankingBrasilConsentV2Permission = "UNARRANGED_ACCOUNTS_OVERDRAFT_WARRANTIES_READ"

	// OpenbankingBrasilConsentV2PermissionRESOURCESREAD captures enum value "RESOURCES_READ"
	OpenbankingBrasilConsentV2PermissionRESOURCESREAD OpenbankingBrasilConsentV2Permission = "RESOURCES_READ"
)

// for schema
var openbankingBrasilConsentV2PermissionEnum []interface{}

func init() {
	var res []OpenbankingBrasilConsentV2Permission
	if err := json.Unmarshal([]byte(`["ACCOUNTS_READ","ACCOUNTS_BALANCES_READ","ACCOUNTS_TRANSACTIONS_READ","ACCOUNTS_OVERDRAFT_LIMITS_READ","CREDIT_CARDS_ACCOUNTS_READ","CREDIT_CARDS_ACCOUNTS_BILLS_READ","CREDIT_CARDS_ACCOUNTS_BILLS_TRANSACTIONS_READ","CREDIT_CARDS_ACCOUNTS_LIMITS_READ","CREDIT_CARDS_ACCOUNTS_TRANSACTIONS_READ","CUSTOMERS_PERSONAL_IDENTIFICATIONS_READ","CUSTOMERS_PERSONAL_ADITTIONALINFO_READ","CUSTOMERS_BUSINESS_IDENTIFICATIONS_READ","CUSTOMERS_BUSINESS_ADITTIONALINFO_READ","FINANCINGS_READ","FINANCINGS_SCHEDULED_INSTALMENTS_READ","FINANCINGS_PAYMENTS_READ","FINANCINGS_WARRANTIES_READ","INVOICE_FINANCINGS_READ","INVOICE_FINANCINGS_SCHEDULED_INSTALMENTS_READ","INVOICE_FINANCINGS_PAYMENTS_READ","INVOICE_FINANCINGS_WARRANTIES_READ","LOANS_READ","LOANS_SCHEDULED_INSTALMENTS_READ","LOANS_PAYMENTS_READ","LOANS_WARRANTIES_READ","UNARRANGED_ACCOUNTS_OVERDRAFT_READ","UNARRANGED_ACCOUNTS_OVERDRAFT_SCHEDULED_INSTALMENTS_READ","UNARRANGED_ACCOUNTS_OVERDRAFT_PAYMENTS_READ","UNARRANGED_ACCOUNTS_OVERDRAFT_WARRANTIES_READ","RESOURCES_READ"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openbankingBrasilConsentV2PermissionEnum = append(openbankingBrasilConsentV2PermissionEnum, v)
	}
}

func (m OpenbankingBrasilConsentV2Permission) validateOpenbankingBrasilConsentV2PermissionEnum(path, location string, value OpenbankingBrasilConsentV2Permission) error {
	if err := validate.EnumCase(path, location, value, openbankingBrasilConsentV2PermissionEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this openbanking brasil consent v2 permission
func (m OpenbankingBrasilConsentV2Permission) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOpenbankingBrasilConsentV2PermissionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this openbanking brasil consent v2 permission based on context it is used
func (m OpenbankingBrasilConsentV2Permission) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
