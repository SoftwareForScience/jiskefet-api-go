// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateLogDto create log dto
// swagger:model CreateLogDto
type CreateLogDto struct {

	// Attachments of this log
	// Required: true
	Attachments []string `json:"attachments"`

	// Where did the log come from?
	// Required: true
	// Enum: [human process]
	Origin *string `json:"origin"`

	// Attached run numbers of this log
	// Required: true
	Runs []string `json:"runs"`

	// What kind of log is it?
	// Required: true
	// Enum: [run subsystem announcement intervention comment]
	Subtype *string `json:"subtype"`

	// describes the log in depth
	// Required: true
	Text *string `json:"text"`

	// describes the log in short
	// Required: true
	Title *string `json:"title"`
}

// Validate validates this create log dto
func (m *CreateLogDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrigin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubtype(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateLogDto) validateAttachments(formats strfmt.Registry) error {

	if err := validate.Required("attachments", "body", m.Attachments); err != nil {
		return err
	}

	return nil
}

var createLogDtoTypeOriginPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["human","process"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createLogDtoTypeOriginPropEnum = append(createLogDtoTypeOriginPropEnum, v)
	}
}

const (

	// CreateLogDtoOriginHuman captures enum value "human"
	CreateLogDtoOriginHuman string = "human"

	// CreateLogDtoOriginProcess captures enum value "process"
	CreateLogDtoOriginProcess string = "process"
)

// prop value enum
func (m *CreateLogDto) validateOriginEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, createLogDtoTypeOriginPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CreateLogDto) validateOrigin(formats strfmt.Registry) error {

	if err := validate.Required("origin", "body", m.Origin); err != nil {
		return err
	}

	// value enum
	if err := m.validateOriginEnum("origin", "body", *m.Origin); err != nil {
		return err
	}

	return nil
}

func (m *CreateLogDto) validateRuns(formats strfmt.Registry) error {

	if err := validate.Required("runs", "body", m.Runs); err != nil {
		return err
	}

	return nil
}

var createLogDtoTypeSubtypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["run","subsystem","announcement","intervention","comment"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createLogDtoTypeSubtypePropEnum = append(createLogDtoTypeSubtypePropEnum, v)
	}
}

const (

	// CreateLogDtoSubtypeRun captures enum value "run"
	CreateLogDtoSubtypeRun string = "run"

	// CreateLogDtoSubtypeSubsystem captures enum value "subsystem"
	CreateLogDtoSubtypeSubsystem string = "subsystem"

	// CreateLogDtoSubtypeAnnouncement captures enum value "announcement"
	CreateLogDtoSubtypeAnnouncement string = "announcement"

	// CreateLogDtoSubtypeIntervention captures enum value "intervention"
	CreateLogDtoSubtypeIntervention string = "intervention"

	// CreateLogDtoSubtypeComment captures enum value "comment"
	CreateLogDtoSubtypeComment string = "comment"
)

// prop value enum
func (m *CreateLogDto) validateSubtypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, createLogDtoTypeSubtypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CreateLogDto) validateSubtype(formats strfmt.Registry) error {

	if err := validate.Required("subtype", "body", m.Subtype); err != nil {
		return err
	}

	// value enum
	if err := m.validateSubtypeEnum("subtype", "body", *m.Subtype); err != nil {
		return err
	}

	return nil
}

func (m *CreateLogDto) validateText(formats strfmt.Registry) error {

	if err := validate.Required("text", "body", m.Text); err != nil {
		return err
	}

	return nil
}

func (m *CreateLogDto) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateLogDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateLogDto) UnmarshalBinary(b []byte) error {
	var res CreateLogDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}