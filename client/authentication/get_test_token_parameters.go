// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetTestTokenParams creates a new GetTestTokenParams object
// with the default values initialized.
func NewGetTestTokenParams() *GetTestTokenParams {
	var ()
	return &GetTestTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTestTokenParamsWithTimeout creates a new GetTestTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTestTokenParamsWithTimeout(timeout time.Duration) *GetTestTokenParams {
	var ()
	return &GetTestTokenParams{

		timeout: timeout,
	}
}

// NewGetTestTokenParamsWithContext creates a new GetTestTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTestTokenParamsWithContext(ctx context.Context) *GetTestTokenParams {
	var ()
	return &GetTestTokenParams{

		Context: ctx,
	}
}

// NewGetTestTokenParamsWithHTTPClient creates a new GetTestTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTestTokenParamsWithHTTPClient(client *http.Client) *GetTestTokenParams {
	var ()
	return &GetTestTokenParams{
		HTTPClient: client,
	}
}

/*GetTestTokenParams contains all the parameters to send to the API endpoint
for the get test token operation typically these are written to a http.Request
*/
type GetTestTokenParams struct {

	/*HashedSecret*/
	HashedSecret *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get test token params
func (o *GetTestTokenParams) WithTimeout(timeout time.Duration) *GetTestTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get test token params
func (o *GetTestTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get test token params
func (o *GetTestTokenParams) WithContext(ctx context.Context) *GetTestTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get test token params
func (o *GetTestTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get test token params
func (o *GetTestTokenParams) WithHTTPClient(client *http.Client) *GetTestTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get test token params
func (o *GetTestTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHashedSecret adds the hashedSecret to the get test token params
func (o *GetTestTokenParams) WithHashedSecret(hashedSecret *string) *GetTestTokenParams {
	o.SetHashedSecret(hashedSecret)
	return o
}

// SetHashedSecret adds the hashedSecret to the get test token params
func (o *GetTestTokenParams) SetHashedSecret(hashedSecret *string) {
	o.HashedSecret = hashedSecret
}

// WriteToRequest writes these params to a swagger request
func (o *GetTestTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.HashedSecret != nil {

		// query param hashedSecret
		var qrHashedSecret string
		if o.HashedSecret != nil {
			qrHashedSecret = *o.HashedSecret
		}
		qHashedSecret := qrHashedSecret
		if qHashedSecret != "" {
			if err := r.SetQueryParam("hashedSecret", qHashedSecret); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
