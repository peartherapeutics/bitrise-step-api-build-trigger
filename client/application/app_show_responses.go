// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/peartherapeutics/bitrise-step-api-build-trigger/models"
)

// AppShowReader is a Reader for the AppShow structure.
type AppShowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AppShowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAppShowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAppShowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAppShowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAppShowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewAppShowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAppShowOK creates a AppShowOK with default headers values
func NewAppShowOK() *AppShowOK {
	return &AppShowOK{}
}

/*AppShowOK handles this case with default header values.

OK
*/
type AppShowOK struct {
	Payload *models.V0AppShowResponseModel
}

func (o *AppShowOK) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}][%d] appShowOK  %+v", 200, o.Payload)
}

func (o *AppShowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0AppShowResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppShowBadRequest creates a AppShowBadRequest with default headers values
func NewAppShowBadRequest() *AppShowBadRequest {
	return &AppShowBadRequest{}
}

/*AppShowBadRequest handles this case with default header values.

Bad Request
*/
type AppShowBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AppShowBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}][%d] appShowBadRequest  %+v", 400, o.Payload)
}

func (o *AppShowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppShowUnauthorized creates a AppShowUnauthorized with default headers values
func NewAppShowUnauthorized() *AppShowUnauthorized {
	return &AppShowUnauthorized{}
}

/*AppShowUnauthorized handles this case with default header values.

Unauthorized
*/
type AppShowUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AppShowUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}][%d] appShowUnauthorized  %+v", 401, o.Payload)
}

func (o *AppShowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppShowNotFound creates a AppShowNotFound with default headers values
func NewAppShowNotFound() *AppShowNotFound {
	return &AppShowNotFound{}
}

/*AppShowNotFound handles this case with default header values.

Not Found
*/
type AppShowNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AppShowNotFound) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}][%d] appShowNotFound  %+v", 404, o.Payload)
}

func (o *AppShowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppShowInternalServerError creates a AppShowInternalServerError with default headers values
func NewAppShowInternalServerError() *AppShowInternalServerError {
	return &AppShowInternalServerError{}
}

/*AppShowInternalServerError handles this case with default header values.

Internal Server Error
*/
type AppShowInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AppShowInternalServerError) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}][%d] appShowInternalServerError  %+v", 500, o.Payload)
}

func (o *AppShowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
