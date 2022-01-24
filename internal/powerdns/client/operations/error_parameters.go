// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewErrorParams creates a new ErrorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewErrorParams() *ErrorParams {
	return &ErrorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewErrorParamsWithTimeout creates a new ErrorParams object
// with the ability to set a timeout on a request.
func NewErrorParamsWithTimeout(timeout time.Duration) *ErrorParams {
	return &ErrorParams{
		timeout: timeout,
	}
}

// NewErrorParamsWithContext creates a new ErrorParams object
// with the ability to set a context for a request.
func NewErrorParamsWithContext(ctx context.Context) *ErrorParams {
	return &ErrorParams{
		Context: ctx,
	}
}

// NewErrorParamsWithHTTPClient creates a new ErrorParams object
// with the ability to set a custom HTTPClient for a request.
func NewErrorParamsWithHTTPClient(client *http.Client) *ErrorParams {
	return &ErrorParams{
		HTTPClient: client,
	}
}

/* ErrorParams contains all the parameters to send to the API endpoint
   for the error operation.

   Typically these are written to a http.Request.
*/
type ErrorParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the error params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ErrorParams) WithDefaults() *ErrorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the error params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ErrorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the error params
func (o *ErrorParams) WithTimeout(timeout time.Duration) *ErrorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the error params
func (o *ErrorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the error params
func (o *ErrorParams) WithContext(ctx context.Context) *ErrorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the error params
func (o *ErrorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the error params
func (o *ErrorParams) WithHTTPClient(client *http.Client) *ErrorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the error params
func (o *ErrorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ErrorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
