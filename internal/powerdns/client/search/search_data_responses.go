// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/k3llla/terraform-provider-powerdns/internal/powerdns/models"
)

// SearchDataReader is a Reader for the SearchData structure.
type SearchDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchDataOK creates a SearchDataOK with default headers values
func NewSearchDataOK() *SearchDataOK {
	return &SearchDataOK{}
}

/* SearchDataOK describes a response with status code 200, with default header values.

Returns a JSON array with results
*/
type SearchDataOK struct {
	Payload models.SearchResults
}

func (o *SearchDataOK) Error() string {
	return fmt.Sprintf("[GET /servers/{server_id}/search-data][%d] searchDataOK  %+v", 200, o.Payload)
}
func (o *SearchDataOK) GetPayload() models.SearchResults {
	return o.Payload
}

func (o *SearchDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
