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

	"github.com/k3llla/terraform-provider-powerdns/internal/powerdns/models"
)

// NewModifyMetadataParams creates a new ModifyMetadataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewModifyMetadataParams() *ModifyMetadataParams {
	return &ModifyMetadataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewModifyMetadataParamsWithTimeout creates a new ModifyMetadataParams object
// with the ability to set a timeout on a request.
func NewModifyMetadataParamsWithTimeout(timeout time.Duration) *ModifyMetadataParams {
	return &ModifyMetadataParams{
		timeout: timeout,
	}
}

// NewModifyMetadataParamsWithContext creates a new ModifyMetadataParams object
// with the ability to set a context for a request.
func NewModifyMetadataParamsWithContext(ctx context.Context) *ModifyMetadataParams {
	return &ModifyMetadataParams{
		Context: ctx,
	}
}

// NewModifyMetadataParamsWithHTTPClient creates a new ModifyMetadataParams object
// with the ability to set a custom HTTPClient for a request.
func NewModifyMetadataParamsWithHTTPClient(client *http.Client) *ModifyMetadataParams {
	return &ModifyMetadataParams{
		HTTPClient: client,
	}
}

/* ModifyMetadataParams contains all the parameters to send to the API endpoint
   for the modify metadata operation.

   Typically these are written to a http.Request.
*/
type ModifyMetadataParams struct {

	/* Metadata.

	   metadata to add/create
	*/
	Metadata *models.Metadata

	/* MetadataKind.

	   The kind of metadata
	*/
	MetadataKind string

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

// WithDefaults hydrates default values in the modify metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModifyMetadataParams) WithDefaults() *ModifyMetadataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the modify metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModifyMetadataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the modify metadata params
func (o *ModifyMetadataParams) WithTimeout(timeout time.Duration) *ModifyMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify metadata params
func (o *ModifyMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify metadata params
func (o *ModifyMetadataParams) WithContext(ctx context.Context) *ModifyMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify metadata params
func (o *ModifyMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify metadata params
func (o *ModifyMetadataParams) WithHTTPClient(client *http.Client) *ModifyMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify metadata params
func (o *ModifyMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMetadata adds the metadata to the modify metadata params
func (o *ModifyMetadataParams) WithMetadata(metadata *models.Metadata) *ModifyMetadataParams {
	o.SetMetadata(metadata)
	return o
}

// SetMetadata adds the metadata to the modify metadata params
func (o *ModifyMetadataParams) SetMetadata(metadata *models.Metadata) {
	o.Metadata = metadata
}

// WithMetadataKind adds the metadataKind to the modify metadata params
func (o *ModifyMetadataParams) WithMetadataKind(metadataKind string) *ModifyMetadataParams {
	o.SetMetadataKind(metadataKind)
	return o
}

// SetMetadataKind adds the metadataKind to the modify metadata params
func (o *ModifyMetadataParams) SetMetadataKind(metadataKind string) {
	o.MetadataKind = metadataKind
}

// WithServerID adds the serverID to the modify metadata params
func (o *ModifyMetadataParams) WithServerID(serverID string) *ModifyMetadataParams {
	o.SetServerID(serverID)
	return o
}

// SetServerID adds the serverId to the modify metadata params
func (o *ModifyMetadataParams) SetServerID(serverID string) {
	o.ServerID = serverID
}

// WithZoneID adds the zoneID to the modify metadata params
func (o *ModifyMetadataParams) WithZoneID(zoneID string) *ModifyMetadataParams {
	o.SetZoneID(zoneID)
	return o
}

// SetZoneID adds the zoneId to the modify metadata params
func (o *ModifyMetadataParams) SetZoneID(zoneID string) {
	o.ZoneID = zoneID
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Metadata != nil {
		if err := r.SetBodyParam(o.Metadata); err != nil {
			return err
		}
	}

	// path param metadata_kind
	if err := r.SetPathParam("metadata_kind", o.MetadataKind); err != nil {
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
