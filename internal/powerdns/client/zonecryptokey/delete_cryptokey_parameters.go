// Code generated by go-swagger; DO NOT EDIT.

package zonecryptokey

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

// NewDeleteCryptokeyParams creates a new DeleteCryptokeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteCryptokeyParams() *DeleteCryptokeyParams {
	return &DeleteCryptokeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCryptokeyParamsWithTimeout creates a new DeleteCryptokeyParams object
// with the ability to set a timeout on a request.
func NewDeleteCryptokeyParamsWithTimeout(timeout time.Duration) *DeleteCryptokeyParams {
	return &DeleteCryptokeyParams{
		timeout: timeout,
	}
}

// NewDeleteCryptokeyParamsWithContext creates a new DeleteCryptokeyParams object
// with the ability to set a context for a request.
func NewDeleteCryptokeyParamsWithContext(ctx context.Context) *DeleteCryptokeyParams {
	return &DeleteCryptokeyParams{
		Context: ctx,
	}
}

// NewDeleteCryptokeyParamsWithHTTPClient creates a new DeleteCryptokeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteCryptokeyParamsWithHTTPClient(client *http.Client) *DeleteCryptokeyParams {
	return &DeleteCryptokeyParams{
		HTTPClient: client,
	}
}

/* DeleteCryptokeyParams contains all the parameters to send to the API endpoint
   for the delete cryptokey operation.

   Typically these are written to a http.Request.
*/
type DeleteCryptokeyParams struct {

	/* CryptokeyID.

	   The id value of the Cryptokey
	*/
	CryptokeyID string

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

// WithDefaults hydrates default values in the delete cryptokey params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCryptokeyParams) WithDefaults() *DeleteCryptokeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete cryptokey params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCryptokeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete cryptokey params
func (o *DeleteCryptokeyParams) WithTimeout(timeout time.Duration) *DeleteCryptokeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cryptokey params
func (o *DeleteCryptokeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cryptokey params
func (o *DeleteCryptokeyParams) WithContext(ctx context.Context) *DeleteCryptokeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cryptokey params
func (o *DeleteCryptokeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cryptokey params
func (o *DeleteCryptokeyParams) WithHTTPClient(client *http.Client) *DeleteCryptokeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cryptokey params
func (o *DeleteCryptokeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCryptokeyID adds the cryptokeyID to the delete cryptokey params
func (o *DeleteCryptokeyParams) WithCryptokeyID(cryptokeyID string) *DeleteCryptokeyParams {
	o.SetCryptokeyID(cryptokeyID)
	return o
}

// SetCryptokeyID adds the cryptokeyId to the delete cryptokey params
func (o *DeleteCryptokeyParams) SetCryptokeyID(cryptokeyID string) {
	o.CryptokeyID = cryptokeyID
}

// WithServerID adds the serverID to the delete cryptokey params
func (o *DeleteCryptokeyParams) WithServerID(serverID string) *DeleteCryptokeyParams {
	o.SetServerID(serverID)
	return o
}

// SetServerID adds the serverId to the delete cryptokey params
func (o *DeleteCryptokeyParams) SetServerID(serverID string) {
	o.ServerID = serverID
}

// WithZoneID adds the zoneID to the delete cryptokey params
func (o *DeleteCryptokeyParams) WithZoneID(zoneID string) *DeleteCryptokeyParams {
	o.SetZoneID(zoneID)
	return o
}

// SetZoneID adds the zoneId to the delete cryptokey params
func (o *DeleteCryptokeyParams) SetZoneID(zoneID string) {
	o.ZoneID = zoneID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCryptokeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cryptokey_id
	if err := r.SetPathParam("cryptokey_id", o.CryptokeyID); err != nil {
		return err
	}

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
