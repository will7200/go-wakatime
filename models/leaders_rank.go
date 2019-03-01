// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LeadersRank leaders rank
// swagger:model leaders-rank
type LeadersRank struct {

	// rank
	Rank int64 `json:"rank,omitempty"`

	// running total
	RunningTotal *LeadersRankRunningTotal `json:"running_total,omitempty"`

	// user
	User *LeadersRankUser `json:"user,omitempty"`
}

// Validate validates this leaders rank
func (m *LeadersRank) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRunningTotal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LeadersRank) validateRunningTotal(formats strfmt.Registry) error {

	if swag.IsZero(m.RunningTotal) { // not required
		return nil
	}

	if m.RunningTotal != nil {
		if err := m.RunningTotal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("running_total")
			}
			return err
		}
	}

	return nil
}

func (m *LeadersRank) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LeadersRank) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LeadersRank) UnmarshalBinary(b []byte) error {
	var res LeadersRank
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LeadersRankRunningTotal leaders rank running total
// swagger:model LeadersRankRunningTotal
type LeadersRankRunningTotal struct {

	// daily average
	DailyAverage int64 `json:"daily_average,omitempty"`

	// human readable daily average
	HumanReadableDailyAverage string `json:"human_readable_daily_average,omitempty"`

	// human readable total
	HumanReadableTotal string `json:"human_readable_total,omitempty"`

	// languages
	Languages []*LeadersRankRunningTotalLanguagesItems0 `json:"languages"`

	// modified at
	ModifiedAt string `json:"modified_at,omitempty"`

	// total seconds
	TotalSeconds float64 `json:"total_seconds,omitempty"`
}

// Validate validates this leaders rank running total
func (m *LeadersRankRunningTotal) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLanguages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LeadersRankRunningTotal) validateLanguages(formats strfmt.Registry) error {

	if swag.IsZero(m.Languages) { // not required
		return nil
	}

	for i := 0; i < len(m.Languages); i++ {
		if swag.IsZero(m.Languages[i]) { // not required
			continue
		}

		if m.Languages[i] != nil {
			if err := m.Languages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("running_total" + "." + "languages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LeadersRankRunningTotal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LeadersRankRunningTotal) UnmarshalBinary(b []byte) error {
	var res LeadersRankRunningTotal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LeadersRankRunningTotalLanguagesItems0 leaders rank running total languages items0
// swagger:model LeadersRankRunningTotalLanguagesItems0
type LeadersRankRunningTotalLanguagesItems0 struct {

	// name
	Name string `json:"name,omitempty"`

	// total seconds
	TotalSeconds float64 `json:"total_seconds,omitempty"`
}

// Validate validates this leaders rank running total languages items0
func (m *LeadersRankRunningTotalLanguagesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LeadersRankRunningTotalLanguagesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LeadersRankRunningTotalLanguagesItems0) UnmarshalBinary(b []byte) error {
	var res LeadersRankRunningTotalLanguagesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LeadersRankUser leaders rank user
// swagger:model LeadersRankUser
type LeadersRankUser struct {

	// display name
	DisplayName string `json:"display_name,omitempty"`

	// email
	Email *string `json:"email,omitempty"`

	// email public
	EmailPublic bool `json:"email_public,omitempty"`

	// full name
	FullName string `json:"full_name,omitempty"`

	// human readable website
	HumanReadableWebsite *string `json:"human_readable_website,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// location
	Location string `json:"location,omitempty"`

	// photo
	Photo string `json:"photo,omitempty"`

	// photo public
	PhotoPublic bool `json:"photo_public,omitempty"`

	// username
	Username *string `json:"username,omitempty"`

	// website
	Website *string `json:"website,omitempty"`
}

// Validate validates this leaders rank user
func (m *LeadersRankUser) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LeadersRankUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LeadersRankUser) UnmarshalBinary(b []byte) error {
	var res LeadersRankUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
