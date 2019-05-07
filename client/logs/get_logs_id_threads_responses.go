// Code generated by go-swagger; DO NOT EDIT.

package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetLogsIDThreadsReader is a Reader for the GetLogsIDThreads structure.
type GetLogsIDThreadsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogsIDThreadsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLogsIDThreadsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLogsIDThreadsOK creates a GetLogsIDThreadsOK with default headers values
func NewGetLogsIDThreadsOK() *GetLogsIDThreadsOK {
	return &GetLogsIDThreadsOK{}
}

/*GetLogsIDThreadsOK handles this case with default header values.

GetLogsIDThreadsOK get logs Id threads o k
*/
type GetLogsIDThreadsOK struct {
}

func (o *GetLogsIDThreadsOK) Error() string {
	return fmt.Sprintf("[GET /logs/{id}/threads][%d] getLogsIdThreadsOK ", 200)
}

func (o *GetLogsIDThreadsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
