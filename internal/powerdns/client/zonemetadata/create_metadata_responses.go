// Code generated by go-swagger; DO NOT EDIT.

package zonemetadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateMetadataReader is a Reader for the CreateMetadata structure.
type CreateMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCreateMetadataNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateMetadataNoContent creates a CreateMetadataNoContent with default headers values
func NewCreateMetadataNoContent() *CreateMetadataNoContent {
	return &CreateMetadataNoContent{}
}

/* CreateMetadataNoContent describes a response with status code 204, with default header values.

OK
*/
type CreateMetadataNoContent struct {
}

func (o *CreateMetadataNoContent) Error() string {
	return fmt.Sprintf("[POST /servers/{server_id}/zones/{zone_id}/metadata][%d] createMetadataNoContent ", 204)
}

func (o *CreateMetadataNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
