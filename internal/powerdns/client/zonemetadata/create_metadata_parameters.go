// Code generated by go-swagger; DO NOT EDIT.

package zonemetadata

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

	"github.com/gonzolino/terraform-provider-powerdns/internal/powerdns/models"
)

// NewCreateMetadataParams creates a new CreateMetadataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateMetadataParams() *CreateMetadataParams {
	return &CreateMetadataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMetadataParamsWithTimeout creates a new CreateMetadataParams object
// with the ability to set a timeout on a request.
func NewCreateMetadataParamsWithTimeout(timeout time.Duration) *CreateMetadataParams {
	return &CreateMetadataParams{
		timeout: timeout,
	}
}

// NewCreateMetadataParamsWithContext creates a new CreateMetadataParams object
// with the ability to set a context for a request.
func NewCreateMetadataParamsWithContext(ctx context.Context) *CreateMetadataParams {
	return &CreateMetadataParams{
		Context: ctx,
	}
}

// NewCreateMetadataParamsWithHTTPClient creates a new CreateMetadataParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateMetadataParamsWithHTTPClient(client *http.Client) *CreateMetadataParams {
	return &CreateMetadataParams{
		HTTPClient: client,
	}
}

/* CreateMetadataParams contains all the parameters to send to the API endpoint
   for the create metadata operation.

   Typically these are written to a http.Request.
*/
type CreateMetadataParams struct {

	/* Metadata.

	   Metadata object with list of values to create
	*/
	Metadata *models.Metadata

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

// WithDefaults hydrates default values in the create metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMetadataParams) WithDefaults() *CreateMetadataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMetadataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create metadata params
func (o *CreateMetadataParams) WithTimeout(timeout time.Duration) *CreateMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create metadata params
func (o *CreateMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create metadata params
func (o *CreateMetadataParams) WithContext(ctx context.Context) *CreateMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create metadata params
func (o *CreateMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create metadata params
func (o *CreateMetadataParams) WithHTTPClient(client *http.Client) *CreateMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create metadata params
func (o *CreateMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMetadata adds the metadata to the create metadata params
func (o *CreateMetadataParams) WithMetadata(metadata *models.Metadata) *CreateMetadataParams {
	o.SetMetadata(metadata)
	return o
}

// SetMetadata adds the metadata to the create metadata params
func (o *CreateMetadataParams) SetMetadata(metadata *models.Metadata) {
	o.Metadata = metadata
}

// WithServerID adds the serverID to the create metadata params
func (o *CreateMetadataParams) WithServerID(serverID string) *CreateMetadataParams {
	o.SetServerID(serverID)
	return o
}

// SetServerID adds the serverId to the create metadata params
func (o *CreateMetadataParams) SetServerID(serverID string) {
	o.ServerID = serverID
}

// WithZoneID adds the zoneID to the create metadata params
func (o *CreateMetadataParams) WithZoneID(zoneID string) *CreateMetadataParams {
	o.SetZoneID(zoneID)
	return o
}

// SetZoneID adds the zoneId to the create metadata params
func (o *CreateMetadataParams) SetZoneID(zoneID string) {
	o.ZoneID = zoneID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Metadata != nil {
		if err := r.SetBodyParam(o.Metadata); err != nil {
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
