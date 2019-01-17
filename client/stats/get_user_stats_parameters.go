// Code generated by go-swagger; DO NOT EDIT.

package stats

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

// NewGetUserStatsParams creates a new GetUserStatsParams object
// with the default values initialized.
func NewGetUserStatsParams() *GetUserStatsParams {
	var (
		rangeVarDefault = string("last_7_days")
	)
	return &GetUserStatsParams{
		Range: rangeVarDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserStatsParamsWithTimeout creates a new GetUserStatsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserStatsParamsWithTimeout(timeout time.Duration) *GetUserStatsParams {
	var (
		rangeVarDefault = string("last_7_days")
	)
	return &GetUserStatsParams{
		Range: rangeVarDefault,

		timeout: timeout,
	}
}

// NewGetUserStatsParamsWithContext creates a new GetUserStatsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserStatsParamsWithContext(ctx context.Context) *GetUserStatsParams {
	var (
		rangeDefault = string("last_7_days")
	)
	return &GetUserStatsParams{
		Range: rangeDefault,

		Context: ctx,
	}
}

// NewGetUserStatsParamsWithHTTPClient creates a new GetUserStatsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserStatsParamsWithHTTPClient(client *http.Client) *GetUserStatsParams {
	var (
		rangeDefault = string("last_7_days")
	)
	return &GetUserStatsParams{
		Range:      rangeDefault,
		HTTPClient: client,
	}
}

/*GetUserStatsParams contains all the parameters to send to the API endpoint
for the get user stats operation typically these are written to a http.Request
*/
type GetUserStatsParams struct {

	/*Range
	  Range activity

	*/
	Range string
	/*User
	  ID of the user to obtain

	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user stats params
func (o *GetUserStatsParams) WithTimeout(timeout time.Duration) *GetUserStatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user stats params
func (o *GetUserStatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user stats params
func (o *GetUserStatsParams) WithContext(ctx context.Context) *GetUserStatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user stats params
func (o *GetUserStatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user stats params
func (o *GetUserStatsParams) WithHTTPClient(client *http.Client) *GetUserStatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user stats params
func (o *GetUserStatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRange adds the rangeVar to the get user stats params
func (o *GetUserStatsParams) WithRange(rangeVar string) *GetUserStatsParams {
	o.SetRange(rangeVar)
	return o
}

// SetRange adds the range to the get user stats params
func (o *GetUserStatsParams) SetRange(rangeVar string) {
	o.Range = rangeVar
}

// WithUser adds the user to the get user stats params
func (o *GetUserStatsParams) WithUser(user string) *GetUserStatsParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the get user stats params
func (o *GetUserStatsParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserStatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param range
	if err := r.SetPathParam("range", o.Range); err != nil {
		return err
	}

	// path param user
	if err := r.SetPathParam("user", o.User); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
