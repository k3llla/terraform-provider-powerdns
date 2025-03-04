// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/k3llla/terraform-provider-powerdns/internal/powerdns/models"
)

// GetConfigSettingReader is a Reader for the GetConfigSetting structure.
type GetConfigSettingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConfigSettingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConfigSettingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConfigSettingOK creates a GetConfigSettingOK with default headers values
func NewGetConfigSettingOK() *GetConfigSettingOK {
	return &GetConfigSettingOK{}
}

/* GetConfigSettingOK describes a response with status code 200, with default header values.

List of config values
*/
type GetConfigSettingOK struct {
	Payload *models.ConfigSetting
}

func (o *GetConfigSettingOK) Error() string {
	return fmt.Sprintf("[GET /servers/{server_id}/config/{config_setting_name}][%d] getConfigSettingOK  %+v", 200, o.Payload)
}
func (o *GetConfigSettingOK) GetPayload() *models.ConfigSetting {
	return o.Payload
}

func (o *GetConfigSettingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigSetting)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
