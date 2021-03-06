// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/will7200/go-wakatime/models"
)

// UserAgentsReader is a Reader for the UserAgents structure.
type UserAgentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserAgentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserAgentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewUserAgentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUserAgentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserAgentsOK creates a UserAgentsOK with default headers values
func NewUserAgentsOK() *UserAgentsOK {
	return &UserAgentsOK{}
}

/*UserAgentsOK handles this case with default header values.

List of plugins.
*/
type UserAgentsOK struct {
	Payload *UserAgentsOKBody
}

func (o *UserAgentsOK) Error() string {
	return fmt.Sprintf("[GET /users/{user}/user_agents][%d] userAgentsOK  %+v", 200, o.Payload)
}

func (o *UserAgentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UserAgentsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserAgentsUnauthorized creates a UserAgentsUnauthorized with default headers values
func NewUserAgentsUnauthorized() *UserAgentsUnauthorized {
	return &UserAgentsUnauthorized{}
}

/*UserAgentsUnauthorized handles this case with default header values.

The request requires authentication, or your authentication was invalid.
*/
type UserAgentsUnauthorized struct {
	Payload *models.Error
}

func (o *UserAgentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{user}/user_agents][%d] userAgentsUnauthorized  %+v", 401, o.Payload)
}

func (o *UserAgentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserAgentsNotFound creates a UserAgentsNotFound with default headers values
func NewUserAgentsNotFound() *UserAgentsNotFound {
	return &UserAgentsNotFound{}
}

/*UserAgentsNotFound handles this case with default header values.

The resource does not exist.
*/
type UserAgentsNotFound struct {
	Payload *models.Error
}

func (o *UserAgentsNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user}/user_agents][%d] userAgentsNotFound  %+v", 404, o.Payload)
}

func (o *UserAgentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UserAgentsOKBody user agents o k body
swagger:model UserAgentsOKBody
*/
type UserAgentsOKBody struct {

	// data
	Data []*models.UserAgents `json:"data"`
}

// Validate validates this user agents o k body
func (o *UserAgentsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UserAgentsOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userAgentsOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UserAgentsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UserAgentsOKBody) UnmarshalBinary(b []byte) error {
	var res UserAgentsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
