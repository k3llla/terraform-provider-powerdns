// Code generated by go-swagger; DO NOT EDIT.

package stats

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/gonzolino/terraform-provider-powerdns/internal/powerdns/models"
)

// GetStatsReader is a Reader for the GetStats structure.
type GetStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewGetStatsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetStatsOK creates a GetStatsOK with default headers values
func NewGetStatsOK() *GetStatsOK {
	return &GetStatsOK{}
}

/* GetStatsOK describes a response with status code 200, with default header values.

List of Statistic Items
*/
type GetStatsOK struct {
	Payload *GetStatsOKBodyTuple0
}

func (o *GetStatsOK) Error() string {
	return fmt.Sprintf("[GET /servers/{server_id}/statistics][%d] getStatsOK  %+v", 200, o.Payload)
}
func (o *GetStatsOK) GetPayload() *GetStatsOKBodyTuple0 {
	return o.Payload
}

func (o *GetStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetStatsOKBodyTuple0)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatsUnprocessableEntity creates a GetStatsUnprocessableEntity with default headers values
func NewGetStatsUnprocessableEntity() *GetStatsUnprocessableEntity {
	return &GetStatsUnprocessableEntity{}
}

/* GetStatsUnprocessableEntity describes a response with status code 422, with default header values.

Returned when a non-existing statistic name has been requested. Contains an error message
*/
type GetStatsUnprocessableEntity struct {
}

func (o *GetStatsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /servers/{server_id}/statistics][%d] getStatsUnprocessableEntity ", 422)
}

func (o *GetStatsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetStatsOKBodyTuple0 GetStatsOKBodyTuple0 a representation of an anonymous Tuple type
swagger:model GetStatsOKBodyTuple0
*/
type GetStatsOKBodyTuple0 struct {

	// p0
	// Required: true
	P0 *models.StatisticItem `json:"-"` // custom serializer

	// p1
	// Required: true
	P1 *models.MapStatisticItem `json:"-"` // custom serializer

	// p2
	// Required: true
	P2 *models.RingStatisticItem `json:"-"` // custom serializer

}

// UnmarshalJSON unmarshals this tuple type from a JSON array
func (o *GetStatsOKBodyTuple0) UnmarshalJSON(raw []byte) error {
	// stage 1, get the array but just the array
	var stage1 []json.RawMessage
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&stage1); err != nil {
		return err
	}

	// stage 2: hydrates struct members with tuple elements
	if len(stage1) > 0 {
		var dataP0 models.StatisticItem
		buf = bytes.NewBuffer(stage1[0])
		dec := json.NewDecoder(buf)
		dec.UseNumber()
		if err := dec.Decode(&dataP0); err != nil {
			return err
		}
		o.P0 = &dataP0

	}
	if len(stage1) > 1 {
		var dataP1 models.MapStatisticItem
		buf = bytes.NewBuffer(stage1[1])
		dec := json.NewDecoder(buf)
		dec.UseNumber()
		if err := dec.Decode(&dataP1); err != nil {
			return err
		}
		o.P1 = &dataP1

	}
	if len(stage1) > 2 {
		var dataP2 models.RingStatisticItem
		buf = bytes.NewBuffer(stage1[2])
		dec := json.NewDecoder(buf)
		dec.UseNumber()
		if err := dec.Decode(&dataP2); err != nil {
			return err
		}
		o.P2 = &dataP2

	}

	return nil
}

// MarshalJSON marshals this tuple type into a JSON array
func (o GetStatsOKBodyTuple0) MarshalJSON() ([]byte, error) {
	data := []interface{}{
		o.P0, o.P1, o.P2,
	}

	return json.Marshal(data)
}

// Validate validates this get stats o k body tuple0
func (o *GetStatsOKBodyTuple0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateP0(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateP1(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateP2(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetStatsOKBodyTuple0) validateP0(formats strfmt.Registry) error {

	if err := validate.Required("P0", "body", o.P0); err != nil {
		return err
	}

	if o.P0 != nil {
		if err := o.P0.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("P0")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("P0")
			}
			return err
		}
	}

	return nil
}

func (o *GetStatsOKBodyTuple0) validateP1(formats strfmt.Registry) error {

	if err := validate.Required("P1", "body", o.P1); err != nil {
		return err
	}

	if o.P1 != nil {
		if err := o.P1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("P1")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("P1")
			}
			return err
		}
	}

	return nil
}

func (o *GetStatsOKBodyTuple0) validateP2(formats strfmt.Registry) error {

	if err := validate.Required("P2", "body", o.P2); err != nil {
		return err
	}

	if o.P2 != nil {
		if err := o.P2.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("P2")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("P2")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get stats o k body tuple0 based on the context it is used
func (o *GetStatsOKBodyTuple0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateP0(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateP1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateP2(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetStatsOKBodyTuple0) contextValidateP0(ctx context.Context, formats strfmt.Registry) error {

	if o.P0 != nil {
		if err := o.P0.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("P0")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("P0")
			}
			return err
		}
	}

	return nil
}

func (o *GetStatsOKBodyTuple0) contextValidateP1(ctx context.Context, formats strfmt.Registry) error {

	if o.P1 != nil {
		if err := o.P1.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("P1")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("P1")
			}
			return err
		}
	}

	return nil
}

func (o *GetStatsOKBodyTuple0) contextValidateP2(ctx context.Context, formats strfmt.Registry) error {

	if o.P2 != nil {
		if err := o.P2.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("P2")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("P2")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetStatsOKBodyTuple0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStatsOKBodyTuple0) UnmarshalBinary(b []byte) error {
	var res GetStatsOKBodyTuple0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
