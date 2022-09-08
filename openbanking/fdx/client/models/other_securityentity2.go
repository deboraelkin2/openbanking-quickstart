// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OtherSecurityentity2 OtherSecurityentity2
//
// Another type of security
//
// swagger:model OtherSecurityentity2
type OtherSecurityentity2 struct {

	// Description of other security
	TypeDescription string `json:"typeDescription,omitempty"`
}

// Validate validates this other securityentity2
func (m *OtherSecurityentity2) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this other securityentity2 based on context it is used
func (m *OtherSecurityentity2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OtherSecurityentity2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OtherSecurityentity2) UnmarshalBinary(b []byte) error {
	var res OtherSecurityentity2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
