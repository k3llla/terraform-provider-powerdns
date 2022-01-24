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
)

// NewRectifyZoneParams creates a new RectifyZoneParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRectifyZoneParams() *RectifyZoneParams {
	return &RectifyZoneParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRectifyZoneParamsWithTimeout creates a new RectifyZoneParams object
// with the ability to set a timeout on a request.
func NewRectifyZoneParamsWithTimeout(timeout time.Duration) *RectifyZoneParams {
	return &RectifyZoneParams{
		timeout: timeout,
	}
}

// NewRectifyZoneParamsWithContext creates a new RectifyZoneParams object
// with the ability to set a context for a request.
func NewRectifyZoneParamsWithContext(ctx context.Context) *RectifyZoneParams {
	return &RectifyZoneParams{
		Context: ctx,
	}
}

// NewRectifyZoneParamsWithHTTPClient creates a new RectifyZoneParams object
// with the ability to set a custom HTTPClient for a request.
func NewRectifyZoneParamsWithHTTPClient(client *http.Client) *RectifyZoneParams {
	return &RectifyZoneParams{
		HTTPClient: client,
	}
}

/* RectifyZoneParams contains all the parameters to send to the API endpoint
   for the rectify zone operation.

   Typically these are written to a http.Request.
*/
type RectifyZoneParams struct {

	/* ServerID.

	   The id of the server to retrieve
	*/
	ServerID string

	/* ZoneID.

	   The id of the zone to retrieve
	*/
	ZoneID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the rectify zone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RectifyZoneParams) WithDefaults() *RectifyZoneParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the rectify zone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RectifyZoneParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the rectify zone params
func (o *RectifyZoneParams) WithTimeout(timeout time.Duration) *RectifyZoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rectify zone params
func (o *RectifyZoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rectify zone params
func (o *RectifyZoneParams) WithContext(ctx context.Context) *RectifyZoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rectify zone params
func (o *RectifyZoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rectify zone params
func (o *RectifyZoneParams) WithHTTPClient(client *http.Client) *RectifyZoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rectify zone params
func (o *RectifyZoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServerID adds the serverID to the rectify zone params
func (o *RectifyZoneParams) WithServerID(serverID string) *RectifyZoneParams {
	o.SetServerID(serverID)
	return o
}

// SetServerID adds the serverId to the rectify zone params
func (o *RectifyZoneParams) SetServerID(serverID string) {
	o.ServerID = serverID
}

// WithZoneID adds the zoneID to the rectify zone params
func (o *RectifyZoneParams) WithZoneID(zoneID string) *RectifyZoneParams {
	o.SetZoneID(zoneID)
	return o
}

// SetZoneID adds the zoneId to the rectify zone params
func (o *RectifyZoneParams) SetZoneID(zoneID string) {
	o.ZoneID = zoneID
}

// WriteToRequest writes these params to a swagger request
func (o *RectifyZoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param server_id
	if err := r.SetPathParam("server_id", o.ServerID); err != nil {
		return err
	}

	// path param zone_id
	if err := r.SetPathParam("zone_id", o.ZoneID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
