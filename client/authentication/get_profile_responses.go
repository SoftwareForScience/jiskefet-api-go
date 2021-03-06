// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetProfileReader is a Reader for the GetProfile structure.
type GetProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetProfileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProfileOK creates a GetProfileOK with default headers values
func NewGetProfileOK() *GetProfileOK {
	return &GetProfileOK{}
}

/*GetProfileOK handles this case with default header values.

User successfully received profile information.
*/
type GetProfileOK struct {
	Payload interface{}
}

func (o *GetProfileOK) Error() string {
	return fmt.Sprintf("[GET /profile][%d] getProfileOK  %+v", 200, o.Payload)
}

func (o *GetProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfileUnauthorized creates a GetProfileUnauthorized with default headers values
func NewGetProfileUnauthorized() *GetProfileUnauthorized {
	return &GetProfileUnauthorized{}
}

/*GetProfileUnauthorized handles this case with default header values.

User is unauthorized
*/
type GetProfileUnauthorized struct {
}

func (o *GetProfileUnauthorized) Error() string {
	return fmt.Sprintf("[GET /profile][%d] getProfileUnauthorized ", 401)
}

func (o *GetProfileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
