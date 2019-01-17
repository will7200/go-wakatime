// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new user API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CurrentUser Get the current user.
*/
func (a *Client) CurrentUser(params *CurrentUserParams, authInfo runtime.ClientAuthInfoWriter) (*CurrentUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCurrentUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "currentUser",
		Method:             "GET",
		PathPattern:        "/users/current",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CurrentUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CurrentUserOK), nil

}

/*
GetUser Get the user specified in path
*/
func (a *Client) GetUser(params *GetUserParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUser",
		Method:             "GET",
		PathPattern:        "/users/{user}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
