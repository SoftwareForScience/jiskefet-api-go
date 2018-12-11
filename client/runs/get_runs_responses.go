// Code generated by go-swagger; DO NOT EDIT.

package runs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetRunsReader is a Reader for the GetRuns structure.
type GetRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRunsOK creates a GetRunsOK with default headers values
func NewGetRunsOK() *GetRunsOK {
	return &GetRunsOK{}
}

/*GetRunsOK handles this case with default header values.

OK
*/
type GetRunsOK struct {
	Payload string
}

func (o *GetRunsOK) Error() string {
	return fmt.Sprintf("[GET /runs][%d] getRunsOK  %+v", 200, o.Payload)
}

func (o *GetRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
