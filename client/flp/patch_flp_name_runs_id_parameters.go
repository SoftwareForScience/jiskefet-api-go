// Code generated by go-swagger; DO NOT EDIT.

package flp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/SoftwareForScience/jiskefet-api-go/models"
)

// NewPatchFlpNameRunsIDParams creates a new PatchFlpNameRunsIDParams object
// with the default values initialized.
func NewPatchFlpNameRunsIDParams() *PatchFlpNameRunsIDParams {
	var ()
	return &PatchFlpNameRunsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchFlpNameRunsIDParamsWithTimeout creates a new PatchFlpNameRunsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchFlpNameRunsIDParamsWithTimeout(timeout time.Duration) *PatchFlpNameRunsIDParams {
	var ()
	return &PatchFlpNameRunsIDParams{

		timeout: timeout,
	}
}

// NewPatchFlpNameRunsIDParamsWithContext creates a new PatchFlpNameRunsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchFlpNameRunsIDParamsWithContext(ctx context.Context) *PatchFlpNameRunsIDParams {
	var ()
	return &PatchFlpNameRunsIDParams{

		Context: ctx,
	}
}

// NewPatchFlpNameRunsIDParamsWithHTTPClient creates a new PatchFlpNameRunsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchFlpNameRunsIDParamsWithHTTPClient(client *http.Client) *PatchFlpNameRunsIDParams {
	var ()
	return &PatchFlpNameRunsIDParams{
		HTTPClient: client,
	}
}

/*PatchFlpNameRunsIDParams contains all the parameters to send to the API endpoint
for the patch flp name runs ID operation typically these are written to a http.Request
*/
type PatchFlpNameRunsIDParams struct {

	/*PatchFlpDto*/
	PatchFlpDto *models.PatchFlpDto
	/*ID*/
	ID int64
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch flp name runs ID params
func (o *PatchFlpNameRunsIDParams) WithTimeout(timeout time.Duration) *PatchFlpNameRunsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch flp name runs ID params
func (o *PatchFlpNameRunsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch flp name runs ID params
func (o *PatchFlpNameRunsIDParams) WithContext(ctx context.Context) *PatchFlpNameRunsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch flp name runs ID params
func (o *PatchFlpNameRunsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch flp name runs ID params
func (o *PatchFlpNameRunsIDParams) WithHTTPClient(client *http.Client) *PatchFlpNameRunsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch flp name runs ID params
func (o *PatchFlpNameRunsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPatchFlpDto adds the patchFlpDto to the patch flp name runs ID params
func (o *PatchFlpNameRunsIDParams) WithPatchFlpDto(patchFlpDto *models.PatchFlpDto) *PatchFlpNameRunsIDParams {
	o.SetPatchFlpDto(patchFlpDto)
	return o
}

// SetPatchFlpDto adds the patchFlpDto to the patch flp name runs ID params
func (o *PatchFlpNameRunsIDParams) SetPatchFlpDto(patchFlpDto *models.PatchFlpDto) {
	o.PatchFlpDto = patchFlpDto
}

// WithID adds the id to the patch flp name runs ID params
func (o *PatchFlpNameRunsIDParams) WithID(id int64) *PatchFlpNameRunsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch flp name runs ID params
func (o *PatchFlpNameRunsIDParams) SetID(id int64) {
	o.ID = id
}

// WithName adds the name to the patch flp name runs ID params
func (o *PatchFlpNameRunsIDParams) WithName(name string) *PatchFlpNameRunsIDParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the patch flp name runs ID params
func (o *PatchFlpNameRunsIDParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PatchFlpNameRunsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PatchFlpDto != nil {
		if err := r.SetBodyParam(o.PatchFlpDto); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
