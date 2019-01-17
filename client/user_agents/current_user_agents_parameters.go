// Code generated by go-swagger; DO NOT EDIT.

package user_agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCurrentUserAgentsParams creates a new CurrentUserAgentsParams object
// with the default values initialized.
func NewCurrentUserAgentsParams() *CurrentUserAgentsParams {

	return &CurrentUserAgentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCurrentUserAgentsParamsWithTimeout creates a new CurrentUserAgentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCurrentUserAgentsParamsWithTimeout(timeout time.Duration) *CurrentUserAgentsParams {

	return &CurrentUserAgentsParams{

		timeout: timeout,
	}
}

// NewCurrentUserAgentsParamsWithContext creates a new CurrentUserAgentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewCurrentUserAgentsParamsWithContext(ctx context.Context) *CurrentUserAgentsParams {

	return &CurrentUserAgentsParams{

		Context: ctx,
	}
}

// NewCurrentUserAgentsParamsWithHTTPClient creates a new CurrentUserAgentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCurrentUserAgentsParamsWithHTTPClient(client *http.Client) *CurrentUserAgentsParams {

	return &CurrentUserAgentsParams{
		HTTPClient: client,
	}
}

/*CurrentUserAgentsParams contains all the parameters to send to the API endpoint
for the current user agents operation typically these are written to a http.Request
*/
type CurrentUserAgentsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the current user agents params
func (o *CurrentUserAgentsParams) WithTimeout(timeout time.Duration) *CurrentUserAgentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the current user agents params
func (o *CurrentUserAgentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the current user agents params
func (o *CurrentUserAgentsParams) WithContext(ctx context.Context) *CurrentUserAgentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the current user agents params
func (o *CurrentUserAgentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the current user agents params
func (o *CurrentUserAgentsParams) WithHTTPClient(client *http.Client) *CurrentUserAgentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the current user agents params
func (o *CurrentUserAgentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CurrentUserAgentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
