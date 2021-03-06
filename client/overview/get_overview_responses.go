// Code generated by go-swagger; DO NOT EDIT.

package overview

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetOverviewReader is a Reader for the GetOverview structure.
type GetOverviewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOverviewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOverviewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetOverviewNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOverviewOK creates a GetOverviewOK with default headers values
func NewGetOverviewOK() *GetOverviewOK {
	return &GetOverviewOK{}
}

/*GetOverviewOK handles this case with default header values.

Succesfully returned the overview.
*/
type GetOverviewOK struct {
}

func (o *GetOverviewOK) Error() string {
	return fmt.Sprintf("[GET /overview][%d] getOverviewOK ", 200)
}

func (o *GetOverviewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOverviewNotFound creates a GetOverviewNotFound with default headers values
func NewGetOverviewNotFound() *GetOverviewNotFound {
	return &GetOverviewNotFound{}
}

/*GetOverviewNotFound handles this case with default header values.

Unable to find an overview with given input
*/
type GetOverviewNotFound struct {
}

func (o *GetOverviewNotFound) Error() string {
	return fmt.Sprintf("[GET /overview][%d] getOverviewNotFound ", 404)
}

func (o *GetOverviewNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
