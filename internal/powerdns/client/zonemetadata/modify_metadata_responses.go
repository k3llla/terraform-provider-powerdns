// Code generated by go-swagger; DO NOT EDIT.

package zonemetadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/k3llla/terraform-provider-powerdns/internal/powerdns/models"
)

// ModifyMetadataReader is a Reader for the ModifyMetadata structure.
type ModifyMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewModifyMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewModifyMetadataOK creates a ModifyMetadataOK with default headers values
func NewModifyMetadataOK() *ModifyMetadataOK {
	return &ModifyMetadataOK{}
}

/* ModifyMetadataOK describes a response with status code 200, with default header values.

Metadata object with list of values
*/
type ModifyMetadataOK struct {
	Payload *models.Metadata
}

func (o *ModifyMetadataOK) Error() string {
	return fmt.Sprintf("[PUT /servers/{server_id}/zones/{zone_id}/metadata/{metadata_kind}][%d] modifyMetadataOK  %+v", 200, o.Payload)
}
func (o *ModifyMetadataOK) GetPayload() *models.Metadata {
	return o.Payload
}

func (o *ModifyMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Metadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
