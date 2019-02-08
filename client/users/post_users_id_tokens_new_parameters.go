// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewPostUsersIDTokensNewParams creates a new PostUsersIDTokensNewParams object
// with the default values initialized.
func NewPostUsersIDTokensNewParams() *PostUsersIDTokensNewParams {
	var ()
	return &PostUsersIDTokensNewParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostUsersIDTokensNewParamsWithTimeout creates a new PostUsersIDTokensNewParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostUsersIDTokensNewParamsWithTimeout(timeout time.Duration) *PostUsersIDTokensNewParams {
	var ()
	return &PostUsersIDTokensNewParams{

		timeout: timeout,
	}
}

// NewPostUsersIDTokensNewParamsWithContext creates a new PostUsersIDTokensNewParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostUsersIDTokensNewParamsWithContext(ctx context.Context) *PostUsersIDTokensNewParams {
	var ()
	return &PostUsersIDTokensNewParams{

		Context: ctx,
	}
}

// NewPostUsersIDTokensNewParamsWithHTTPClient creates a new PostUsersIDTokensNewParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostUsersIDTokensNewParamsWithHTTPClient(client *http.Client) *PostUsersIDTokensNewParams {
	var ()
	return &PostUsersIDTokensNewParams{
		HTTPClient: client,
	}
}

/*PostUsersIDTokensNewParams contains all the parameters to send to the API endpoint
for the post users ID tokens new operation typically these are written to a http.Request
*/
type PostUsersIDTokensNewParams struct {

	/*CreateSubSystemPermissionDto*/
	CreateSubSystemPermissionDto *models.CreateSubSystemPermissionDto
	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post users ID tokens new params
func (o *PostUsersIDTokensNewParams) WithTimeout(timeout time.Duration) *PostUsersIDTokensNewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post users ID tokens new params
func (o *PostUsersIDTokensNewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post users ID tokens new params
func (o *PostUsersIDTokensNewParams) WithContext(ctx context.Context) *PostUsersIDTokensNewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post users ID tokens new params
func (o *PostUsersIDTokensNewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post users ID tokens new params
func (o *PostUsersIDTokensNewParams) WithHTTPClient(client *http.Client) *PostUsersIDTokensNewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post users ID tokens new params
func (o *PostUsersIDTokensNewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateSubSystemPermissionDto adds the createSubSystemPermissionDto to the post users ID tokens new params
func (o *PostUsersIDTokensNewParams) WithCreateSubSystemPermissionDto(createSubSystemPermissionDto *models.CreateSubSystemPermissionDto) *PostUsersIDTokensNewParams {
	o.SetCreateSubSystemPermissionDto(createSubSystemPermissionDto)
	return o
}

// SetCreateSubSystemPermissionDto adds the createSubSystemPermissionDto to the post users ID tokens new params
func (o *PostUsersIDTokensNewParams) SetCreateSubSystemPermissionDto(createSubSystemPermissionDto *models.CreateSubSystemPermissionDto) {
	o.CreateSubSystemPermissionDto = createSubSystemPermissionDto
}

// WithID adds the id to the post users ID tokens new params
func (o *PostUsersIDTokensNewParams) WithID(id int64) *PostUsersIDTokensNewParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post users ID tokens new params
func (o *PostUsersIDTokensNewParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostUsersIDTokensNewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateSubSystemPermissionDto != nil {
		if err := r.SetBodyParam(o.CreateSubSystemPermissionDto); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
