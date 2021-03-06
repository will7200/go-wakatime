// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/will7200/go-wakatime/models"
)

// DurationReader is a Reader for the Duration structure.
type DurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDurationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDurationOK creates a DurationOK with default headers values
func NewDurationOK() *DurationOK {
	return &DurationOK{}
}

/*DurationOK handles this case with default header values.

The request has succeeded.
*/
type DurationOK struct {
	Payload *models.Durations
}

func (o *DurationOK) Error() string {
	return fmt.Sprintf("[GET /users/{user}/durations][%d] durationOK  %+v", 200, o.Payload)
}

func (o *DurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Durations)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDurationUnauthorized creates a DurationUnauthorized with default headers values
func NewDurationUnauthorized() *DurationUnauthorized {
	return &DurationUnauthorized{}
}

/*DurationUnauthorized handles this case with default header values.

The request requires authentication, or your authentication was invalid.
*/
type DurationUnauthorized struct {
	Payload *models.Error
}

func (o *DurationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{user}/durations][%d] durationUnauthorized  %+v", 401, o.Payload)
}

func (o *DurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDurationNotFound creates a DurationNotFound with default headers values
func NewDurationNotFound() *DurationNotFound {
	return &DurationNotFound{}
}

/*DurationNotFound handles this case with default header values.

The resource does not exist.
*/
type DurationNotFound struct {
	Payload *models.Error
}

func (o *DurationNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user}/durations][%d] durationNotFound  %+v", 404, o.Payload)
}

func (o *DurationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
