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

// User user
// swagger:model user
type User struct {

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// display name
	DisplayName string `json:"display_name,omitempty"`

	// email
	Email *string `json:"email,omitempty"`

	// email public
	EmailPublic bool `json:"email_public,omitempty"`

	// full name
	FullName *string `json:"full_name,omitempty"`

	// has premium features
	HasPremiumFeatures bool `json:"has_premium_features,omitempty"`

	// human readable website
	HumanReadableWebsite *string `json:"human_readable_website,omitempty"`

	// id
	// Read Only: true
	ID string `json:"id,omitempty"`

	// is email confirmed
	IsEmailConfirmed bool `json:"is_email_confirmed,omitempty"`

	// is hireable
	IsHireable bool `json:"is_hireable,omitempty"`

	// languages used public
	LanguagesUsedPublic bool `json:"languages_used_public,omitempty"`

	// last heartbeat
	// Format: date-time
	LastHeartbeat *strfmt.DateTime `json:"last_heartbeat,omitempty"`

	// last plugin
	LastPlugin *string `json:"last_plugin,omitempty"`

	// last plugin name
	LastPluginName *string `json:"last_plugin_name,omitempty"`

	// last project
	LastProject *string `json:"last_project,omitempty"`

	// location
	Location *string `json:"location,omitempty"`

	// logged time public
	LoggedTimePublic bool `json:"logged_time_public,omitempty"`

	// modified at
	// Format: date-time
	ModifiedAt *strfmt.DateTime `json:"modified_at,omitempty"`

	// photo
	Photo string `json:"photo,omitempty"`

	// photo public
	PhotoPublic bool `json:"photo_public,omitempty"`

	// plan
	Plan string `json:"plan,omitempty"`

	// timezone
	Timezone string `json:"timezone,omitempty"`

	// username
	Username *string `json:"username,omitempty"`

	// website
	Website *string `json:"website,omitempty"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastHeartbeat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *User) validateLastHeartbeat(formats strfmt.Registry) error {

	if swag.IsZero(m.LastHeartbeat) { // not required
		return nil
	}

	if err := validate.FormatOf("last_heartbeat", "body", "date-time", m.LastHeartbeat.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *User) validateModifiedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_at", "body", "date-time", m.ModifiedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
