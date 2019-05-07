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
)

// NewGetFlpNameRunsIDParams creates a new GetFlpNameRunsIDParams object
// with the default values initialized.
func NewGetFlpNameRunsIDParams() *GetFlpNameRunsIDParams {
	var ()
	return &GetFlpNameRunsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFlpNameRunsIDParamsWithTimeout creates a new GetFlpNameRunsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFlpNameRunsIDParamsWithTimeout(timeout time.Duration) *GetFlpNameRunsIDParams {
	var ()
	return &GetFlpNameRunsIDParams{

		timeout: timeout,
	}
}

// NewGetFlpNameRunsIDParamsWithContext creates a new GetFlpNameRunsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFlpNameRunsIDParamsWithContext(ctx context.Context) *GetFlpNameRunsIDParams {
	var ()
	return &GetFlpNameRunsIDParams{

		Context: ctx,
	}
}

// NewGetFlpNameRunsIDParamsWithHTTPClient creates a new GetFlpNameRunsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFlpNameRunsIDParamsWithHTTPClient(client *http.Client) *GetFlpNameRunsIDParams {
	var ()
	return &GetFlpNameRunsIDParams{
		HTTPClient: client,
	}
}

/*GetFlpNameRunsIDParams contains all the parameters to send to the API endpoint
for the get flp name runs ID operation typically these are written to a http.Request
*/
type GetFlpNameRunsIDParams struct {

	/*ID*/
	ID int64
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get flp name runs ID params
func (o *GetFlpNameRunsIDParams) WithTimeout(timeout time.Duration) *GetFlpNameRunsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get flp name runs ID params
func (o *GetFlpNameRunsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get flp name runs ID params
func (o *GetFlpNameRunsIDParams) WithContext(ctx context.Context) *GetFlpNameRunsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get flp name runs ID params
func (o *GetFlpNameRunsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get flp name runs ID params
func (o *GetFlpNameRunsIDParams) WithHTTPClient(client *http.Client) *GetFlpNameRunsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get flp name runs ID params
func (o *GetFlpNameRunsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get flp name runs ID params
func (o *GetFlpNameRunsIDParams) WithID(id int64) *GetFlpNameRunsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get flp name runs ID params
func (o *GetFlpNameRunsIDParams) SetID(id int64) {
	o.ID = id
}

// WithName adds the name to the get flp name runs ID params
func (o *GetFlpNameRunsIDParams) WithName(name string) *GetFlpNameRunsIDParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get flp name runs ID params
func (o *GetFlpNameRunsIDParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetFlpNameRunsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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