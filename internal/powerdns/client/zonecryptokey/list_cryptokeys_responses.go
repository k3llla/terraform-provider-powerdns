// Code generated by go-swagger; DO NOT EDIT.

package zonecryptokey

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/gonzolino/terraform-provider-powerdns/internal/powerdns/models"
)

// ListCryptokeysReader is a Reader for the ListCryptokeys structure.
type ListCryptokeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCryptokeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCryptokeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListCryptokeysOK creates a ListCryptokeysOK with default headers values
func NewListCryptokeysOK() *ListCryptokeysOK {
	return &ListCryptokeysOK{}
}

/* ListCryptokeysOK describes a response with status code 200, with default header values.

List of Cryptokey objects
*/
type ListCryptokeysOK struct {
	Payload []*models.Cryptokey
}

func (o *ListCryptokeysOK) Error() string {
	return fmt.Sprintf("[GET /servers/{server_id}/zones/{zone_id}/cryptokeys][%d] listCryptokeysOK  %+v", 200, o.Payload)
}
func (o *ListCryptokeysOK) GetPayload() []*models.Cryptokey {
	return o.Payload
}

func (o *ListCryptokeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
