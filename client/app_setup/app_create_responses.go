// Code generated by go-swagger; DO NOT EDIT.

package app_setup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/peartherapeutics/bitrise-step-api-build-trigger/models"
)

// AppCreateReader is a Reader for the AppCreate structure.
type AppCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AppCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAppCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAppCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAppCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAppCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewAppCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAppCreateOK creates a AppCreateOK with default headers values
func NewAppCreateOK() *AppCreateOK {
	return &AppCreateOK{}
}

/*AppCreateOK handles this case with default header values.

OK
*/
type AppCreateOK struct {
	Payload *models.V0AppRespModel
}

func (o *AppCreateOK) Error() string {
	return fmt.Sprintf("[POST /apps/register][%d] appCreateOK  %+v", 200, o.Payload)
}

func (o *AppCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0AppRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppCreateBadRequest creates a AppCreateBadRequest with default headers values
func NewAppCreateBadRequest() *AppCreateBadRequest {
	return &AppCreateBadRequest{}
}

/*AppCreateBadRequest handles this case with default header values.

Bad Request
*/
type AppCreateBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AppCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/register][%d] appCreateBadRequest  %+v", 400, o.Payload)
}

func (o *AppCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppCreateUnauthorized creates a AppCreateUnauthorized with default headers values
func NewAppCreateUnauthorized() *AppCreateUnauthorized {
	return &AppCreateUnauthorized{}
}

/*AppCreateUnauthorized handles this case with default header values.

Unauthorized
*/
type AppCreateUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AppCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apps/register][%d] appCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *AppCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppCreateNotFound creates a AppCreateNotFound with default headers values
func NewAppCreateNotFound() *AppCreateNotFound {
	return &AppCreateNotFound{}
}

/*AppCreateNotFound handles this case with default header values.

Not Found
*/
type AppCreateNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AppCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /apps/register][%d] appCreateNotFound  %+v", 404, o.Payload)
}

func (o *AppCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppCreateInternalServerError creates a AppCreateInternalServerError with default headers values
func NewAppCreateInternalServerError() *AppCreateInternalServerError {
	return &AppCreateInternalServerError{}
}

/*AppCreateInternalServerError handles this case with default header values.

Internal Server Error
*/
type AppCreateInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AppCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /apps/register][%d] appCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *AppCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
