// Code generated by go-swagger; DO NOT EDIT.

package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PatchTagsIDRunsReader is a Reader for the PatchTagsIDRuns structure.
type PatchTagsIDRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchTagsIDRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPatchTagsIDRunsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPatchTagsIDRunsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPatchTagsIDRunsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchTagsIDRunsNoContent creates a PatchTagsIDRunsNoContent with default headers values
func NewPatchTagsIDRunsNoContent() *PatchTagsIDRunsNoContent {
	return &PatchTagsIDRunsNoContent{}
}

/*PatchTagsIDRunsNoContent handles this case with default header values.

Succesfully linked a Run to a Tag.
*/
type PatchTagsIDRunsNoContent struct {
}

func (o *PatchTagsIDRunsNoContent) Error() string {
	return fmt.Sprintf("[PATCH /tags/{id}/runs][%d] patchTagsIdRunsNoContent ", 204)
}

func (o *PatchTagsIDRunsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchTagsIDRunsNotFound creates a PatchTagsIDRunsNotFound with default headers values
func NewPatchTagsIDRunsNotFound() *PatchTagsIDRunsNotFound {
	return &PatchTagsIDRunsNotFound{}
}

/*PatchTagsIDRunsNotFound handles this case with default header values.

The Run or Tag does not exist.
*/
type PatchTagsIDRunsNotFound struct {
}

func (o *PatchTagsIDRunsNotFound) Error() string {
	return fmt.Sprintf("[PATCH /tags/{id}/runs][%d] patchTagsIdRunsNotFound ", 404)
}

func (o *PatchTagsIDRunsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchTagsIDRunsConflict creates a PatchTagsIDRunsConflict with default headers values
func NewPatchTagsIDRunsConflict() *PatchTagsIDRunsConflict {
	return &PatchTagsIDRunsConflict{}
}

/*PatchTagsIDRunsConflict handles this case with default header values.

The Run is already linked to the Tag.
*/
type PatchTagsIDRunsConflict struct {
}

func (o *PatchTagsIDRunsConflict) Error() string {
	return fmt.Sprintf("[PATCH /tags/{id}/runs][%d] patchTagsIdRunsConflict ", 409)
}

func (o *PatchTagsIDRunsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}