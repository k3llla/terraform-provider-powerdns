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

// NewModifyCryptokeyParams creates a new ModifyCryptokeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewModifyCryptokeyParams() *ModifyCryptokeyParams {
	return &ModifyCryptokeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewModifyCryptokeyParamsWithTimeout creates a new ModifyCryptokeyParams object
// with the ability to set a timeout on a request.
func NewModifyCryptokeyParamsWithTimeout(timeout time.Duration) *ModifyCryptokeyParams {
	return &ModifyCryptokeyParams{
		timeout: timeout,
	}
}

// NewModifyCryptokeyParamsWithContext creates a new ModifyCryptokeyParams object
// with the ability to set a context for a request.
func NewModifyCryptokeyParamsWithContext(ctx context.Context) *ModifyCryptokeyParams {
	return &ModifyCryptokeyParams{
		Context: ctx,
	}
}

// NewModifyCryptokeyParamsWithHTTPClient creates a new ModifyCryptokeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewModifyCryptokeyParamsWithHTTPClient(client *http.Client) *ModifyCryptokeyParams {
	return &ModifyCryptokeyParams{
		HTTPClient: client,
	}
}

/* ModifyCryptokeyParams contains all the parameters to send to the API endpoint
   for the modify cryptokey operation.

   Typically these are written to a http.Request.
*/
type ModifyCryptokeyParams struct {

	/* Cryptokey.

	   the Cryptokey
	*/
	Cryptokey *models.Cryptokey

	/* CryptokeyID.

	   Cryptokey to manipulate
	*/
	CryptokeyID string

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

// WithDefaults hydrates default values in the modify cryptokey params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModifyCryptokeyParams) WithDefaults() *ModifyCryptokeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the modify cryptokey params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModifyCryptokeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the modify cryptokey params
func (o *ModifyCryptokeyParams) WithTimeout(timeout time.Duration) *ModifyCryptokeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify cryptokey params
func (o *ModifyCryptokeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify cryptokey params
func (o *ModifyCryptokeyParams) WithContext(ctx context.Context) *ModifyCryptokeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify cryptokey params
func (o *ModifyCryptokeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify cryptokey params
func (o *ModifyCryptokeyParams) WithHTTPClient(client *http.Client) *ModifyCryptokeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify cryptokey params
func (o *ModifyCryptokeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCryptokey adds the cryptokey to the modify cryptokey params
func (o *ModifyCryptokeyParams) WithCryptokey(cryptokey *models.Cryptokey) *ModifyCryptokeyParams {
	o.SetCryptokey(cryptokey)
	return o
}

// SetCryptokey adds the cryptokey to the modify cryptokey params
func (o *ModifyCryptokeyParams) SetCryptokey(cryptokey *models.Cryptokey) {
	o.Cryptokey = cryptokey
}

// WithCryptokeyID adds the cryptokeyID to the modify cryptokey params
func (o *ModifyCryptokeyParams) WithCryptokeyID(cryptokeyID string) *ModifyCryptokeyParams {
	o.SetCryptokeyID(cryptokeyID)
	return o
}

// SetCryptokeyID adds the cryptokeyId to the modify cryptokey params
func (o *ModifyCryptokeyParams) SetCryptokeyID(cryptokeyID string) {
	o.CryptokeyID = cryptokeyID
}

// WithServerID adds the serverID to the modify cryptokey params
func (o *ModifyCryptokeyParams) WithServerID(serverID string) *ModifyCryptokeyParams {
	o.SetServerID(serverID)
	return o
}

// SetServerID adds the serverId to the modify cryptokey params
func (o *ModifyCryptokeyParams) SetServerID(serverID string) {
	o.ServerID = serverID
}

// WithZoneID adds the zoneID to the modify cryptokey params
func (o *ModifyCryptokeyParams) WithZoneID(zoneID string) *ModifyCryptokeyParams {
	o.SetZoneID(zoneID)
	return o
}

// SetZoneID adds the zoneId to the modify cryptokey params
func (o *ModifyCryptokeyParams) SetZoneID(zoneID string) {
	o.ZoneID = zoneID
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyCryptokeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Cryptokey != nil {
		if err := r.SetBodyParam(o.Cryptokey); err != nil {
			return err
		}
	}

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
