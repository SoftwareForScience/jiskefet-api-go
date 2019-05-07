// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostUsersIDTokensReader is a Reader for the PostUsersIDTokens structure.
type PostUsersIDTokensReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUsersIDTokensReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostUsersIDTokensOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostUsersIDTokensOK creates a PostUsersIDTokensOK with default headers values
func NewPostUsersIDTokensOK() *PostUsersIDTokensOK {
	return &PostUsersIDTokensOK{}
}

/*PostUsersIDTokensOK handles this case with default header values.

Succesfully created a Token.
*/
type PostUsersIDTokensOK struct {
}

func (o *PostUsersIDTokensOK) Error() string {
	return fmt.Sprintf("[POST /users/{id}/tokens][%d] postUsersIdTokensOK ", 200)
}

func (o *PostUsersIDTokensOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}