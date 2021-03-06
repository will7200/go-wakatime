// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewGoalsParams creates a new GoalsParams object
// with the default values initialized.
func NewGoalsParams() *GoalsParams {
	var (
		userDefault = string("current")
	)
	return &GoalsParams{
		User: userDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGoalsParamsWithTimeout creates a new GoalsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGoalsParamsWithTimeout(timeout time.Duration) *GoalsParams {
	var (
		userDefault = string("current")
	)
	return &GoalsParams{
		User: userDefault,

		timeout: timeout,
	}
}

// NewGoalsParamsWithContext creates a new GoalsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGoalsParamsWithContext(ctx context.Context) *GoalsParams {
	var (
		userDefault = string("current")
	)
	return &GoalsParams{
		User: userDefault,

		Context: ctx,
	}
}

// NewGoalsParamsWithHTTPClient creates a new GoalsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGoalsParamsWithHTTPClient(client *http.Client) *GoalsParams {
	var (
		userDefault = string("current")
	)
	return &GoalsParams{
		User:       userDefault,
		HTTPClient: client,
	}
}

/*GoalsParams contains all the parameters to send to the API endpoint
for the goals operation typically these are written to a http.Request
*/
type GoalsParams struct {

	/*User
	  ID of the user to obtain

	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the goals params
func (o *GoalsParams) WithTimeout(timeout time.Duration) *GoalsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the goals params
func (o *GoalsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the goals params
func (o *GoalsParams) WithContext(ctx context.Context) *GoalsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the goals params
func (o *GoalsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the goals params
func (o *GoalsParams) WithHTTPClient(client *http.Client) *GoalsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the goals params
func (o *GoalsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUser adds the user to the goals params
func (o *GoalsParams) WithUser(user string) *GoalsParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the goals params
func (o *GoalsParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *GoalsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user
	if err := r.SetPathParam("user", o.User); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
