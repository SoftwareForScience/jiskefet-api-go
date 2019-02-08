// Code generated by go-swagger; DO NOT EDIT.

package attachments

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

	models "github.com/SoftwareForScience/jiskefet-api-go/models"
)

// NewPostAttachmentsParams creates a new PostAttachmentsParams object
// with the default values initialized.
func NewPostAttachmentsParams() *PostAttachmentsParams {
	var ()
	return &PostAttachmentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAttachmentsParamsWithTimeout creates a new PostAttachmentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAttachmentsParamsWithTimeout(timeout time.Duration) *PostAttachmentsParams {
	var ()
	return &PostAttachmentsParams{

		timeout: timeout,
	}
}

// NewPostAttachmentsParamsWithContext creates a new PostAttachmentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAttachmentsParamsWithContext(ctx context.Context) *PostAttachmentsParams {
	var ()
	return &PostAttachmentsParams{

		Context: ctx,
	}
}

// NewPostAttachmentsParamsWithHTTPClient creates a new PostAttachmentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAttachmentsParamsWithHTTPClient(client *http.Client) *PostAttachmentsParams {
	var ()
	return &PostAttachmentsParams{
		HTTPClient: client,
	}
}

/*PostAttachmentsParams contains all the parameters to send to the API endpoint
for the post attachments operation typically these are written to a http.Request
*/
type PostAttachmentsParams struct {

	/*CreateAttachmentDto*/
	CreateAttachmentDto *models.CreateAttachmentDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post attachments params
func (o *PostAttachmentsParams) WithTimeout(timeout time.Duration) *PostAttachmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post attachments params
func (o *PostAttachmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post attachments params
func (o *PostAttachmentsParams) WithContext(ctx context.Context) *PostAttachmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post attachments params
func (o *PostAttachmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post attachments params
func (o *PostAttachmentsParams) WithHTTPClient(client *http.Client) *PostAttachmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post attachments params
func (o *PostAttachmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateAttachmentDto adds the createAttachmentDto to the post attachments params
func (o *PostAttachmentsParams) WithCreateAttachmentDto(createAttachmentDto *models.CreateAttachmentDto) *PostAttachmentsParams {
	o.SetCreateAttachmentDto(createAttachmentDto)
	return o
}

// SetCreateAttachmentDto adds the createAttachmentDto to the post attachments params
func (o *PostAttachmentsParams) SetCreateAttachmentDto(createAttachmentDto *models.CreateAttachmentDto) {
	o.CreateAttachmentDto = createAttachmentDto
}

// WriteToRequest writes these params to a swagger request
func (o *PostAttachmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateAttachmentDto != nil {
		if err := r.SetBodyParam(o.CreateAttachmentDto); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
