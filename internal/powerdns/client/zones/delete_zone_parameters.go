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

// NewDeleteZoneParams creates a new DeleteZoneParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteZoneParams() *DeleteZoneParams {
	return &DeleteZoneParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteZoneParamsWithTimeout creates a new DeleteZoneParams object
// with the ability to set a timeout on a request.
func NewDeleteZoneParamsWithTimeout(timeout time.Duration) *DeleteZoneParams {
	return &DeleteZoneParams{
		timeout: timeout,
	}
}

// NewDeleteZoneParamsWithContext creates a new DeleteZoneParams object
// with the ability to set a context for a request.
func NewDeleteZoneParamsWithContext(ctx context.Context) *DeleteZoneParams {
	return &DeleteZoneParams{
		Context: ctx,
	}
}

// NewDeleteZoneParamsWithHTTPClient creates a new DeleteZoneParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteZoneParamsWithHTTPClient(client *http.Client) *DeleteZoneParams {
	return &DeleteZoneParams{
		HTTPClient: client,
	}
}

/* DeleteZoneParams contains all the parameters to send to the API endpoint
   for the delete zone operation.

   Typically these are written to a http.Request.
*/
type DeleteZoneParams struct {

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

// WithDefaults hydrates default values in the delete zone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteZoneParams) WithDefaults() *DeleteZoneParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete zone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteZoneParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete zone params
func (o *DeleteZoneParams) WithTimeout(timeout time.Duration) *DeleteZoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete zone params
func (o *DeleteZoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete zone params
func (o *DeleteZoneParams) WithContext(ctx context.Context) *DeleteZoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete zone params
func (o *DeleteZoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete zone params
func (o *DeleteZoneParams) WithHTTPClient(client *http.Client) *DeleteZoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete zone params
func (o *DeleteZoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServerID adds the serverID to the delete zone params
func (o *DeleteZoneParams) WithServerID(serverID string) *DeleteZoneParams {
	o.SetServerID(serverID)
	return o
}

// SetServerID adds the serverId to the delete zone params
func (o *DeleteZoneParams) SetServerID(serverID string) {
	o.ServerID = serverID
}

// WithZoneID adds the zoneID to the delete zone params
func (o *DeleteZoneParams) WithZoneID(zoneID string) *DeleteZoneParams {
	o.SetZoneID(zoneID)
	return o
}

// SetZoneID adds the zoneId to the delete zone params
func (o *DeleteZoneParams) SetZoneID(zoneID string) {
	o.ZoneID = zoneID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteZoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
