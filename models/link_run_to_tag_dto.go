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

// LinkRunToTagDto link run to tag dto
// swagger:model LinkRunToTagDto
type LinkRunToTagDto struct {

	// The id of the run to link to the tag.
	// Required: true
	RunNumber *int64 `json:"runNumber"`
}

// Validate validates this link run to tag dto
func (m *LinkRunToTagDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRunNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LinkRunToTagDto) validateRunNumber(formats strfmt.Registry) error {

	if err := validate.Required("runNumber", "body", m.RunNumber); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LinkRunToTagDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinkRunToTagDto) UnmarshalBinary(b []byte) error {
	var res LinkRunToTagDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
