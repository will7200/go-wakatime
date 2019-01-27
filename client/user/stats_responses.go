// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/will7200/go-wakatime/models"
)

// StatsReader is a Reader for the Stats structure.
type StatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewStatsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewStatsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewStatsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStatsOK creates a StatsOK with default headers values
func NewStatsOK() *StatsOK {
	return &StatsOK{}
}

/*StatsOK handles this case with default header values.

Statistical Coding Activity
*/
type StatsOK struct {
	Payload *StatsOKBody
}

func (o *StatsOK) Error() string {
	return fmt.Sprintf("[GET /users/{user}/stats/{range}][%d] statsOK  %+v", 200, o.Payload)
}

func (o *StatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StatsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatsAccepted creates a StatsAccepted with default headers values
func NewStatsAccepted() *StatsAccepted {
	return &StatsAccepted{}
}

/*StatsAccepted handles this case with default header values.

Stats Accepted to Compute
*/
type StatsAccepted struct {
	Payload *StatsAcceptedBody
}

func (o *StatsAccepted) Error() string {
	return fmt.Sprintf("[GET /users/{user}/stats/{range}][%d] statsAccepted  %+v", 202, o.Payload)
}

func (o *StatsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StatsAcceptedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatsUnauthorized creates a StatsUnauthorized with default headers values
func NewStatsUnauthorized() *StatsUnauthorized {
	return &StatsUnauthorized{}
}

/*StatsUnauthorized handles this case with default header values.

The request requires authentication, or your authentication was invalid.
*/
type StatsUnauthorized struct {
	Payload *models.Error
}

func (o *StatsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{user}/stats/{range}][%d] statsUnauthorized  %+v", 401, o.Payload)
}

func (o *StatsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatsNotFound creates a StatsNotFound with default headers values
func NewStatsNotFound() *StatsNotFound {
	return &StatsNotFound{}
}

/*StatsNotFound handles this case with default header values.

The resource does not exist.
*/
type StatsNotFound struct {
}

func (o *StatsNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user}/stats/{range}][%d] statsNotFound ", 404)
}

func (o *StatsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*StatsAcceptedBody stats accepted body
swagger:model StatsAcceptedBody
*/
type StatsAcceptedBody struct {

	// data
	Data *models.StatsPending `json:"data,omitempty"`
}

// Validate validates this stats accepted body
func (o *StatsAcceptedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StatsAcceptedBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statsAccepted" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *StatsAcceptedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StatsAcceptedBody) UnmarshalBinary(b []byte) error {
	var res StatsAcceptedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StatsOKBody stats o k body
swagger:model StatsOKBody
*/
type StatsOKBody struct {

	// data
	Data *models.Stats `json:"data,omitempty"`
}

// Validate validates this stats o k body
func (o *StatsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StatsOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statsOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *StatsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StatsOKBody) UnmarshalBinary(b []byte) error {
	var res StatsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
