// Code generated by go-swagger; DO NOT EDIT.

package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetTagsReader is a Reader for the GetTags structure.
type GetTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetTagsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTagsOK creates a GetTagsOK with default headers values
func NewGetTagsOK() *GetTagsOK {
	return &GetTagsOK{}
}

/*GetTagsOK handles this case with default header values.

Successfully returned all Tags.
*/
type GetTagsOK struct {
}

func (o *GetTagsOK) Error() string {
	return fmt.Sprintf("[GET /tags][%d] getTagsOK ", 200)
}

func (o *GetTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTagsNotFound creates a GetTagsNotFound with default headers values
func NewGetTagsNotFound() *GetTagsNotFound {
	return &GetTagsNotFound{}
}

/*GetTagsNotFound handles this case with default header values.

There are no Tags.
*/
type GetTagsNotFound struct {
}

func (o *GetTagsNotFound) Error() string {
	return fmt.Sprintf("[GET /tags][%d] getTagsNotFound ", 404)
}

func (o *GetTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
