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

	"github.com/k3llla/terraform-provider-powerdns/internal/powerdns/models"
)

// NewCreateCryptokeyParams creates a new CreateCryptokeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateCryptokeyParams() *CreateCryptokeyParams {
	return &CreateCryptokeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCryptokeyParamsWithTimeout creates a new CreateCryptokeyParams object
// with the ability to set a timeout on a request.
func NewCreateCryptokeyParamsWithTimeout(timeout time.Duration) *CreateCryptokeyParams {
	return &CreateCryptokeyParams{
		timeout: timeout,
	}
}

// NewCreateCryptokeyParamsWithContext creates a new CreateCryptokeyParams object
// with the ability to set a context for a request.
func NewCreateCryptokeyParamsWithContext(ctx context.Context) *CreateCryptokeyParams {
	return &CreateCryptokeyParams{
		Context: ctx,
	}
}

// NewCreateCryptokeyParamsWithHTTPClient creates a new CreateCryptokeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateCryptokeyParamsWithHTTPClient(client *http.Client) *CreateCryptokeyParams {
	return &CreateCryptokeyParams{
		HTTPClient: client,
	}
}

/* CreateCryptokeyParams contains all the parameters to send to the API endpoint
   for the create cryptokey operation.

   Typically these are written to a http.Request.
*/
type CreateCryptokeyParams struct {

	/* Cryptokey.

	   Add a Cryptokey
	*/
	Cryptokey *models.Cryptokey

	/* ServerID.

	   The id of the server to retrieve
	*/
	ServerID string

	// ZoneID.
	ZoneID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create cryptokey params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCryptokeyParams) WithDefaults() *CreateCryptokeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create cryptokey params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCryptokeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create cryptokey params
func (o *CreateCryptokeyParams) WithTimeout(timeout time.Duration) *CreateCryptokeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cryptokey params
func (o *CreateCryptokeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cryptokey params
func (o *CreateCryptokeyParams) WithContext(ctx context.Context) *CreateCryptokeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cryptokey params
func (o *CreateCryptokeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cryptokey params
func (o *CreateCryptokeyParams) WithHTTPClient(client *http.Client) *CreateCryptokeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cryptokey params
func (o *CreateCryptokeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCryptokey adds the cryptokey to the create cryptokey params
func (o *CreateCryptokeyParams) WithCryptokey(cryptokey *models.Cryptokey) *CreateCryptokeyParams {
	o.SetCryptokey(cryptokey)
	return o
}

// SetCryptokey adds the cryptokey to the create cryptokey params
func (o *CreateCryptokeyParams) SetCryptokey(cryptokey *models.Cryptokey) {
	o.Cryptokey = cryptokey
}

// WithServerID adds the serverID to the create cryptokey params
func (o *CreateCryptokeyParams) WithServerID(serverID string) *CreateCryptokeyParams {
	o.SetServerID(serverID)
	return o
}

// SetServerID adds the serverId to the create cryptokey params
func (o *CreateCryptokeyParams) SetServerID(serverID string) {
	o.ServerID = serverID
}

// WithZoneID adds the zoneID to the create cryptokey params
func (o *CreateCryptokeyParams) WithZoneID(zoneID string) *CreateCryptokeyParams {
	o.SetZoneID(zoneID)
	return o
}

// SetZoneID adds the zoneId to the create cryptokey params
func (o *CreateCryptokeyParams) SetZoneID(zoneID string) {
	o.ZoneID = zoneID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCryptokeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Cryptokey != nil {
		if err := r.SetBodyParam(o.Cryptokey); err != nil {
			return err
		}
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
