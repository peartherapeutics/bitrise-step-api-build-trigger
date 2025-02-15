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

// AppConfigCreateReader is a Reader for the AppConfigCreate structure.
type AppConfigCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AppConfigCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAppConfigCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAppConfigCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAppConfigCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAppConfigCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewAppConfigCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAppConfigCreateOK creates a AppConfigCreateOK with default headers values
func NewAppConfigCreateOK() *AppConfigCreateOK {
	return &AppConfigCreateOK{}
}

/*AppConfigCreateOK handles this case with default header values.

OK
*/
type AppConfigCreateOK struct {
	Payload models.V0AppConfigRespModel
}

func (o *AppConfigCreateOK) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/bitrise.yml][%d] appConfigCreateOK  %+v", 200, o.Payload)
}

func (o *AppConfigCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppConfigCreateBadRequest creates a AppConfigCreateBadRequest with default headers values
func NewAppConfigCreateBadRequest() *AppConfigCreateBadRequest {
	return &AppConfigCreateBadRequest{}
}

/*AppConfigCreateBadRequest handles this case with default header values.

Bad Request
*/
type AppConfigCreateBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AppConfigCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/bitrise.yml][%d] appConfigCreateBadRequest  %+v", 400, o.Payload)
}

func (o *AppConfigCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppConfigCreateUnauthorized creates a AppConfigCreateUnauthorized with default headers values
func NewAppConfigCreateUnauthorized() *AppConfigCreateUnauthorized {
	return &AppConfigCreateUnauthorized{}
}

/*AppConfigCreateUnauthorized handles this case with default header values.

Unauthorized
*/
type AppConfigCreateUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AppConfigCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/bitrise.yml][%d] appConfigCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *AppConfigCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppConfigCreateNotFound creates a AppConfigCreateNotFound with default headers values
func NewAppConfigCreateNotFound() *AppConfigCreateNotFound {
	return &AppConfigCreateNotFound{}
}

/*AppConfigCreateNotFound handles this case with default header values.

Not Found
*/
type AppConfigCreateNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AppConfigCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/bitrise.yml][%d] appConfigCreateNotFound  %+v", 404, o.Payload)
}

func (o *AppConfigCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppConfigCreateInternalServerError creates a AppConfigCreateInternalServerError with default headers values
func NewAppConfigCreateInternalServerError() *AppConfigCreateInternalServerError {
	return &AppConfigCreateInternalServerError{}
}

/*AppConfigCreateInternalServerError handles this case with default header values.

Internal Server Error
*/
type AppConfigCreateInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AppConfigCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/bitrise.yml][%d] appConfigCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *AppConfigCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
