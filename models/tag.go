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

// Tag tag
// swagger:model Tag
type Tag struct {

	// logs
	// Required: true
	Logs []string `json:"logs"`

	// runs
	// Required: true
	Runs []string `json:"runs"`

	// tag Id
	// Required: true
	TagID *int64 `json:"tagId"`

	// tag text
	// Required: true
	TagText *string `json:"tagText"`
}

// Validate validates this tag
func (m *Tag) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTagID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTagText(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Tag) validateLogs(formats strfmt.Registry) error {

	if err := validate.Required("logs", "body", m.Logs); err != nil {
		return err
	}

	return nil
}

func (m *Tag) validateRuns(formats strfmt.Registry) error {

	if err := validate.Required("runs", "body", m.Runs); err != nil {
		return err
	}

	return nil
}

func (m *Tag) validateTagID(formats strfmt.Registry) error {

	if err := validate.Required("tagId", "body", m.TagID); err != nil {
		return err
	}

	return nil
}

func (m *Tag) validateTagText(formats strfmt.Registry) error {

	if err := validate.Required("tagText", "body", m.TagText); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Tag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Tag) UnmarshalBinary(b []byte) error {
	var res Tag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
