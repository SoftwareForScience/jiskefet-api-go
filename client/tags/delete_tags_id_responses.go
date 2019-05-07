// Code generated by go-swagger; DO NOT EDIT.

package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteTagsIDReader is a Reader for the DeleteTagsID structure.
type DeleteTagsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTagsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteTagsIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteTagsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteTagsIDNoContent creates a DeleteTagsIDNoContent with default headers values
func NewDeleteTagsIDNoContent() *DeleteTagsIDNoContent {
	return &DeleteTagsIDNoContent{}
}

/*DeleteTagsIDNoContent handles this case with default header values.

Succesfully deleted a Tag.
*/
type DeleteTagsIDNoContent struct {
}

func (o *DeleteTagsIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /tags/{id}][%d] deleteTagsIdNoContent ", 204)
}

func (o *DeleteTagsIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTagsIDNotFound creates a DeleteTagsIDNotFound with default headers values
func NewDeleteTagsIDNotFound() *DeleteTagsIDNotFound {
	return &DeleteTagsIDNotFound{}
}

/*DeleteTagsIDNotFound handles this case with default header values.

The Tag with the given ID does not exist.
*/
type DeleteTagsIDNotFound struct {
}

func (o *DeleteTagsIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /tags/{id}][%d] deleteTagsIdNotFound ", 404)
}

func (o *DeleteTagsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}