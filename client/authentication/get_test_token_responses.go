// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetTestTokenReader is a Reader for the GetTestToken structure.
type GetTestTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTestTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTestTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetTestTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetTestTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTestTokenOK creates a GetTestTokenOK with default headers values
func NewGetTestTokenOK() *GetTestTokenOK {
	return &GetTestTokenOK{}
}

/*GetTestTokenOK handles this case with default header values.

The hashed secret given matches the secret in the environment and a JWT is returned.
*/
type GetTestTokenOK struct {
}

func (o *GetTestTokenOK) Error() string {
	return fmt.Sprintf("[GET /test-token][%d] getTestTokenOK ", 200)
}

func (o *GetTestTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTestTokenBadRequest creates a GetTestTokenBadRequest with default headers values
func NewGetTestTokenBadRequest() *GetTestTokenBadRequest {
	return &GetTestTokenBadRequest{}
}

/*GetTestTokenBadRequest handles this case with default header values.

The required query parameter 'hashedSecret' is missing.
*/
type GetTestTokenBadRequest struct {
}

func (o *GetTestTokenBadRequest) Error() string {
	return fmt.Sprintf("[GET /test-token][%d] getTestTokenBadRequest ", 400)
}

func (o *GetTestTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTestTokenUnauthorized creates a GetTestTokenUnauthorized with default headers values
func NewGetTestTokenUnauthorized() *GetTestTokenUnauthorized {
	return &GetTestTokenUnauthorized{}
}

/*GetTestTokenUnauthorized handles this case with default header values.

Hashed secret was not accepted
*/
type GetTestTokenUnauthorized struct {
}

func (o *GetTestTokenUnauthorized) Error() string {
	return fmt.Sprintf("[GET /test-token][%d] getTestTokenUnauthorized ", 401)
}

func (o *GetTestTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}