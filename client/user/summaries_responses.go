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

// SummariesReader is a Reader for the Summaries structure.
type SummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSummariesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSummariesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSummariesOK creates a SummariesOK with default headers values
func NewSummariesOK() *SummariesOK {
	return &SummariesOK{}
}

/*SummariesOK handles this case with default header values.

The request has succeeded.
*/
type SummariesOK struct {
	Payload *models.Summaries
}

func (o *SummariesOK) Error() string {
	return fmt.Sprintf("[GET /users/{user}/summaries][%d] summariesOK  %+v", 200, o.Payload)
}

func (o *SummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Summaries)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSummariesUnauthorized creates a SummariesUnauthorized with default headers values
func NewSummariesUnauthorized() *SummariesUnauthorized {
	return &SummariesUnauthorized{}
}

/*SummariesUnauthorized handles this case with default header values.

The request requires authentication, or your authentication was invalid.
*/
type SummariesUnauthorized struct {
	Payload *models.Error
}

func (o *SummariesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{user}/summaries][%d] summariesUnauthorized  %+v", 401, o.Payload)
}

func (o *SummariesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSummariesNotFound creates a SummariesNotFound with default headers values
func NewSummariesNotFound() *SummariesNotFound {
	return &SummariesNotFound{}
}

/*SummariesNotFound handles this case with default header values.

The resource does not exist.
*/
type SummariesNotFound struct {
	Payload *models.Error
}

func (o *SummariesNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user}/summaries][%d] summariesNotFound  %+v", 404, o.Payload)
}

func (o *SummariesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
