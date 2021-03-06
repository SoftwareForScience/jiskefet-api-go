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

// LinkLogToTagDto link log to tag dto
// swagger:model LinkLogToTagDto
type LinkLogToTagDto struct {

	// The id of the log to link to the tag.
	// Required: true
	LogID *int64 `json:"logId"`
}

// Validate validates this link log to tag dto
func (m *LinkLogToTagDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LinkLogToTagDto) validateLogID(formats strfmt.Registry) error {

	if err := validate.Required("logId", "body", m.LogID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LinkLogToTagDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinkLogToTagDto) UnmarshalBinary(b []byte) error {
	var res LinkLogToTagDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
