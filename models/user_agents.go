// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserAgents user agents
// swagger:model user_agents
type UserAgents struct {

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// editor
	// Read Only: true
	Editor string `json:"editor,omitempty"`

	// id
	// Read Only: true
	ID string `json:"id,omitempty"`

	// last seen
	// Format: date-time
	LastSeen strfmt.DateTime `json:"last_seen,omitempty"`

	// os
	// Read Only: true
	Os string `json:"os,omitempty"`

	// value
	// Read Only: true
	Value string `json:"value,omitempty"`

	// version
	// Read Only: true
	Version string `json:"version,omitempty"`
}

// Validate validates this user agents
func (m *UserAgents) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastSeen(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserAgents) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserAgents) validateLastSeen(formats strfmt.Registry) error {

	if swag.IsZero(m.LastSeen) { // not required
		return nil
	}

	if err := validate.FormatOf("last_seen", "body", "date-time", m.LastSeen.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserAgents) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserAgents) UnmarshalBinary(b []byte) error {
	var res UserAgents
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
