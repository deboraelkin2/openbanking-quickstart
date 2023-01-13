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

// SynchronizableArrayLinksentity SynchronizableArrayLinksentity
//
// Resource URLs for retrieving changes, next or previous datasets
//
// swagger:model SynchronizableArrayLinksentity
type SynchronizableArrayLinksentity struct {

	// next
	Next *HATEOASLink1 `json:"next,omitempty"`

	// prev
	Prev *HATEOASLink2 `json:"prev,omitempty"`

	// updates
	Updates *HATEOASLink3 `json:"updates,omitempty"`
}

// Validate validates this synchronizable array linksentity
func (m *SynchronizableArrayLinksentity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrev(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SynchronizableArrayLinksentity) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("next")
			}
			return err
		}
	}

	return nil
}

func (m *SynchronizableArrayLinksentity) validatePrev(formats strfmt.Registry) error {
	if swag.IsZero(m.Prev) { // not required
		return nil
	}

	if m.Prev != nil {
		if err := m.Prev.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("prev")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("prev")
			}
			return err
		}
	}

	return nil
}

func (m *SynchronizableArrayLinksentity) validateUpdates(formats strfmt.Registry) error {
	if swag.IsZero(m.Updates) { // not required
		return nil
	}

	if m.Updates != nil {
		if err := m.Updates.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updates")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this synchronizable array linksentity based on the context it is used
func (m *SynchronizableArrayLinksentity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrev(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SynchronizableArrayLinksentity) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

	if m.Next != nil {
		if err := m.Next.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("next")
			}
			return err
		}
	}

	return nil
}

func (m *SynchronizableArrayLinksentity) contextValidatePrev(ctx context.Context, formats strfmt.Registry) error {

	if m.Prev != nil {
		if err := m.Prev.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("prev")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("prev")
			}
			return err
		}
	}

	return nil
}

func (m *SynchronizableArrayLinksentity) contextValidateUpdates(ctx context.Context, formats strfmt.Registry) error {

	if m.Updates != nil {
		if err := m.Updates.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updates")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SynchronizableArrayLinksentity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SynchronizableArrayLinksentity) UnmarshalBinary(b []byte) error {
	var res SynchronizableArrayLinksentity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
