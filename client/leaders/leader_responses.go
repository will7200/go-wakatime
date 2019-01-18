// Code generated by go-swagger; DO NOT EDIT.

package leaders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/will7200/wakatime/models"
)

// LeaderReader is a Reader for the Leader structure.
type LeaderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LeaderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewLeaderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewLeaderUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewLeaderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLeaderOK creates a LeaderOK with default headers values
func NewLeaderOK() *LeaderOK {
	return &LeaderOK{}
}

/*LeaderOK handles this case with default header values.

The request has succeeded.
*/
type LeaderOK struct {
	Payload *models.Leaders
}

func (o *LeaderOK) Error() string {
	return fmt.Sprintf("[GET /leaders][%d] leaderOK  %+v", 200, o.Payload)
}

func (o *LeaderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Leaders)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLeaderUnauthorized creates a LeaderUnauthorized with default headers values
func NewLeaderUnauthorized() *LeaderUnauthorized {
	return &LeaderUnauthorized{}
}

/*LeaderUnauthorized handles this case with default header values.

The request requires authentication, or your authentication was invalid.
*/
type LeaderUnauthorized struct {
	Payload *models.Error
}

func (o *LeaderUnauthorized) Error() string {
	return fmt.Sprintf("[GET /leaders][%d] leaderUnauthorized  %+v", 401, o.Payload)
}

func (o *LeaderUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLeaderNotFound creates a LeaderNotFound with default headers values
func NewLeaderNotFound() *LeaderNotFound {
	return &LeaderNotFound{}
}

/*LeaderNotFound handles this case with default header values.

The resource does not exist.
*/
type LeaderNotFound struct {
}

func (o *LeaderNotFound) Error() string {
	return fmt.Sprintf("[GET /leaders][%d] leaderNotFound ", 404)
}

func (o *LeaderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
