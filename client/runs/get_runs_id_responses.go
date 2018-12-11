// Code generated by go-swagger; DO NOT EDIT.

package runs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetRunsIDReader is a Reader for the GetRunsID structure.
type GetRunsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRunsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRunsIDOK creates a GetRunsIDOK with default headers values
func NewGetRunsIDOK() *GetRunsIDOK {
	return &GetRunsIDOK{}
}

/*GetRunsIDOK handles this case with default header values.

GetRunsIDOK get runs Id o k
*/
type GetRunsIDOK struct {
}

func (o *GetRunsIDOK) Error() string {
	return fmt.Sprintf("[GET /runs/{id}][%d] getRunsIdOK ", 200)
}

func (o *GetRunsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}