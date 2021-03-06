// Code generated by go-swagger; DO NOT EDIT.

package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/SoftwareForScience/jiskefet-api-go/models"
)

// PostLogsIDAttachmentsReader is a Reader for the PostLogsIDAttachments structure.
type PostLogsIDAttachmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLogsIDAttachmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostLogsIDAttachmentsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLogsIDAttachmentsCreated creates a PostLogsIDAttachmentsCreated with default headers values
func NewPostLogsIDAttachmentsCreated() *PostLogsIDAttachmentsCreated {
	return &PostLogsIDAttachmentsCreated{}
}

/*PostLogsIDAttachmentsCreated handles this case with default header values.

Succesfully created an Attachment.
*/
type PostLogsIDAttachmentsCreated struct {
	Payload *models.Attachment
}

func (o *PostLogsIDAttachmentsCreated) Error() string {
	return fmt.Sprintf("[POST /logs/{id}/attachments][%d] postLogsIdAttachmentsCreated  %+v", 201, o.Payload)
}

func (o *PostLogsIDAttachmentsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Attachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
