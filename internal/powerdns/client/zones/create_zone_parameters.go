// Code generated by go-swagger; DO NOT EDIT.

package zones

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/k3llla/terraform-provider-powerdns/internal/powerdns/models"
)

// NewCreateZoneParams creates a new CreateZoneParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateZoneParams() *CreateZoneParams {
	return &CreateZoneParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateZoneParamsWithTimeout creates a new CreateZoneParams object
// with the ability to set a timeout on a request.
func NewCreateZoneParamsWithTimeout(timeout time.Duration) *CreateZoneParams {
	return &CreateZoneParams{
		timeout: timeout,
	}
}

// NewCreateZoneParamsWithContext creates a new CreateZoneParams object
// with the ability to set a context for a request.
func NewCreateZoneParamsWithContext(ctx context.Context) *CreateZoneParams {
	return &CreateZoneParams{
		Context: ctx,
	}
}

// NewCreateZoneParamsWithHTTPClient creates a new CreateZoneParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateZoneParamsWithHTTPClient(client *http.Client) *CreateZoneParams {
	return &CreateZoneParams{
		HTTPClient: client,
	}
}

/* CreateZoneParams contains all the parameters to send to the API endpoint
   for the create zone operation.

   Typically these are written to a http.Request.
*/
type CreateZoneParams struct {

	/* Rrsets.

	   “true” (default) or “false”, whether to include the “rrsets” in the response Zone object.

	   Default: true
	*/
	Rrsets *bool

	/* ServerID.

	   The id of the server to retrieve
	*/
	ServerID string

	/* ZoneStruct.

	   The zone struct to patch with
	*/
	ZoneStruct *models.Zone

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create zone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateZoneParams) WithDefaults() *CreateZoneParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create zone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateZoneParams) SetDefaults() {
	var (
		rrsetsDefault = bool(true)
	)

	val := CreateZoneParams{
		Rrsets: &rrsetsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create zone params
func (o *CreateZoneParams) WithTimeout(timeout time.Duration) *CreateZoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create zone params
func (o *CreateZoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create zone params
func (o *CreateZoneParams) WithContext(ctx context.Context) *CreateZoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create zone params
func (o *CreateZoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create zone params
func (o *CreateZoneParams) WithHTTPClient(client *http.Client) *CreateZoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create zone params
func (o *CreateZoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRrsets adds the rrsets to the create zone params
func (o *CreateZoneParams) WithRrsets(rrsets *bool) *CreateZoneParams {
	o.SetRrsets(rrsets)
	return o
}

// SetRrsets adds the rrsets to the create zone params
func (o *CreateZoneParams) SetRrsets(rrsets *bool) {
	o.Rrsets = rrsets
}

// WithServerID adds the serverID to the create zone params
func (o *CreateZoneParams) WithServerID(serverID string) *CreateZoneParams {
	o.SetServerID(serverID)
	return o
}

// SetServerID adds the serverId to the create zone params
func (o *CreateZoneParams) SetServerID(serverID string) {
	o.ServerID = serverID
}

// WithZoneStruct adds the zoneStruct to the create zone params
func (o *CreateZoneParams) WithZoneStruct(zoneStruct *models.Zone) *CreateZoneParams {
	o.SetZoneStruct(zoneStruct)
	return o
}

// SetZoneStruct adds the zoneStruct to the create zone params
func (o *CreateZoneParams) SetZoneStruct(zoneStruct *models.Zone) {
	o.ZoneStruct = zoneStruct
}

// WriteToRequest writes these params to a swagger request
func (o *CreateZoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Rrsets != nil {

		// query param rrsets
		var qrRrsets bool

		if o.Rrsets != nil {
			qrRrsets = *o.Rrsets
		}
		qRrsets := swag.FormatBool(qrRrsets)
		if qRrsets != "" {

			if err := r.SetQueryParam("rrsets", qRrsets); err != nil {
				return err
			}
		}
	}

	// path param server_id
	if err := r.SetPathParam("server_id", o.ServerID); err != nil {
		return err
	}
	if o.ZoneStruct != nil {
		if err := r.SetBodyParam(o.ZoneStruct); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
