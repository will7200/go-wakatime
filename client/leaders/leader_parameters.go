// Code generated by go-swagger; DO NOT EDIT.

package leaders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewLeaderParams creates a new LeaderParams object
// with the default values initialized.
func NewLeaderParams() *LeaderParams {
	var ()
	return &LeaderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLeaderParamsWithTimeout creates a new LeaderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLeaderParamsWithTimeout(timeout time.Duration) *LeaderParams {
	var ()
	return &LeaderParams{

		timeout: timeout,
	}
}

// NewLeaderParamsWithContext creates a new LeaderParams object
// with the default values initialized, and the ability to set a context for a request
func NewLeaderParamsWithContext(ctx context.Context) *LeaderParams {
	var ()
	return &LeaderParams{

		Context: ctx,
	}
}

// NewLeaderParamsWithHTTPClient creates a new LeaderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLeaderParamsWithHTTPClient(client *http.Client) *LeaderParams {
	var ()
	return &LeaderParams{
		HTTPClient: client,
	}
}

/*LeaderParams contains all the parameters to send to the API endpoint
for the leader operation typically these are written to a http.Request
*/
type LeaderParams struct {

	/*Language
	  Filter leaders by a specific language.

	*/
	Language *string
	/*Page
	  Page number of leaderboard.

	*/
	Page *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the leader params
func (o *LeaderParams) WithTimeout(timeout time.Duration) *LeaderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the leader params
func (o *LeaderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the leader params
func (o *LeaderParams) WithContext(ctx context.Context) *LeaderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the leader params
func (o *LeaderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the leader params
func (o *LeaderParams) WithHTTPClient(client *http.Client) *LeaderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the leader params
func (o *LeaderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLanguage adds the language to the leader params
func (o *LeaderParams) WithLanguage(language *string) *LeaderParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the leader params
func (o *LeaderParams) SetLanguage(language *string) {
	o.Language = language
}

// WithPage adds the page to the leader params
func (o *LeaderParams) WithPage(page *int64) *LeaderParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the leader params
func (o *LeaderParams) SetPage(page *int64) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *LeaderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Language != nil {

		// query param language
		var qrLanguage string
		if o.Language != nil {
			qrLanguage = *o.Language
		}
		qLanguage := qrLanguage
		if qLanguage != "" {
			if err := r.SetQueryParam("language", qLanguage); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
