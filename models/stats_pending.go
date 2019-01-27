// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StatsPending stats pending
// swagger:model stats-pending
type StatsPending struct {

	// data
	Data *StatsPendingData `json:"data,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this stats pending
func (m *StatsPending) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatsPending) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StatsPending) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatsPending) UnmarshalBinary(b []byte) error {
	var res StatsPending
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// StatsPendingData stats pending data
// swagger:model StatsPendingData
type StatsPendingData struct {

	// id
	ID string `json:"id,omitempty"`

	// is already updating
	IsAlreadyUpdating bool `json:"is_already_updating,omitempty"`

	// is coding activity visible
	IsCodingActivityVisible bool `json:"is_coding_activity_visible,omitempty"`

	// is other usage visible
	IsOtherUsageVisible bool `json:"is_other_usage_visible,omitempty"`

	// is stuck
	IsStuck bool `json:"is_stuck,omitempty"`

	// is up to date
	IsUpToDate bool `json:"is_up_to_date,omitempty"`

	// range
	Range string `json:"range,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// timeout
	Timeout int64 `json:"timeout,omitempty"`

	// user id
	UserID string `json:"user_id,omitempty"`

	// username
	Username string `json:"username,omitempty"`

	// writes only
	WritesOnly bool `json:"writes_only,omitempty"`
}

// Validate validates this stats pending data
func (m *StatsPendingData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StatsPendingData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatsPendingData) UnmarshalBinary(b []byte) error {
	var res StatsPendingData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}