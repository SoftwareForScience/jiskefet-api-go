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

// CreateTagDto create tag dto
// swagger:model CreateTagDto
type CreateTagDto struct {

	// The text of the tag
	// Required: true
	TagText *string `json:"tagText"`
}

// Validate validates this create tag dto
func (m *CreateTagDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTagText(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateTagDto) validateTagText(formats strfmt.Registry) error {

	if err := validate.Required("tagText", "body", m.TagText); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateTagDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateTagDto) UnmarshalBinary(b []byte) error {
	var res CreateTagDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}